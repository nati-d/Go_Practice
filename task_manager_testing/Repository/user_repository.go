package repository

import (
	"context"
	infrastructure "task_manager_refactored/Infrastructure"
	"task_manager_refactored/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepository{collection: collection}
}

// RegisterUser adds a new user to the database.
func (ur *UserRepository) RegisterUser(username, password, role string) error {

	id := primitive.NewObjectID()
	user := domain.User{ID: id, Username: username, Password: password, Role: role}

	_, err := ur.collection.InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

// Login authenticates a user.
func (ur *UserRepository) Login(username, password string) (domain.User, error) {
	var user domain.User

	err := ur.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	err = infrastructure.ComparePasswords(user.Password, password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// get user by id
func (ur *UserRepository) GetUserById(id primitive.ObjectID) (domain.User, error) {
	var user domain.User

	err := ur.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetAllUsers returns all users from the database.
func (ur *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := ur.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates a user in the database.
func (ur *UserRepository) UpdateUser(oid primitive.ObjectID, user domain.User) error {

	_, err := ur.collection.ReplaceOne(context.TODO(), bson.M{"_id": oid}, &user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database.
func (ur *UserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := ur.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
