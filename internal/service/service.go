package service

import (
    "github.com/google/uuid"
)

// TaskService представляет сервис для работы с задачами
type TaskService interface {
    CreateTask(taskName string) (Task, error)
    GetAllTasks() ([]Task, error)
    GetTaskByID(id string) (Task, error)
    UpdateTask(id, taskName string, isDone bool) (Task, error)
    DeleteTask(id string) error
}

type TaskRequest struct {
    Task   string `json:"task"`    // Описание задачи
    IsDone bool   `json:"is_done"` // Флаг выполнения задачи
}

// Структура taskService
type taskService struct {
    repo TaskRepository
}

// Конструктор TaskService
func NewTaskService(repo TaskRepository) TaskService {
    return &taskService{repo: repo}
}

// Методы сервиса

func (s *taskService) CreateTask(taskName string) (Task, error) {
    task := Task{
        ID:     uuid.NewString(),
        Task:   taskName,
        IsDone: false, // Новая задача всегда не выполнена
    }

    if err := s.repo.CreateTask(task); err != nil {
        return Task{}, err
    }

    return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
    return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
    return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id, taskName string, isDone bool) (Task, error) {
    task, err := s.repo.GetTaskByID(id)
    if err != nil {
        return Task{}, err
    }

    task.Task = taskName
    task.IsDone = isDone

    if err := s.repo.UpdateTask(task); err != nil {
        return Task{}, err
    }

    return task, nil
}

func (s *taskService) DeleteTask(id string) error {
    return s.repo.DeleteTask(id)
}