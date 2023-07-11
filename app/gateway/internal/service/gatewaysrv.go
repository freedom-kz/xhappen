package service

import (
	"context"

	pb "xhappen/api/gateway/v1"
)

type GatewaySrvService struct {
	pb.UnimplementedGatewaySrvServer
}

func NewGatewaySrvService() *GatewaySrvService {
	return &GatewaySrvService{}
}

func (s *GatewaySrvService) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncReply, error) {
	return &pb.SyncReply{}, nil
}

func (s *GatewaySrvService) Deliver(ctx context.Context, req *pb.DeliverRequest) (*pb.DeliverReply, error) {
	return &pb.DeliverReply{}, nil
}

//广播
func Broadcast(ctx context.Context, req *pb.BroadcastRequest) (*pb.BroadcastReply, error) {
	return &pb.BroadcastReply{}, nil
}

//指令
func Action(ctx context.Context, req *pb.ActionRequest) (*pb.ActionReply, error) {
	return &pb.ActionReply{}, nil
}

func (s *GatewaySrvService) Disconnectedforce(ctx context.Context, req *pb.DisconnectForceRequest) (*pb.DisconnectForceReply, error) {
	return &pb.DisconnectForceReply{}, nil
}
