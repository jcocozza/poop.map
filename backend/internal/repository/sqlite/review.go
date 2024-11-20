package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jcocozza/poop.map/backend/internal/model"
)

type SQLiteReviewRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewSQLiteReviewRepository(db *sql.DB, logger *slog.Logger) *SQLiteReviewRepository {
	return &SQLiteReviewRepository{
		db:     db,
		logger: logger.WithGroup("review repo"),
	}
}

func (srr *SQLiteReviewRepository) Create(ctx context.Context, review model.Review) error {
	sql := `
INSERT INTO review
(uuid, poop_location_uuid, rating, comment, time, upvotes, downvotes)
VALUES
(?,?,?,?,?,?,?);
`
	_, err := srr.db.ExecContext(ctx, sql,
		review.UUID,
		review.PoopLocationUUID,
		review.Rating,
		review.Comment,
		timeToString(review.Time),
		review.Upvotes,
		review.DownVotes)
	return err
}

func (srr *SQLiteReviewRepository) Read(ctx context.Context, UUID string) (model.Review, error) {
	return model.Review{}, errors.New("not implemented")
}

func (srr *SQLiteReviewRepository) ReadByPoopLocation(ctx context.Context, poopLocationUUID string) ([]model.Review, error) {
	sql := `
SELECT
	uuid,
	rating,
	comment,
	time,
	upvotes,
	downvotes
FROM review
WHERE poop_location_uuid = ?;
`
	rows, err := srr.db.QueryContext(ctx, sql, poopLocationUUID)
	if err != nil {
		return nil, err
	}
	reviewList := []model.Review{}
	for rows.Next() {
		review := model.Review{}
		var timeStr string
		err := rows.Scan(
			&review.UUID,
			&review.Rating,
			&review.Comment,
			&timeStr,
			&review.Upvotes,
			&review.DownVotes,
		)
		if err != nil {
			return nil, err
		}
		tm, err := parseTime(timeStr)
		if err != nil {
			return nil, err
		}
		review.Time = tm
		review.PoopLocationUUID = poopLocationUUID
		reviewList = append(reviewList, review)
	}
	return reviewList, nil
}

func (srr *SQLiteReviewRepository) Update(ctx context.Context, review model.Review) error {
	return errors.New("not implemented")
}

func (srr *SQLiteReviewRepository) Upvote(ctx context.Context, UUID string) error {
	sql := "UPDATE review SET upvotes = upvotes + 1 WHERE uuid = ?;"
	_, err := srr.db.ExecContext(ctx, sql, UUID)
	return err
}

func (srr *SQLiteReviewRepository) Downvote(ctx context.Context, UUID string) error {
	sql := "UPDATE review SET downvotes = downvotes + 1 WHERE uuid = ?;"
	_, err := srr.db.ExecContext(ctx, sql, UUID)
	return err
}

func (srr *SQLiteReviewRepository) Delete(ctx context.Context, UUID string) error {
	return errors.New("not implemented")
}
