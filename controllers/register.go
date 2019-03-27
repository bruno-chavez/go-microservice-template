package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Username string          `json:"user"`
	Password json.RawMessage `json:"password"`
	Email    string          `json:"email"`
}

func (c *Controller) PostRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// decodes the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr user
	err = json.Unmarshal(body, &usr)
	if err != nil {
		log.Fatal(err)
	}

	// hashes the password
	hash, err := bcrypt.GenerateFromPassword(usr.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// inserts the new user
	query := `INSERT INTO "user" (email, username, password) VALUES ($1, $2, $3);`
	_, err = c.Db.Exec(query, usr.Email, usr.Username, hash)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
