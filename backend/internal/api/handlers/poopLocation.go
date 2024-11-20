package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api/responder"
	"github.com/jcocozza/poop.map/backend/internal/model"
	"github.com/jcocozza/poop.map/backend/internal/service"
)

type PoopLocationUUIDHandler struct {
	plService *service.PoopLocationService
	logger    *slog.Logger
}

func NewPoopLocationUUIDHandler(plService *service.PoopLocationService, logger *slog.Logger) *PoopLocationUUIDHandler {
	return &PoopLocationUUIDHandler{
		plService: plService,
		logger:    logger.WithGroup("poop location uuid handler"),
	}
}

func (plh *PoopLocationUUIDHandler) get(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("uuid")
	responder.RespondError(w, http.StatusNotImplemented, "get by uuid is not implemented", nil)
}

func (plh *PoopLocationUUIDHandler) PoopLocationUUIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		plh.get(w, r)
		return
	default:
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
}

func (plh *PoopLocationUUIDHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	uuid := r.PathValue("uuid")
	if len(uuid) == 0 {
		responder.RespondError(w, http.StatusBadRequest, "uuid required for upvote", nil)
		return
	}
	err := plh.plService.Upvote(context.TODO(), uuid)
	if err != nil {
		responder.RespondError(w, http.StatusInternalServerError, "unable to upvote", nil)
		return
	}
	responder.RespondSuccess(w, nil)
}

func (plh *PoopLocationUUIDHandler) Downvote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	uuid := r.PathValue("uuid")
	if len(uuid) == 0 {
		responder.RespondError(w, http.StatusBadRequest, "uuid required for downvote", nil)
		return
	}
	err := plh.plService.Downvote(context.TODO(), uuid)
	if err != nil {
		responder.RespondError(w, http.StatusInternalServerError, "unable to downvote", nil)
		return
	}
	responder.RespondSuccess(w, nil)
}

type PoopLocationHandler struct {
	plService *service.PoopLocationService
	logger    *slog.Logger
}

func NewPoopLocationHandler(plService *service.PoopLocationService, logger *slog.Logger) *PoopLocationHandler {
	return &PoopLocationHandler{
		plService: plService,
		logger:    logger.WithGroup("poop location handler"),
	}
}

func (plh *PoopLocationHandler) getall(w http.ResponseWriter, r *http.Request) {
	poopLocations, err := plh.plService.ReadAllPoopLocations(context.TODO())
	if err != nil {
		err := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "unable to get all poop locations", []responder.Error{err})
		return
	}
	responder.RespondSuccess(w, poopLocations)
}

func (plh *PoopLocationHandler) put(w http.ResponseWriter, r *http.Request) {
	var pl model.PoopLocation
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusBadRequest, "invalid request payload", []responder.Error{e})
		return
	}
	err := plh.plService.CreatePoopLocation(context.TODO(), pl)
	if err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "unable to create poop location", []responder.Error{e})
		return
	}
	responder.RespondSuccessWithStatus(w, http.StatusCreated, nil)
}

func (plh *PoopLocationHandler) PoopLocationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		plh.getall(w, r)
		return
	case http.MethodPut:
		plh.put(w, r)
		return
	default:
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
	}
}
