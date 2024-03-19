package service

var (
	servingState = false
	serveIndex   = -1
)

func StateModify(state bool, index int) {
	servingState = state
	serveIndex = index
}
