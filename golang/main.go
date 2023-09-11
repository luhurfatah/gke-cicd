package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks []Task

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.Run(":8080")
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	tasks = append(tasks, task)

	c.JSON(http.StatusCreated, task)
}
