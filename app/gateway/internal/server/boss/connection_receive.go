package boss

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	basic "xhappen/api/basic/v1"
	protocol "xhappen/api/protocol/v1"
	pb "xhappen/api/transfer/v1"
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
		//走到这里的一定是经过bind和auth业务的
		if connection.KeepAlive > 0 {
			connection.SetReadDeadline(time.Now().Add(connection.KeepAlive))
		} else {
			connection.SetReadDeadline(time.Now().Add(DEFAULT_KEEP_ALIVE))
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
	//从受boss管理的连接中移除数据，不会再产生下行业务数据
	connection.Boss.RemoveConnFromHub(connection)
	//正常关闭
	connection.Shutdown(true)
	return err
}

func (connection *Connection) exec(packet packets.ControlPacket) error {
	//走的这里的逻辑状态必定为同步中或者正常状态
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

/*
bind仅为临时状态，此状态下无业务逻辑可运行
客户端可以作为网络的验证
服务端验证客户端合法性
*/
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
		connection.logger.Log(log.LevelError, "msg", "1st read packet not bind", "body", packet.String())
		return err
	}

	bindAck := &packets.BindAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.BINDACK},
		BindAck: protocol.BindAck{
			BindRet:         true,
			ServerTimeStamp: uint64(time.Now().UnixMilli()),
		},
	}

	//版本校验
	if bind.CurVersion < connection.Boss.minSupportProtoVersion {
		bindAck.BindRet = false
		bindAck.Err = &basic.ErrorUpgrade("version %d not support.", bind.CurVersion).Status

		err = connection.Write(bindAck)
		if err != nil {
			return err
		}
		err = connection.Flush()
		if err != nil {
			return err
		} else {
			return fmt.Errorf(bindAck.Err.Reason)
		}
	}
	//客户端ID校验
	if IsValidDeviceId(bind.DeviceID) {
		bindAck.BindRet = false
		bindAck.Err = &basic.ErrorClientidRejected("invalid deviceId:%s.", bind.DeviceID).Status

		err = connection.Write(bindAck)
		if err != nil {
			return err
		}
		err = connection.Flush()
		if err != nil {
			return err
		} else {
			return fmt.Errorf(bindAck.Err.Reason)
		}
	}

	//业务调用
	in := &pb.BindRequest{
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
	//返回错误，关闭连接
	if !bindAck.BindRet {
		connection.logger.Log(log.LevelDebug, "msg", "Bind failure", "deviceId", connection.DeviceId, "err", reply.Err)
		return fmt.Errorf(bindAck.Err.Reason)
	}
	//填充信息
	connection.DeviceId = bind.DeviceID

	if bind.KeepAlive > uint64(connection.Boss.GetConfig().Main.MinKeepAlive.Seconds) &&
		bind.KeepAlive < uint64(connection.Boss.GetConfig().Main.MaxKeepAlive.Seconds) {
		connection.KeepAlive = time.Duration(bind.KeepAlive) * time.Second
	}
	connection.Version = int(bind.CurVersion)
	connection.Os = bind.DeviceType
	if bind.QueueSize > connection.Boss.GetConfig().Queue.MaxRdyCount {
		connection.ReadyCount = int64(connection.Boss.GetConfig().Queue.MaxRdyCount)
	} else {
		connection.ReadyCount = int64(bind.QueueSize)
	}
	connection.ReadyCount = int64(bind.QueueSize)
	connection.LoginType = bind.LoginType
	return nil
}

// 授权处理，此后会进入正常逻辑流程
// 目前的设计，匿名用户也会临时分配一个虚拟用户来进入读取业务，不允许写入逻辑
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
		DeviceID:       connection.DeviceId,
		ServerID:       connection.Boss.serverId,
		ConnectSequece: connection.connectSequence,
		LoginType:      connection.LoginType,
		DeviceType:     connection.Os,
		CurVersion:     int32(connection.Version),
		AuthInfo:       &auth.Auth,
	}

	//直进行验证后端调用
	reply, err := connection.Boss.passClient.Auth(context.Background(), in)
	if err != nil {
		ack.AuthRet = false
		ack.Err = &errors.FromError(err).Status
	} else {
		ack.AuthRet = reply.Ret
		ack.UID = reply.UID
		ack.Err = reply.Err
	}

	err = connection.Write(ack)
	if err != nil {
		return err
	}
	err = connection.Flush()
	if err != nil {
		return err
	}
	if !ack.AuthRet {
		connection.logger.Log(log.LevelDebug, "msg", "auth failure", "deviceId", connection.DeviceId, "err", reply.Err)
		return fmt.Errorf(ack.Err.Reason)
	}
	//信息填充
	connection.UserId = reply.UID
	connection.RoleType = auth.RoleType //用户认证角色
	connection.UserType = reply.UType
	connection.tokenExpire = reply.TokenExpire.AsTime()

	if len(connection.syncSessions) == 0 {
		connection.sendConnState(STATE_NORMAL)
	} else {
		for _, sid := range reply.Sessions {
			connection.syncSessions[sid] = struct{}{}
		}
		connection.sendConnState(STATE_SYNC)
	}
	//尝试关闭已有连接
	if conn := connection.Boss.GetConnFromHub(connection.DeviceId); conn != nil {
		conn.Close()
	}

	//纳入hub管理
	connection.Boss.AddConnToHub(connection)
	return nil
}

