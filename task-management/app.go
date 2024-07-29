package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Task struct {
	ID          string    
	Title       string    
	Description string    
	DueDate     time.Time 
	Status      string    
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {

	//Start the server
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	defer router.Run()


	//Get All Tasks

	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})


	//Get Specific Task

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, task := range tasks {
			if task.ID == id {
				ctx.JSON(http.StatusOK, task)
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})


	//Update Task

	router.PUT("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
	
		var updatedTask Task
	
		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		for i, task := range tasks {
			if task.ID == id {
				// Update only the specified fields
				if updatedTask.Title != "" {
					tasks[i].Title = updatedTask.Title
				}
				if updatedTask.Description != "" {
					tasks[i].Description = updatedTask.Description
				}
				ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
				return
			}
		}
	
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	//Delete Task

	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
	
		for i, val := range tasks {
			if val.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
				return
			}
		}
	
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	//Add New Tasks

	router.POST("/tasks", func(ctx *gin.Context) {
		var newTask Task
	
		if err := ctx.ShouldBindJSON(&newTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		tasks = append(tasks, newTask)
		ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
	})
	fmt.Println(tasks)
}
