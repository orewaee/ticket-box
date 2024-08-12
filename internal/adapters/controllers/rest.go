package controllers

import (
	"github.com/orewaee/ticket-box/internal/adapters/handlers"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"net/http"
	"time"
)

type RestController struct {
	httpServer     *http.Server
	topicService   ports.TopicService
	discordService ports.DiscordService
}

func NewRestController(addr string, topicService ports.TopicService, discordService ports.DiscordService) *RestController {
	mux := http.NewServeMux()

	mux.Handle("GET /token", handlers.NewGetTokenHandler(discordService))

	mux.Handle("GET /user", AuthMiddleware(handlers.NewGetUserHandler(discordService), discordService))
	mux.Handle("OPTIONS /user", handlers.NewOptionsUserHandler())

	mux.Handle("GET /guild/{guild_id}", AuthMiddleware(handlers.NewGetGuildHandler(discordService), discordService))
	mux.Handle("OPTIONS /guild/{guild_id}", handlers.NewOptionsGuildHandler())

	mux.Handle("GET /guilds", AuthMiddleware(handlers.NewGetGuildsHandler(discordService), discordService))
	mux.Handle("OPTIONS /guilds", handlers.NewOptionsGuildsHandler())

	mux.Handle("POST /topic", handlers.NewPostTopicHandler(topicService, discordService))
	mux.Handle("GET /topic/{topic_id}", handlers.NewGetTopicHandler(topicService, discordService))
	mux.Handle("GET /topics/{guild_id}", handlers.NewGetTopicsHandler(topicService, discordService))

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &RestController{
		httpServer:     httpServer,
		topicService:   topicService,
		discordService: discordService,
	}
}

func (controller *RestController) Run() error {
	return controller.httpServer.ListenAndServe()
}
