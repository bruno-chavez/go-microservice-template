package database

import (
	"github.com/jmoiron/sqlx"
	"os"
)

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
