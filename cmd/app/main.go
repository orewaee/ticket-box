package main

import (
	"github.com/orewaee/ticket-box/internal/adapters/controllers"
	"github.com/orewaee/ticket-box/internal/adapters/discord"
	"github.com/orewaee/ticket-box/internal/adapters/repos"
	"github.com/orewaee/ticket-box/internal/app/services"
	"github.com/orewaee/ticket-box/internal/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	cfg := config.Get()

	discordBot, err := discord.NewBotBridge(cfg.Token)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := discordBot.Run(); err != nil {
			log.Fatalln(err)
		}
	}()

	discordApi := discord.NewApiBridge()
	discordService := services.NewDiscordService(discordApi, discordBot)

	topicRepo := repos.NewTopicRepo()
	topicService := services.NewTopicService(topicRepo)

	controller := controllers.NewRestController(cfg.Addr, topicService, discordService)

	go func() {
		if err := controller.Run(); err != nil {
			log.Fatalln(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
