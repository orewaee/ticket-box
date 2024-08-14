package handlers

import (
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
)

type DeleteTopicHandler struct {
	topicService   ports.TopicService
	discordService ports.DiscordService
}

func NewDeleteTopicHandler(topicService ports.TopicService, discordService ports.DiscordService) *DeleteTopicHandler {
	return &DeleteTopicHandler{
		topicService:   topicService,
		discordService: discordService,
	}
}

// DELETE /topic/{topic_id}
func (handler *DeleteTopicHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	topicId := request.PathValue("topic_id")
	if topicId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "missing topic_id")
		return
	}

	topic, err := handler.topicService.GetTopicById(topicId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		utils.WriteString(writer, "topic not found")
		return
	}

	accessToken := request.Context().Value("accessToken").(string)
	isAdmin := handler.discordService.CurrentUserIsGuildAdmin(accessToken, topic.GuildId)

	if !isAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "you can't delete this guild's topic")
		return
	}

	if err := handler.topicService.RemoveTopicById(topicId); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, "failed to delete topic")
		return
	}

	writer.WriteHeader(http.StatusOK)
}
