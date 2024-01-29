package boss

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	pb "xhappen/api/business/v1"
	protocol "xhappen/api/protocol/v1"
	"xhappen/app/gateway/internal/packets"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

func (connection *Connection) packetProcess() error {
	var err error
	var packet packets.ControlPacket
	for {
		if connection.state == STATE_QUIT {
			break
		}

		packet, err = connection.ReadPacket()
		if err != nil {
			break
		}

		err = connection.exec(packet)
		if err != nil {
			break
		}

	}
	connection.Shutdown(true)
	return err
}

func (connection *Connection) exec(packet packets.ControlPacket) error {
	var err error
	switch pkt := packet.(type) {
	case *packets.SubmitPacket:
		err = connection.processSubmit(&pkt.Submit)
	case *packets.SyncAckPacket:
		err = connection.processSyncAck(&pkt.SyncAck)
	case *packets.DeliverAckPacket:
		err = connection.processDeliverAck(&pkt.DeliverAck)
	case *packets.ActionPacket:
		err = connection.processAction(&pkt.Action)
	case *packets.ActionAckPacket:
		err = connection.processActionAck(&pkt.ActionAck)
	case *packets.PingPacket:
		err = connection.processPing(&pkt.Ping)
	case *packets.QuitPacket:
		err = connection.processQuit(&pkt.Quit)
	default:
		err = fmt.Errorf("invalid message type %s", pkt.String())
	}
	return err
}

func (connection *Connection) ReadPacket() (cp packets.ControlPacket, err error) {
	return packets.ReadPacket(connection)
}

func (connection *Connection) processBind() error {
	connection.sendConnState(STATE_BIND)
	if connection.DailTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.DailTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_DAIL_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.logger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	bind, ok := packet.(*packets.BindPacket)
	if !ok {
		connection.logger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	bindAck := &packets.BindAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.BINDACK},
		BindAck: protocol.BindAck{
			BindRet:         true,
			ServerTimeStamp: uint64(time.Now().UnixMilli()),
		},
	}

	in := &pb.BindRequest{
		ServerID: connection.Boss.serverId,
		BindInfo: &bind.Bind,
	}

	reply, err := connection.Boss.passClient.Bind(context.Background(), in)
	if err != nil {
		bindAck.BindRet = false
		bindAck.Err = &errors.FromError(err).Status
	} else {
		bindAck.BindRet = reply.Ret
		bindAck.Err = reply.Err
	}

	err = connection.Write(bindAck)
	if err != nil {
		return err
	}
	err = connection.Flush()
	if err != nil {
		return err
	}
	if !bindAck.BindRet {
		connection.logger.Log(log.LevelDebug, "msg", "Bind reply fail", "clientId", connection.ClientId, "err", reply.Err)
		return fmt.Errorf(bindAck.Err.Reason)
	}
	//填充信息
	connection.ClientId = bind.ClientID
	connection.KeepAlive = time.Duration(bind.KeepAlive)
	connection.Version = strconv.Itoa(int(bind.CurVersion))
	connection.Os = bind.DeviceType
	if bind.QueueSize > connection.Boss.GetConfig().Queue.MaxRdyCount {
		bind.QueueSize = connection.Boss.GetConfig().Queue.MaxRdyCount
	}
	connection.ReadyCount = int64(bind.QueueSize)
	connection.LoginType = bind.LoginType
	return nil
}

func (connection *Connection) processAuth() error {
	connection.sendConnState(STATE_AUTH)
	if connection.ReadTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.ReadTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_READ_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.Boss.logger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	auth, ok := packet.(*packets.AUTHPacket)
	if !ok {
		connection.Boss.logger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	ack := &packets.AuthAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.AUTHACK},
		AuthAck: protocol.AuthAck{
			AuthRet: true,
		},
	}

	in := &pb.AuthRequest{
		ClientId:    connection.ClientId,
		BindVersion: uint64(connection.ConnectTime.UnixNano()),
		AuthInfo:    &auth.Auth,
	}

	reply, err := connection.Boss.passClient.Auth(context.Background(), in)
	if err != nil {
		ack.AuthRet = false
		ack.Err = &errors.FromError(err).Status
	} else {
		ack.AuthRet = reply.Ret
		ack.Uid = reply.Uid
		ack.Err = reply.Err
	}

	err = connection.Write(auth)
	if err != nil {
		return err
	}
	err = connection.Flush()
	if err != nil {
		return err
	}
	if !ack.AuthRet {
		return fmt.Errorf(ack.Err.Reason)
	}
	//信息填充
	connection.UserId = reply.Uid
	connection.RoleType = reply.Role
	connection.UserType = reply.UType
	connection.tokenExpire = reply.TokenExpire.AsTime()
	for _, sid := range reply.Sessions {
		connection.syncSessions[sid] = struct{}{}
	}
	if len(connection.syncSessions) == 0 {
		connection.sendConnState(STATE_NORMAL)
	} else {
		connection.sendConnState(STATE_SYNC)
	}
	connection.Boss.AddConnToHub(connection)
	return nil
}

