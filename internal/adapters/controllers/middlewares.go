package controllers

import (
	"context"
	"fmt"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"github.com/orewaee/ticket-box/internal/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, discordService ports.DiscordService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		if authorization == "" {
			writer.WriteHeader(http.StatusUnauthorized)
			utils.WriteString(writer, "missing authorization header")
			return
		}

		accessToken := strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer "))
		if accessToken == "" {
			writer.WriteHeader(http.StatusUnauthorized)
			utils.WriteString(writer, "missing token")
			return
		}

		authInfo, err := discordService.GetCurrentAuthInfo(accessToken)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			utils.WriteString(writer, err.Error())
			fmt.Println(err)
			return
		}

		ctx := context.WithValue(
			context.WithValue(request.Context(), "id", authInfo.User.Id),
			"accessToken", accessToken,
		)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
