package service

import "gorm.io/gorm"

// Task представляет задачу в базе данных
type Task struct {
    ID     string `gorm:"primaryKey" json:"id"`  // Уникальный ID задачи
    Task   string `json:"task"`                  // Описание задачи
    IsDone bool   `json:"is_done"`              // Флаг выполнения задачи
}

// InitDB инициализирует миграции
func InitDB(db *gorm.DB) error {
    return db.AutoMigrate(&Task{})
}