// Package main initializes and starts an HTTP server for the moviewatchers application.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jacobkly/moviewatchers/server/internal/routes"
	"github.com/jacobkly/moviewatchers/server/internal/services"
)

// main is the entry point for the MovieWatchers application, fully compatible with Windows.
// It populates data used by the client, sets up the HTTP router, and starts the server to listen
// on localhost at port 8080. Note that it may handle hidden files differently on Mac/Linux.
func main() {
	err := services.PopulateJSON("F:\\")
	if err != nil {
		log.Fatal("Issue when populating JSON data: ", err)
	}

	router := routes.NewRouter()
	addr := fmt.Sprintf("localhost:%d", 8080)
	fmt.Printf("Server listening on http://%s\n", addr) // http while in development
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
