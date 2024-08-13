package handlers

import "net/http"

type OptionsTopicsHandler struct {
}

func NewOptionsTopicsHandler() *OptionsTopicsHandler {
	return &OptionsTopicsHandler{}
}

func (handler *OptionsTopicsHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")
	writer.WriteHeader(http.StatusOK)
}
