package service

import (
	"context"

	pb "xhappen/api/gateway/v1"
	"xhappen/app/gateway/internal/server/boss"

	"github.com/go-kratos/kratos/v2/errors"
)

type GatewaySrvService struct {
	pb.UnimplementedGatewaySrvServer
	boss *boss.Boss
}

func NewGatewaySrvService(boss *boss.Boss) *GatewaySrvService {
	return &GatewaySrvService{
		boss: boss,
	}
}

// 接收同步消息
func (s *GatewaySrvService) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncReply, error) {
	done := make(chan *errors.Error)
	s.boss.SendSyncToHubConn(done, req)
	select {
	case err := <-done:
		if err != nil {
			return &pb.SyncReply{
				Ret: false,
				Err: &err.Status,
			}, nil
		} else {
			return &pb.SyncReply{
				Ret: true,
			}, nil
		}
	case <-ctx.Done(): //超时
		return &pb.SyncReply{
			Ret: false,
			Err: &errors.New(413, "SYNC_HANDLE_TIME_OUT", "ctx deadline").Status,
		}, nil
	}
}

// 接收下行消息
func (s *GatewaySrvService) Deliver(ctx context.Context, req *pb.DeliverRequest) (*pb.DeliverReply, error) {
	done := make(chan *errors.Error)
	s.boss.SendDeliverToHubConn(done, req)
	select {
	case err := <-done:
		if err != nil {
			return &pb.DeliverReply{
				Ret: false,
				Err: &err.Status,
			}, nil
		} else {
			return &pb.DeliverReply{
				Ret: true,
			}, nil
		}
	case <-ctx.Done():
		return &pb.DeliverReply{
			Ret: false,
			Err: &errors.Status{
				Code:    413,
				Reason:  "DELIVAE_HANDLE_TIME_OUT",
				Message: "ctx deadline",
			},
		}, nil
	}
}

// 广播
func (s *GatewaySrvService) Broadcast(ctx context.Context, req *pb.BroadcastRequest) (*pb.BroadcastReply, error) {
	done := make(chan *errors.Error)
	s.boss.SendBroadcastToHubConn(done, req)
	select {
	case err := <-done:
		if err != nil {
			return &pb.BroadcastReply{
				Ret: false,
				Err: &err.Status,
			}, nil
		} else {
			return &pb.BroadcastReply{
				Ret: true,
			}, nil
		}
	case <-ctx.Done():
		return &pb.BroadcastReply{
			Ret: false,
			Err: &errors.Status{
				Code:    413,
				Reason:  "BROADCAST_HANDLE_TIME_OUT",
				Message: "ctx deadline",
			},
		}, nil
	}
}

// 指令下发
func (s *GatewaySrvService) Action(ctx context.Context, req *pb.ActionRequest) (*pb.ActionReply, error) {
	done := make(chan *errors.Error)
	s.boss.SendActionToHubConn(done, req)
	select {
	case err := <-done:
		if err != nil {
			return &pb.ActionReply{
				Ret: false,
				Err: &err.Status,
			}, nil
		} else {
			return &pb.ActionReply{
				Ret: true,
			}, nil
		}
	case <-ctx.Done():
		return &pb.ActionReply{
			Ret: false,
			Err: &errors.Status{
				Code:    413,
				Reason:  "ACTION_HANDLE_TIME_OUT",
				Message: "ctx deadline",
			},
		}, nil
	}
}

// 强制下线
func (s *GatewaySrvService) Disconnectedforce(ctx context.Context, req *pb.DisconnectForceRequest) (*pb.DisconnectForceReply, error) {
	done := make(chan *errors.Error)
	s.boss.DisconnectedConn(done, req)
	select {
	case err := <-done:
		if err != nil {
			return &pb.DisconnectForceReply{
				Ret: false,
				Err: &err.Status,
			}, nil
		} else {
			return &pb.DisconnectForceReply{
				Ret: true,
			}, nil
		}
	case <-ctx.Done():
		return &pb.DisconnectForceReply{
			Ret: false,
			Err: &errors.Status{
				Code:    413,
				Reason:  "DISCONNECT_HANDLE_TIME_OUT",
				Message: "ctx deadline",
			},
		}, nil
	}
}
