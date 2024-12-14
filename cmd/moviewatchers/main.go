package main

import (
	"fmt"
	"net/http"

	"github.com/jacobkly/moviewatchers/internal/routes"
	"github.com/jacobkly/moviewatchers/internal/services"
)

func main() {
	// populate json first for instant feedback later
	services.PopulateJSON("F:\\")

	router := routes.NewRouter()

	addr := fmt.Sprintf(":%d", 8080)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
