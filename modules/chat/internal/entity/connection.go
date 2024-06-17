package entity

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ID           string
	Conn         *websocket.Conn
	_participant Participant
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		ID:   uuid.NewString(),
		Conn: conn,
	}
}

func (c *Connection) RegisterParticipant(sender Participant) {
	c._participant = sender
}

func (c *Connection) GetParticipant() Participant {
	return c._participant
}
