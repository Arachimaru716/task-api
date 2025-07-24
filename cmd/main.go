package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task-api/internal/database"
	"task-api/internal/handlers"
	"task-api/internal/tasksService"
	"task-api/internal/userService"
	"task-api/internal/web/tasks"
	"task-api/internal/web/users"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&tasksService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	taskRepo := tasksService.NewTaskRepository(database.DB)
	taskSvc := tasksService.NewService(taskRepo)

	userRepo := userService.NewUserRepository(database.DB)
	userSvc := userService.NewService(userRepo)

	userH := handlers.NewUserHandlers(userSvc)
	taskH := handlers.NewHandler(taskSvc, userSvc)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksH := tasks.NewStrictHandler(taskH, nil)
	tasks.RegisterHandlers(e, strictTasksH)

	users.RegisterHandlers(e, userH)

	log.Println("Server started on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
