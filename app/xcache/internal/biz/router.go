package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Router struct {
}

// GreeterUsecase is a Greeter usecase.
type RouterUsecase struct {
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{log: log.NewHelper(logger)}
}
