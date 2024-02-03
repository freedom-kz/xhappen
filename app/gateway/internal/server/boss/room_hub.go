package boss

import "github.com/go-kratos/kratos/v2/log"

type RoomHub struct {
	logger log.Logger

	boss  *Boss
	index int

	exitCh chan struct{}
}
