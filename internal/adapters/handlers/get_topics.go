package handlers

import (
	"encoding/json"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type GetTopicsHandler struct {
	topicService   ports.TopicService
	discordService ports.DiscordService
}

func NewGetTopicsHandler(topicService ports.TopicService, discordService ports.DiscordService) *GetTopicsHandler {
	return &GetTopicsHandler{
		topicService:   topicService,
		discordService: discordService,
	}
}

// GET /topics/{guild_id}
func (handler *GetTopicsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	guildId := request.PathValue("guild_id")
	if guildId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "missing guild_id")
		return
	}

	accessToken := request.Context().Value("accessToken").(string)
	isAdmin := handler.discordService.CurrentUserIsGuildAdmin(accessToken, guildId)

	if !isAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "you can't get the list of topics of this guild")
		return
	}

	topics, err := handler.topicService.GetTopicsByGuildId(guildId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, "error getting list of topics: "+err.Error())
		return
	}

	data := make([]*dto.Topic, len(topics))
	for i, topic := range topics {
		data[i] = &dto.Topic{
			Id:          topic.Id,
			GuildId:     topic.GuildId,
			Emoji:       topic.Emoji,
			Name:        topic.Name,
			Description: topic.Description,
		}
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, "error getting list of topics: "+err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	utils.WriteBytes(writer, bytes)
}
