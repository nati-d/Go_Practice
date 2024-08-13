package repository_test

import (
	"task_manager_testing/config/database"
	repository "task_manager_testing/Repository"
	"task_manager_testing/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositorySuite struct {
	suite.Suite
	repository repository.TaskRepository
	collection *mongo.Collection
	client     *mongo.Client
}

func (suite *TaskRepositorySuite) SetupTest() {
	client, _ := database.ConnectToMongoDB("mongodb+srv://nathnaeldes:12345678n@cluster0.w8bpdtf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	suite.client = client
	db := "taskdb"
	repository := repository.NewTaskRepository(client, db, "taskstest")
	suite.repository = *repository

}

func (suite *TaskRepositorySuite) TestAddTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Pending",
		CreatedBy:   primitive.NewObjectID(),
	}

	err := suite.repository.AddTask(task)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TestGetAllTasks() {
	tasks, err := suite.repository.GetAllTasks()
	suite.NoError(err)
	suite.Empty(tasks)
}

func (suite *TaskRepositorySuite) TestGetTaskById() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Pending",
		CreatedBy:   primitive.NewObjectID(),
	}

	err := suite.repository.AddTask(task)
	suite.NoError(err)

	task, err = suite.repository.GetTaskById(task.ID)
	suite.NoError(err)
	suite.NotEmpty(task)
}

func (suite *TaskRepositorySuite) TestGetMyTasks() {
	tasks, err := suite.repository.GetMyTasks(primitive.NewObjectID())
	suite.NoError(err)
	suite.Empty(tasks)
}

func (suite *TaskRepositorySuite) TestUpdateFullTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Pending",
		CreatedBy:   primitive.NewObjectID(),
	}

	err := suite.repository.AddTask(task)
	suite.NoError(err)

	task.Status = "Completed"
	err = suite.repository.UpdateFullTask(task.ID, task)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TestUpdateSomeTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Pending",
		CreatedBy:   primitive.NewObjectID(),
	}

	err := suite.repository.AddTask(task)
	suite.NoError(err)

	update := map[string]interface{}{
		"status": "Completed",
	}

	err = suite.repository.UpdateSomeTask(task.ID, update)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Pending",
		CreatedBy:   primitive.NewObjectID(),
	}

	err := suite.repository.AddTask(task)
	suite.NoError(err)

	err = suite.repository.DeleteTask(task.ID)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TearDownTest() {
	suite.client.Database("taskdb").Collection("taskstest").Drop(nil)
}

func (suite *TaskRepositorySuite) TearDownSuite() {
	suite.client.Disconnect(nil)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
