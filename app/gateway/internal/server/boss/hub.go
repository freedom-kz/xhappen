package boss

import (
	protocol "xhappen/api/protocol/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type Hub struct {
	logger log.Logger

	boss  *Boss
	index int

	directDeliverMessage chan *DirectDeliverMessage
}

type DirectDeliverMessage struct {
	conn    *Connection
	deliver *protocol.Deliver
}

func newHub(boss *Boss) *Hub {
	return &Hub{
		boss: boss,
	}
}

func (h *Hub) Start() {
	connIndex := newConnectionIndex()

	for {
		select {
		case directDeliver := <-h.directDeliverMessage:
			if !connIndex.Has(directDeliver.conn) {
				continue
			}

			select {
			// case directMsg.conn.send <- directDeliver.deliver:
			// default:
			// 	close(directMsg.conn.send)
			// 	connIndex.Remove(directMsg.conn)
			}
		}
	}
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
	i.byUserId[connection.UserId] = append(i.byUserId[connection.UserId], connection)
	i.byConnection[connection] = len(i.byUserId[connection.UserId]) - 1
	i.byClientId[connection.ClientId] = connection
}

func (i *ConnectionIndex) Remove(connection *Connection) {
	userConnIndex, ok := i.byConnection[connection]
	if !ok {
		return
	}

	userConnections := i.byUserId[connection.UserId]
	last := userConnections[len(userConnections)-1]
	userConnections[userConnIndex] = last
	i.byUserId[connection.UserId] = userConnections[:len(userConnections)-1]
	i.byConnection[last] = userConnIndex

	delete(i.byConnection, connection)
	delete(i.byClientId, connection.ClientId)
}

func (i *ConnectionIndex) Has(connection *Connection) bool {
	_, ok := i.byConnection[connection]
	return ok
}

// ForUser returns all connections for a user ID.
func (i *ConnectionIndex) ForUser(id string) []*Connection {
	return i.byUserId[id]
}

// All returns the full webConn index.
func (i *ConnectionIndex) All() map[*Connection]int {
	return i.byConnection
}
