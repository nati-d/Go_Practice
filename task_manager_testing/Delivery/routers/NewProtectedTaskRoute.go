package routers

import (
	"task_manager_testing/Delivery/controllers"
	repository "task_manager_testing/Repository"
	usecase "task_manager_testing/Usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func NewProtectedTaskRouter(client *mongo.Client , db string, group *gin.RouterGroup) {
	
	taskRepository := repository.NewTaskRepository(client, db, "tasks")
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userRepository := repository.NewUserRepository(client, db, "users")
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskController := controllers.NewTaskController(taskUsecase, userUsecase)

	group.POST("/tasks", taskController.AddTask)
	// Route to get tasks created by the logged-in user
	group.GET("/tasks", taskController.GetMyTasks)
	// Route to get a specific task by ID (requires authentication)
	group.GET("/tasks/:id", taskController.GetTaskById)
	// Route to update a task fully by ID (requires authentication)
	group.PATCH("/tasks/:id", taskController.UpdateFullTask)
	// Route to update a task partially by ID (requires authentication)
	group.PUT("/tasks/:id", taskController.UpdateSomeTask)
	// Route to delete a task by ID (requires authentication)
	group.DELETE("/tasks/:id", taskController.DeleteTask)
	
}