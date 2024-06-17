package handler

import (
	"context"

	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/users/internal/usecase"
)

type UsersExternalResourceHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewUsersExternalResourceHandler(authUsecase usecase.AuthUsecase) *UsersExternalResourceHandler {
	return &UsersExternalResourceHandler{
		authUsecase: authUsecase,
	}
}

func (ex *UsersExternalResourceHandler) IsTokenValid(ctx context.Context, token string) bool {
	return ex.authUsecase.IsTokenValid(ctx, token)
}
func (ex *UsersExternalResourceHandler) GetUserFromToken(ctx context.Context, token string) (entity.User, error) {
	return ex.authUsecase.GetUserFromToken(ctx, token)
}
