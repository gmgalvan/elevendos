package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	client *sql.DB
}

func NewStore(client *sql.DB) *Store {
	return &Store{
		client: client,
	}
}
