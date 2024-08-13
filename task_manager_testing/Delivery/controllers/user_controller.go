package controllers

import (
	"net/http"
	infrastructure "task_manager_testing/Infrastructure"
	"task_manager_testing/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserController handles user-related requests and interactions.
type UserController struct {
	UserUsecase domain.UserUsecase
}

// NewUserController initializes a new UserController.
func NewUserController(userusecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userusecase}
}

// RegisterUser handles user registration requests.
func (uc *UserController) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	// Validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide a valid username, password, and role."})
		return
	}

	// Attempt to register the new user
	if err := uc.UserUsecase.RegisterUser(req.Username, req.Password, req.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed. Please try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered."})
}

// Login handles user authentication and token generation.
func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide a valid username and password."})
		return
	}

	// Attempt to authenticate the user
	user, err := uc.UserUsecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. Please check your username and password."})
		return
	}

	// Generate JWT token
	token, err := infrastructure.GenerateJWT(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}

// GetAllUsers retrieves all registered users.
func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Retrieve users from the use case layer
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Users retrieved successfully.", "data": users})
}

// GetUserById retrieves a user by their ID.
func (uc *UserController) GetUserById(c *gin.Context) {
	paramId := c.Param("id")
	newParamId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format."})
		return
	}

	// Retrieve the user by ID from the use case layer
	user, err := uc.UserUsecase.GetUserById(newParamId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User retrieved successfully.", "data": user})
}

// UpdateUser updates the profile of an existing user.
func (uc *UserController) UpdateUser(c *gin.Context) {
	paramId := c.Param("id")
	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)
	var user domain.User

	// Validate incoming JSON request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide valid user data."})
		return
	}

	newParamId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format."})
		return
	}

	// Retrieve the user to be updated
	otherUser, err := uc.UserUsecase.GetUserById(newParamId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found."})
		return
	}

	user.ID, _ = primitive.ObjectIDFromHex(userClaims.UserID)

	// Role-based access control checks
	if userClaims.Role == "user" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own profile."})
		return
	}
	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot edit other admins or root users."})
		return
	}
	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root users cannot edit other root users."})
		return
	}

	// Prevent unauthorized password changes
	if userClaims.UserID != paramId {
		user.Password = ""
	}

	// Attempt to update the user's profile
	if err := uc.UserUsecase.UpdateUser(newParamId, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully.", "data": user})
}

// DeleteUser deletes a user from the system.
func (uc *UserController) DeleteUser(c *gin.Context) {
	paramId := c.Param("id")
	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)
	newParamId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format."})
		return
	}

	// Retrieve the user to be deleted
	otherUser, err := uc.UserUsecase.GetUserById(newParamId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found."})
		return
	}

	// Role-based access control checks
	if userClaims.Role == "user" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own profile."})
		return
	}
	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot delete other admins or root users."})
		return
	}
	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root users cannot delete other root users."})
		return
	}

	// Attempt to delete the user
	if err := uc.UserUsecase.DeleteUser(newParamId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User with ID " + paramId + " has been successfully deleted."})
}
