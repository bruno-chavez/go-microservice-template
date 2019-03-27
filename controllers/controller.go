package controllers

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/boj/redistore.v1"
	"log"
)

type Controller struct {
	Db           *sqlx.DB
	SessionStore *redistore.RediStore
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
