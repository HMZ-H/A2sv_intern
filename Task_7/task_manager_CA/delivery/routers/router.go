package router

import (
	"log"
	"task_manager/delivery/controllers"
	"task_manager/repositories"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRouting(client *mongo.Client) {
	router := gin.Default()

	// Initialize repositories with the MongoDB client
	taskRepo := repositories.NewTaskRepository(client)
	userRepo := repositories.NewUserRepository(client)

	// Initialize use cases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserHandler(userUsecase)

	// Define routes
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks/:id", taskController.GetTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func NewUserRepository(client *mongo.Client) {
	panic("unimplemented")
}

func NewTaskRepository(client *mongo.Client) {
	panic("unimplemented")
}
