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

type ReviewHandler struct {
	reviewService *service.ReviewService
	logger *slog.Logger
}

func NewReviewHandler(reviewService *service.ReviewService, logger *slog.Logger) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
		logger: logger.WithGroup("review handler"),
	}
}

func (rh *ReviewHandler) getAll(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	if len(uuid) == 0 {
		responder.RespondError(w, http.StatusBadRequest, "poop location uuid required", nil)
		return
	}
	reviews, err := rh.reviewService.ReadAllByPoopLocation(context.TODO(), uuid)
	if err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "error reading reviews by poop location", []responder.Error{e})
		return
	}
	responder.RespondSuccess(w, reviews)
}

func (rh *ReviewHandler) put(w http.ResponseWriter, r *http.Request) {
	var review model.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusBadRequest, "invalid request payload", []responder.Error{e})
		return
	}
	err := rh.reviewService.CreateReview(context.TODO(), review)
	if err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "unable to create review", []responder.Error{e})
		return
	}
	responder.RespondSuccessWithStatus(w, http.StatusCreated, nil)
}

func (rh *ReviewHandler) ReviewHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rh.getAll(w,r)
		return
	case http.MethodPut:
		rh.put(w, r)
	default:
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
}

func (rh *ReviewHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	uuid := r.PathValue("uuid")
	if len(uuid) == 0 {
		responder.RespondError(w, http.StatusBadRequest, "uuid required for upvote", nil)
		return
	}
	err := rh.reviewService.Upvote(context.TODO(), uuid)
	if err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "unable to upvote", []responder.Error{e})
		return
	}
	responder.RespondSuccess(w, nil)
}

func (rh *ReviewHandler) Downvote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		responder.RespondError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	uuid := r.PathValue("uuid")
	if len(uuid) == 0 {
		responder.RespondError(w, http.StatusBadRequest, "uuid required for downvote", nil)
		return
	}
	err := rh.reviewService.Downvote(context.TODO(), uuid)
	if err != nil {
		e := responder.Error{Field: "err", Message: err.Error()}
		responder.RespondError(w, http.StatusInternalServerError, "unable to downvote", []responder.Error{e})
		return
	}
	responder.RespondSuccess(w, nil)
}
