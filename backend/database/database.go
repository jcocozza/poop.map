package database

import (
	"context"
	"database/sql"
	"log/slog"
)

type Database struct {
	Logger *slog.Logger
	DB *sql.DB
}

// Execute sql without returning rows.
//
// Pass in args for placeholders in query.
func (db *Database) Execute(ctx context.Context, sql string, args ...any) error {
	db.Logger.DebugContext(ctx, "running sql", slog.String("sql", sql), slog.Group("args", args...))
	_, err := db.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		db.Logger.ErrorContext(ctx, "sql failed to execute", slog.String("error", err.Error()))
		return err
	}
	return nil
}

// Execute sql and return (several) rows.
//
// Pass in args for placeholders in query.
//
// Make sure to use defer rows.Close().
func (db *Database) Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error) {
	db.Logger.DebugContext(ctx, "running sql", slog.String("sql", sql), slog.Group("args", args...))
	result, err := db.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		db.Logger.ErrorContext(ctx, "sql failed to query", slog.String("error", err.Error()))
		return nil, err
	}
	return result, nil
}
