package config

import (
	"github.com/dimitriirfan/chat-2/modules/chat/internal/handler"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/repository"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/usecase"
	"github.com/gorilla/websocket"
)

func NewChatWebsocketHandler(upgrader *websocket.Upgrader) *handler.ChatWebsocketHandler {
	sessionsRepository := repository.NewSessionsRepository()
	sessionUsecase := usecase.NewSessionUsecase(sessionsRepository)
	userUsecase := usecase.NewUserUsecase()
	return handler.NewChatWebsocketHandler(upgrader, sessionUsecase, userUsecase)
}
