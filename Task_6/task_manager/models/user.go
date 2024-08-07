package models

type User struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
	Active   bool   `bson:"active"` // Changed from string to bool
}
