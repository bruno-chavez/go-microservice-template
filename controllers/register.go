package controllers

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type body struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (c *Controller) PostRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	body := decodeBody(r)

	// hashes the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// inserts the new body
	query := `INSERT INTO "user" (email, username, password) VALUES ($1, $2, $3);`
	_, err = c.Db.Exec(query, body.Email, body.Username, hash)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
