package handlers

import (
	"backend/internal/models"
	"backend/internal/store"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(router *gin.Engine, db *sql.DB) {
	router.GET("/tasks", func(c *gin.Context) {
		tasks, err := store.GetTasks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})
}

func CreateTask(router *gin.Engine, db *sql.DB) {
	router.POST("/tasks", func(c *gin.Context) {
		var task models.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdTask, err := store.CreateTask(db, task.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, createdTask)
	})
}
