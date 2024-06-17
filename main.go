package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketServer struct {
	Upgrader websocket.Upgrader
	Sessions map[string]Session
}

type WebsocketConnection struct {
	ID   string
	Conn *websocket.Conn
}

type Session struct {
	Clients map[string]WebsocketConnection
}

func NewWebSocketServer() *WebsocketServer {
	return &WebsocketServer{
		Sessions: map[string]Session{
			"session-1": Session{
				Clients: map[string]WebsocketConnection{},
			},
			"session-2": Session{
				Clients: map[string]WebsocketConnection{},
			},
			"session-3": Session{
				Clients: map[string]WebsocketConnection{},
			},
			"session-4": Session{
				Clients: map[string]WebsocketConnection{},
			},
		},
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (s *WebsocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionId"]

	session, ok := s.Sessions[sessionID]
	if !ok {
		w.Write([]byte("server busy"))
	}

	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer conn.Close()

	clientID := uuid.NewString()
	session.Clients[clientID] = WebsocketConnection{
		ID:   clientID,
		Conn: conn,
	}

	for {
		_, buf, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		for _, client := range session.Clients {
			if client.ID == clientID {
				continue
			}

			err = client.Conn.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}

		fmt.Printf("msg: %v\n", string(buf))

	}
}

func main() {
	ws := NewWebSocketServer()

	r := mux.NewRouter()
	r.HandleFunc("/chat/{sessionId}", ws.HandleConnections)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("server started on port: %v", 8080)
	httpServer.ListenAndServe()

}
