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

func (s *GatewaySrvService) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
	return &pb.SyncResponse{}, nil
}
func (s *GatewaySrvService) Deliver(ctx context.Context, req *pb.DeliverRequest) (*pb.DeliverResponse, error) {
	return &pb.DeliverResponse{}, nil
}
func (s *GatewaySrvService) Disconnectedforce(ctx context.Context, req *pb.DisconnectForceRequest) (*pb.DisconnectForceResponse, error) {
	return &pb.DisconnectForceResponse{}, nil
}
