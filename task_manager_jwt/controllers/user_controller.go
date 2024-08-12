package controllers

import (
	"net/http"
	"task_manager_jwt/data"
	"task_manager_jwt/middleware"
	"task_manager_jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service *data.UserService
}

// NewUserController creates a new user controller.
func NewUserController(service *data.UserService) *UserController {
	return &UserController{service: service}
}

// RegisterUser registers a new user.
func (uc *UserController) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide a valid username, password, and role"})
		return
	}

	err := uc.service.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed. Please try again later"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
}

// Login authenticates a user and generates a JWT token.
func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide a valid username and password"})
		return
	}

	user, err := uc.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. Please check your username and password"})
		return
	}

	token, err := middleware.GenerateJWT(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token. Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

// GetAllUsers returns all users.
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users. Please try again later"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users retrieved successfully", "data": users})
}

// GetUserById returns a user by ID.
func (uc *UserController) GetUserById(c *gin.Context) {
	paramId := c.Param("id")
	user, err := uc.service.GetUserById(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User retrieved successfully", "data": user})
}

// UpdateUser updates a user's profile.
func (uc *UserController) UpdateUser(c *gin.Context) {
	paramId := c.Param("id")
	claims, _ := c.Get("user")
	userClaims := claims.(*models.Claims)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Please provide valid user data"})
		return
	}

	otherUser, err := uc.service.GetUserById(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found"})
		return
	}

	user.ID, _ = primitive.ObjectIDFromHex(userClaims.UserID)

	// Check if user is allowed to edit the profile
	if userClaims.Role == "user" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own profile"})
		return
	}
	// Admin cannot edit other admins or root
	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot edit other admins or root"})
		return
	}
	// Root user cannot edit other root user
	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root user cannot edit other root user"})
		return
	}
	// Prevent password changes if user is not allowed
	if userClaims.UserID != paramId {
		user.Password = ""
	}

	if err := uc.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user. Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully", "data": user})
}

// DeleteUser deletes a user.
func (uc *UserController) DeleteUser(c *gin.Context) {
	paramId := c.Param("id")
	claims, _ := c.Get("user")
	userClaims := claims.(*models.Claims)

	// Retrieve the user to be deleted
	otherUser, err := uc.service.GetUserById(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found"})
		return
	}

	// Check if user is allowed to delete the profile
	if userClaims.Role == "user" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own profile"})
		return
	}
	// Admin cannot delete other admins or root
	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot delete other admins or root"})
		return
	}
	// Root user cannot delete other root user
	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root user cannot delete other root user"})
		return
	}

	// Perform the delete operation
	if err := uc.service.DeleteUser(paramId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user. Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User with ID " + paramId + " has been successfully deleted"})
}


func(uc *UserController) PromoteUser(c *gin.Context) {
	paramId := c.Param("id")
	claims, _ := c.Get("user")
	userClaims := claims.(*models.Claims)

	// Retrieve the user to be promoted
	otherUser, err := uc.service.GetUserById(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with the given ID not found"})
		return
	}

	// Check if user is allowed to promote the profile
	if userClaims.Role == "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to promote users"})
		return
	}
	// Admin cannot promote other admins or root
	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot promote other admins or root"})
		return
	}
	// Root user cannot promote other root user
	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != paramId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root user cannot promote other root user"})
		return
	}
	changedParamId, _ := primitive.ObjectIDFromHex(paramId)
	// Perform the promote operation
	if err := uc.service.PromoteUser(changedParamId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to promote user. Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User with ID " + paramId + " has been successfully promoted"})
}
