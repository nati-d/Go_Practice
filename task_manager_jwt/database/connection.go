package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)


// ConnectToMongoDB establishes a connection to a MongoDB database.
func ConnectToMongoDB(uri string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}


	fmt.Println("Connected to MongoDB!")
	return client, nil
}


// DisconnectFromMongoDB closes the connection to a MongoDB database.
func DisconnectFromMongoDB(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB!")
	return nil
}