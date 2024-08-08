package controllers

import (
	"net/http"
	usecase "task_manager_refactored/Usecase"
	"task_manager_refactored/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskController handles HTTP requests for task operations
type TaskController struct {
	TaskUsecase *usecase.TaskUsecase
	UserUsecase *usecase.UserUsecase
}

// NewTaskController creates a new TaskController
func NewTaskController(taskUsecase *usecase.TaskUsecase, userUsecase *usecase.UserUsecase) *TaskController {
	return &TaskController{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}
}

// AddTask handles the creation of a new task
func (tc *TaskController) AddTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the task data is correct: " + err.Error()})
		return
	}

	// Set the task ID
	task.ID = primitive.NewObjectID()
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to add a task."})
		return
	}

	userClaims := claims.(*domain.Claims)
	createdByID, err := primitive.ObjectIDFromHex(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID. Please try again."})
		return
	}
	task.CreatedBy = createdByID

	if err := tc.TaskUsecase.AddTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add task. Please ensure all required fields are filled: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully!", "task": task})
}

// GetMyTasks retrieves tasks created by the logged-in user
func (tc *TaskController) GetMyTasks(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to view your tasks."})
		return
	}

	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims. Please try again."})
		return
	}

	tasks, err := tc.TaskUsecase.GetMyTasks(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve your tasks. Please try again later: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Your tasks retrieved successfully!", "tasks": tasks})
}

// GetAllTasks handles the retrieval of all tasks
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve tasks. Please try again later: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All tasks retrieved successfully!", "tasks": tasks})
}

// GetTaskById handles the retrieval of a task by ID
func (tc *TaskController) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}
	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully!", "task": task})
}

// UpdateFullTask handles full updates to a task by ID
func (tc *TaskController) UpdateFullTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	var task domain.Task
	task.ID = id
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the task data is correct: " + err.Error()})
		return
	}

	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy.Hex())

	// Authorization checks
	if userClaims.Role == "user" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own tasks."})
		return
	}

	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot edit tasks of other admins or root users."})
		return
	}

	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root users cannot edit tasks of other root users."})
		return
	}

	if err := tc.TaskUsecase.UpdateFullTask(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task. Please ensure all required fields are filled: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// UpdateSomeTask handles partial updates to a task by ID
func (tc *TaskController) UpdateSomeTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}

	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy.Hex())

	// Authorization checks
	if userClaims.Role == "user" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own tasks."})
		return
	}

	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot edit tasks of other admins or root users."})
		return
	}

	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root users cannot edit tasks of other root users."})
		return
	}

	var update bson.M
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the update data is correct: " + err.Error()})
		return
	}
	if err := tc.TaskUsecase.UpdateSomeTask(id, update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task. Please try again: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// DeleteTask handles the deletion of a task by ID
func (tc *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}

	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy.Hex())

	// Authorization checks
	if userClaims.Role == "user" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own tasks."})
		return
	}

	if userClaims.Role == "admin" && (otherUser.Role == "admin" || otherUser.Role == "root") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins cannot delete tasks of other admins or root users."})
		return
	}

	if userClaims.Role == "root" && otherUser.Role == "root" && userClaims.UserID != task.CreatedBy.Hex() {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root users cannot delete tasks of other root users."})
		return
	}

	if err := tc.TaskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete task. Please try again: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
