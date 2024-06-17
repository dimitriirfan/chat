package handler

import (
	"net/http"

	"github.com/dimitriirfan/chat-2/modules/chat/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/chat/internal/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ChatWebsocketHandler struct {
	upgrader       *websocket.Upgrader
	sessionUsecase usecase.SessionUsecase
	userUsecase    usecase.UserUsecase
}

func NewChatWebsocketHandler(upgrader *websocket.Upgrader, sessionUsecase usecase.SessionUsecase, userUsecase usecase.UserUsecase) *ChatWebsocketHandler {
	return &ChatWebsocketHandler{
		upgrader:       upgrader,
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
	}
}

func (h *ChatWebsocketHandler) HandleSessionConnections(w http.ResponseWriter, r *http.Request) {
	sessionID := mux.Vars(r)["sessionId"]
	token := r.URL.Query().Get("token")
	ctx := r.Context()

	participant, err := h.userUsecase.GetParticipantFromToken(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	validParticipant := h.sessionUsecase.IsValidParticipant(ctx, sessionID, participant)
	if !validParticipant {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	connection := entity.NewConnection(conn)

	// TODO: Register client to a given session
	err = h.sessionUsecase.RegisterSessionConnection(ctx, sessionID, token, connection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Wait for a message from client
	for {
		// TODO: Broadcast a message to all clients in a given session
		_, msg, err := conn.ReadMessage()
		if err != nil {
			continue
		}

		h.sessionUsecase.BroadcastMessage(ctx, sessionID, connection, entity.Message{
			ID:      uuid.NewString(),
			Content: string(msg),
			Sender:  connection.GetParticipant(),
		})

	}

}
