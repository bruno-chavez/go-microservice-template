package database

import (
	"github.com/jmoiron/sqlx"
	"os"
)

func NewDB() (*sqlx.DB, error) {
	// connects to the db
	db, err := sqlx.Connect("postgres", os.Getenv("DB"))
	if err != nil {
		return nil, err
	}

	// runs migrations
	err = runMigrations(db.DB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
