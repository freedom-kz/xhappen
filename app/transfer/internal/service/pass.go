package service

import (
	"context"
	v1 "xhappen/api/transfer/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type PassService struct {
	v1.UnimplementedPassServer

	log *log.Helper
}

func NewPassService(logger log.Logger) *PassService {
	return &PassService{
		log: log.NewHelper(logger),
	}
}

func (s *PassService) Bind(ctx context.Context, in *v1.BindRequest) (*v1.BindReply, error) {
	
	return &v1.BindReply{}, nil
}

func (s *PassService) Auth(ctx context.Context, in *v1.AuthRequest) (*v1.AuthReply, error) {
	return &v1.AuthReply{}, nil
}

func (s *PassService) Submit(ctx context.Context, in *v1.SubmitRequest) (*v1.SubmitReply, error) {
	return nil, nil
}
func (s *PassService) Action(ctx context.Context, in *v1.ActionRequest) (*v1.ActionReply, error) {
	return nil, nil
}
func (s *PassService) Quit(ctx context.Context, in *v1.QuitRequest) (*v1.QuitReply, error) {
	return nil, nil
}


