package routes

import (
	"github.com/gorilla/mux"
	"golang-todo-api/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Rotas CRUD
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	return router
}
