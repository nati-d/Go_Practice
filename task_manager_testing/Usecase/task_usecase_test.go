package usecase_test

import (
	usecase "task_manager_testing/Usecase"
	"task_manager_testing/domain"
	"task_manager_testing/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
    suite.Suite
    taskUsecase *usecase.TaskUsecase
    mockRepo    *mocks.TaskRepository
}

func (suite *TaskUsecaseSuite) SetupTest() {
    suite.mockRepo = new(mocks.TaskRepository)
    suite.taskUsecase = usecase.NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseSuite) TestAddTask() {
    task := domain.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     primitive.NewDateTimeFromTime(time.Now()),
        Status:      "Completed",
        CreatedBy:   primitive.NewObjectID(),
    }

    suite.mockRepo.On("AddTask", task).Return(nil)

    err := suite.taskUsecase.AddTask(task)

    assert.NoError(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTaskById() {
    task := domain.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     primitive.NewDateTimeFromTime(time.Now()),
        Status:      "Completed",
        CreatedBy:   primitive.NewObjectID(),
    }

    suite.mockRepo.On("GetTaskById", task.ID).Return(task, nil)

    result, err := suite.taskUsecase.GetTaskById(task.ID)

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), task, result)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetAllTasks() {
    tasks := []domain.Task{
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 1",
            Description: "This is task 1",
            DueDate:     primitive.NewDateTimeFromTime(time.Now()),
            Status:      "Completed",
            CreatedBy:   primitive.NewObjectID(),
        },
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 2",
            Description: "This is task 2",
            DueDate:     primitive.NewDateTimeFromTime(time.Now()),
            Status:      "Completed",
            CreatedBy:   primitive.NewObjectID(),
        },
    }

    suite.mockRepo.On("GetAllTasks").Return(tasks, nil)

    result, err := suite.taskUsecase.GetAllTasks()

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), tasks, result)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetMyTasks() {
    userId := primitive.NewObjectID()
    tasks := []domain.Task{
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 1",
            Description: "This is task 1",
            DueDate:     primitive.NewDateTimeFromTime(time.Now()),
            Status:      "Completed",
            CreatedBy:   userId,
        },
    }

    suite.mockRepo.On("GetMyTasks", userId).Return(tasks, nil)

    result, err := suite.taskUsecase.GetMyTasks(userId)

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), tasks, result)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateFullTask() {
    task := domain.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     primitive.NewDateTimeFromTime(time.Now()),
        Status:      "Completed",
        CreatedBy:   primitive.NewObjectID(),
    }

    suite.mockRepo.On("UpdateFullTask", task.ID, task).Return(nil)

    err := suite.taskUsecase.UpdateFullTask(task.ID, task)

    assert.NoError(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateSomeTask() {
    taskId := primitive.NewObjectID()
    update := map[string]interface{}{
        "status": "Completed",
    }

    suite.mockRepo.On("UpdateSomeTask", taskId, update).Return(nil)

    err := suite.taskUsecase.UpdateSomeTask(taskId, update)

    assert.NoError(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDeleteTask() {
    taskId := primitive.NewObjectID()

    suite.mockRepo.On("DeleteTask", taskId).Return(nil)

    err := suite.taskUsecase.DeleteTask(taskId)

    assert.NoError(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
    suite.Run(t, new(TaskUsecaseSuite))
}
