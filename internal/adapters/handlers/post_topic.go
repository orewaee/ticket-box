package handlers

import (
	"encoding/json"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/ticket-box/internal/app/domain"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/dto"
	"github.com/orewaee/ticket-box/internal/utils"
	"io"
	"net/http"
	"strings"
)

type PostTopicHandler struct {
	topicService   ports.TopicService
	discordService ports.DiscordService
}

func NewPostTopicHandler(topicService ports.TopicService, discordService ports.DiscordService) *PostTopicHandler {
	return &PostTopicHandler{
		topicService:   topicService,
		discordService: discordService,
	}
}

// POST /topic
func (handler *PostTopicHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "missing token")
		return
	}

	accessToken := strings.TrimPrefix(authorization, "Bearer ")

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "failed to read the request body: "+err.Error())
		return
	}

	addTopicRequest := new(dto.AddTopicRequest)
	if err := json.Unmarshal(bytes, addTopicRequest); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		utils.WriteString(writer, "failed to read the request body: "+err.Error())
		return
	}

	// validate body

	isAdmin := handler.discordService.CurrentUserIsGuildAdmin(accessToken, addTopicRequest.GuildId)
	if !isAdmin {
		writer.WriteHeader(http.StatusUnauthorized)
		utils.WriteString(writer, "you can't create a topic for this guild")
		return
	}

	id := gonanoid.Must(8)
	newTopic := &domain.Topic{
		Id:          id,
		GuildId:     addTopicRequest.GuildId,
		Emoji:       addTopicRequest.Emoji,
		Name:        addTopicRequest.Name,
		Description: addTopicRequest.Description,
	}

	if err := handler.topicService.AddTopic(newTopic); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.WriteString(writer, "failed to create topic: "+err.Error())
		return
	}

	writer.WriteHeader(http.StatusCreated)
	utils.WriteString(writer, id)
}
