package main

import (
	"fmt"
	"log"
	"net/http"
	"salary-management-system/database"
	"salary-management-system/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
		return
	}

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
