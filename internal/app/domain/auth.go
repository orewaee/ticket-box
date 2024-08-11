package domain

type AuthInfo struct {
	Scopes []string
	User   *User
}
