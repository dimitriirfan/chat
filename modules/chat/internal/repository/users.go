package repository

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
	usersModuleConfig "github.com/dimitriirfan/chat-2/modules/users/config"
)

type UsersRepository interface {
	GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error)
}

type UsersRepo struct {
	usersModuleResource usersModuleConfig.UsersExternalResource
}

func NewUsersRepository(usersModuleResource usersModuleConfig.UsersExternalResource) *UsersRepo {
	return &UsersRepo{
		usersModuleResource: usersModuleResource,
	}
}

func (r *UsersRepo) GetParticipantFromToken(ctx context.Context, token string) (entity.Participant, error) {
	user, err := r.usersModuleResource.GetUserFromToken(ctx, token)
	if err != nil {
		return entity.Participant{}, err
	}

	return entity.Participant{
		ID:       user.ID,
		Username: user.Username,
	}, nil

}
