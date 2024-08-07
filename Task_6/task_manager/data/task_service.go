package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"task_manager/models"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/options"
)

// var taskCollection *mongo.Collection

// func InitTaskService() {
// 	client := ConnectToDB()
// 	taskCollection = TaskCollection(client)
// }

func ConnectToDB() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	mongoURI := os.Getenv("MONGO_URI") //loading the mongo uri from the .env file
	// mongoURI := "mongodb://localhost:27017"
	// mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error during connecting to mongodb: ", err)
	}

	err = client.Ping(context.TODO(), nil) // Testing the connection

	if err != nil {
		log.Fatal("Error during connecting to mongodb: ", err)
	}
	fmt.Print("successfully connected!")
	return client //sends client that is the connected db location
}

func TaskCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("task_manager").Collection("tasks")
}

func GetAllTasks() []models.Task {
	var client = ConnectToDB()
	var collection = TaskCollection(client)

	filterOp := options.Find()

	var tasks []models.Task

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, filterOp)

	if err != nil {
		log.Fatal("error searching for the task.", err)
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal("error retrieving")
		}
		tasks = append(tasks, task)
	}
	return tasks

}

func GetTaskByID(id string) (models.Task, error) {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	var task *models.Task
	filter := bson.D{{Key: "id", Value: id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return *task, err
	}
	return *task, nil

}

func CreateTask(newTask models.Task) error {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	filterOp := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	var lastTask models.Task
	err := collection.FindOne(context.TODO(), bson.D{}, filterOp).Decode(&lastTask)
	// fmt.Println("one||||||\n")
	if err != nil && err != mongo.ErrNoDocuments {
		log.Fatal("error, ", err)
	}
	newTask.Status = "Pending"

	// taking the last id from the db and converting to int.
	LastID, err := strconv.Atoi(lastTask.ID)
	if err != nil && err == mongo.ErrNoDocuments {
		LastID = 0
	} else if err != nil {
		log.Fatal(err)
	}
	LastID++
	t := LastID
	newTask.ID = strconv.Itoa(t)
	newTask.DueDate = time.Now()
	_, err = collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		log.Fatal("Error during Inserting")
	}
	// fmt.Println("\n Insert ID: ", insertOne.InsertedID, "\n")
	return nil

}

func UpdateTask(id string, updateTask models.Task) error {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	filter := bson.D{{Key: "id", Value: id}}
	var tasks *models.Task
	err := collection.FindOne(context.TODO(), filter).Decode(&tasks)
	if err != nil && err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		log.Fatal(err)
	}

	update := bson.D{{Key: "$set", Value: bson.M{"title": updateTask.Title}}}

	task, err := collection.UpdateByID(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(task)
	return nil

}

func DeleteTask(id string) error {
	var client = ConnectToDB()
	var collection = TaskCollection(client)
	deletedTask, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		return err
	}
	fmt.Println("Deleted task: ", deletedTask)
	return nil
}
