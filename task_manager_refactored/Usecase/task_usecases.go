package usecase

import (
	"fmt"
	"task_manager_refactored/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//apply use cases for all epositories
type TaskUsecase struct {
	TaskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepository: taskRepository}
}

func (tu *TaskUsecase) AddTask(task domain.Task) error {
	if task.Title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if task.Description == "" {
		return fmt.Errorf("task description cannot be empty")
	}
	if task.Status == "" {
		return fmt.Errorf("task status cannot be empty")
	}

	if task.Status != "Not Started" && task.Status != "In Progress" && task.Status != "Completed" {
		return fmt.Errorf("task status must be one of 'Not Started', 'In Progress', or 'Completed'")
	}

	return tu.TaskRepository.AddTask(task)
}

func (tu *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	return tu.TaskRepository.GetAllTasks()
}

func (tu *TaskUsecase) GetMyTasks(userId string) ([]domain.Task, error) {
	newUserId,_ := primitive.ObjectIDFromHex(userId)
	return tu.TaskRepository.GetMyTasks(newUserId)
}

func (tu *TaskUsecase) GetTaskById(id primitive.ObjectID) (domain.Task, error) {
	return tu.TaskRepository.GetTaskById(id)
}

func (tu *TaskUsecase) UpdateFullTask(id primitive.ObjectID, task domain.Task) error {
	return tu.TaskRepository.UpdateFullTask(id,task)
}

func (tu *TaskUsecase) UpdateSomeTask(id primitive.ObjectID,task map[string]interface{}) error {
	return tu.TaskRepository.UpdateSomeTask(id,task)
}

func (tu *TaskUsecase) DeleteTask(id primitive.ObjectID) error {
	return tu.TaskRepository.DeleteTask(id)
}