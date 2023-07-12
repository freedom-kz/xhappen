package boss

import (
	pb "xhappen/api/gateway/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Hub struct {
	logger log.Logger

	boss  *Boss
	index int

	deliverToHub           chan *deliverToHub
	syncToHub              chan *syncToHub
	broadcastToHub         chan *broadcastToHub
	actionToHub            chan *actionToHub
	disconnectedforceToHub chan *disconnectedforceToHub

	addConn    chan *Connection
	removeConn chan *Connection

	exitCh chan struct{}
}

type deliverToHub struct {
	done           chan *errors.Error
	deliverMessage *pb.DeliverRequest
}

type syncToHub struct {
	done        chan *errors.Error
	syncMessage *pb.SyncRequest
}

type broadcastToHub struct {
	done             chan *errors.Error
	broadcastMessage *pb.BroadcastRequest
}

type actionToHub struct {
	done          chan *errors.Error
	actionMessage *pb.ActionRequest
}

type disconnectedforceToHub struct {
	done                   chan *errors.Error
	disconnectForceMessage *pb.DisconnectForceRequest
}

func newHub(boss *Boss) *Hub {
	return &Hub{
		boss:                   boss,
		addConn:                make(chan *Connection, 0),
		removeConn:             make(chan *Connection, 0),
		syncToHub:              make(chan *syncToHub, 1000),
		deliverToHub:           make(chan *deliverToHub, 1000),
		broadcastToHub:         make(chan *broadcastToHub, 1000),
		disconnectedforceToHub: make(chan *disconnectedforceToHub, 1000),
	}
}

func (h *Hub) Start() {

	go func() {
		connIndex := newConnectionIndex()

		for {
			select {
			case connection := <-h.addConn:
				connIndex.Add(connection)
			case connection := <-h.removeConn:
				connIndex.Remove(connection)
			case <-h.deliverToHub:

			case <-h.syncToHub:

			case <-h.actionToHub:

			case <-h.broadcastToHub:

			case dicconnectforce := <-h.disconnectedforceToHub:
				conns := connIndex.ForUser(dicconnectforce.disconnectForceMessage.Userid)
				if dicconnectforce.disconnectForceMessage.Clientid != "" {
					for _, conn := range conns {
						if conn.ClientId == dicconnectforce.disconnectForceMessage.Clientid {
							conn.Close()
						}
					}
				} else {
					for _, conn := range conns {
						conn.Close()
					}
				}
				dicconnectforce.done <- nil
			case <-h.exitCh:
				goto exit
			}

		exit:
			h.logger.Log(log.LevelInfo, "msg", "hub exit", "index", h.index)
		}
	}()

}

func (h *Hub) Stop() {

}

func (h *Hub) AddConn(conn *Connection) {
	select {
	case h.addConn <- conn:
	case <-h.exitCh:
	}
}

func (h *Hub) RemoveConn(conn *Connection) {
	select {
	case h.removeConn <- conn:
	case <-h.exitCh:
	}
}

func (h *Hub) SendDeliverToConn(done chan *errors.Error, deliver *pb.DeliverRequest) {
	deliverToHub := &deliverToHub{
		done:           done,
		deliverMessage: deliver,
	}
	select {
	case h.deliverToHub <- deliverToHub:
	case <-h.exitCh:
	}
}

func (h *Hub) SendSyncToConn(done chan *errors.Error, sync *pb.SyncRequest) {
	syncToHub := &syncToHub{
		done:        done,
		syncMessage: sync,
	}
	select {
	case h.syncToHub <- syncToHub:
	case <-h.exitCh:
	}
}

func (h *Hub) SendBroadcastToConn(done chan *errors.Error, broadcast *pb.BroadcastRequest) {

	broadcastToHub := &broadcastToHub{
		done:             done,
		broadcastMessage: broadcast,
	}
	select {
	case h.broadcastToHub <- broadcastToHub:
	case <-h.exitCh:
	}
}

func (h *Hub) SendActionToConn(done chan *errors.Error, action *pb.ActionRequest) {
	actionToHub := &actionToHub{
		done:          done,
		actionMessage: action,
	}
	select {
	case h.actionToHub <- actionToHub:
	case <-h.exitCh:
	}
}

func (h *Hub) DisconnectedConn(done chan *errors.Error, disconnectConn *pb.DisconnectForceRequest) {
	disconnectToHub := &disconnectedforceToHub{
		done:                   done,
		disconnectForceMessage: disconnectConn,
	}
	select {
	case h.disconnectedforceToHub <- disconnectToHub:
	case <-h.exitCh:
	}
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
