package httpresponse

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JSONWriter interface {
	WriteJSON(w http.ResponseWriter)
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (r *Response) WriteJSON(w http.ResponseWriter) {
	writeJSON(w, r, r.Code)
}

func (e *Error) WriteJSON(w http.ResponseWriter) {
	writeJSON(w, e, e.Code)
}

func writeJSON(w http.ResponseWriter, v interface{}, code int) {
	b, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
