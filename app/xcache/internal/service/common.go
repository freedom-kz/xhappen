package service

var (
	servingState    = 0
	loaclServeIndex = -1
)

type ServingStaeListen struct {
}

func stateModify(state int, index int) {
	servingState = int(state)
	loaclServeIndex = int(index)
}
