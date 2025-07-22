package userService

import (
	"task-api/internal/tasksService"
	"time"
)

type User struct {
	ID        uint                `db:"id" json:"id"`
	Email     string              `db:"email" json:"email"`
	Password  string              `db:"password" json:"-"`
	CreatedAt time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt time.Time           `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time          `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	Tasks     []tasksService.Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
