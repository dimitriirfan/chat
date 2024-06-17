package config

import (
	"context"
	"database/sql"

	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/users/internal/handler"
	"github.com/dimitriirfan/chat-2/modules/users/internal/repository"
	"github.com/dimitriirfan/chat-2/modules/users/internal/usecase"
)

type UsersExternalResource interface {
	IsTokenValid(ctx context.Context, token string) bool
	GetUserFromToken(ctx context.Context, token string) (entity.User, error)
}

func NewUsersRESTHandler(db *sql.DB) *handler.UsersRESTHandler {
	usersRepository := repository.NewUsersRepository(db)
	authUsecase := usecase.NewAuthUsecase(usersRepository)
	return handler.NewUsersRESTHandler(authUsecase)
}

func NewUsersExternalResourceHandler(db *sql.DB) *handler.UsersExternalResourceHandler {
	usersRepository := repository.NewUsersRepository(db)
	authUsecase := usecase.NewAuthUsecase(usersRepository)
	return handler.NewUsersExternalResourceHandler(authUsecase)
}
