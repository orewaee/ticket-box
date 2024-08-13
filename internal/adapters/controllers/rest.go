package controllers

import (
	"github.com/orewaee/ticket-box/internal/adapters/cors"
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

	mux.Handle("GET /user", cors.Middleware(AuthMiddleware(
		handlers.NewGetUserHandler(discordService), discordService)))
	mux.Handle("OPTIONS /user", cors.NewOptionsHandler())

	mux.Handle("GET /guild/{guild_id}", cors.Middleware(AuthMiddleware(
		handlers.NewGetGuildHandler(discordService), discordService)))
	mux.Handle("OPTIONS /guild/{guild_id}", cors.NewOptionsHandler())

	mux.Handle("GET /guilds", cors.Middleware(AuthMiddleware(
		handlers.NewGetGuildsHandler(discordService), discordService)))
	mux.Handle("OPTIONS /guilds", cors.NewOptionsHandler())

	mux.Handle("GET /topic/{topic_id}", cors.Middleware(AuthMiddleware(
		handlers.NewGetTopicHandler(topicService, discordService), discordService)))
	mux.Handle("OPTIONS /topic/{topic_id}", cors.NewOptionsHandler())

	mux.Handle("GET /topics/{guild_id}", cors.Middleware(AuthMiddleware(
		handlers.NewGetTopicsHandler(topicService, discordService), discordService)))
	mux.Handle("OPTIONS /topics/{guild_id}", cors.NewOptionsHandler())

	mux.Handle("POST /topic", cors.Middleware(AuthMiddleware(
		handlers.NewPostTopicHandler(topicService, discordService), discordService)))
	mux.Handle("OPTIONS /topic", cors.NewOptionsHandler())

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
