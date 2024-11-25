package routes

import (
	"github.com/gorilla/mux"
	"salary-management-system/controllers"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
	
	r.HandleFunc("/api/worklogs", controllers.CreateWorkLog).Methods("POST")
	r.HandleFunc("/api/worklogs", controllers.GetWorkLogs).Methods("GET")
	r.HandleFunc("/api/worklogs/{id}", controllers.GetWorkLog).Methods("GET")
	r.HandleFunc("/api/worklogs/{id}", controllers.DeleteWorkLog).Methods("DELETE")
}
