// package router

// import (
// 	"task_manager/controllers"
// 	"task_manager/middleware"

// 	// "task_manager/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func SetUpRouter() {
// 	router := gin.Default()

// 	// User routes
// 	router.POST("/register", controllers.Register)
// 	router.POST("/login", controllers.Login)

// 	allow := router.Group("")
// 	allow.Use(middleware.AuthUser())
// 	allow.GET("/tasks", controllers.GetTasks)
// 	allow.GET("/tasks/:id", controllers.GetTask)
// 	protected := router.Group("/admin")
// 	protected.Use(middleware.AuthMiddleware("admin"))
// 	protected.PUT("/tasks/:id", controllers.UpdateTask)
// 	protected.DELETE("/tasks/:id", controllers.DeleteTask)
// 	protected.POST("/tasks", controllers.CreateTask)
// 	protected.POST("/register", controllers.RegisterAdmin)

// 	protected.POST("/activate/:username", controllers.Activate)
// 	protected.POST("/deactivate/:username", controllers.DeActivate)
// 	protected.GET("/promote/:username", controllers.Promote)

// 	// Start the server
// 	router.Run("localhost:8080")
// }

package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	router := gin.Default()

	// User routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Routes accessible to authenticated users
	allow := router.Group("")
	allow.Use(middleware.AuthUser())
	allow.GET("/tasks", controllers.GetTasks)
	allow.GET("/tasks/:id", controllers.GetTask)

	// Routes accessible to admin users only
	protected := router.Group("/admin")
	protected.Use(middleware.AuthMiddleware("admin"))
	protected.PUT("/tasks/:id", controllers.UpdateTask)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)
	protected.POST("/tasks", controllers.CreateTask)
	protected.POST("/register", controllers.RegisterAdmin)
	protected.POST("/activate/:username", controllers.Activate)
	protected.POST("/deactivate/:username", controllers.DeActivate)
	protected.GET("/promote/:username", controllers.Promote)

	// Start the server
	router.Run("localhost:8080")
}
