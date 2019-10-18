package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// requestUser is used to map the request body to a struct
type requestUser struct {
	Username string          `json:"username"`
	Password json.RawMessage `json:"password"`
	Email    string          `json:"email"`
}

// Creates a new user
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var requestUsr requestUser
	err := json.NewDecoder(r.Body).Decode(&requestUsr)
	if err != nil {
		log.Println(err)
	}

	// hashes the password
	hash, err := bcrypt.GenerateFromPassword(requestUsr.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	// inserts the new user in the DB
	query := `INSERT INTO "requestUser" (email, username, password) VALUES ($1, $2, $3);`
	_, err = h.Db.Exec(query, requestUsr.Email, requestUsr.Username, hash)
	if err != nil {
		log.Println(err)
		err = writeResponse(w, "username or email not unique", http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = writeResponse(w, "user created", http.StatusCreated)
	if err != nil {
		log.Println(err)
	}
}
