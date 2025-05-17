package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "task-api/internal/service"
)

// TaskHandler представляет обработчики для работы с задачами
type TaskHandler struct {
    service service.TaskService
}

// Конструктор TaskHandler
func NewTaskHandler(s service.TaskService) *TaskHandler {
    return &TaskHandler{service: s}
}

// Методы обработчиков

// GetTasks возвращает список всех задач
func (h *TaskHandler) GetTasks(c echo.Context) error {
    tasks, err := h.service.GetAllTasks()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
    }
    return c.JSON(http.StatusOK, tasks)
}

// PostTask создаёт новую задачу
func (h *TaskHandler) PostTask(c echo.Context) error {
    var req service.TaskRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    task, err := h.service.CreateTask(req.Task)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
    }

    return c.JSON(http.StatusCreated, task)
}

// PatchTask обновляет задачу
func (h *TaskHandler) PatchTask(c echo.Context) error {
    id := c.Param("id")

    var req service.TaskRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    updatedTask, err := h.service.UpdateTask(id, req.Task, req.IsDone)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
    }

    return c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask удаляет задачу
func (h *TaskHandler) DeleteTask(c echo.Context) error {
    id := c.Param("id")

    if err := h.service.DeleteTask(id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
    }
    return c.NoContent(http.StatusNoContent)
}