package discord

import "github.com/bwmarrin/discordgo"

type BotBridge struct {
	session *discordgo.Session
}

func NewBotBridge(token string) (*BotBridge, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents |= discordgo.IntentGuilds
	session.Identify.Intents |= discordgo.IntentGuildMembers

	return &BotBridge{
		session: session,
	}, nil
}

func (bridge *BotBridge) Run() error {
	if err := bridge.session.Open(); err != nil {
		return err
	}

	return nil
}

func (bridge *BotBridge) GuildExists(guildId string) bool {
	_, err := bridge.session.Guild(guildId)
	if err != nil {
		return false
	}

	return true
}

func (bridge *BotBridge) BotExists(guildId string) bool {
	_, err := bridge.session.Guild(guildId)
	if err != nil {
		return false
	}

	return true
}
