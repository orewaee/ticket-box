package services

import (
	"github.com/orewaee/ticket-box/internal/app/domain"
	"github.com/orewaee/ticket-box/internal/app/ports"
)

type DiscordService struct {
	discordApiBridge ports.DiscordApiBridge
	discordBotBridge ports.DiscordBotBridge
}

func NewDiscordService(discordApiBridge ports.DiscordApiBridge, discordBotBridge ports.DiscordBotBridge) *DiscordService {
	return &DiscordService{
		discordApiBridge: discordApiBridge,
		discordBotBridge: discordBotBridge,
	}
}

func (service *DiscordService) IsAdmin(permissions int) bool {
	return service.discordApiBridge.IsAdmin(permissions)
}

func (service *DiscordService) CurrentUserIsGuildAdmin(accessToken, guildId string) bool {
	return service.discordApiBridge.CurrentUserIsGuildAdmin(accessToken, guildId)
}

func (service *DiscordService) GetCurrentUser(accessToken string) (*domain.User, error) {
	return service.discordApiBridge.GetCurrentUser(accessToken)
}

func (service *DiscordService) GetCurrentUserGuilds(accessToken string) ([]*domain.UserGuild, error) {
	return service.discordApiBridge.GetCurrentUserGuilds(accessToken)
}

func (service *DiscordService) BotExists(guildId string) bool {
	return service.discordBotBridge.BotExists(guildId)
}

func (service *DiscordService) GetCurrentAuthInfo(accessToken string) (*domain.AuthInfo, error) {
	return service.discordApiBridge.GetCurrentAuthInfo(accessToken)
}

func (service *DiscordService) GetTokenByCode(code string) (string, error) {
	return service.discordApiBridge.GetTokenByCode(code)
}
