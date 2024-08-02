package router

import (
	"task_manager_jwt/controllers"
	"task_manager_jwt/data"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Initialize the TaskService
	taskService := data.NewTaskService(client , "taskdb", "tasks")
	// Initialize the TaskController with the TaskService
	taskController := controllers.NewTaskController(taskService)

	//Initialize User Service
	userService := data.NewUserService(client, "taskdb", "users")

	//Initialize User Controller
	userController := controllers.NewUserController(userService)

		

	// Define the routes for task management

	r.POST("/register", userController.Register)

	// r.POST("/login", userController.Login)

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