// Package main takes care of creating the db connection pool, opening the Redis session store,
// load environment variables, register routes and handlers, enables CORS and initiates the server
package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"go-web-template/handlers"
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// Makes code line show on errors
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// loads env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	storeSize, err := strconv.Atoi(os.Getenv("REDIS_STORE_SIZE"))
	if err != nil {
		log.Println(err)
	}

	// connects to redis session store
	store, err := redistore.NewRediStore(storeSize,
		"tcp",
		os.Getenv("REDIS_STORE_ADDRESS"),
		os.Getenv("REDIS_STORE_PASSWORD"),
		[]byte(os.Getenv("REDIS_SESSION_KEY")))
	if err != nil {
		log.Println(err)
	}

	// connects to db
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Println(err)
	}

	// closes db connection
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// wraps the session store and db pool in a struct to be passed to the handlers
	controller := &handlers.Controller{
		Db:           db,
		SessionStore: store,
	}

	// initiates router
	router := httprouter.New()

	// lists routes with the controller methods
	router.POST("/register", controller.PostRegister)
	router.POST("/login", controller.PostLogin)
	router.DELETE("/logout", controller.DeleteLogout)

	// binds cors options to the router
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONT-END-ADDRESS")},
		AllowedMethods:   []string{"OPTIONS", "POST", "GET", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// starts the server
	log.Println(http.ListenAndServe(":8080", handler))
}
