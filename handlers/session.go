package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Returns current session type or "none" if it doesnt exists
func (c *Controller) GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := c.SessionStore.Get(r, "user")
	if err != nil {
		log.Println(err)
	}

	sessionType := session.Values["type"]

	if sessionType != nil {
		// Type assertion needed to escape from type interface{}
		err = writeResponse(w, sessionType.(string), http.StatusOK)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = writeResponse(w, "none", http.StatusOK)
		if err != nil {
			log.Println(err)
		}
	}
}
