package main

import (
	"log"
	"os"
	"task_manager_refactored/Delivery/routers"
	"task_manager_refactored/config/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Connect to the MongoDB database
	uri := os.Getenv("MONGO_DB_URI")
	client, err := database.ConnectToMongoDB(uri)
	if err != nil {
		log.Fatal(err)
	}
		
		


	// Disconnect from the MongoDB database when the application closes
	defer database.DisconnectFromMongoDB(client)


	// Set up the router and start the application
	r := routers.SetupRouter(client)
	r.Run(":8080")
}
