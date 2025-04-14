package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	var body requestBody

	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})

	}

	task = body.Task

	response := fmt.Sprintf("Task updated: %s", task)

	return c.JSON(http.StatusOK, map[string]string{"message": response})
}

func getHandler(c echo.Context) error {

	response := fmt.Sprintf("hello, %s", task)

	return c.JSON(http.StatusOK, map[string]string{"message": response})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/post", postHandler)
	e.GET("/get", getHandler)

	e.Start("localhost:8080")
}
