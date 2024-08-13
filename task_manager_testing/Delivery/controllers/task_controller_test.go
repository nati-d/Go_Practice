package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_manager_testing/Delivery/controllers"
	"task_manager_testing/domain"
	"task_manager_testing/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	taskUsecase   mocks.TaskUsecase
	taskCtrl      controllers.TaskController
	testingServer *httptest.Server
	userUsecase   mocks.UserUsecase
	userCtrl      controllers.UserController
}

func (suite *TaskControllerSuite) SetupTest() {
	taskUsecase := &mocks.TaskUsecase{}
	userUsecase := &mocks.UserUsecase{}
	handler := controllers.NewTaskController(taskUsecase, userUsecase)

	router := gin.Default()
	router.POST("/task", handler.AddTask)
	router.GET("/task/:id", handler.GetMyTasks)
	router.GET("/tasks", handler.GetAllTasks)
	router.PUT("/task/:id", handler.UpdateFullTask)
	router.PATCH("/task/:id", handler.UpdateSomeTask)
	router.DELETE("/task/:id", handler.DeleteTask)

	testingServer := httptest.NewServer(router)
	suite.testingServer = testingServer
	suite.taskUsecase = *taskUsecase
	suite.taskCtrl = *handler
}

func (suite *TaskControllerSuite) TearDownTest() {
	defer suite.testingServer.Close()
}

func (suite *TaskControllerSuite) TestAddTask() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	suite.taskUsecase.On("AddTask", task).Return(nil)

	requestBody, err := json.Marshal(&task)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/task", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")

	responseBody := domain.Response{}

	json.NewDecoder(response.Body).Decode(&responseBody)

}

func (suite *TaskControllerSuite) TestGetMyTasks() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	suite.taskUsecase.On("GetMyTasks", task.CreatedBy).Return([]domain.Task{task}, nil)

	response, err := http.Get(fmt.Sprintf("%s/task/%s", suite.testingServer.URL, task.CreatedBy.Hex()))
	suite.NoError(err, "no error when calling the endpoint")

	responseBody := domain.Response{}

	json.NewDecoder(response.Body).Decode(&responseBody)
}

func (suite *TaskControllerSuite) TestGetAllTasks() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	suite.taskUsecase.On("GetAllTasks").Return([]domain.Task{task}, nil)

	response, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")

	responseBody := domain.Response{}

	json.NewDecoder(response.Body).Decode(&responseBody)
}

func (suite *TaskControllerSuite) TestUpdateFullTask() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	suite.taskUsecase.On("UpdateFullTask", task.ID, task).Return(nil)

	requestBody, err := json.Marshal(&task)
	suite.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/task/%s", suite.testingServer.URL, task.ID.Hex()), bytes.NewBuffer(requestBody))
	suite.NoError(err, "can not create PUT request")

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// You can add further assertions here if needed
}

func (suite *TaskControllerSuite) TestUpdateSomeTask() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	update := map[string]interface{}{
		"status": "In Progress",
	}

	suite.taskUsecase.On("UpdateSomeTask", task.ID, update).Return(nil)

	requestBody, err := json.Marshal(&update)
	suite.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/task/%s", suite.testingServer.URL, task.ID.Hex()), bytes.NewBuffer(requestBody))
	suite.NoError(err, "can not create PATCH request")

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// You can add further assertions here if needed
}

func (suite *TaskControllerSuite) TestDeleteTask() {
	task := domain.Task{
		ID: 		primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     primitive.NewDateTimeFromTime(time.Now()),
		Status:      "Completed",
		CreatedBy:   primitive.NewObjectID(),
	}

	suite.taskUsecase.On("DeleteTask", task.ID).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/task/%s", suite.testingServer.URL, task.ID.Hex()), nil)
	suite.NoError(err, "can not create DELETE request")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}