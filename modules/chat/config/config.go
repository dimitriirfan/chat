package config

import (
	"database/sql"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/handler"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/repository"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/usecase"
	usersModule "github.com/dimitriirfan/chat-2/modules/users/config"
	"github.com/gorilla/websocket"
)

func NewChatWebsocketHandler(upgrader *websocket.Upgrader, db *sql.DB) *handler.ChatWebsocketHandler {
	sessionsRepository := repository.NewSessionsRepository()
	sessionUsecase := usecase.NewSessionUsecase(sessionsRepository)

	usersModuleResource := usersModule.NewUsersExternalResourceHandler(db)
	usersRepository := repository.NewUsersRepository(usersModuleResource)
	userUsecase := usecase.NewUserUsecase(usersRepository)
	return handler.NewChatWebsocketHandler(upgrader, sessionUsecase, userUsecase)
}
