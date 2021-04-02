// go-microservice-template main purpose is to be a starting point
// for a REST API in Go. Reducing boilerplate writing and speeding up development of new microservices
package main

import (
	"context"
	"github.com/bruno-chavez/go-microservice-template/handlers"
	"github.com/bruno-chavez/go-microservice-template/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Makes error line show for log.Println()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	handler := handlers.NewHandler()

	srvr := server.NewServer(handler)

	// Graceful Shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srvr.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started in port " + os.Getenv("PORT"))

	<-done
	log.Println("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Close connections and stuff before shutting down
	defer func() {
		cancel()
	}()

	if err := srvr.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("server exited properly")
}
