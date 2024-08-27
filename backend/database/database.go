package database

import (
	"context"
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

// Execute sql without returning rows.
//
// Pass in args for placeholders in query.
func (db *Database) Execute(ctx context.Context, sql string, args ...any) error {
	_, err := db.DB.ExecContext(ctx, sql, args...)
	if err != nil {
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
	result, err := db.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
