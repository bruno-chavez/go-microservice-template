package controllers

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/boj/redistore.v1"
)

type Controller struct {
	Db *sqlx.DB
	SessionStore *redistore.RediStore
}
