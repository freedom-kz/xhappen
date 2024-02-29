package cluster

import "github.com/google/wire"

// import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewCluster)
