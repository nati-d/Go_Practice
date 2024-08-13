package controllers

import (
	"net/http"
	"task_manager_testing/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskController handles HTTP requests for task operations.
// It interacts with the TaskUsecase and UserUsecase to manage tasks.
type TaskController struct {
	TaskUsecase domain.TaskUsecase // Usecase layer for task-related operations.
	UserUsecase domain.UserUsecase // Usecase layer for user-related operations.
}

// NewTaskController creates a new instance of TaskController.
// It initializes the controller with the given TaskUsecase and UserUsecase.
func NewTaskController(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *TaskController {
	return &TaskController{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}
}

// AddTask handles the creation of a new task.
// It binds the JSON input to a Task object, validates the user,
// and then delegates the task creation to the TaskUsecase.
func (tc *TaskController) AddTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the task data is correct: " + err.Error()})
		return
	}

	// Set the task ID to a new ObjectID
	task.ID = primitive.NewObjectID()

	// Retrieve user claims from the context (set by middleware)
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to add a task."})
		return
	}

	// Extract user ID from claims and assign to CreatedBy field
	userClaims := claims.(*domain.Claims)
	createdByID, err := primitive.ObjectIDFromHex(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID. Please try again."})
		return
	}
	task.CreatedBy = createdByID

	// Delegate task creation to the TaskUsecase
	if err := tc.TaskUsecase.AddTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add task. Please ensure all required fields are filled: " + err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully!", "task": task})
}

// GetMyTasks retrieves tasks created by the logged-in user.
// It validates the user and fetches their tasks via the TaskUsecase.
func (tc *TaskController) GetMyTasks(c *gin.Context) {
	// Retrieve user claims from the context
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to view your tasks."})
		return
	}

	// Extract user ID from claims
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims. Please try again."})
		return
	}

	userId, _ := primitive.ObjectIDFromHex(userClaims.UserID)

	// Fetch tasks created by the user from the TaskUsecase
	tasks, err := tc.TaskUsecase.GetMyTasks(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve your tasks. Please try again later: " + err.Error()})
		return
	}

	// Respond with retrieved tasks
	c.JSON(http.StatusOK, gin.H{"message": "Your tasks retrieved successfully!", "tasks": tasks})
}

// GetAllTasks handles the retrieval of all tasks.
// It delegates the task retrieval to the TaskUsecase.
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	// Fetch all tasks from the TaskUsecase
	tasks, err := tc.TaskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve tasks. Please try again later: " + err.Error()})
		return
	}

	// Respond with all tasks
	c.JSON(http.StatusOK, gin.H{"message": "All tasks retrieved successfully!", "tasks": tasks})
}

// GetTaskById handles the retrieval of a task by its ID.
// It validates the task ID and fetches the task via the TaskUsecase.
func (tc *TaskController) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the task ID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	// Fetch the task by ID from the TaskUsecase
	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}

	// Respond with the retrieved task
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully!", "task": task})
}

// UpdateFullTask handles full updates to a task by its ID.
// It performs authorization checks and updates the task via the TaskUsecase.
func (tc *TaskController) UpdateFullTask(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the task ID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	// Retrieve user claims from the context
	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	// Bind the incoming JSON to a Task object
	var task domain.Task
	task.ID = id
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the task data is correct: " + err.Error()})
		return
	}

	// Fetch the user who created the task
	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy)

	// Authorization checks based on user roles
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

	// Delegate the full task update to the TaskUsecase
	if err := tc.TaskUsecase.UpdateFullTask(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task. Please ensure all required fields are filled: " + err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// UpdateSomeTask handles partial updates to a task by its ID.
// It validates the task ID, performs authorization checks,
// and delegates the partial update to the TaskUsecase.
func (tc *TaskController) UpdateSomeTask(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the task ID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	// Retrieve user claims from the context
	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	// Fetch the task by ID from the TaskUsecase
	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}

	// Fetch the user who created the task
	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy)

	// Authorization checks based on user roles
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

	// Bind the incoming JSON to a map for partial update
	var updatedFields bson.M
	if err := c.BindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the task data is correct: " + err.Error()})
		return
	}

	// Delegate the partial update to the TaskUsecase
	if err := tc.TaskUsecase.UpdateSomeTask(id, updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task. Please ensure all required fields are filled: " + err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// DeleteTask handles the deletion of a task by its ID.
// It validates the task ID, performs authorization checks,
// and delegates the deletion to the TaskUsecase.
func (tc *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the task ID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid task ID."})
		return
	}

	// Retrieve user claims from the context
	claims, _ := c.Get("user")
	userClaims := claims.(*domain.Claims)

	// Fetch the task by ID from the TaskUsecase
	task, err := tc.TaskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found. Please ensure the task ID is correct: " + err.Error()})
		return
	}

	// Fetch the user who created the task
	otherUser, _ := tc.UserUsecase.GetUserById(task.CreatedBy)

	// Authorization checks based on user roles
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

	// Delegate the task deletion to the TaskUsecase
	if err := tc.TaskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete task. Please try again: " + err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
