package data

import (
	"context"
	"fmt"
	"log"
	"task_manager_jwt/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

// NewUserService creates a new instance of UserService.

func NewUserService(client *mongo.Client, dbName, collectionName string) *UserService {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserService{collection}
}

func (us *UserService) Register(user *models.User) error {
    // Check if username is empty
    if user.Username == "" {
        return fmt.Errorf("username cannot be empty")
    }
    
    // Check if password is empty
    if user.Password == "" {
        return fmt.Errorf("password cannot be empty")
    }
    
    // Set default role if not provided
    if user.Role == "" {
        user.Role = "user"
    }

    // Check if username already exists
    count, err := us.collection.CountDocuments(context.TODO(), bson.M{"username": user.Username})
    if err != nil {
        return fmt.Errorf("failed to check if username exists: %w", err)
    }

    if count > 0 {
        return fmt.Errorf("username already exists")
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }

    // Set the hashed password back to the user object
    user.Password = string(hashedPassword)

    // Insert the user into the database
    _, err = us.collection.InsertOne(context.TODO(), user)
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }

    log.Println(user) // Print the user for debugging
    return nil
}


func (us *UserService) GetUser(username string) (*models.User, error) {
	//hash the password authenticate the user using jwt token
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil

}

func (us *UserService) Login(username, password string) (string, error) {
	user, err := us.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}
	return user.Role, nil
}

//delete all users
func (us *UserService) DeleteAllUsers() error {
	_, err := us.collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete users: %w", err)
	}
	return nil
}
