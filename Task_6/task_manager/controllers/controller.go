package controllers

import (
	// "fmt"

	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var tasks []models.Task = []models.Task{}

// Get All Tasks
func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()

	c.IndentedJSON(http.StatusOK, tasks)
}

// Get Task By ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// Create Tasks
func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// Update Tasks
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updateTask models.Task

	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTask(id, updateTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// Delete Tasks
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// Register a new user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := c.Param("role")
	err := data.RegisterUser(role, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

// Login a user
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Incorrect input"})
		return
	}
	token, err := data.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "token": token})
}

func ProtectedHandler(c *gin.Context) {
	claims, _ := c.Get("user")
	userClaim := claims.(jwt.MapClaims)
	userID := userClaim["id"]
	username := userClaim["username"]

	response := gin.H{
		"message":  "Successfully accessed the protected route",
		"id":       userID,
		"username": username,
	}
	c.JSON(200, response)
}

func RegisterAdmin(c *gin.Context) {
	var newAdmin models.User
	err := c.ShouldBindJSON(&newAdmin)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
	err = data.RegisterAdmin(newAdmin)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully registered!"})
}

func Activate(c *gin.Context) {
	username := c.Param("username")
	err := data.Activate(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "successfully activated!"})
}

func DeActivate(c *gin.Context) {
	username := c.Param("username")
	err := data.DeActivate(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "successfully deactivated!"})
}

func Promote(c *gin.Context) {
	username := c.Param("username")
	var UpdatedUser models.User
	if err := c.ShouldBind(&UpdatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(tasks)
	err := data.UpdateUser(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"}) // indicates the task with given id is not found in the db
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"}) // updates successfully
}
