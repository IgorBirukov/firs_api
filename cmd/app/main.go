package main

import (
	"first_api/internal/database"
	"first_api/internal/handlers"
	taskservice "first_api/internal/taskService"
	"first_api/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.InitDB()

	repo := taskservice.NewTaskRepository(database.DB)
	service := taskservice.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	stricHandker := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, stricHandker)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

	//database.DB.AutoMigrate(&taskservice.Task{})

	// repo := taskservice.NewTaskRepository(database.DB)
	// servise := taskservice.NewService(repo)

	// handler := handlers.NewHandler(servise)

	// router := mux.NewRouter()
	// router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	// router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	// router.HandleFunc("/api/patch/{id}", handler.PatchTaskHadler).Methods("PATCH")
	// router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	// http.ListenAndServe(":8080", router)
}
