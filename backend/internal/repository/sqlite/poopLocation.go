package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jcocozza/poop.map/backend/internal/model"
)

type SQLitePoopLocationRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewSQLitePoopLocationRepository(db *sql.DB, logger *slog.Logger) *SQLitePoopLocationRepository {
	return &SQLitePoopLocationRepository{
		db:     db,
		logger: logger,
	}
}

func (splr *SQLitePoopLocationRepository) Create(ctx context.Context, pl model.PoopLocation) error {
	sql := `
INSERT INTO poop_location 
(uuid, name, latitude, longitude, first_created, last_modified, seasonal, seasons_mask, accessible, upvotes, downvotes) 
VALUES
(?,?,?,?,?,?,?,?,?,?,?)
`
	_, err := splr.db.ExecContext(ctx, sql,
		pl.UUID,
		pl.Name,
		pl.Latitude,
		pl.Longitude,
		timeToString(pl.FirstCreated),
		timeToString(pl.LastModified),
		pl.Seasonal,
		model.SeasonMask(pl.Seasons),
		pl.Accessible,
		pl.Upvotes,
		pl.DownVotes,
	)
	return err
}

func (splr *SQLitePoopLocationRepository) Read(ctx context.Context, UUID string) (model.PoopLocation, error) {
	return model.PoopLocation{}, errors.New("not implemented")
}

func (splr *SQLitePoopLocationRepository) ReadAll(ctx context.Context) ([]model.PoopLocation, error) {
	sql := `
SELECT
	uuid,
	name,
	latitude,
	longitude,
	first_created,
	last_created,
	seasonal,
	seasons_mask,
	accessible,
	upvotes,
	downvotes
FROM poop_location
`
	rows, err := splr.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	poopLocationList := []model.PoopLocation{}
	for rows.Next() {
		poopLocation := model.PoopLocation{}
		var firstCreatedStr string
		var lastModifiedStr string
		var seasonsMask int
		err := rows.Scan(
			&poopLocation.UUID,
			&poopLocation.Name,
			&poopLocation.Latitude,
			&poopLocation.Longitude,
			&firstCreatedStr,
			&lastModifiedStr,
			&seasonsMask,
			&poopLocation.Accessible,
			&poopLocation.Upvotes,
			&poopLocation.DownVotes,
		)
		if err != nil {
			return nil, err
		}
		firstCreated, err := parseTime(firstCreatedStr)
		if err != nil {
			return nil, err
		}
		poopLocation.FirstCreated = firstCreated
		lastModified, err := parseTime(lastModifiedStr)
		if err != nil {
			return nil, err
		}
		poopLocation.LastModified = lastModified
		seasons := model.GetSeasons(seasonsMask)
		poopLocation.Seasons = seasons
		poopLocationList = append(poopLocationList, poopLocation)
	}
	return poopLocationList, nil
}

func (splr *SQLitePoopLocationRepository) Update(ctx context.Context, pl model.PoopLocation) error {
	return errors.New("not implemented")
}

func (splr *SQLitePoopLocationRepository) Upvote(ctx context.Context, UUID string) error {
	sql := "UPDATE poop_location SET upvotes = upvotes + 1 WHERE uuid = ?;"
	_, err := splr.db.ExecContext(ctx, sql, UUID)
	return err
}

func (splr *SQLitePoopLocationRepository) Downvote(ctx context.Context, UUID string) error {
	sql := "UPDATE poop_location SET downvotes = downvotes + 1 WHERE uuid = ?;"
	_, err := splr.db.ExecContext(ctx, sql, UUID)
	return err
}

func (splr *SQLitePoopLocationRepository) Delete(ctx context.Context, UUID string) error {
	return errors.New("not implemented")
}
