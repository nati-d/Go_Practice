package domain

import "go.mongodb.org/mongo-driver/bson/primitive"



type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     primitive.DateTime `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
}

type TaskRepository interface {
	AddTask(task Task) error
	GetTaskById(id primitive.ObjectID) (Task, error)
	GetAllTasks() ([]Task, error)
	GetMyTasks(userId primitive.ObjectID) ([]Task, error)
	UpdateFullTask(id primitive.ObjectID, task Task) error
	UpdateSomeTask(id primitive.ObjectID, task map[string]interface{}) error
	DeleteTask(id primitive.ObjectID) error
}

type TaskUsecase interface {
	AddTask(task Task) error
	GetTaskById(id primitive.ObjectID) (Task, error)
	GetAllTasks() ([]Task, error)
	GetMyTasks(userId primitive.ObjectID) ([]Task, error)
	UpdateFullTask(id primitive.ObjectID, task Task) error
	UpdateSomeTask(id primitive.ObjectID, task map[string]interface{}) error
	DeleteTask(id primitive.ObjectID) error
}