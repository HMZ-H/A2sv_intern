package repositories

import (
	"context"
	"errors"
	"log"
	"os"
	"task_manager/domains"
	"task_manager/infrastructure"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Register(user domains.User) error
	LoginUser(user domains.User) (string, error)
	RegisterAdmin(user domains.User) error
	UpdateUser(username string) error
	Activate(username string) error
	DeActivate(username string) error
}

type userRepository struct {
	collection *mongo.Collection
	jwtSecret  []byte
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(client *mongo.Client) UserRepository {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &userRepository{

		collection: client.Database("task_manager").Collection("users"),
		jwtSecret:  []byte(os.Getenv("jWT_SECRET")),
	}

}

// Register registers a new user in the database
func (r *userRepository) Register(user domains.User) error {
	exists, err := r.usernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}

	findOption := options.Find()
	cursor, err := r.collection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	// Assign role based on existing users
	if !cursor.Next(context.TODO()) {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	user.Password = hashed
	user.Activate = "true"

	_, err = r.collection.InsertOne(context.TODO(), user)
	return err
}

// LoginUser authenticates the user and returns a JWT token
func (r *userRepository) LoginUser(user domains.User) (string, error) {
	var existingUser domains.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	err = infrastructure.CheckPasswordHash(user.Password, existingUser.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	jwtToken, err := infrastructure.GenerateToken(existingUser, r.jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

// usernameExists checks if a username already exists in the database
func (r *userRepository) usernameExists(username string) (bool, error) {
	var user domains.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}
	return err == nil, nil
}

// RegisterAdmin registers a new admin user
func (r *userRepository) RegisterAdmin(user domains.User) error {
	exists, err := r.usernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Role = "admin"
	user.Password = hashed
	user.Activate = "true"

	_, err = r.collection.InsertOne(context.TODO(), user)
	return err
}

// Activate activates a user account
func (r *userRepository) Activate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"activate": "true"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeActivate deactivates a user account
func (r *userRepository) DeActivate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"activate": "false"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// UpdateUser updates a user's role
func (r *userRepository) UpdateUser(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
