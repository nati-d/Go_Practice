package controllers

import (
	"net/http"
	"task_manager_jwt/data"
	"task_manager_jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service *data.UserService
}

func NewUserController(service *data.UserService) *UserController {
	return &UserController{service: service}
}

// Register is a controller function to register a new user.
func (uc *UserController) Register(c *gin.Context) {
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
func (uc *UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	token, err := uc.service.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

// DeleteAll deletes all users.
func (uc *UserController) DeleteAll(c *gin.Context) {
	err := uc.service.DeleteAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users deleted successfully"})
}
