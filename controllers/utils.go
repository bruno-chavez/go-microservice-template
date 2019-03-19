package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func decodeBody(r *http.Request) *body {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := new(body)
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
