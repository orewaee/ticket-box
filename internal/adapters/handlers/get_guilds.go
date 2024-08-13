package handlers

import (
	"encoding/json"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type GetGuildsHandler struct {
	discordService ports.DiscordService
}

func NewGetGuildsHandler(discordService ports.DiscordService) *GetGuildsHandler {
	return &GetGuildsHandler{
		discordService: discordService,
	}
}

// GET /guilds
func (handler *GetGuildsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	accessToken := request.Context().Value("accessToken").(string)

	userGuilds, err := handler.discordService.GetCurrentUserGuilds(accessToken)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	guilds := make([]*dto.Guild, 0, len(userGuilds))
	for _, guild := range userGuilds {
		if !handler.discordService.IsAdmin(guild.Permissions) {
			continue
		}

		guilds = append(guilds, &dto.Guild{
			Id:      guild.Id,
			Name:    guild.Name,
			IconUrl: guild.IconUrl,
			WithBot: handler.discordService.BotExists(guild.Id),
		})
	}

	bytes, err := json.Marshal(guilds)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Cache-Control", "max-age=300")
	writer.Header().Set("Content-Type", "application/json")
	utils.WriteBytes(writer, bytes)
}
