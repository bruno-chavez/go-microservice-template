package database

import (
	"database/sql"
	"github.com/lopezator/migrator"
)

// runs all migrations in order
func runMigrations(db *sql.DB) error {

	// to add a new migration simply add a new struct with the format
	// Name, Func, where func is a sql transaction as a string
	m, err := migrator.New(migrator.Migrations(&migrator.Migration{
		Name: "Create users table",
		Func: func(tx *sql.Tx) error {
			_, err := tx.Exec(
				`create table "user"
						(
							user_id  serial not null
								constraint user_pk
									primary key,
							password varchar(250),
							username varchar(250),
							email    varchar(250)
						);
						
						create unique index "User_user-id_uindex"
							on "user" (user_id);
						
						create unique index user_username_uindex
							on "user" (username);
						
						create unique index user_email_uindex
							on "user" (email);`)
			if err != nil {
				return err
			}
			return nil
		},
	}))
	if err != nil {
		return err
	}

	err = m.Migrate(db)
	if err != nil {
		return err
	}
	return nil
}
