package middlewares

import (
	"net/http"
	"onlineshop/pkg/config"
)

type AuthMiddleware struct {
	App *config.AppConfig
}

func NewAuthMiddleware(app *config.AppConfig) *AuthMiddleware {
	return &AuthMiddleware{App: app}
}

func (auth *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := auth.App.Session.GetString(r.Context(), "user")
		if username == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
