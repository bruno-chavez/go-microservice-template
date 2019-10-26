package database

import (
	"database/sql"
	"github.com/lopezator/migrator"
)

// runs all migrations in the order they are added
func runMigrations(db *sql.DB) error {

	// To add a new migration simply add a new migrator.Migration struct with the format
	// Name, Func, where func is a sql transaction as a string
	migrations := migrator.Migrations(
		&migrator.Migration{
			Name: "Create users table",
			Func: func(tx *sql.Tx) error {
				_, err := tx.Exec(
					`create table "users"
							(
								user_id  serial not null constraint user_pk primary key,
								password varchar(250),
								username varchar(250),
								email    varchar(250)
							);
						
							create unique index "User_user-id_uindex" on "users" (user_id);
						
							create unique index user_username_uindex on "users" (username);
						
							create unique index user_email_uindex on "users" (email);`)
				if err != nil {
					return err
				}
				return nil
			},
		},
	)

	m, err := migrator.New(migrations)
	if err != nil {
		return err
	}

	err = m.Migrate(db)
	if err != nil {
		return err
	}
	return nil
}
