package cors

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next.ServeHTTP(writer, request)
	})
}
