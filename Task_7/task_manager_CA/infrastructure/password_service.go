package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt and returns the hashed password as a string
func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password with a cost of 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Return an empty string and the error if hashing fails
	}
	return string(bytes), nil // Return the hashed password
}

// CheckPasswordHash compares a hashed password with a plain password and returns an error if they do not match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err // Return nil if the password matches, otherwise return an error
}
