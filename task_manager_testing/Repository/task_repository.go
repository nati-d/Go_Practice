package repository

import (
	"context"
	"task_manager_testing/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client, dbName, collectionName string) *TaskRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &TaskRepository{collection: collection}
}

func (tr *TaskRepository) AddTask(task domain.Task) error {
	_, err := tr.collection.InsertOne(context.TODO(), task)
	return err
}

func (tr *TaskRepository) GetMyTasks(userID primitive.ObjectID) ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := tr.collection.Find(context.TODO(), bson.M{"created_by": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := tr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *TaskRepository) GetTaskById(id primitive.ObjectID) (domain.Task, error) {
	var task domain.Task
	err := tr.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func (tr *TaskRepository) UpdateFullTask(id primitive.ObjectID, task domain.Task) error {
	_, err := tr.collection.ReplaceOne(context.TODO(), bson.M{"_id": id}, &task)
	return err
}

func (tr *TaskRepository) UpdateSomeTask(id primitive.ObjectID, update map[string]interface{}) error {
	_, err := tr.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (tr *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	_, err := tr.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
