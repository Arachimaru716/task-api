package userService

import (
	"errors"
	"task-api/internal/tasksService"
)

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateUser(email, password string) (User, error) {
	if email == "" {
		return User{}, errors.New("email is required")
	}
	if password == "" {
		return User{}, errors.New("password is required")
	}
	u := &User{Email: email, Password: password}
	if err := s.repo.Create(u); err != nil {
		return User{}, err
	}
	return *u, nil
}

func (s *Service) UpdateUser(id uint, email, password string) (User, error) {
	if id == 0 {
		return User{}, errors.New("invalid id")
	}
	if email == "" && password == "" {
		return User{}, errors.New("nothing to update")
	}
	u := &User{ID: id, Email: email, Password: password}
	if err := s.repo.Update(u); err != nil {
		return User{}, err
	}
	return *u, nil
}

func (s *Service) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(id)
}

func (s *Service) GetTasksForUser(userID uint) ([]tasksService.Task, error) {
    return s.repo.GetTasksByUser(userID)
}
