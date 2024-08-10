package infrastructure

import (
	"task_manager/domains"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(existingUser domains.User, jwtSecret []byte) (string, error) {
	claims := jwt.MapClaims{
		"id":       existingUser.ID,
		"username": existingUser.Username,
		"role":     existingUser.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
