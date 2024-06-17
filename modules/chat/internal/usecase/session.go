package usecase

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/repository"
)

type SessionUsecase interface {
	IsValidParticipant(ctx context.Context, sessionID string, participant entity.Participant) bool
	RegisterSessionConnection(ctx context.Context, sessionID string, participant entity.Participant, connection *entity.Connection) error
	BroadcastMessage(ctx context.Context, sessionID string, senderConnection *entity.Connection, message entity.Message)
}

type SessionUc struct {
	sessionsRepository repository.SessionsRepository
}

func NewSessionUsecase(sessionsRepository repository.SessionsRepository) *SessionUc {
	return &SessionUc{
		sessionsRepository: sessionsRepository,
	}
}

func (u *SessionUc) IsValidParticipant(ctx context.Context, sessionID string, participant entity.Participant) bool {
	return u.sessionsRepository.IsValidParticipant(ctx, sessionID, participant)
}

func (u *SessionUc) RegisterSessionConnection(ctx context.Context, sessionID string, participant entity.Participant, connection *entity.Connection) error {

	connection.RegisterParticipant(participant)

	u.sessionsRepository.AddConnection(ctx, sessionID, connection)

	return nil
}

func (u *SessionUc) BroadcastMessage(ctx context.Context, sessionID string, senderConnection *entity.Connection, message entity.Message) {
	// TODO: Validate sender session ID

	sessionConnections := u.sessionsRepository.GetSessionConnections(ctx, sessionID)

	for _, connection := range sessionConnections {
		if connection.ID == senderConnection.ID {
			continue
		}

		connection.Conn.WriteJSON(message)
	}

}
