package routers

import (
	"task_manager_refactored/Delivery/controllers"
	repository "task_manager_refactored/Repository"
	usecase "task_manager_refactored/Usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func NewProtectedUserRouter(client *mongo.Client , db string, group *gin.RouterGroup) {
	
	userRepository := repository.NewUserRepository(client, db, "users")
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)


	// Route to update a user's details (requires admin role)
	group.PATCH("/users/:id", userController.UpdateUser)
	// Route to delete a user (requires admin role)
	group.DELETE("/users/:id", userController.DeleteUser)


	
	
	
}