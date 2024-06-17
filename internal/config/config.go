package config

import (
	"net/http"

	"github.com/dimitriirfan/chat-2/modules/chat/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewWebsocketRouter() *mux.Router {
	r := mux.NewRouter()

	chatHandler := config.NewChatWebsocketHandler(&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	})

	r.HandleFunc("/ws/chat/{sessionId}", chatHandler.HandleSessionConnections)

	return r
}
