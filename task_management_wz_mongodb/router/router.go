package router

import (
	"task_management_wz_mongodb/controllers"
	"task_management_wz_mongodb/data"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Initialize the TaskService
	service := data.NewTaskService(client.Database("taskdb").Collection("tasks"))

	// Initialize the TaskController with the TaskService
	taskController := controllers.NewTaskController(service)

	// Define the routes for task management

	// Route for adding a new task
	r.POST("/tasks", taskController.AddTask)

	// Route for getting all tasks
	r.GET("/tasks", taskController.GetAllTasks)

	// Route for getting a task by ID
	r.GET("/tasks/:id", taskController.GetTaskById)

	// Route for updating a task fully by ID
	r.PUT("/tasks/:id", taskController.UpdateFullTask)

	// Route for updating a task partially by ID
	r.PATCH("/tasks/:id", taskController.UpdateSomeTask)

	// Route for deleting a task by ID
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	// Return the configured router
	return r
}