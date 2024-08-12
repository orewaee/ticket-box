package handlers

import "net/http"

type OptionsGuildHandler struct {
}

func NewOptionsGuildHandler() *OptionsGuildHandler {
	return &OptionsGuildHandler{}
}

// OPTIONS /guild
func (handler *OptionsGuildHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")
	writer.WriteHeader(http.StatusOK)
}
