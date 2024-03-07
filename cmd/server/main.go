// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"github.com/Clankcoll/WorkersCompanion/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Register the handler for the work time calculator
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/worktime", handlers.WorkTimeCalculator)
	mux.HandleFunc("/union", handlers.UnionCalculator)
	mux.HandleFunc("/homeoffice", handlers.HomeOfficeCalculator)
	mux.HandleFunc("/clockinout", handlers.ClockInOut)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
