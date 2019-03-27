package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

type dbUser struct {
	Id       int    `db:"user-id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func (c *Controller) PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr user
	err = json.Unmarshal(b, &usr)
	if err != nil {
		log.Fatal(err)
	}

	var queryResult dbUser
	query := `SELECT * FROM "user" WHERE email = $1;`

	err = c.Db.Get(&queryResult, query, usr.Email)
	if err != nil {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(queryResult.Password), []byte(usr.Password))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write([]byte("false"))
		if err != nil {
			log.Fatal(err)
		}
	}

	session, err := c.SessionStore.Get(r, "user")
	if err != nil {
		fmt.Println(err, ":(")

	}

	session.Values["type"] = "user"
	session.Values["id"] = queryResult.Id

	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte("true"))
	if err != nil {
		log.Fatal(err)
	}
}
