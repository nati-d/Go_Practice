package domain

import "go.mongodb.org/mongo-driver/bson/primitive"



type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type UserRepository interface {
	RegisterUser(username, password, role string) error
	Login(username, password string) (User, error)
	UpdateUser(id primitive.ObjectID, user User) error
	GetUserById(id primitive.ObjectID) (User, error)
	DeleteUser(id primitive.ObjectID) error
	GetAllUsers() ([]User, error)
}


type UserUsecase interface {
	RegisterUser(username, password, role string) error
	Login(username, password string) (User, error)
	UpdateUser(id primitive.ObjectID, user User) error
	GetUserById(id primitive.ObjectID) (User, error)
	DeleteUser(id primitive.ObjectID) error
	GetAllUsers() ([]User, error)
}