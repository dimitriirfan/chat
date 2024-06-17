package entity

import (
	"sync"

	"github.com/google/uuid"
)

type Session struct {
	mu           sync.Mutex
	ID           string
	Participants map[string]Participant
	Connections  map[string]*Connection
}

func NewSession() *Session {
	return &Session{
		ID:           uuid.NewString(),
		Participants: make(map[string]Participant),
		Connections:  make(map[string]*Connection),
	}
}

func (s *Session) AddParticipant(participant Participant) {
	s.mu.Lock()
	s.Participants[participant.ID] = participant
	s.mu.Unlock()
}

func (s *Session) AddConnection(connection *Connection) {
	s.mu.Lock()
	s.Connections[connection.ID] = connection
	s.mu.Unlock()
}

func (s *Session) RemoveConnection(connection Connection) {
	s.mu.Lock()
	delete(s.Connections, connection.ID)
	s.mu.Unlock()
}

func (s *Session) RemoveParticipant(participant Participant) {
	s.mu.Lock()
	delete(s.Participants, participant.ID)
	s.mu.Unlock()
}

func (s *Session) GetConnections() []*Connection {
	s.mu.Lock()
	defer s.mu.Unlock()

	var connections []*Connection
	for _, connection := range s.Connections {
		connections = append(connections, connection)
	}

	return connections
}
