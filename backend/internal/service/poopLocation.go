package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/jcocozza/poop.map/backend/internal/model"
	"github.com/jcocozza/poop.map/backend/internal/repository"
	"github.com/jcocozza/poop.map/backend/internal/utils"
)

type PoopLocationService struct {
	poopLocationRepo repository.PoopLocationRepository
	logger *slog.Logger
}

func NewPoopLocationService(plRepo repository.PoopLocationRepository, logger *slog.Logger) *PoopLocationService {
	return &PoopLocationService{
		poopLocationRepo: plRepo,
		logger: logger.WithGroup("poop location service"),
	}
}

func (pls *PoopLocationService) CreatePoopLocation(ctx context.Context, pl model.PoopLocation) error {
	pls.logger.InfoContext(ctx, "creating poop location")
	uuid := utils.GenerateUUID()
	pl.UUID = uuid
	now := time.Now().UTC()
	pl.FirstCreated = now
	pl.LastModified = now
	return pls.poopLocationRepo.Create(ctx, pl)
}

func (pls *PoopLocationService) ReadAllPoopLocations(ctx context.Context) ([]model.PoopLocation, error) {
	pls.logger.InfoContext(ctx, "reading all poop locations")
	return pls.poopLocationRepo.ReadAll(ctx)
}

func (pls *PoopLocationService) Upvote(ctx context.Context, UUID string) error {
	pls.logger.InfoContext(ctx, "upvoting", slog.String("uuid", UUID))
	return pls.poopLocationRepo.Upvote(ctx, UUID)
}

func (pls *PoopLocationService) Downvote(ctx context.Context, UUID string) error {
	pls.logger.InfoContext(ctx, "downvoting", slog.String("uuid", UUID))
	return pls.poopLocationRepo.Downvote(ctx, UUID)
}
