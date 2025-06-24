package tasksService

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Task, error)
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(id uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *TaskRepository) Create(task Task) (Task, error) {
	result := r.db.Create(&task)
	return task, result.Error
}

func (r *TaskRepository) Update(task Task) (Task, error) {
	result := r.db.Save(&task)
	return task, result.Error
}

func (r *TaskRepository) Delete(id uint) error {
	result := r.db.Delete(&Task{}, id)
	return result.Error
}