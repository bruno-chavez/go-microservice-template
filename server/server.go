package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go-microservice-template/handlers"
	"net/http"
	"os"
	"time"
)

func NewServer(h *handlers.Handler) *http.Server {

	router := httprouter.New()

	// health check
	router.GET("/info", h.Info())

	// enables CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
		AllowCredentials: true,
	})
	handlerWithCORS := corsConfig.Handler(router)

	srvr := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      handlerWithCORS,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return srvr
}
