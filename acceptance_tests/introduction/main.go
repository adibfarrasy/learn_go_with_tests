package main

import (
	"context"
	"log"
	"net/http"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
	"github.com/quii/go-graceful-shutdown/acceptancetests"
)

func main() {
	httpServer := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	server := gracefulshutdown.NewServer(httpServer)

	if err := server.ListenAndServe(context.TODO()); err != nil {
		log.Fatalf("server did not shutdown gracefully %v", err)
	}

	log.Println("server shutdown gracefully.")
}
