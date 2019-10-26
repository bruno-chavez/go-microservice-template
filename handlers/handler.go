// Package handlers contains the handlers for each endpoint of the app
package handlers

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/boj/redistore.v1"
	"log"
)

// Handler is used to expose dependencies to the handlers
type Handler struct {
	Db           *sqlx.DB
	SessionStore *redistore.RediStore
	Logger       *log.Logger
}
