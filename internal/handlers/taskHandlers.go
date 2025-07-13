package handlers

import (
	"context"
	"gorm.io/gorm"
	"task-api/internal/tasksService"
	"task-api/internal/web/tasks"
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

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskToUpdate := tasksService.Task{
		Model:  gorm.Model{ID: uint(request.Id)},
		Text:   request.Body.Task,   // Убрали * - это уже string
		IsDone: request.Body.IsDone, // Убрали * - это уже bool
	}

	updatedTask, err := h.Service.UpdateTask(taskToUpdate)
	if err != nil {
		return nil, err
	}

	id := uint64(updatedTask.ID)
	return tasks.PatchTasksId200JSONResponse{
		Id:     &id,
		Task:   &updatedTask.Text,   // Здесь & нужно, так как ответ требует *string
		IsDone: &updatedTask.IsDone, // И здесь & для *bool
	}, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTask(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
