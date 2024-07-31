package data

import (
	"errors"
	"fmt"
	"task_management/models"
)

type TaskService struct {
	tasks  map[int]models.Task
	NextID int
}

// NewTaskService creates a new instance of TaskService
func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]models.Task),
		NextID: 1,
	}
}

// AddTask adds a new task to the task list
// Returns an error if the task already exists or if required fields are missing
func (ts *TaskService) AddTask(task models.Task) error {
	_, exists := ts.tasks[task.ID]
	if exists {
		return errors.New("oops! a task with this id already exists. please try a different id.")
	}

	if task.Status != "Not Started" && task.Status != "In Progress" && task.Status != "Completed" {
		return errors.New("invalid status! please use 'not started', 'in progress', or 'completed'.")
	}

	if task.Title == "" {
		return errors.New("the task title is required. please provide a title.")
	}

	if task.Description == "" {
		return errors.New("the task description is required. please provide a description.")
	}

	ts.tasks[task.ID] = task
	return nil
}

// GetAllTasks returns all the tasks
func (ts *TaskService) GetAllTasks() []models.Task {
	var tasks []models.Task
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskById returns a task with the given id
// Returns an error if the task does not exist
func (ts *TaskService) GetTaskById(id int) (models.Task, error) {
	task, exists := ts.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("oh no! task with id %d does not exist. please check the id and try again.", id)
	}

	return task, nil
}

// UpdateSomeTask updates the task with the given id
// Allows partial updates and returns an error if the task does not exist or if the status is invalid
func (ts *TaskService) UpdateSomeTask(id int, updatedTask models.Task) error {
	task, exists := ts.tasks[id]
	if !exists {
		return errors.New("oh no! task with this id does not exist. please check the id and try again.")
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}
	if updatedTask.Status != "" {
		if updatedTask.Status != "Not Started" && updatedTask.Status != "In Progress" && updatedTask.Status != "Completed" {
			return errors.New("invalid status! please use 'not started', 'in progress', or 'completed'.")
		}
		task.Status = updatedTask.Status
	}

	ts.tasks[id] = task
	return nil
}

// UpdateFullTask updates the task with the given id with all fields
// Returns an error if the task does not exist or if required fields are missing
func (ts *TaskService) UpdateFullTask(id int, updatedTask models.Task) error {
	_, exists := ts.tasks[id]
	if !exists {
		return errors.New("oh no! task with this id does not exist. please check the id and try again.")
	}

	if updatedTask.Title == "" {
		return errors.New("the task title is required. please provide a title.")
	}
	if updatedTask.Description == "" {
		return errors.New("the task description is required. please provide a description.")
	}

	if updatedTask.Status == "" {
		return errors.New("the task status is required. please provide a status.")
	} else if updatedTask.Status != "Not Started" && updatedTask.Status != "In Progress" && updatedTask.Status != "Completed" {
		return errors.New("invalid status! please use 'not started', 'in progress', or 'completed'.")
	}

	ts.tasks[id] = updatedTask
	return nil
}

// DeleteTask deletes the task with the given id
// Returns an error if the task does not exist
func (ts *TaskService) DeleteTask(id int) error {
	_, exists := ts.tasks[id]
	if !exists {
		return errors.New("oh no! task with this id does not exist. please check the id and try again.")
	}

	delete(ts.tasks, id)
	return nil
}
