package service

import (
	"context"

	v1 "xhappen/api/sequence/v1"
	"xhappen/app/xcache/internal/biz"
)

type SequenceService struct {
	v1.UnimplementedSequenceServer

	uc *biz.SequenceUsecase
}

// NewGreeterService new a greeter service.
func NewSequenceService(uc *biz.SequenceUsecase) *SequenceService {
	return &SequenceService{uc: uc}
}

func (s *SequenceService) GenSequenceByUserIds(context.Context, *v1.GenSequenceByUserIdsRequest) (*v1.GenSequenceByUserIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GetCurrentSequenceByUserIds(ctx context.Context, in *v1.GetCurrentSequenceByUserIdsRequest) (*v1.GetCurrentSequenceByUserIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GenRoomSequenceByRoomIds(ctx context.Context, in *v1.GenRoomSequenceByRoomIdsRequest) (*v1.GenRoomSequenceByRoomIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GetCurrentRoomSequenceByRoomIds(ctx context.Context, in *v1.GetCurrentSequenceByRoomIdsRequest) (*v1.GetCurrentSequenceByRoomIdsReply, error) {
	return nil, nil
}
