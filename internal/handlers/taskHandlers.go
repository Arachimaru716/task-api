package handlers

import (
	"container/list"
	"context"
	"task-api/internal/tasksService"
	"task-api/internal/userService"
	"task-api/internal/web/tasks"

	"gorm.io/gorm"
)

type Handler struct {
	Service *tasksService.Service
	userService *userService.Service
}

func NewHandler(taskSvc *tasksService.Service, userSvc *userService.Service) *Handler {
	return &Handler{TaskService: taskSvc, UserService: userSvc,}
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

func (h *Handler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject,) (tasks.GetUsersIdTasksRequestObject, error) {
	userID := uint(request.Id)

	list, err := h.userService.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}

	resp := tasks.GetUsersIdTasks200JSONResponse{}
	for _, t := range list {
		id := uint64(t.ID)
		resp = append(resp, tasks.Task{
			Id: &id,
			Task: &t.Text,
			IsDone: &t.IsDone,
			UserId: &id,
		})
	}
	return resp, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTaskRequestObject) (tasks.PostTaskResponseObject, error) {
	body := request.Body

	text := body.Task
	isDone := body.IsDone
	userID := uint(body.UserId)

	createdTask, err := h.Service.CreateTask(text, isDone, userID)
	if err != nil {
		return nil, err
	}

	id := uint64(createdTask.ID)
	return tasks.PostTask201JSONResponse{
		Id:     &id,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
		UserId: &id,
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
