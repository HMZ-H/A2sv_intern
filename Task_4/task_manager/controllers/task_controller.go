package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask := data.CreateTask(task)
	c.JSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := data.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
