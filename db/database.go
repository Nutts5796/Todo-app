package db

import (
	"github.com/Nutts5796/todo-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}
	DB.AutoMigrate(&models.Task{})
}
