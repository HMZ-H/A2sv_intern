package usecases

import (
	"task_manager/domains"
	"task_manager/repositories"
)

// UserUsecase defines the methods for user-related use cases
type UserUsecase interface {
	Register(user domains.User) error
	LoginUser(user domains.User) (string, error)
	RegisterAdmin(user domains.User) error
	UpdateUser(username string) error
	Activate(username string) error
	DeActivate(username string) error
}

// userUsecase implements the UserUsecase interface
type userUsecase struct {
	repo repositories.UserRepository // Repository for data access
}

// NewUserUsecase creates a new instance of userUsecase with the provided repositories
func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

// Register handles the user registration process
func (u *userUsecase) Register(user domains.User) error {
	return u.repo.Register(user) // Call the repository method to register the user
}

// LoginUser handles user login and returns a JWT token
func (u *userUsecase) LoginUser(user domains.User) (string, error) {
	return u.repo.LoginUser(user) // Call the repository method to log in the user
}

// RegisterAdmin handles admin registration
func (u *userUsecase) RegisterAdmin(user domains.User) error {
	return u.repo.RegisterAdmin(user) // Call the repository method to register an admin
}

// UpdateUser updates user information
func (u *userUsecase) UpdateUser(username string) error {
	return u.repo.UpdateUser(username) // Call the repository method to update the user
}

// Activate activates a user account
func (u *userUsecase) Activate(username string) error {
	return u.repo.Activate(username) // Call the repository method to activate the user
}

// DeActivate deactivates a user account
func (u *userUsecase) DeActivate(username string) error {
	return u.repo.DeActivate(username) // Call the repository method to deactivate the user
}
