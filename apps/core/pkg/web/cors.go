package web

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func CorsMiddleware(next http.Handler) http.Handler {
	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}

	return cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowCredentials: true,
	}).Handler(next)
}
