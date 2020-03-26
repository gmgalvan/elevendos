package database

import (
	"context"
	// Driver for postgres
	"database/sql"

	_ "github.com/lib/pq"
)

// NewClientDB create a new connection to the postgres server
func NewClientDB(ctx context.Context, dbDriver, dns string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
