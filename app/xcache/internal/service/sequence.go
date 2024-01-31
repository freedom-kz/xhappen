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

func (s *SequenceService) GetSequenceByIds(ctx context.Context, in *v1.GetSequenceByIdsRequest) (*v1.GetSequenceByIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GetLocalSequenceByIds(ctx context.Context, in *v1.GetLocalSequenceByIdsRequest) (*v1.GetLocalSequenceByIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GetCurrentSequenceByIds(ctx context.Context, in *v1.GetCurrentSequenceByIdsRequest) (*v1.GetCurrentSequenceByIdsReply, error) {
	return nil, nil
}

func (s *SequenceService) GetLocalCurrentSequenceByIds(ctx context.Context, in *v1.GetLocalCurrentSequenceByIdsRequest) (*v1.GetLocalCurrentSequenceByIdsReply, error) {
	return nil, nil
}
