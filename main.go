// Package main takes care of creating the db connection pool, opening the Redis session store,
// load environment variables, register routes and handlers, enabling CORS requests and initiating the server
package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"go-web-template/controllers"
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// Makes error line appear when an error happens
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// loads env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storeSize, err := strconv.Atoi(os.Getenv("REDIS_STORE_SIZE"))

	// connects to redis session store
	store, err := redistore.NewRediStore(storeSize,
		os.Getenv("REDIS_STORE_NETWORK"),
		os.Getenv("REDIS_STORE_ADDRESS"),
		os.Getenv("REDIS_STORE_PASSWORD"),
		[]byte(os.Getenv("REDIS_SESSION_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	// connects to db
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Fatal(err)
	}

	// closes db connection
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// wraps the session store and db pool in a struct to be passed to the handlers
	controller := &controllers.Controller{
		Db:           db,
		SessionStore: store,
	}

	// initiates router
	router := httprouter.New()

	// lists routes with the controller methods
	router.POST("/register", controller.PostRegister)
	router.POST("/login", controller.PostLogin)

	// binds cors options to the router
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONT-END-ADDRESS")},
		AllowedMethods:   []string{"OPTIONS", "POST", "GET"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// starts the server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
