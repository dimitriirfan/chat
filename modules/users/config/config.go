package config

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
)

type UsersRepositoryInterface interface {
	GetUserFromToken(ctx context.Context, token string) (entity.User, error)
}
