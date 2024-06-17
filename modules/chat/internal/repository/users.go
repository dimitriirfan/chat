package repository

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
	usersModuleConfig "github.com/dimitriirfan/chat-2/modules/users/config"
)

type UsersRepositoryInterface interface {
	GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error)
}

type UsersRepo struct {
	usersModule usersModuleConfig.UsersRepositoryInterface
}

func NewUsersRepository(usersModule usersModuleConfig.UsersRepositoryInterface) *UsersRepo {
	return &UsersRepo{
		usersModule: usersModule,
	}
}

func (r *UsersRepo) GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error) {
	user, err := r.usersModule.GetUserFromToken(ctx, token)
	if err != nil {
		return entity.Participant{}, err
	}

	return entity.Participant{
		ID:       user.ID,
		Username: user.Username,
	}, nil

}
