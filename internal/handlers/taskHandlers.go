package handlers

import (
	"context"
	"task-api/internal/tasksService"
	"task-api/internal/userService"
	"task-api/internal/web/tasks"

	"gorm.io/gorm"
)

type Handler struct {
	Service     *tasksService.Service
	userService *userService.Service
}

func NewHandler(taskSvc *tasksService.Service, userSvc *userService.Service) *Handler {
	return &Handler{Service: taskSvc, userService: userSvc}
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

func (h *Handler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userID := uint(request.Id)

	list, err := h.userService.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}

	resp := make(tasks.GetTasksByUserID200JSONResponse, 0, len(list))
	for _, t := range list {
		taskID := uint64(t.ID)
		text := t.Text
		isDone := t.IsDone
		var userIDval uint64
		if t.UserID != nil {
			userIDval = uint64(*t.UserID)
		}

		resp = append(resp, tasks.Task{
			Id:     &taskID,
			Task:   &text,
			IsDone: &isDone,
			UserId: &userIDval,
		})
	}
	return resp, nil
}

func (h *Handler) PostTask(ctx context.Context, request tasks.PostTaskRequestObject) (tasks.PostTaskResponseObject, error) {
	body := request.Body

	text := body.Task
	isDone := body.IsDone
	userID := body.UserId

	createdTask, err := h.Service.CreateTask(text, isDone, uint(userID))
	if err != nil {
		return nil, err
	}

	taskID := uint64(createdTask.ID)
	return tasks.PostTask201JSONResponse{
		Id:     &taskID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
		UserId: &userID,
	}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskToUpdate := tasksService.Task{
		Model:  gorm.Model{ID: uint(request.Id)},
		Text:   request.Body.Task,
		IsDone: request.Body.IsDone,
	}

	updatedTask, err := h.Service.UpdateTask(taskToUpdate)
	if err != nil {
		return nil, err
	}

	id := uint64(updatedTask.ID)
	return tasks.PatchTasksId200JSONResponse{
		Id:     &id,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTask(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
