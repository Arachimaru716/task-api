package tasksService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
