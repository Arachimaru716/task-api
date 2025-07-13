package userService

import (
	"time"

	"gorm.io/gorm"
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

func (r *UserRepository) Delete(id uint) error {
	return r.db.Model(&User{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
