// Package handlers contains the controllers for each route
package handlers

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/boj/redistore.v1"
	"log"
)

// Controller is used to expose the db connection pool and the session store to the controllers.
type Controller struct {
	Db           *sqlx.DB
	SessionStore *redistore.RediStore
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
