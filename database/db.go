// Package database takes care of setting up the sql connection, running migrations, etc.
package database

import (
	"github.com/jmoiron/sqlx"
	"os"
)

// NewDB creates a new sql connection and returns it
func NewDB() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", os.Getenv("DB"))
	if err != nil {
		return nil, err
	}

	err = runMigrations(db.DB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
