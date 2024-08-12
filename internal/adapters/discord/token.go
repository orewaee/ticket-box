package discord

type tokenResponse struct {
	Error       string `json:"error"`
	AccessToken string `json:"access_token"`
}
