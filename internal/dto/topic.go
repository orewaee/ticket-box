package dto

type Topic struct {
	Id          string `json:"id"`
	GuildId     string `json:"guild_id"`
	Emoji       string `json:"emoji"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddTopicRequest struct {
	GuildId     string `json:"guild_id"`
	Emoji       string `json:"emoji"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
