package handlers

import (
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type GetTokenHandler struct {
	discordService ports.DiscordService
}

func NewGetTokenHandler(discordService ports.DiscordService) *GetTokenHandler {
	return &GetTokenHandler{
		discordService: discordService,
	}
}

// GET /token?code=abc...
func (handler *GetTokenHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	code := request.URL.Query().Get("code")

	token, err := handler.discordService.GetTokenByCode(code)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, err.Error())
		return
	}

	utils.WriteString(writer, token)
}
