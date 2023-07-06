package boss

type Hub struct {
	boss  *Boss
	index int
}

func newHub(boss *Boss) *Hub {
	return &Hub{
		boss: boss,
	}
}

func (h *Hub) Start() {

}

func (h *Hub) Stop() {

}

type ConnectionIndex struct {
	byUserId     map[string][]*Connection
	byConnection map[*Connection]int
	byClientId   map[string]*Connection
}

func newConnectionIndex() *ConnectionIndex {
	return &ConnectionIndex{
		byUserId:     make(map[string][]*Connection),
		byConnection: make(map[*Connection]int),
		byClientId:   make(map[string]*Connection),
	}
}

func (i *ConnectionIndex) Add(connection *Connection) {
	i.byUserId[connection.userId] = append(i.byUserId[connection.userId], connection)
	i.byConnection[connection] = len(i.byUserId[connection.userId]) - 1
	i.byClientId[connection.clientId] = connection
}

func (i *ConnectionIndex) Remove(connection *Connection) {
	userConnIndex, ok := i.byConnection[connection]
	if !ok {
		return
	}

	userConnections := i.byUserId[connection.userId]
	last := userConnections[len(userConnections)-1]
	userConnections[userConnIndex] = last
	i.byUserId[connection.userId] = userConnections[:len(userConnections)-1]
	i.byConnection[last] = userConnIndex

	delete(i.byConnection, connection)
	delete(i.byClientId, connection.clientId)
}

func (i *ConnectionIndex) Has(connection *Connection) bool {
	_, ok := i.byConnection[connection]
	return ok
}
