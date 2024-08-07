package data

import (
	"context"
	"fmt"
	"task_manager_jwt/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService provides CRUD operations for tasks.
type TaskService struct {
	collection *mongo.Collection
}

// NewTaskService creates a new TaskService with the given MongoDB collection.
func NewTaskService(client *mongo.Client, dbName, collectionName string) *TaskService {
	collection := client.Database(dbName).Collection(collectionName)
	return &TaskService{collection: collection}
}

// AddTask adds a new task to the database.
func (ts *TaskService) AddTask(task *models.Task) error {
	if task.Title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if task.Description == "" {
		return fmt.Errorf("task description cannot be empty")
	}
	if task.Status == "" {
		return fmt.Errorf("task status cannot be empty")
	}

	if task.Status != "Not Started" && task.Status != "In Progress" && task.Status != "Completed" {
		return fmt.Errorf("task status must be one of 'Not Started', 'In Progress', or 'Completed'")
	}

	// Insert the task into the database
	_, err := ts.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	return nil
}

//get my tasks
func (ts *TaskService) GetMyTasks(userId string) ([]models.Task, error) {
	// Retrieve all tasks from the database
	dm,_ := primitive.ObjectIDFromHex(userId)
	result, err := ts.collection.Find(context.TODO(), bson.M{"created_by": dm})
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer result.Close(context.TODO())

	// Decode the tasks into a slice
	var tasks []models.Task
	err = result.All(context.TODO(), &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tasks: %w", err)
	}
	return tasks, nil
}

// GetAllTasks retrieves all tasks from the database.
func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	// Retrieve all tasks from the database
	result, err := ts.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer result.Close(context.TODO())

	// Decode the tasks into a slice
	var tasks []models.Task
	err = result.All(context.TODO(), &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tasks: %w", err)
	}
	return tasks, nil
}

// GetTaskById retrieves a task with the given ID from the database.
func (ts *TaskService) GetTaskById(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	err := ts.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, fmt.Errorf("no task found with the given ID: %w", err)
		}
		return models.Task{}, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

// UpdateFullTask updates a task with the given ID with the provided task data.
func (ts *TaskService) UpdateFullTask(id primitive.ObjectID, task models.Task) error {
	if task.Title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if task.Description == "" {
		return fmt.Errorf("task description cannot be empty")
	}
	if task.Status == "" {
		return fmt.Errorf("task status cannot be empty")
	}

	if task.Status != "Not Started" && task.Status != "In Progress" && task.Status != "Completed" {
		return fmt.Errorf("task status must be one of 'Not Started', 'In Progress', or 'Completed'")
	}
	//if the loggedin user in he can update others user but not admin and root user
	


	// Replace the task with the given ID with the provided task data
	result, err := ts.collection.ReplaceOne(context.TODO(), bson.M{"_id": id}, task)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	// Check if a task was found with the given ID
	if result.MatchedCount == 0 {
		return fmt.Errorf("no task found with the given ID: %w", err)
	}
	return nil
}

// Update the task with the given ID with the provided update data
func (ts *TaskService) UpdateSomeTask(id primitive.ObjectID, update bson.M) error {
	result, err := ts.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	// Check if a task was found with the given ID
	if result.MatchedCount == 0 {
		return fmt.Errorf("no task found with the given ID: %w", err)
	}
	return nil
}

// DeleteTask deletes a task with the given ID.
func (ts *TaskService) DeleteTask(id primitive.ObjectID) error {
	result, err := ts.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	// Check if a task was found with the given ID
	if result.DeletedCount == 0 {
		return fmt.Errorf("no task found with the given ID: %w", err)
	}

	return nil
}
