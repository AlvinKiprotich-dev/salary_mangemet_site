package routes

import (
	"github.com/gorilla/mux"
	"salary-management-system/controllers"
)

func WorkLogRoutes(r *mux.Router) {
	r.HandleFunc("/worklogs", controllers.GetWorkLogs).Methods("GET")
	r.HandleFunc("/worklogs/{id}", controllers.GetWorkLog).Methods("GET")
	r.HandleFunc("/worklogs", controllers.CreateWorkLog).Methods("POST")
	r.HandleFunc("/worklogs/{id}", controllers.UpdateWorkLog).Methods("PUT")
	r.HandleFunc("/worklogs/{id}", controllers.DeleteWorkLog).Methods("DELETE")
}
