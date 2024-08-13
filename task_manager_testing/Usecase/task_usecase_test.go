package usecase_test

import (
	"testing"

	usecase "task_manager_testing/Usecase"
	"task_manager_testing/domain"
	"task_manager_testing/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecaseSuite defines the suite for task usecase tests
type TaskUsecaseSuite struct {
	suite.Suite
	taskRepo   *mocks.TaskRepository
	taskUsecase *usecase.TaskUsecase
}
// SetupTest sets up the necessary resources before each test
func (suite *TaskUsecaseSuite) SetupTest() {
	suite.taskRepo = &mocks.TaskRepository{}
	suite.taskUsecase = usecase.NewTaskUsecase(suite.taskRepo)
}

// TearDownTest clears resources after each test
func (suite *TaskUsecaseSuite) TearDownTest() {
	// Reset the mock expectations
	suite.taskRepo.AssertExpectations(suite.T())
}

// TestAddTask tests the AddTask use case
func (suite *TaskUsecaseSuite) TestAddTask() {
	task := domain.Task{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1", Status: "Completed", CreatedBy: primitive.NewObjectID()}

	suite.taskRepo.On("AddTask", task).Return(nil)

	err := suite.taskUsecase.AddTask(task)

	assert.Nil(suite.T(), err)
}

// TestDeleteTask tests the DeleteTask use case
func (suite *TaskUsecaseSuite) TestDeleteTask() {
	id := primitive.NewObjectID()

	suite.taskRepo.On("DeleteTask", id).Return(nil)

	err := suite.taskUsecase.DeleteTask(id)

	assert.Nil(suite.T(), err)
}

// TestGetAllTasks tests the GetAllTasks use case
func (suite *TaskUsecaseSuite) TestGetAllTasks() {
	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1", Status: "Completed", CreatedBy: primitive.NewObjectID()},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2", Status: "Pending", CreatedBy: primitive.NewObjectID()},
	}

	suite.taskRepo.On("GetAllTasks").Return(tasks, nil)

	result, err := suite.taskUsecase.GetAllTasks()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tasks, result)
}

// TestGetTaskById tests the GetTaskById use case
func (suite *TaskUsecaseSuite) TestGetTaskById() {
	id := primitive.NewObjectID()
	task := domain.Task{ID: id, Title: "Task 1", Description: "Description 1", Status: "Completed", CreatedBy: primitive.NewObjectID()}

	suite.taskRepo.On("GetTaskById", id).Return(task, nil)

	result, err := suite.taskUsecase.GetTaskById(id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), task, result)
}

// TestUpdateFullTask tests the UpdateFullTask use case
func (suite *TaskUsecaseSuite) TestUpdateFullTask() {
	id := primitive.NewObjectID()
	task := domain.Task{ID: id, Title: "Task 1", Description: "Description 1", Status: "Completed", CreatedBy: primitive.NewObjectID()}

	suite.taskRepo.On("UpdateFullTask", id, task).Return(nil)

	err := suite.taskUsecase.UpdateFullTask(id, task)

	assert.Nil(suite.T(), err)
}

// TestUpdateSomeTask tests the UpdateSomeTask use case
func (suite *TaskUsecaseSuite) TestUpdateSomeTask() {
	id := primitive.NewObjectID()
	task := map[string]interface{}{"status": "Completed"}

	suite.taskRepo.On("UpdateSomeTask", id, task).Return(nil)

	err := suite.taskUsecase.UpdateSomeTask(id, task)

	assert.Nil(suite.T(), err)
}

// TestGetMyTasks tests the GetMyTasks use case
func (suite *TaskUsecaseSuite) TestGetMyTasks() {
	userId := primitive.NewObjectID()
	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1", Status: "Completed", CreatedBy: userId},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2", Status: "Pending", CreatedBy: userId},
	}

	suite.taskRepo.On("GetMyTasks", userId).Return(tasks, nil)

	result, err := suite.taskUsecase.GetMyTasks(userId)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tasks, result)
}

// TestTaskUsecaseSuite is the entry point for running the suite tests
func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}