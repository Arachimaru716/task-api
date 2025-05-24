package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm" 
)

var db *gorm.DB

// InitDB инициализирует подключение к базе данных
func InitDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=yourpassword dbname=taskdb port=5432 sslmode=disable"
    var err error

    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    return db, nil
}