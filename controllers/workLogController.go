package controllers

import (
	"encoding/json"
	 
	"net/http"
	"salary-management-system/database"
	"salary-management-system/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateWorkLog handles creating a new work log
func CreateWorkLog(w http.ResponseWriter, r *http.Request) {
	var workLog models.WorkLog
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workLog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := database.Client.Database("salary_management").Collection("work_logs")
	_, err := collection.InsertOne(r.Context(), workLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workLog)
}

// GetWorkLogs handles fetching all work logs
func GetWorkLogs(w http.ResponseWriter, r *http.Request) {
	var workLogs []models.WorkLog
	collection := database.Client.Database("salary_management").Collection("work_logs")
	cursor, err := collection.Find(r.Context(), bson.D{{}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	for cursor.Next(r.Context()) {
		var workLog models.WorkLog
		if err := cursor.Decode(&workLog); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		workLogs = append(workLogs, workLog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workLogs)
}

// GetWorkLog handles fetching a specific work log by ID
func GetWorkLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workLogID := vars["id"]

	var workLog models.WorkLog
	collection := database.Client.Database("salary_management").Collection("work_logs")
	err := collection.FindOne(r.Context(), bson.D{{"id", workLogID}}).Decode(&workLog)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Work log not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workLog)
}

// DeleteWorkLog handles deleting a work log by ID
func DeleteWorkLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workLogID := vars["id"]

	collection := database.Client.Database("salary_management").Collection("work_logs")
	_, err := collection.DeleteOne(r.Context(), bson.D{{"id", workLogID}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
