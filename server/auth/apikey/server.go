package apikey

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Config holds the configuration for the API key authentication middleware
type Config struct {
	HeaderName  string
	ApiKeyValue string
}

func Setup(router *mux.Router, cfg *Config) {
	router.Use(authMiddleware(cfg))
}

func authMiddleware(cfg *Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(cfg.HeaderName)
			if apiKey != cfg.ApiKeyValue {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
