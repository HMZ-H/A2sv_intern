package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection
var jwtSecret []byte

func init() {
	client := ConnectToDB()
	collection = UserCollection(client)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func UserCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("task_manager").Collection("users")
}

func usernameExists(collection *mongo.Collection, username string) (bool, error) {
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}
	return err == nil, nil
}

func RegisterUser(role string, user models.User) error {
	got, _ := usernameExists(collection, user.Username)
	if got {
		return errors.New("username exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	filterOp := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	var lastUser models.User
	err = collection.FindOne(context.TODO(), bson.D{}, filterOp).Decode(&lastUser)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Fatal("error, ", err)
	}

	newID := "1"
	if lastUser.ID != "" {
		lastID, err := strconv.Atoi(lastUser.ID)
		if err != nil {
			log.Fatal(err)
		}
		newID = strconv.Itoa(lastID + 1)
	}

	user.Role = role
	user.ID = newID
	user.Password = string(hashedPassword)
	user.Active = true // Set active to true by default during registration

	insertOne, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Print("\n inserted Id: ", insertOne.InsertedID, "\n")
	return nil
}

func LoginUser(user models.User) (string, error) {
	var foundUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("invalid username or password")
	} else if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	claims := jwt.MapClaims{
		"user_id":  foundUser.ID,
		"username": foundUser.Username,
		"role":     foundUser.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Use HS256
	jwtoken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	fmt.Println("message: ", "User successfully logged in, ", "token:", jwtoken)
	return jwtoken, nil
}

func RegisterAdmin(user models.User) error {
	got, _ := usernameExists(collection, user.Username)
	if got {
		return errors.New("username exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Role = "admin"
	user.Password = string(hashedPassword)
	user.Active = true // Admins are active by default

	insertOne, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Print("\n inserted Id: ", insertOne.InsertedID, "\n")
	return nil
}

func Activate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"active": true}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeActivate(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"active": false}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}
