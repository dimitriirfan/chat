package usecase

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
)

type UserUsecase interface {
	GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error)
}

type UserUc struct {
}

func NewUserUsecase() *UserUc {
	return &UserUc{}
}

func (u *UserUc) GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error) {
	return entity.Participant{
		Username: token,
	}, nil
}