func (connection *Connection) processSubmit(submit *protocol.Submit) error {
	in := &pb.SubmitRequest{
		Userid:   connection.UserId,
		Clientid: connection.ClientId,
		Submit:   submit,
	}
	reply, err := connection.Boss.passClient.Submit(context.Background(), in)

	ack := &packets.SubmitAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.SUBMITACK},
		SubmitAck: protocol.SubmitAck{
			SubmitRet: true,
			Id:        submit.Id,
			Timestamp: uint64(time.Now().UnixMilli()),
		},
	}
	if err != nil {
		ack.SubmitRet = false
		ack.Err = &errors.FromError(err).Status
	} else if !reply.Ret {
		ack.SubmitRet = false
		ack.Err = reply.Err
	} else {
		ack.SubmitRet = reply.Ret
		ack.Sequence = reply.Sequence
		ack.SessionId = reply.SessionId
	}

	err = connection.Write(ack)
	if err != nil {
		return err
	}
	//这里往发送协程发送一个假的deliver,用来处理按sequence发送deliver逻辑
	deliverGhost := &protocol.Deliver{
		Sequence: ack.Sequence,
	}
	connection.SendDeliverCh(deliverGhost)
	return nil
}

func (connection *Connection) processSyncAck(syncAck *protocol.SyncAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process SyncAck", "id", syncAck.Id)
	delete(connection.syncSessions, uint64(syncAck.Id))
	if len(connection.syncSessions) == 0 {
		packet := &packets.SyncConfirmPacket{FixedHeader: packets.FixedHeader{MessageType: packets.SYNCCONFIRM}}
		err := connection.Write(packet)
		if err != nil {
			return err
		}
		err = connection.Flush()
		if err != nil {
			return err
		}
	}
	if len(connection.syncSessions) == 0 {
		connection.sendConnState(STATE_NORMAL)
	}
	return nil
}

func (connection *Connection) processDeliverAck(deliverAck *protocol.DeliverAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process deliver ack")
	err := connection.FinishDeliver(deliverAck.Sequence)
	connection.logger.Log(log.LevelError, "msg", "process deliver ack", "error", err)
	return nil
}

func (connection *Connection) processAction(action *protocol.Action) error {
	return nil
}

func (connection *Connection) processActionAck(actionAck *protocol.ActionAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process action ack")
	err := connection.FinishAction(uint64(actionAck.Id))
	connection.logger.Log(log.LevelError, "msg", "process action ack", "error", err)
	return nil
}

func (connection *Connection) processPing(ping *protocol.Ping) error {
	pong := &packets.PongPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.SUBMITACK},
		Pong:        protocol.Pong{},
	}
	err := connection.Write(pong)
	if err != nil {
		return err
	}
	err = connection.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (connection *Connection) processQuit(quit *protocol.Quit) error {
	return fmt.Errorf("normal exit")
}

func (connection *Connection) FinishDeliver(sequence uint64) error {
	msg, err := connection.popInFlightMessage(sequence)
	if err != nil {
		return err
	}
	connection.removeFromInFlight(msg)
	connection.FinishedMessage()
	return nil
}
func (connection *Connection) FinishAction(sequence uint64) error {
	msg, err := connection.popActionInFlightMessage(sequence)
	if err != nil {
		return err
	}
	connection.removeActionFromInFlight(msg)
	connection.FinishedMessage()
	return nil
}

func (connection *Connection) FinishedMessage() {
	atomic.AddUint64(&connection.FinishCount, 1)
	atomic.AddInt64(&connection.InFlightCount, -1)
	connection.tryUpdateReadyState()
}
