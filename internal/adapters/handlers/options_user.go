package handlers

import (
	"net/http"
)

type OptionsUserHandler struct {
}

func NewOptionsUserHandler() *OptionsUserHandler {
	return &OptionsUserHandler{}
}

// OPTIONS /user
func (handler *OptionsUserHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")
	writer.WriteHeader(http.StatusOK)
}
