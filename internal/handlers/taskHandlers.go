package handlers

import (
	"context"
	"task-api/internal/tasksService"
	"task-api/internal/web/tasks"
	"gorm.io/gorm"
)

type Handler struct {
	Service *tasksService.Service
}

func NewHandler(service *tasksService.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
for _, t := range allTasks {
    id := uint64(t.ID)
    task := tasks.Task{
        Id:     &id,
        Task:   &t.Text,
        IsDone: &t.IsDone,
    }
    response = append(response, task)
}
return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskToCreate := tasksService.Task{
		Text:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
if err != nil {
    return nil, err
}

id := uint64(createdTask.ID)
return tasks.PostTasks201JSONResponse{
    Id:     &id,
    Task:   &createdTask.Text,
    IsDone: &createdTask.IsDone,
}, nil
}

func (h *Handler) PatchTasks(ctx context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	taskToUpdate := tasksService.Task{
		Model:  gorm.Model{ID: uint(*request.Body.Id)},
		Text:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	updatedTask, err := h.Service.UpdateTask(taskToUpdate)
	if err != nil {
		return nil, err
	}

	id := uint64(updatedTask.ID)

response := tasks.PatchTasks200JSONResponse{
    Id:     &id,
    Task:   &updatedTask.Text,
    IsDone: &updatedTask.IsDone,
}
return response, nil
}

func (h *Handler) DeleteTasks(ctx context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	err := h.Service.DeleteTask(uint(*request.Body.Id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasks204Response{}, nil
}