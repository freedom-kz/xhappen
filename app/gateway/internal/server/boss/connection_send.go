package boss

import (
	"fmt"
	"math"
	"time"

	protocol "xhappen/api/protocol/v1"
	"xhappen/app/gateway/internal/packets"

	"github.com/go-kratos/kratos/v2/log"
)

func (connection *Connection) messagePump(startedChan chan bool) {
	var (
		err         error
		flusherChan <-chan time.Time
		reSendChan  <-chan time.Time
		timeoutChan <-chan time.Time
		dChan       chan *protocol.Deliver = nil
		aChan       chan *protocol.Action  = nil

		timeoutTicker *time.Ticker
	)
	sChan := connection.syncCh

	reSendTicker := time.NewTicker(DEFAULT_RESEND_TICKER)
	outputBufferTicker := time.NewTicker(connection.FlushEvery)
	reSendChan = reSendTicker.C
	flushed := true
	close(startedChan)
	for {
		//同步中状态不走这些逻辑
		if !connection.IsReadyForMessages() {
			dChan = nil
			aChan = nil
		} else if flushed {
			flusherChan = nil
			dChan = connection.deliverCh
			aChan = connection.actionCh
		} else {
			flusherChan = outputBufferTicker.C
			dChan = connection.deliverCh
			aChan = connection.actionCh
		}

		select {
		case <-flusherChan:
			err = connection.Flush()
			connection.logger.Log(log.LevelError, "msg", "socket flush err")
		case d := <-dChan:
			//如果是期待sequence，则进行发送处理，否则进入等待队列
			if d.Sequence == connection.expectNextSequence {
				//判断真实deliver，如是则进行进行发送
				if len(d.Payload) != 0 {
					msg := NewMessage(d, int64(d.Sequence))
					err := connection.StartInflight(msg)
					if err != nil {
						connection.logger.Log(log.LevelError, "msg", "in startInflight", "err", err)
					}
					err = connection.WriteDeliver(msg)
					if err != nil {
						goto exit
					}
					flushed = false
				}
				connection.expectNextSequence++
				for i := 0; ; i++ {
					msg, _ := connection.toFlightPQ.PeekAndShift(int64(connection.expectNextSequence))
					if msg == nil {
						break
					}
					connection.removeFromToFlight(msg)
					connection.expectNextSequence++
					//过滤掉ghost deliver
					if len(msg.Payload) == 0 {
						continue
					}
					err := connection.StartInflight(msg)
					if err != nil {
						connection.logger.Log(log.LevelError, "msg", "in startInflight", "err", err)
					}
					err = connection.WriteDeliver(msg)
					if err != nil {
						goto exit
					}
					flushed = false
				}
				//这里在连续出现空缺的时候，可能超时不准确，待优化
				if len(connection.toFlightMessages) == 0 {
					timeoutChan = nil
					if timeoutTicker != nil {
						timeoutTicker.Stop()
					}
				} else {
					timeoutTicker.Stop()
					msg := connection.toFlightPQ[0]
					gap := time.Until(msg.deliveryTS.Add(-connection.MsgTimeout))
					if gap > connection.MsgTimeout {
						connection.logger.Log(log.LevelError, "msg", "wait expect sequence timeout")
						goto exit
					}
					timeoutTicker = time.NewTicker(gap)
					timeoutChan = timeoutTicker.C
				}
			} else {
				msg := NewMessage(d, int64(d.Sequence))
				err := connection.StartToflight(msg)
				if err != nil {
					connection.logger.Log(log.LevelError, "msg", "in startToflight", "err", err)
				}

				if timeoutTicker == nil {
					timeoutTicker = time.NewTicker(connection.MsgTimeout * 3)
					timeoutChan = timeoutTicker.C
				}
			}

		case a := <-aChan:
			msg := NewAMessage(a, time.Now().UnixNano())
			err := connection.StartActionInflight(msg)
			if err != nil {
				connection.logger.Log(log.LevelError, "msg", "in startInflight", "err", err)
			}
			err = connection.WriteAction(msg)
			if err != nil {
				goto exit
			}
			flushed = false
		case sync := <-sChan:
			ln := len(sync.Delivers)
			connection.processExpectSequence(sync.Delivers[ln].Sequence)
			packet := &packets.SyncPacket{
				FixedHeader: packets.FixedHeader{MessageType: packets.SYNC},
				Sync:        *sync,
			}
			err = connection.Write(packet)
			if err != nil {
				goto exit
			}
			flushed = false
		case <-reSendChan:
			err = connection.processDeliverReSend()
			if err != nil {
				goto exit
			}
			err = connection.processActionReSend()
			if err != nil {
				goto exit
			}
		case <-timeoutChan:
			if len(connection.toFlightMessages) != 0 {
				msg := connection.toFlightPQ[0]
				gap := time.Until(msg.deliveryTS.Add(-connection.MsgTimeout))
				if gap > connection.MsgTimeout {
					connection.logger.Log(log.LevelError, "msg", "wait expect sequence timeout")
					goto exit
				}
			}
		case <-connection.ReadyStateChan:
		case state, closed := <-connection.StateChan:
			if !closed {
				connection.state = STATE_QUIT
				goto exit
			}

			connection.state = state
			switch state {
			case STATE_NORMAL:
				sChan = nil
			case STATE_QUIT:
				goto exit
			}
		}
	}
exit:
	reSendTicker.Stop()
	outputBufferTicker.Stop()
	if timeoutTicker != nil {
		timeoutTicker.Stop()
	}
	connection.Shutdown(true)
	if err != nil {
		connection.logger.Log(log.LevelError, "msg", "send goroutine exit", "clientId", connection.ClientId, "userId", connection.UserId, "err", err)
	}
}

