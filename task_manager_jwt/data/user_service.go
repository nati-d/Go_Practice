package data

import (
	"context"
	"fmt"
	"task_manager_jwt/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client, dbName, collectionName string) *UserService {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserService{collection: collection}
}

// RegisterUser adds a new user to the database.
func (us *UserService) RegisterUser(username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	id := primitive.NewObjectID()
	user := models.User{ID: id,Username: username, Password: string(hashedPassword), Role: role}

	_, err = us.collection.InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

// Login authenticates a user.
func (us *UserService) Login(username, password string) (models.User, error) {
	var user models.User

	err := us.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

//get user by id
func (us *UserService) GetUserById(id string) (models.User, error) {
	var user models.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	err = us.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetAllUsers returns all users from the database.
func (us *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := us.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates a user's profile.
// UpdateUser updates a user's profile.
func (us *UserService) UpdateUser(user *models.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{}}

	if user.Username != "" {
		// Check if the given username is already in the database
		existingUser := models.User{
			ID:       [12]byte{},
			Username: "",
			Password: "",
			Role:     "",
		}
		err := us.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
		if err == nil {
			return fmt.Errorf("username already exists")
		}
		update["$set"].(bson.M)["username"] = user.Username
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		update["$set"].(bson.M)["password"] = string(hashedPassword)
	}

	if user.Role != "" {
		update["$set"].(bson.M)["role"] = user.Role
	}

	_, err := us.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteUser deletes a user from the database.
func (us *UserService) DeleteUser(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = us.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (ts *UserService) PromoteUser(id primitive.ObjectID) error {
	result, err := ts.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return fmt.Errorf("failed to promote user: %w", err)
	}

	// Check if a task was found with the given ID
	if result.MatchedCount == 0 {
		return fmt.Errorf("no user found with the given ID: %w", err)
	}
	return nil
}



