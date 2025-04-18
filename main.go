package main

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID string `json:"id"`
	Task string `json:"task"`
}

var tasks []Task

type requestBody struct {
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	var body requestBody

	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})

	}

	newTask := Task {
		ID: uuid.NewString(),
		Task: body.Task,
	}

	tasks = append(tasks, newTask)

	return c.JSON(http.StatusCreated, newTask)
}

func getHandler(c echo.Context) error {
    if len(tasks) == 0 {
        return c.JSON(http.StatusOK, []Task{})
    }
    return c.JSON(http.StatusOK, tasks)
}

func pathHandler(c echo.Context) error {
	id := c.Param("id")
    
	var body requestBody
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	for i, task := range tasks {
		if task.ID == id {
			if body.Task != "" {
				tasks[i].Task = body.Task
			}
			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/tasks", postHandler)
	e.GET("/tasks", getHandler)
	e.PATCH("/tasks/:id", pathHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Start("localhost:8080")
}
