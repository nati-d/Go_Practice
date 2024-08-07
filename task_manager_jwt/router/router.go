package router

import (
	"task_manager_jwt/controllers"
	"task_manager_jwt/data"
	"task_manager_jwt/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter initializes the Gin router with the given MongoDB client.
func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Initialize the user service.
	userService := data.NewUserService(client, "taskdb", "users")
	userController := controllers.NewUserController(userService)

	// Initialize the task service.
	taskService := data.NewTaskService(client, "taskdb", "tasks")
	taskController := controllers.NewTaskController(taskService, userService)

	// Register the user routes.
	// Route for user registration
	r.POST("/register", userController.RegisterUser)
	// Route for user login
	r.POST("/login", userController.Login)
	// Route to get a list of all users (admin only)
	r.GET("/users", userController.GetAllUsers)
	// Route to get a list of all tasks (public access)
	r.GET("/alltasks", taskController.GetAllTasks)

	// Register the task routes under authentication middleware.
	// Group routes that require authentication
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Route to update a user's details (requires admin role)
		authorized.PATCH("/users/:id", userController.UpdateUser)
		// Route to delete a user (requires admin role)
		authorized.DELETE("/users/:id", userController.DeleteUser)
		// Route to add a new task (requires authentication)
		authorized.POST("/tasks", taskController.AddTask)
		// Route to get tasks created by the logged-in user
		authorized.GET("/tasks", taskController.GetMyTasks)
		// Route to get a specific task by ID (requires authentication)
		authorized.GET("/tasks/:id", taskController.GetTaskById)
		// Route to update a task fully by ID (requires authentication)
		authorized.PATCH("/tasks/:id", taskController.UpdateFullTask)
		// Route to update a task partially by ID (requires authentication)
		authorized.PUT("/tasks/:id", taskController.UpdateSomeTask)
		// Route to delete a task by ID (requires authentication)
		authorized.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	return r
}
