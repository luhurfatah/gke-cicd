package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/healthz", healthCheck)
	router.POST("/tasks", createTask)
	return router
}

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, tasks, response)
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	task := Task{Title: "New Task", Description: "New Description"}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, task, response)
	assert.Contains(t, tasks, task)
}
