package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Returns current session type or "none" if it doesnt exists
func (h *Handler) session(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := h.SessionStore.Get(r, "user")
	if err != nil {
		log.Println(err)
		err = writeResponse(w, "no valid session was found", http.StatusNotFound)
		if err != nil {
			log.Println(err)
		}
		return
	}

	sessionType := session.Values["type"]
	if sessionType == nil {
		err = writeResponse(w, "no valid session was found", http.StatusNotFound)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Type assertion needed to escape from type interface{}
	err = writeResponse(w, sessionType.(string), http.StatusOK)
	if err != nil {
		log.Println(err)
	}
}
