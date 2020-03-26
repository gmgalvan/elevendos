package database

import (
	"database/sql"
	// Driver for postgres
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

// MigrationStore struct for a goose migration
type MigrationStore struct {
	db           *sql.DB
	dirMigration string
}

// NewMigration start a new goose implementation
func NewMigration(client *sql.DB, dirMigration string) *MigrationStore {
	return &MigrationStore{
		db:           client,
		dirMigration: dirMigration,
	}
}

// StartMigration start a new migration
func (m MigrationStore) StartMigration() error {
	return m.MigrationUp()
}

// MigrationUp start migration Up
func (m MigrationStore) MigrationUp() error {
	err := goose.Up(m.db, m.dirMigration)
	if err != nil {
		return err
	}
	return goose.Status(m.db, m.dirMigration)
}

// MigrationDown start migration Down
func (m MigrationStore) MigrationDown() error {
	err := goose.Down(m.db, m.dirMigration)
	if err != nil {
		return err
	}
	return goose.Status(m.db, m.dirMigration)
}

// MigrationStatus prints migration status
func (m MigrationStore) MigrationStatus() error {
	return goose.Status(m.db, m.dirMigration)
}

// MigrationRollBack rolls back the most recently applied migration, then runs it again.
func (m MigrationStore) MigrationRollBack() error {
	err := goose.Redo(m.db, m.dirMigration)
	if err != nil {
		return err
	}
	return goose.Status(m.db, m.dirMigration)
}
