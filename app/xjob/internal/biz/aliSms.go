package biz

import "github.com/go-kratos/kratos/v2/log"

type AliSMSUseCase struct {
	log *log.Helper
}

func NewAliSMSUseCase(logger log.Logger) *AliSMSUseCase {
	return &AliSMSUseCase{
		log: log.NewHelper(log.With(logger, "module", "usecase/alisms")),
	}
}
