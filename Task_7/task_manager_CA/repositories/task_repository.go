package repositories

import (
	"context"
	"strconv"
	"task_manager/domains"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetTaskByID(id string) (domains.Task, error)
	GetAllTasks() ([]domains.Task, error)
	CreateTask(task domains.Task) error
	DeleteTask(id string) error
	UpdateTask(id string, task domains.Task) error
}

type taskRepository struct {
	collection *mongo.Collection
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository(client *mongo.Client) TaskRepository {
	return &taskRepository{
		collection: client.Database("task_manager").Collection("tasks"),
	}
}

// Get task by ID
func (r *taskRepository) GetTaskByID(id string) (domains.Task, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var res domains.Task
	err := r.collection.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}

// GetAll retrieves all tasks
func (r *taskRepository) GetAllTasks() ([]domains.Task, error) {
	findOption := options.Find()
	var tasks []domains.Task
	curr, err := r.collection.Find(context.TODO(), bson.D{{}}, findOption)
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.TODO())

	for curr.Next(context.TODO()) {
		var element domains.Task
		if err := curr.Decode(&element); err != nil {
			return nil, err
		}
		tasks = append(tasks, element)
	}

	if err := curr.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Add inserts a new task into the collection
func (r *taskRepository) CreateTask(task domains.Task) error {
	opts := options.Find().SetSort(bson.D{{Key: "id", Value: -1}})
	cursor, err := r.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	LastID := 0
	for cursor.Next(context.TODO()) {
		var existingTask domains.Task
		if err := cursor.Decode(&existingTask); err != nil {
			return err
		}
		id, err := strconv.Atoi(existingTask.ID)
		if err != nil {
			return err
		}
		if id > LastID {
			LastID = id
		}
	}

	LastID++
	task.ID = strconv.Itoa(LastID)
	task.Status = "Pending"
	task.DueDate = time.Now()

	_, err = r.collection.InsertOne(context.TODO(), task)
	return err
}

// Delete removes a task by ID
func (r *taskRepository) DeleteTask(id string) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	return err
}

// Update modifies an existing task
func (r *taskRepository) UpdateTask(id string, task domains.Task) error {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.M{"title": task.Title, "description": task.Description}}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
