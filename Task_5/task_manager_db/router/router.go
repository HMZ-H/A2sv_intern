package router

import (
	"task_manager_db/controllers"

	"github.com/gin-gonic/gin"
)

// DeleteTask
func CreateRouting() {
	router := gin.Default()
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("tasks/:id", controllers.DeleteTask)
	router.POST("tasks/", controllers.CreateTask)
	router.Run("localhost:8080")
}
