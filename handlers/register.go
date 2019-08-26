package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// user is used to map the request body to a struct
type user struct {
	Username string          `json:"username"`
	Password json.RawMessage `json:"password"`
	Email    string          `json:"email"`
}

// PostRegister handles POST request to /register
func (c *Controller) PostRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var usr user
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println(err)
	}

	// hashes the password
	hash, err := bcrypt.GenerateFromPassword(usr.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	// inserts the new user into the db
	query := `INSERT INTO "user" (email, username, password) VALUES ($1, $2, $3);`
	_, err = c.Db.Exec(query, usr.Email, usr.Username, hash)
	if err != nil {
		log.Println(err)
		err = writeJSON(w, "username or email not unique", http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = writeJSON(w, "user created", http.StatusCreated)
	if err != nil {
		log.Println(err)
	}
}
