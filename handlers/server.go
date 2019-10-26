package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"os"
)

// NewServer sets a http router and returns a concrete type to be used as a server handler in package main
func (h Handler) NewServer() (http.Handler, error) {

	router := httprouter.New()

	router.POST("/register", h.createUser)
	router.POST("/login", h.createSession)
	router.DELETE("/logout", h.deleteSession)
	router.GET("/session", h.session)

	// enables CORS requests
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONT-END-ADDRESS")},
		AllowedMethods:   []string{"OPTIONS", "POST", "GET", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	return handler, nil
}
