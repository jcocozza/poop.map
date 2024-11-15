package repository

import (
	"context"

	"github.com/jcocozza/poop.map/backend/internal/model"
)

type PoopLocationRepository interface {
	Create(ctx context.Context, pl model.PoopLocation) error
	Read(ctx context.Context, UUID string) (model.PoopLocation, error)
	ReadAll(ctx context.Context) ([]model.PoopLocation, error)
	Update(ctx context.Context, pl model.PoopLocation) error
	Upvote(ctx context.Context, UUID string) error
	Downvote(ctx context.Context, UUID string) error
	Delete(ctx context.Context, UUID string) error
}

type ReviewRepository interface {
	Create(ctx context.Context, review model.Review) error
	Read(ctx context.Context, UUID string) (model.Review, error)
	ReadByPoopLocation(ctx context.Context, poopLocationUUID string) ([]model.Review, error)
	Update(ctx context.Context, review model.Review) error
	Upvote(ctx context.Context, UUID string) error
	Downvote(ctx context.Context, UUID string) error
	Delete(ctx context.Context, UUID string) error
}
