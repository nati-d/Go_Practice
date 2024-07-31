package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task_management/data"
	"task_management/models"
)

// TaskController handles HTTP requests for task operations
type TaskController struct {
	service *data.TaskService
}

// NewTaskController creates a new TaskController
func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

// AddTask handles the creation of a new task
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Task

	// Bind JSON data to the task model
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Set the task ID and increment the NextID
	task.ID = tc.service.NextID
	tc.service.NextID++

	// Add the task to the service
	err = tc.service.AddTask(task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, task)
}

// GetAllTasks handles the retrieval of all tasks
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks := tc.service.GetAllTasks()
	c.JSON(200, gin.H{"tasks": tasks})
}

// GetTaskById handles the retrieval of a task by ID
func (tc *TaskController) GetTaskById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := tc.service.GetTaskById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

// UpdateSomeTask handles partial updates to a task by ID
func (tc *TaskController) UpdateSomeTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	// Bind JSON data to the task model
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the task partially
	err = tc.service.UpdateSomeTask(id, task)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task updated"})
}

// UpdateFullTask handles full updates to a task by ID
func (tc *TaskController) UpdateFullTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	// Bind JSON data to the task model
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the task fully
	err = tc.service.UpdateFullTask(id, task)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task updated"})
}

// DeleteTask handles the deletion of a task by ID
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Delete the task
	err := tc.service.DeleteTask(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted"})
}
