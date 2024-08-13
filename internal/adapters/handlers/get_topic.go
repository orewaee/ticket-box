package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
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
	topicId := request.PathValue("topic_id")
	if topicId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "missing topic_id")
		return
	}

	fmt.Println(topicId)

	topic, err := handler.topicService.GetTopicById(topicId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		utils.WriteString(writer, "topic not found")
		return
	}

	fmt.Println(topic)

	accessToken := request.Context().Value("accessToken").(string)
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
