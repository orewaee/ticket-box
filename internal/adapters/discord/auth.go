package discord

type authInfo struct {
	Scopes []string `json:"scopes"`
	User   struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	}
}
