package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"gopkg.in/boj/redistore.v1"
	"log"
	"miller/controllers"
	"net/http"
	"os"
)

func main() {


	// loads env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connects to redis session store
	store, err := redistore.NewRediStore(10,
		"tcp",
		":6379",
		"",
		[]byte(os.Getenv("SESSION_KEY")))
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
		Db:db,
		SessionStore:store,
	}

	// initiates router
	router := httprouter.New()

	// lists routes with the controller methods
	/*router.POST("/register", controller.PostRegister)
	router.POST("/login", controller.PostLogin)*/

	// binds cors options to the router
	c := cors.New(cors.Options{
		AllowedOrigins:		[]string{"http://127.0.0.1:8081"},
		AllowedMethods:     []string{"OPTIONS", "POST", "GET"},
		AllowedHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
