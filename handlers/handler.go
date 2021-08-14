// Package handlers contains all the routes for the API
package handlers

// Handler contains all the routes as methods.
// This makes it easy to spread api keys and secrets between your routes.
// In case you need to add one of those said common parts, you just need to add them to your struct definition.
type Handler struct{}

// NewHandler creates and returns a Handler struct
func NewHandler() *Handler {
	return &Handler{}
}
