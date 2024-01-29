package service

import (
	"context"

	v1 "xhappen/api/helloworld/v1"
	"xhappen/app/transfer/internal/biz"
)

// GreeterService is a greeter service.
type AuthService struct {
	
}

// NewGreeterService new a greeter service.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// SayHello implements helloworld.GreeterServer.
func (s *AuthService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
