package data

import (
	"github.com/go-kratos/kratos/v2/log"
)

type messageRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageRepo(data *Data, logger log.Logger) *messageRepo {
	return &messageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
