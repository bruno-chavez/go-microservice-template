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

// PostLogin handles POST request to /login
func (c *Controller) PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// a requestUser struct is used for simplicity even though the username field won't be filled
	var usr requestUser
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println(err)
	}

	// the query expects a single requestUser and maps it into a dbUser
	var queryResult dbUser
	query := `SELECT * FROM "requestUser" WHERE email = $1;`

	err = c.Db.Get(&queryResult, query, usr.Email)
	if err != nil {
		err = writeJSON(w, "wrong email or password", http.StatusForbidden)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// compares the request password with the one stored in the db
	err = bcrypt.CompareHashAndPassword([]byte(queryResult.Password), usr.Password)
	if err != nil {
		err = writeJSON(w, "wrong email or password", http.StatusForbidden)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// retrieves the session if it exists or creates a new one if it doesn't
	session, err := c.SessionStore.Get(r, "requestUser")
	if err != nil {
		log.Println(err)
	}

	// sets a type and id to the session
	session.Values["type"] = "requestUser"
	session.Values["id"] = queryResult.Id

	// saves session data
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}

	// writes response message
	err = writeJSON(w, "authenticated", http.StatusCreated)
	if err != nil {
		log.Println(err)
	}
}
