package handler

import (
	"encoding/json"
	"net/http"

	httpresponse "github.com/dimitriirfan/chat-2/internal/http_response"
	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/users/internal/usecase"
)

type UsersRESTHandler struct {
	AuthUsecase usecase.AuthUsecase
}

func NewUsersRESTHandler(authUsecase usecase.AuthUsecase) *UsersRESTHandler {
	return &UsersRESTHandler{
		AuthUsecase: authUsecase,
	}
}

func (h *UsersRESTHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerPayload := new(entity.AuthRegisterBody)
	err := json.NewDecoder(r.Body).Decode(&registerPayload)
	if err != nil {
		httpresponse.NewError(http.StatusBadRequest, err.Error()).WriteJSON(w)
		return
	}

	ctx := r.Context()
	result, err := h.AuthUsecase.Register(ctx, registerPayload.Username, registerPayload.Password)
	if err != nil {
		httpresponse.NewError(http.StatusBadRequest, err.Error()).WriteJSON(w)
		return
	}

	httpresponse.NewResponse(http.StatusCreated, "ok", result).WriteJSON(w)
}

func (h *UsersRESTHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginPayload := new(entity.AuthLoginBody)
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		httpresponse.NewError(http.StatusBadRequest, err.Error()).WriteJSON(w)
		return
	}

	ctx := r.Context()
	result, err := h.AuthUsecase.Login(ctx, loginPayload.Username, loginPayload.Password)
	if err != nil {
		httpresponse.NewError(http.StatusBadRequest, err.Error()).WriteJSON(w)
		return
	}

	httpresponse.NewResponse(http.StatusOK, "ok", result).WriteJSON(w)
}
