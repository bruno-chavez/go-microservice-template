package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
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

	// decodes the request body into a user struct
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usr user
	err = json.Unmarshal(b, &usr)
	if err != nil {
		log.Fatal(err)
	}

	// the query expects a single user and maps it into a dbUser
	var queryResult dbUser
	query := `SELECT * FROM "user" WHERE email = $1;`

	err = c.Db.Get(&queryResult, query, usr.Email)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write([]byte("false"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// compares the request password with the one stored in the db
	err = bcrypt.CompareHashAndPassword([]byte(queryResult.Password), usr.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write([]byte("false"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// retrieves the session if it exists or creates a new one if there isn't one already
	session, err := c.SessionStore.Get(r, "user")
	if err != nil {
		log.Fatal(err)
	}

	// sets a type and id to the session
	session.Values["type"] = "user"
	session.Values["id"] = queryResult.Id

	// saves the session data
	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	// writes a response if all went ok
	_, err = w.Write([]byte("true"))
	if err != nil {
		log.Fatal(err)
	}
}
