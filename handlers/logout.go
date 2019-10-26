package handlers

import (
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Deletes current user session
func (h *Handler) deleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := h.SessionStore.Get(r, "user")
	if err != nil {
		err = writeResponse(w, "no valid session was found", http.StatusNotFound)
		if err != nil {
			h.Logger.Println(err)
		}
		return
	}

	// Deletes session on the Redis store.
	session.Options.MaxAge = -1
	err = sessions.Save(r, w)
	if err != nil {
		h.Logger.Println(err)
		return
	}

	err = writeResponse(w, "logged out", http.StatusOK)
	if err != nil {
		h.Logger.Println(err)
	}
}
