package handlers

import (
	"encoding/json"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type GetGuildHandler struct {
	discordService ports.DiscordService
}

func NewGetGuildHandler(discordService ports.DiscordService) *GetGuildHandler {
	return &GetGuildHandler{
		discordService: discordService,
	}
}

// GET /guild/{guild_id}
func (handler *GetGuildHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	guildId := request.PathValue("guild_id")
	if guildId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "missing guild_id")
		return
	}

	accessToken := request.Context().Value("accessToken").(string)

	userGuilds, err := handler.discordService.GetCurrentUserGuilds(accessToken)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var guild *dto.Guild = nil
	ok := false
	for _, userGuild := range userGuilds {
		if userGuild.Id != guildId {
			continue
		}

		guild = &dto.Guild{
			Id:      userGuild.Id,
			Name:    userGuild.Name,
			IconUrl: userGuild.IconUrl,
			WithBot: handler.discordService.BotExists(guildId),
		}

		ok = true

		break
	}

	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		utils.WriteString(writer, "guild not found")
		return
	}

	bytes, err := json.Marshal(guild)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Cache-Control", "max-age=300")
	writer.Header().Set("Content-Type", "application/json")
	utils.WriteBytes(writer, bytes)
}
