package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=taskdb port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Автомиграция для создания таблицы tasks
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func getHandler(c echo.Context) error {
	var tasks []Task

	// Получаем все задачи из БД
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func postHandler(c echo.Context) error {
	var req TaskRequest

	// Декодируем JSON-тело запроса
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Создаём новую задачу
	task := Task{
		ID:     uuid.NewString(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	// Сохраняем задачу в БД
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Находим задачу по ID
	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	// Обновляем поля задачи
	task.Task = req.Task
	task.IsDone = req.IsDone

	// Сохраняем изменения в БД
	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}
func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	// Удаляем задачу из БД
	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/tasks", postHandler)
	e.GET("/tasks", getHandler)
	e.PATCH("/tasks/:id", patchHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Start("localhost:8080")
}
