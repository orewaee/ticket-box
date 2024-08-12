package handlers

import (
	"encoding/json"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
	"strings"
)

type GetTopicHandler struct {
	topicService   ports.TopicService
	discordService ports.DiscordService
}

func NewGetTopicHandler(topicService ports.TopicService, discordService ports.DiscordService) *GetTopicHandler {
	return &GetTopicHandler{
		topicService:   topicService,
		discordService: discordService,
	}
}

// GET /topic/{topic_id}
func (handler *GetTopicHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	topicId := request.PathValue("topic_id")
	if topicId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "missing topic_id")
		return
	}

	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "missing token")
		return
	}

	topic, err := handler.topicService.GetTopicById(topicId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		utils.WriteString(writer, "topic not found")
		return
	}

	accessToken := strings.TrimPrefix(authorization, "Bearer ")
	isAdmin := handler.discordService.CurrentUserIsGuildAdmin(accessToken, topic.GuildId)

	if !isAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "you can't get this guild's topic")
		return
	}

	data := &dto.Topic{
		Id:          topic.Id,
		GuildId:     topic.GuildId,
		Emoji:       topic.Emoji,
		Name:        topic.Name,
		Description: topic.Description,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, "error getting topics: "+err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	utils.WriteBytes(writer, bytes)
}
