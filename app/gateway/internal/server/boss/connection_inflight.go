package boss

import (
	"fmt"
	"time"
)

func (connection *Connection) StartInflight(msg *Message) error {
	msg.pri = time.Now().UnixNano()
	err := connection.pushInFlightMessage(msg)
	if err != nil {
		return err
	}
	connection.addToInFlightPQ(msg)
	return nil
}

func (connection *Connection) addToInFlightPQ(msg *Message) {
	connection.inFlightMutex.Lock()
	connection.inFlightPQ.Push(msg)
	connection.inFlightMutex.Unlock()
}

func (connection *Connection) popInFlightMessage(sequence uint64) (*Message, error) {
	connection.inFlightMutex.Lock()
	msg, ok := connection.inFlightMessages[sequence]
	if !ok {
		connection.inFlightMutex.Unlock()
		return nil, fmt.Errorf("sequence not in flight")
	}
	delete(connection.inFlightMessages, sequence)
	connection.inFlightMutex.Unlock()
	return msg, nil
}

func (connection *Connection) removeFromInFlight(msg *Message) {
	connection.inFlightMutex.Lock()
	if msg.index == -1 {
		// this item has already been popped off the pqueue
		connection.inFlightMutex.Unlock()
		return
	}
	connection.inFlightPQ.Remove(msg.index)
	connection.inFlightMutex.Unlock()
}

func (connection *Connection) pushInFlightMessage(msg *Message) error {
	connection.inFlightMutex.Lock()
	_, ok := connection.inFlightMessages[msg.Sequence]
	if ok {
		connection.inFlightMutex.Unlock()
		return fmt.Errorf("sequence already in flight")
	}
	connection.inFlightMessages[msg.Sequence] = msg
	connection.inFlightMutex.Unlock()
	return nil
}

func (connection *Connection) StartActionInflight(msg *AMessage) error {
	msg.pri = time.Now().UnixNano()
	err := connection.pushActionInFlightMessage(msg)
	if err != nil {
		return err
	}
	connection.addToActionInFlightPQ(msg)
	return nil
}

func (connection *Connection) addToActionInFlightPQ(msg *AMessage) {
	connection.inFlightAMutex.Lock()
	connection.inFlightAPQ.Push(msg)
	connection.inFlightAMutex.Unlock()
}

func (connection *Connection) popActionInFlightMessage(id uint64) (*AMessage, error) {
	connection.inFlightAMutex.Lock()
	msg, ok := connection.inFlightAMessages[id]
	if !ok {
		connection.inFlightAMutex.Unlock()
		return nil, fmt.Errorf("id not in flight")
	}
	delete(connection.inFlightAMessages, id)
	connection.inFlightAMutex.Unlock()
	return msg, nil
}

func (connection *Connection) removeActionFromInFlight(msg *AMessage) {
	connection.inFlightAMutex.Lock()
	if msg.index == -1 {
		// this item has already been popped off the pqueue
		connection.inFlightAMutex.Unlock()
		return
	}
	connection.inFlightAPQ.Remove(msg.index)
	connection.inFlightAMutex.Unlock()
}

func (connection *Connection) pushActionInFlightMessage(msg *AMessage) error {
	connection.inFlightAMutex.Lock()
	_, ok := connection.inFlightMessages[uint64(msg.Id)]
	if ok {
		connection.inFlightAMutex.Unlock()
		return fmt.Errorf("pri already in flight")
	}
	connection.inFlightAMessages[uint64(msg.Id)] = msg
	connection.inFlightAMutex.Unlock()
	return nil
}
