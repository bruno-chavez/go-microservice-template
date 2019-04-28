package handlers

import (
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (c *Controller) DeleteLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := c.SessionStore.Get(r, "user")
	if err != nil {
		log.Fatal(err)
	}

	// Deletes session on the Redis store.
	session.Options.MaxAge = -1
	err = sessions.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	err = writeJSON(w, "logged out", http.StatusOK)
	if err != nil {
		log.Fatal(err)
	}
}
