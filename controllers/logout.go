package controllers

import (
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (c *Controller) PostLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := c.SessionStore.Get(r, "user")
	if err != nil {
		log.Fatal(err)
	}

	// Delete session.
	session.Options.MaxAge = -1
	err = sessions.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	// writes a response if all went ok
	_, err = w.Write([]byte("session deleted successfully"))
	if err != nil {
		log.Fatal(err)
	}
}
