package handlers

import (
	"encoding/json"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type GetUserHandler struct {
	discordService ports.DiscordService
}

func NewGetUserHandler(discordService ports.DiscordService) *GetUserHandler {
	return &GetUserHandler{
		discordService: discordService,
	}
}

// GET /user
func (handler *GetUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id").(string)
	accessToken := request.Context().Value("accessToken").(string)

	currentUser, err := handler.discordService.GetCurrentUser(accessToken)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, err.Error())
		return
	}

	if currentUser.Id != id {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "identifier mismatch")
		return
	}

	user := &dto.User{
		Id:        currentUser.Id,
		Username:  currentUser.Username,
		AvatarUrl: currentUser.AvatarUrl,
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	utils.WriteBytes(writer, bytes)
}
