package controllers

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type dbResult struct {
	Id int `db:"user-id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email string `db:"email"`
}

func (c *Controller) PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	query := `SELECT * FROM "user" WHERE email = $1;`

	var result dbResult

	body := decodeBody(r)

	err := c.Db.Get(&result, query, body.Email)
	if err != nil {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusMovedPermanently)
}
