package routers

import (
	"task_manager_testing/Delivery/controllers"
	repository "task_manager_testing/Repository"
	usecase "task_manager_testing/Usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPublicTaskRouter( client *mongo.Client ,db string, group *gin.RouterGroup) {
	
	taskRepository := repository.NewTaskRepository(client, db, "tasks")
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userRepository := repository.NewUserRepository(client, db, "users")
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskController := controllers.NewTaskController(taskUsecase, userUsecase)


	group.GET("/alltasks", taskController.GetAllTasks)
}