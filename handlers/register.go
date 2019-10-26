package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
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

	var user requestUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.Logger.Println(err)
	}

	// hashes and salts the password
	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Println(err)
	}

	query := `INSERT INTO "users" (email, username, password) VALUES ($1, $2, $3);`
	_, err = h.Db.Exec(query, user.Email, user.Username, hash)
	if err != nil {
		h.Logger.Println(err)
		err = writeResponse(w, "username or email not unique", http.StatusBadRequest)
		if err != nil {
			h.Logger.Println(err)
		}
		return
	}

	err = writeResponse(w, "user created", http.StatusCreated)
	if err != nil {
		h.Logger.Println(err)
	}
}
