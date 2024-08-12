package dto

type Guild struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
	WithBot bool   `json:"with_bot"`
}
