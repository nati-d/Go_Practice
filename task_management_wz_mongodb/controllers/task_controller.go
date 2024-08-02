package controllers

import (
	"net/http"
	"task_management_wz_mongodb/data"
	"task_management_wz_mongodb/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskController handles HTTP requests for task operations
type TaskController struct {
	service *data.TaskService
}

// NewTaskController creates a new TaskController
func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{service: service}
}

// AddTask handles the creation of a new task
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	// Set the task ID
	task.ID = primitive.NewObjectID()

	if err := tc.service.AddTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add task: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully", "task": task})
}

// GetAllTasks handles the retrieval of all tasks
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve tasks: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tasks retrieved successfully", "tasks": tasks})
}

// GetTaskById handles the retrieval of a task by ID
func (tc *TaskController) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	task, err := tc.service.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve task: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully", "task": task})
}

// UpdateFullTask handles full updates to a task by ID
func (tc *TaskController) UpdateFullTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var task models.Task

	task.ID = id
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}

	if err := tc.service.UpdateFullTask(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// UpdateSomeTask handles partial updates to a task by ID
func (tc *TaskController) UpdateSomeTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var update bson.M
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		return
	}
	if err := tc.service.UpdateSomeTask(id, update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask handles the deletion of a task by ID
func (tc *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := tc.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete task: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
