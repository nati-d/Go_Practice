package routers

import (
	"task_manager_refactored/Delivery/controllers"
	repository "task_manager_refactored/Repository"
	usecase "task_manager_refactored/Usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPublicUserRouter( client *mongo.Client ,db string, group *gin.RouterGroup) {
	
	userRepository := repository.NewUserRepository(client, db, "users")
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)


	group.POST("/register", userController.RegisterUser)
	group.POST("/login", userController.Login)
	group.GET("/users", userController.GetAllUsers)
}