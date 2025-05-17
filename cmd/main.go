package main

import (
    "log"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "task-api/internal/db"
    "task-api/internal/service"
    "task-api/internal/handlers"
)

func main() {
    //  Инициализация базы данных
    database, err := db.InitDB()
    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    //  Создание экземпляров слоёв
    taskRepo := service.NewTaskRepository(database) 
    taskService := service.NewTaskService(taskRepo) 
    taskHandler := handlers.NewTaskHandler(taskService) 

    //  Настройка Echo (маршрутизатор)
    e := echo.New()
    e.Use(middleware.CORS()) 
    e.Use(middleware.Logger())

    // Маршруты
    e.GET("/tasks", taskHandler.GetTasks)           
    e.POST("/tasks", taskHandler.PostTask)          
    e.PATCH("/tasks/:id", taskHandler.PatchTask)    
    e.DELETE("/tasks/:id", taskHandler.DeleteTask)  
    // Запуск сервера
    e.Start(":8080")
}