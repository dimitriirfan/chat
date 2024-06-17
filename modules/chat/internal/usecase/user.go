package usecase

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/repository"
)

type UserUsecase interface {
	GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error)
}

type UserUc struct {
	usersRepository repository.UsersRepository
}

func NewUserUsecase(usersRepository repository.UsersRepository) *UserUc {
	return &UserUc{
		usersRepository: usersRepository,
	}
}

func (u *UserUc) GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error) {
	return u.usersRepository.GetParticipantFromToken(ctx, token)
}
