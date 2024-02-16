package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Router struct {
}

type RouterUsecase struct {
	log *log.Helper
}

func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{log: log.NewHelper(logger)}
}
