package config

import (
	"net/http"

	chatConfig "github.com/dimitriirfan/chat-2/modules/chat/config"
	usersConfig "github.com/dimitriirfan/chat-2/modules/users/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewWebsocketRouter() *mux.Router {
	dependencies := InitializeAllDependencies()
	r := mux.NewRouter()

	chatHandler := chatConfig.NewChatWebsocketHandler(&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}, dependencies.DB)

	r.HandleFunc("/ws/chat/{sessionId}", chatHandler.HandleSessionConnections)

	return r
}

func NewRESTRouter() *mux.Router {
	dependencies := InitializeAllDependencies()
	r := mux.NewRouter()
	usersHandler := usersConfig.NewUsersRESTHandler(dependencies.DB)
	r.HandleFunc("/v1/users/auth/register", usersHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/users/auth/login", usersHandler.Login).Methods(http.MethodPost)
	return r
}
