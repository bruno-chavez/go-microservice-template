// Package session takes care of setting up the session store
package session

import (
	"gopkg.in/boj/redistore.v1"
	"os"
	"strconv"
)

// NewSession creates a new Redis session store connection and returns it.
func NewSession() (*redistore.RediStore, error) {

	storeSize, err := strconv.Atoi(os.Getenv("SESSION_STORE_SIZE"))
	if err != nil {
		return nil, err
	}

	// connects to the Redis session store
	store, err := redistore.NewRediStore(
		storeSize,
		"tcp",
		os.Getenv("SESSION_STORE_ADDRESS"),
		os.Getenv("SESSION_STORE_PASSWORD"),
		[]byte(os.Getenv("SESSION_STORE_KEY")))
	if err != nil {
		return nil, err
	}

	return store, nil
}
