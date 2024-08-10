package controllers

import (
	"net/http"
	"task_manager/usecases"

	"task_manager/domains"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase usecases.TaskUsecase
}

func NewTaskController(useCase usecases.TaskUsecase) *TaskController {
	return &TaskController{taskUseCase: useCase}
}

// Create Task handles
func (ct *TaskController) CreateTask(c *gin.Context) {
	var task domains.Task

	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ct.taskUseCase.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully."})

}

// Get Task by ID

func (ct *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := ct.taskUseCase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Get All Task
func (ct *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := ct.taskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Update Task

func (ct *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domains.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := ct.taskUseCase.UpdateTask(id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, task)
}

//Delete Task

func (ct *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := ct.taskUseCase.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
