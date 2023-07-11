package server

import (
	"xhappen/app/gateway/internal/server/boss"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, boss.NewBoss)
