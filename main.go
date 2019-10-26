// Package main takes care of gluing all the elements of the app
package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-web-template/database"
	"go-web-template/handlers"
	"go-web-template/session"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	store, err := session.NewSession()
	if err != nil {
		log.Println(err)
	}

	defer func() {
		err := store.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	db, err := database.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// wraps data to a struct to be passed to the handlers
	controller := &handlers.Handler{
		Db:           db,
		SessionStore: store,
		Logger:       logger,
	}

	handler, err := controller.NewServer()
	if err != nil {
		log.Println(err)
	}

	log.Println(http.ListenAndServe(":8080", handler))
}
