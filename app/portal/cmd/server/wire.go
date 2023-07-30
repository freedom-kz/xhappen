//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"xhappen/app/portal/internal/biz"
	"xhappen/app/portal/internal/conf"
	"xhappen/app/portal/internal/data"
	"xhappen/app/portal/internal/event"
	"xhappen/app/portal/internal/server"
	"xhappen/app/portal/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, registry.Registrar, event.Sender) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