func (connection *Connection) processDeliverReSend() error {
	var msgs []*Message
	var err error
	for i := 0; ; i++ {
		msg, _ := connection.inFlightPQ.PeekAndShift(math.MaxInt64)
		if msg == nil {
			break
		}
		if msg.deliveryTS.Add(connection.MsgTimeout).Before(time.Now()) {
			break
		}
	}
	if len(msgs) != 0 {
		for _, msg := range msgs {
			connection.inFlightPQ.Push(msg)
			err = connection.WriteDeliver(msg)
			if err != nil {
				break
			}
		}
	}
	return err
}

func (connection *Connection) processActionReSend() error {
	var msgs []*AMessage
	var err error
	for i := 0; ; i++ {
		msg, _ := connection.inFlightAPQ.PeekAndShift(math.MaxInt64)
		if msg == nil {
			break
		}
		if msg.deliveryTS.Add(connection.MsgTimeout).Before(time.Now()) {
			break
		}
	}
	if len(msgs) != 0 {
		for _, msg := range msgs {
			//重新设置优先级
			msg.pri = time.Now().UnixNano()
			connection.inFlightAPQ.Push(msg)
			err = connection.WriteAction(msg)
			if err != nil {
				break
			}
		}
	}
	return err
}

func (connection *Connection) WriteDeliver(message *Message) error {
	message.Attempts++
	message.deliveryTS = time.Now()
	if message.Attempts > MESSAGE_RETRY_MAX {
		return fmt.Errorf("messgae send > %d", MESSAGE_RETRY_MAX)
	}

	if message.Attempts > 1 {
		connection.logger.Log(log.LevelInfo, "msg", "message retry")
	}

	packet := &packets.DeliverPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.DELIVER},
		Deliver:     *message.Deliver,
	}
	return connection.Write(packet)
}

func (connection *Connection) WriteAction(message *AMessage) error {
	message.Attempts++
	message.deliveryTS = time.Now()
	if message.Attempts > MESSAGE_RETRY_MAX {
		return fmt.Errorf("messgae send > %d", MESSAGE_RETRY_MAX)
	}

	if message.Attempts > 1 {
		connection.logger.Log(log.LevelInfo, "msg", "message retry")
	}

	packet := &packets.ActionPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.ACTION},
		Action:      *message.Action,
	}
	return connection.Write(packet)
}
