package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, value string, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	message := response{Message: value}
	marshaledData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = w.Write(marshaledData)
	if err != nil {
		return err
	}

	return nil
}
