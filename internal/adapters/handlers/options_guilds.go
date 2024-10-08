package handlers

import "net/http"

type OptionsGuildsHandler struct {
}

func NewOptionsGuildsHandler() *OptionsGuildsHandler {
	return &OptionsGuildsHandler{}
}

// OPTIONS /guilds
func (handler *OptionsGuildsHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")
	writer.WriteHeader(http.StatusOK)
}
