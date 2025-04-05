package main

import (
	"first_api/internal/database"
	"first_api/internal/handlers"
	taskservice "first_api/internal/taskService"
	userservice "first_api/internal/userService"
	"first_api/internal/web/tasks"
	"first_api/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.InitDB()

	taskRepo := taskservice.NewTaskRepository(database.DB)
	taskService := taskservice.NewService(taskRepo)

	taskHandler := handlers.NewHandler(taskService)

	userRepo := userservice.NewUserRepository(database.DB)
	userService := userservice.NewService(userRepo)

	userHandler := handlers.NewHandlerUser(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	stricHandker := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, stricHandker)

	stricHandkerUser := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, stricHandkerUser)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

}
