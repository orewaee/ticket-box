package discord

type user struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type userGuild struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	Owner          bool   `json:"owner"`
	PermissionsNew string `json:"permissions_new"`
}
