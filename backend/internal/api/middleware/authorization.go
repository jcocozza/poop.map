package middleware

import (
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api/responder"
)

const authorization = "Authorization"

func AuthorizationMiddleware(next http.Handler, logger *slog.Logger, apiKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authKey := r.Header.Get(authorization)
		if authKey != apiKey {
			responder.RespondError(w, http.StatusForbidden, "invalid api credentials", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
