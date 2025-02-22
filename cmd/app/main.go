package main

import (
	"first_api/internal/database"
	"first_api/internal/handlers"
	taskservice "first_api/internal/taskService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskservice.Task{})

	repo := taskservice.NewTaskRepository(database.DB)
	servise := taskservice.NewService(repo)

	handler := handlers.NewHandler(servise)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHadler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
