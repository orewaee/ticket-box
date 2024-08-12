package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orewaee/ticket-box/internal/app/domain"
	"github.com/orewaee/ticket-box/internal/config"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ApiBridge struct {
	httpClient *http.Client
}

func NewApiBridge() *ApiBridge {
	return &ApiBridge{
		httpClient: &http.Client{},
	}
}

func (bridge *ApiBridge) IsAdmin(permissions int) bool {
	return (permissions & (1 << 3)) == 8
}

func (bridge *ApiBridge) CurrentUserIsGuildAdmin(accessToken, guildId string) bool {
	userGuilds, err := bridge.GetCurrentUserGuilds(accessToken)
	if err != nil {
		return false
	}

	isAdmin := false
	for _, guild := range userGuilds {
		if guild.Id != guildId {
			continue
		}

		isAdmin = bridge.IsAdmin(guild.Permissions)
	}

	return isAdmin
}

func (bridge *ApiBridge) GetCurrentUser(accessToken string) (*domain.User, error) {
	request, err := http.NewRequest(http.MethodGet, "https://discord.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := bridge.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("unauthorized")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error occurred")
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data = new(user)
	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	avatarUrl := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", data.Id, data.Avatar)

	return &domain.User{
		Id:        data.Id,
		Username:  data.Username,
		AvatarUrl: avatarUrl,
	}, nil
}

func (bridge *ApiBridge) GetCurrentUserGuilds(accessToken string) ([]*domain.UserGuild, error) {
	request, err := http.NewRequest(http.MethodGet, "https://discord.com/api/users/@me/guilds", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := bridge.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("unauthorized")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error occurred")
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data = make([]*userGuild, 0)
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	guilds := make([]*domain.UserGuild, len(data))
	for i, guild := range data {
		permissions, err := strconv.Atoi(guild.PermissionsNew)
		if err != nil {
			return nil, err
		}

		iconUrl := fmt.Sprintf("https://cdn.discordapp.com/icons/%s/%s.png", guild.Id, guild.Icon)

		guilds[i] = &domain.UserGuild{
			Id:          guild.Id,
			Name:        guild.Name,
			IconUrl:     iconUrl,
			Owner:       guild.Owner,
			Permissions: permissions,
		}
	}

	return guilds, nil
}

func (bridge *ApiBridge) GetCurrentAuthInfo(accessToken string) (*domain.AuthInfo, error) {
	request, err := http.NewRequest(http.MethodGet, "https://discord.com/api/oauth2/@me", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := bridge.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("unauthorized")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("error occurred")
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data := new(authInfo)
	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	avatarUrl := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", data.User.Id, data.User.Avatar)

	return &domain.AuthInfo{
		Scopes: data.Scopes,
		User: &domain.User{
			Id:        data.User.Id,
			Username:  data.User.Username,
			AvatarUrl: avatarUrl,
		},
	}, nil
}

func (bridge *ApiBridge) GetTokenByCode(code string) (string, error) {
	cfg := config.Get()

	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("redirect_uri", cfg.RedirectUri)
	values.Set("client_id", cfg.ClientId)
	values.Set("client_secret", cfg.ClientSecret)

	body := strings.NewReader(values.Encode())
	request, err := http.NewRequest(http.MethodPost, "https://discord.com/api/oauth2/token", body)
	if err != nil {
		return "", nil
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := bridge.httpClient.Do(request)
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	data := new(tokenResponse)
	if err := json.Unmarshal(bytes, data); err != nil {
		return "", err
	}

	if data.Error != "" {
		return "", errors.New(data.Error)
	}

	return data.AccessToken, nil
}
