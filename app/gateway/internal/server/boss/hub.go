package boss

import (
	pb "xhappen/api/gateway/v1"
	"xhappen/pkg/utils"

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
		addConn:                make(chan *Connection),
		removeConn:             make(chan *Connection),
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
			case deliverToHub := <-h.deliverToHub:
				if deliverToHub.deliverMessage.Clientid != "" {
					conn := connIndex.ForClientId(deliverToHub.deliverMessage.Clientid)
					if conn.UserId != deliverToHub.deliverMessage.Userid {
						deliverToHub.done <- errors.New(460, "NO_DEVICE_ONLINE", "NO_DEVICE_ONLINE")
					} else {
						conn.SendDeliverCh(deliverToHub.deliverMessage.Deliver)
					}
				} else {
					conns := connIndex.ForUser(deliverToHub.deliverMessage.Userid)
					for _, conn := range conns {
						if utils.StringInSlice(conn.ClientId, deliverToHub.deliverMessage.OmitClientids) {
							continue
						}
						conn.SendDeliverCh(deliverToHub.deliverMessage.Deliver)
					}
				}
			case syncToHub := <-h.syncToHub:
				conn := connIndex.ForClientId(syncToHub.syncMessage.Clientid)
				if conn.UserId != syncToHub.syncMessage.Userid || uint64(conn.ConnectTime.UnixNano()) != syncToHub.syncMessage.BindVersion {
					syncToHub.done <- errors.New(461, "DEVICE_NO_PAIR", "DEVICE_NO_PAIR")
				}
				conn.SendSyncCh(syncToHub.syncMessage.Sync)
			case actionToHub := <-h.actionToHub:
				if actionToHub.actionMessage.ClientId != "" {
					conn := connIndex.ForClientId(actionToHub.actionMessage.ClientId)
					if conn.UserId != actionToHub.actionMessage.Uid {
						actionToHub.done <- errors.New(460, "NO_DEVICE_ONLINE", "NO_DEVICE_ONLINE")
					} else {
						conn.SendActionCh(actionToHub.actionMessage.Action)
					}
				} else {
					conns := connIndex.ForUser(actionToHub.actionMessage.Uid)
					for _, conn := range conns {
						if utils.StringInSlice(conn.ClientId, actionToHub.actionMessage.OmitClientids) {
							continue
						}
						conn.SendActionCh(actionToHub.actionMessage.Action)
					}
				}
			case broadcastToHub := <-h.broadcastToHub:
				for conn := range connIndex.All() {
					if utils.StringInSlice(conn.UserId, broadcastToHub.broadcastMessage.OmitUserIds) ||
						utils.StringInSlice(conn.ClientId, broadcastToHub.broadcastMessage.OmitClientids) {
						continue
					}
					conn.SendDeliverCh(broadcastToHub.broadcastMessage.Deliver)
				}

			case dicconnectforce := <-h.disconnectedforceToHub:
				conns := connIndex.ForUser(dicconnectforce.disconnectForceMessage.Userid)
				if dicconnectforce.disconnectForceMessage.Clientid != "" {
					for _, conn := range conns {
						if conn.ClientId == dicconnectforce.disconnectForceMessage.Clientid {
							conn.Shutdown(false)
						}
					}
				} else {
					for _, conn := range conns {
						conn.Shutdown(false)
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

func (i *ConnectionIndex) ForUser(id string) []*Connection {
	return i.byUserId[id]
}

func (i *ConnectionIndex) ForClientId(clientId string) *Connection {
	return i.byClientId[clientId]
}

func (i *ConnectionIndex) All() map[*Connection]int {
	return i.byConnection
}
