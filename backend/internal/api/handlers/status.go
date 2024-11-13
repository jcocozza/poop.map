package handlers

import (
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api/responder"
)

func Status(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{"status": "ok"}
	responder.RespondSuccess(w, payload)
}
