// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"xhappen/app/xcache/internal/biz"
	"xhappen/app/xcache/internal/conf"
	"xhappen/app/xcache/internal/data"
	"xhappen/app/xcache/internal/server"
	"xhappen/app/xcache/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, registrar registry.Registrar, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	sequenceRepo := data.NewGreeterRepo(dataData, logger)
	sequenceUsecase := biz.NewSequenceUsecase(bootstrap, sequenceRepo, logger)
	sequenceService := service.NewSequenceService(sequenceUsecase)
	routerUsecase := biz.NewRouterUsecase(logger)
	routerService := service.NewRouterService(routerUsecase)
	grpcServer := server.NewGRPCServer(bootstrap, sequenceService, routerService, logger)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
