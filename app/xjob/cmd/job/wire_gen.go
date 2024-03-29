// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"xhappen/app/xjob/internal/conf"
	"xhappen/app/xjob/internal/server"
	"xhappen/app/xjob/internal/service"
	"xhappen/pkg/event"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, registrar registry.Registrar, receiver event.Receiver, logger log.Logger) (*kratos.App, func(), error) {
	xJobService := service.NewXJobService(logger)
	grpcServer := server.NewGRPCServer(confServer, xJobService, logger)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
	}, nil
}
