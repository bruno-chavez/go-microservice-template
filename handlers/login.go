package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// dbUser is used to map the query result to a struct
type dbUser struct {
	Id       int    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func (h *Handler) createSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var requestUsr requestUser
	err := json.NewDecoder(r.Body).Decode(&requestUsr)
	if err != nil {
		log.Println(err)
		err = writeResponse(w, "wrong format", http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	var dbUsr dbUser
	query := `SELECT * FROM "users" WHERE email = $1;`
	err = h.Db.Get(&dbUsr, query, requestUsr.Email)
	if err != nil {
		err = writeResponse(w, "wrong email or password", http.StatusBadRequest)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// compares the request password with the one stored in the db
	err = bcrypt.CompareHashAndPassword([]byte(dbUsr.Password), requestUsr.Password)
	if err != nil {
		err = writeResponse(w, "wrong email or password", http.StatusUnauthorized)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// retrieves the session if it exists or creates a new one if it doesn't
	session, err := h.SessionStore.Get(r, "user")
	if err != nil {
		log.Println(err)
	}

	session.Values["type"] = "user"
	session.Values["id"] = dbUsr.Id

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}

	err = writeResponse(w, "authenticated", http.StatusCreated)
	if err != nil {
		log.Println(err)
	}
}
