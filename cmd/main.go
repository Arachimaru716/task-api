package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task-api/internal/database"
	"task-api/internal/handlers"
	"task-api/internal/tasksService"
	"task-api/internal/web/tasks"
)

func main() {
	database.InitDB()
	
	if err := database.DB.AutoMigrate(&tasksService.Task{}); err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}