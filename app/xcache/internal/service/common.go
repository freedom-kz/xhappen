package service

var (
	servingState = false
	serveIndex   = -1
)

type ServingStaeListen struct {
}

func StateModify(state bool, index int) {
	servingState = state
	serveIndex = index
}
