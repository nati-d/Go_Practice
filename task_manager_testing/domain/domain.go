package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     primitive.DateTime `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
}

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}




type TaskRepository interface {
	AddTask(task Task) error
	GetTaskById(id primitive.ObjectID) (Task, error)
	GetAllTasks() ([]Task, error)
	GetMyTasks(userId primitive.ObjectID) ([]Task, error)
	UpdateFullTask(id primitive.ObjectID,task Task) error
	UpdateSomeTask(id primitive.ObjectID,task map[string]interface{}) error
	DeleteTask(id primitive.ObjectID) error
}


type UserRepository interface {
	RegisterUser(username, password, role string) error
	Login(username, password string) (User, error)
	UpdateUser(id primitive.ObjectID,user User) error
	GetUserById(id primitive.ObjectID) (User, error)
	DeleteUser(id primitive.ObjectID) error
	GetAllUsers() ([]User, error)
}
type TaskUsecase interface {
	AddTask(task Task) error
	GetTaskById(id primitive.ObjectID) (Task, error)
	GetAllTasks() ([]Task, error)
	GetMyTasks(userId primitive.ObjectID) ([]Task, error)
	UpdateFullTask(id primitive.ObjectID,task Task) error
	UpdateSomeTask(id primitive.ObjectID,task map[string]interface{}) error
	DeleteTask(id primitive.ObjectID) error
}


type UserUsecase interface {
	RegisterUser(username, password, role string) error
	Login(username, password string) (User, error)
	UpdateUser(id primitive.ObjectID,user User) error
	GetUserById(id primitive.ObjectID) (User, error)
	DeleteUser(id primitive.ObjectID) error
	GetAllUsers() ([]User, error)
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}