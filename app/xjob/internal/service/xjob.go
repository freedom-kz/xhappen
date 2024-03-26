package service

import (
	pb_xjob "xhappen/api/xjob/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type XJobService struct {
	pb_xjob.UnimplementedXJobServer
	log *log.Helper
}

func NewXJobService(
	logger log.Logger,
) *XJobService {
	return &XJobService{
		log: log.NewHelper(logger),
	}
}
