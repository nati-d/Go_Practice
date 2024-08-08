package routers

import (
	"task_manager_refactored/Delivery/controllers"
	infrastructure "task_manager_refactored/Infrastructure"
	repository "task_manager_refactored/Repository"
	usecase "task_manager_refactored/Usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	taskRepository := repository.NewTaskRepository(client, "taskdb", "tasks")
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userRepository := repository.NewUserRepository(client, "taskdb", "users")
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskController := controllers.NewTaskController(taskUsecase, userUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := gin.Default()

	publicRouter := r.Group("/")

	publicRouter.POST("/register", userController.RegisterUser)
	publicRouter.POST("/login", userController.Login)
	publicRouter.GET("/alltasks", taskController.GetAllTasks)
	publicRouter.GET("/users", userController.GetAllUsers)

	protectedRoute := r.Group("/")
	protectedRoute.Use(infrastructure.AuthMiddleware())

	// Route to update a user's details (requires admin role)
	protectedRoute.PATCH("/users/:id", userController.UpdateUser)
	// Route to delete a user (requires admin role)
	protectedRoute.DELETE("/users/:id", userController.DeleteUser)
	// Route to add a new task (requires authentication)
	protectedRoute.POST("/tasks", taskController.AddTask)
	// Route to get tasks created by the logged-in user
	protectedRoute.GET("/tasks", taskController.GetMyTasks)
	// Route to get a specific task by ID (requires authentication)
	protectedRoute.GET("/tasks/:id", taskController.GetTaskById)
	// Route to update a task fully by ID (requires authentication)
	protectedRoute.PATCH("/tasks/:id", taskController.UpdateFullTask)
	// Route to update a task partially by ID (requires authentication)
	protectedRoute.PUT("/tasks/:id", taskController.UpdateSomeTask)
	// Route to delete a task by ID (requires authentication)
	protectedRoute.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}


	


