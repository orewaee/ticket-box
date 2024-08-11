package ports

import (
	"github.com/orewaee/ticket-box/internal/app/domain"
)

type DiscordService interface {
	IsAdmin(permissions int) bool
	CurrentUserIsGuildAdmin(accessToken, guildId string) bool
	GetCurrentUser(accessToken string) (*domain.User, error)
	GetCurrentUserGuilds(accessToken string) ([]*domain.UserGuild, error)
	BotExists(guildId string) bool
	GetCurrentAuthInfo(accessToken string) (*domain.AuthInfo, error)
}

type DiscordApiBridge interface {
	IsAdmin(permissions int) bool
	GetCurrentUser(accessToken string) (*domain.User, error)
	GetCurrentUserGuilds(accessToken string) ([]*domain.UserGuild, error)
	CurrentUserIsGuildAdmin(accessToken, guildId string) bool
	GetCurrentAuthInfo(accessToken string) (*domain.AuthInfo, error)
}

type DiscordBotBridge interface {
	BotExists(guildId string) bool
}
