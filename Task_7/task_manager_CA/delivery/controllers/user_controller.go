package controllers

import (
	"net/http"
	// "task_manager/delivery/domain"
	"task_manager/domains"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usecase usecases.UserUsecase
}

// NewUserHandler creates a new UserHandler with the provided UserUsecase
func NewUserHandler(u usecases.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user domains.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// LoginUser handles user login
func (h *UserHandler) LoginUser(c *gin.Context) {
	var user domains.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Usecase.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// RegisterAdmin handles admin registration
func (h *UserHandler) RegisterAdmin(c *gin.Context) {
	var newAdmin domains.User
	if err := c.ShouldBindJSON(&newAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin data"})
		return
	}

	if err := h.Usecase.RegisterAdmin(newAdmin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered admin!"})
}

// Promote updates user roles or details
func (h *UserHandler) Promote(c *gin.Context) {
	username := c.Param("username")
	var updatedUser domains.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.UpdateUser(username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// Activate activates a user account
func (h *UserHandler) Activate(c *gin.Context) {
	username := c.Param("username")

	if err := h.Usecase.Activate(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully activated!"})
}

// DeActivate deactivates a user account
func (h *UserHandler) DeActivate(c *gin.Context) {
	username := c.Param("username")

	if err := h.Usecase.DeActivate(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deactivated!"})
}
