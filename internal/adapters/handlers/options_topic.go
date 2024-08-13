package handlers

import "net/http"

type OptionsTopicHandler struct {
}

func NewOptionsTopicHandler() *OptionsTopicHandler {
	return &OptionsTopicHandler{}
}

func (handler *OptionsTopicHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")
	writer.WriteHeader(http.StatusOK)
}
