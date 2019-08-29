package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *Controller) GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := c.SessionStore.Get(r, "requestUser")
	if err != nil {
		log.Println(err)
	}

	sessionType := session.Values["type"]

	if sessionType != nil {
		// Type assertion needed to escape from type interface{}
		err = writeJSON(w, sessionType.(string), http.StatusOK)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = writeJSON(w, "none", http.StatusOK)
		if err != nil {
			log.Println(err)
		}
	}
}