func (connection *Connection) processSubmit(submit *protocol.Submit) error {
	in := &pb.SubmitRequest{
		UserID:   connection.UserId,
		DeviceID: connection.DeviceId,
		Submit:   submit,
	}
	reply, err := connection.Boss.passClient.Submit(context.Background(), in)

	ack := &packets.SubmitAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.SUBMITACK},
		SubmitAck: protocol.SubmitAck{
			SubmitRet: true,
			ID:        submit.ID,
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
		ack.SessionID = reply.SessionID
	}

	//响应数据进行抢占写入
	err = connection.Write(ack)
	if err != nil {
		return err
	}
	err = connection.Flush()

	if err != nil {
		return err
	}
	//这里的失败为业务失败，不做处理，对当前业务无影响
	if !ack.SubmitAck.SubmitRet {
		return nil
	}

	//这里往发送协程发送一个假的deliver,用来触发按sequence发送deliver逻辑
	deliverGhost := &protocol.Deliver{
		Sequence: ack.Sequence,
	}
	connection.SendDeliverCh(deliverGhost)
	return nil

}

func (connection *Connection) processSyncAck(syncAck *protocol.SyncAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process SyncAck", "id", syncAck.ID)
	delete(connection.syncSessions, uint64(syncAck.ID))
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
		connection.sendConnState(STATE_NORMAL)
	}

	return nil
}

// 处理下行响应
func (connection *Connection) processDeliverAck(deliverAck *protocol.DeliverAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process deliver ack")
	err := connection.FinishDeliver(deliverAck.Sequence)
	if err != nil {
		connection.logger.Log(log.LevelError,
			"msg", "process deliver ack,maybe resend same sequecne",
			"user", connection.UserId,
			"sequecne", deliverAck.Sequence,
			"error", err)
	}
	//这里业务逻辑不存在关闭异常错误，仅打印日志
	return nil
}

func (connection *Connection) processAction(action *protocol.Action) error {
	//业务操作类型，待完善
	return nil
}

func (connection *Connection) processActionAck(actionAck *protocol.ActionAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process action ack")
	err := connection.FinishAction(uint64(actionAck.ID))
	if err != nil {
		//错误仅打印日志，非业务影响
		connection.logger.Log(log.LevelError,
			"msg", "process action ack",
			"userId", connection.UserId,
			"ackMsgId", actionAck.ID,
			"error", err)
	}

	return nil
}

// 处理ping消息
func (connection *Connection) processPing(ping *protocol.Ping) error {
	pong := &packets.PongPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.PONG},
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

// 标准退出
func (connection *Connection) processQuit(quit *protocol.Quit) error {
	return fmt.Errorf("standard exit")
}

// 清理deliver缓存
func (connection *Connection) FinishDeliver(sequence uint64) error {
	msg, err := connection.popInFlightMessage(sequence)
	if err != nil {
		return err
	}
	connection.removeFromInFlight(msg)
	connection.FinishedMessage()
	return nil
}

// 清理action缓存
func (connection *Connection) FinishAction(sequence uint64) error {
	msg, err := connection.popActionInFlightMessage(sequence)
	if err != nil {
		return err
	}
	connection.removeActionFromInFlight(msg)
	connection.FinishedMessage()
	return nil
}

// 消息计数和更新准备状态
func (connection *Connection) FinishedMessage() {
	atomic.AddUint64(&connection.FinishCount, 1)
	atomic.AddInt64(&connection.InFlightCount, -1)
	connection.tryUpdateReadyState()
}

func IsValidDeviceId(deviceId string) bool {
	return len(deviceId) >= 24 && len(deviceId) <= 36
}
