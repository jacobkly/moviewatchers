package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jacobkly/moviewatchers/internal/routes"
	"github.com/jacobkly/moviewatchers/internal/services"
)

func main() {
	// populate json object first for instant feedback later
	err := services.PopulateJSON("F:\\")
	if err != nil {
		log.Fatal(err)
	}

	router := routes.NewRouter()

	addr := fmt.Sprintf("localhost:%d", 8080)
	fmt.Printf("Server listening on http://%s\n", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
