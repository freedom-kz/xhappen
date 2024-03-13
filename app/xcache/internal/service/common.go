package service

var (
	servingState = 0
	serveIndex   = -1
)

type ServingStaeListen struct {
}

func stateModify(state int, index int) {
	servingState = int(state)
	serveIndex = int(index)
}
