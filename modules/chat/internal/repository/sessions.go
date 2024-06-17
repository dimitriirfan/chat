package repository

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
)

type SessionsRepository interface {
	IsValidParticipant(ctx context.Context, sessionID string, participant entity.Participant) bool
	AddConnection(ctx context.Context, sessionID string, connection *entity.Connection)
	GetSessionConnections(ctx context.Context, sessionID string) []*entity.Connection
}

type SessionsRepo struct {
	// live sessions
	sessions map[string]*entity.Session
}

func NewSessionsRepository() *SessionsRepo {
	return &SessionsRepo{
		sessions: make(map[string]*entity.Session),
	}
}

func (r *SessionsRepo) IsValidParticipant(ctx context.Context, sessionID string, participant entity.Participant) bool {
	if sessionID == "session-2" {
		return false
	}
	return true
}

func (r *SessionsRepo) AddConnection(ctx context.Context, sessionID string, connection *entity.Connection) {
	session, ok := r.sessions[sessionID]
	if !ok {
		session = entity.NewSession()
		r.sessions[sessionID] = session
	}

	session.AddConnection(connection)
	session.AddParticipant(connection.GetParticipant())
}

func (r *SessionsRepo) GetSessionConnections(ctx context.Context, sessionID string) []*entity.Connection {
	return r.sessions[sessionID].GetConnections()
}
