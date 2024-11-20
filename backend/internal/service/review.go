package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/jcocozza/poop.map/backend/internal/model"
	"github.com/jcocozza/poop.map/backend/internal/repository"
	"github.com/jcocozza/poop.map/backend/internal/utils"
)

type ReviewService struct {
	reviewRepo repository.ReviewRepository
	logger *slog.Logger
}

func NewReviewService(reviewRepo repository.ReviewRepository, logger *slog.Logger) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
		logger: logger.WithGroup("review service"),
	}
}

func (rs *ReviewService) CreateReview(ctx context.Context, review model.Review) error {
	rs.logger.InfoContext(ctx,  "creating review")
	uuid := utils.GenerateUUID()
	review.UUID = uuid
	review.Time = time.Now().UTC()
	return rs.reviewRepo.Create(ctx, review)
}

func (rs *ReviewService) ReadAllByPoopLocation(ctx context.Context, poopLocationUUID string) ([]model.Review, error) {
	rs.logger.InfoContext(ctx, "reading by poop location", slog.String("uuid", poopLocationUUID))
	return rs.reviewRepo.ReadByPoopLocation(ctx, poopLocationUUID)
}

func (rs *ReviewService) Upvote(ctx context.Context, UUID string) error {
	rs.logger.InfoContext(ctx, "upvoting", slog.String("uuid", UUID))
	return rs.reviewRepo.Upvote(ctx, UUID)
}

func (rs *ReviewService) Downvote(ctx context.Context, UUID string) error {
	rs.logger.InfoContext(ctx, "downvoting", slog.String("uuid", UUID))
	return rs.reviewRepo.Downvote(ctx, UUID)
}
