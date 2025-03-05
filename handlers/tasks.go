package handlers

import (
	"net/http"

	"github.com/Nutts5796/todo-app/db"
	"github.com/Nutts5796/todo-app/models"
	"github.com/gin-gonic/gin"
)

// Получение всех задач
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := db.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить задачи"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Создание новой задачи
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать задачу"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Обновление задачи
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача успешно обновлена", "task": task})
}

// Удаление задачи
func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	if err := db.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить задачу"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача успешно удалена"})
}
