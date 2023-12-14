package oauth2

import (
	"log"
	"net/http"
	"slices"

	"github.com/gorilla/mux"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

type Config struct {
	ID     string
	Secret string
	Domain string
}

// Setup the OAuth2 server in the same server that is serving the API
// In a real world scenario, this would be a separate server
// so that /authorize and /token are in the auth server
// and validation is in the API/resource server
// Please note that /credentials is a stub for generating client credentials
// and should not be used in production
func Setup(router *mux.Router, cfg *Config) {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	err := clientStore.Set(cfg.ID, &models.Client{
		ID:     cfg.ID,
		Secret: cfg.Secret,
		Domain: cfg.Domain,
	})
	if err != nil {
		log.Fatal("Error setting client store:", err)
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	router.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	router.Use(authMiddleware(srv))
}

var skipRoutes []string = []string{"/token", "/credentials", "/authorize"}

func authMiddleware(srv *server.Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if slices.Contains(skipRoutes, r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := srv.ValidationBearerToken(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
