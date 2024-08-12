package routers

import (
	infrastructure "task_manager_testing/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	publicRouter := r.Group("/")

	NewPublicTaskRouter(client,"taskdb",publicRouter)
	NewPublicUserRouter(client,"taskdb",publicRouter)
	
	protectedRoute := r.Group("/")
	protectedRoute.Use(infrastructure.AuthMiddleware())


	NewProtectedTaskRouter(client,"taskdb",protectedRoute)
	NewProtectedUserRouter(client,"taskdb",protectedRoute)
	

	return r
}


	


