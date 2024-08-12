package domain

type User struct {
	Id        string
	Username  string
	AvatarUrl string
}

type UserGuild struct {
	Id          string
	Name        string
	IconUrl     string
	Owner       bool
	Permissions int
}
