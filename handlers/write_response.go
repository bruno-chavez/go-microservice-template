package handlers

import (
	"fmt"
	"net/http"
)

// Helper function for easily writing response messages
func writeResponse(w http.ResponseWriter, status int, key, value string) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	info := []byte(fmt.Sprintf(`{"%v": "%v"}`, key, value))

	_, err := w.Write(info)
	if err != nil {
		return err
	}

	return nil
}
