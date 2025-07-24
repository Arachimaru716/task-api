package userService

import (
	"gorm.io/gorm"
	"task-api/internal/tasksService"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) Update(u *User) error {
	return r.db.Save(u).Error
}

func (r *UserRepository) DeleteHard(id uint) error {
	return r.db.Unscoped().
        Where("id = ?", id).
        Delete(&User{}).
        Error
}

func (r *UserRepository) GetTasksByUser(userID uint) ([]tasksService.Task, error) {
	var list []tasksService.Task
	if err := r.db.Where("user_id = ?", userID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
