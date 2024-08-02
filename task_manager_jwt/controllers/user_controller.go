package controllers

import (
	"net/http"
	"task_manager_jwt/data"
	"task_manager_jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userController struct {
	service *data.UserService
}

func NewUserController(service *data.UserService) *userController {
	return &userController{service: service}
}


// Register is a controller function to register a new user.
func (uc *userController) Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	user.ID = primitive.NewObjectID()
	if err := uc.service.Register(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})

}


// Login is a controller function to login a user.
func (uc *userController) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user": user})
}

//delete all usera
func (uc *userController) DeleteAll(c *gin.Context) {
	err := uc.service.DeleteAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users deleted successfully"})
}