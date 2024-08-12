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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	userUsecase   mocks.UserUsecase
	userCtrl      controllers.UserController
	testingServer *httptest.Server
}

func (suite *UserControllerSuite) SetupTest() {
	userUsecase := &mocks.UserUsecase{}
	handler := controllers.NewUserController(userUsecase)

	router := gin.Default()
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.Login)

	testingServer := httptest.NewServer(router)
	suite.testingServer = testingServer
	suite.userUsecase = *userUsecase
	suite.userCtrl = *handler // Assigning UserController correctly
}

func (suite *UserControllerSuite) TearDownTest() {
	defer suite.testingServer.Close()
}

func (suite *UserControllerSuite) TestRegisterUser() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
		Role:     "user",
	}
	suite.userUsecase.On("RegisterUser", user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal("Success to create user", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLogin() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
	}
	suite.userUsecase.On("Login", user.Username, user.Password).Return(user, nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Success to login", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLoginInvalidRequest() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
	}

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *UserControllerSuite) TestRegisterUserInvalidRequest() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
		Role:     "user",
	}

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *UserControllerSuite) TestLoginInvalidCredentials() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
	}

	suite.userUsecase.On("Login", user.Username, user.Password).Return(domain.User{}, fmt.Errorf("Invalid credentials"))

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)

}

// get all users
func (suite *UserControllerSuite) TestGetAllTasks() {
	user := []domain.User{
		{
			Username: "tester1",
			Password: "password",
			Role:     "user",
		},
		{
			Username: "tester2",
			Password: "password",
			Role:     "user",
		},
	}
	suite.userUsecase.On("GetAllUsers").Return(user, nil)

	response, err := http.Get(fmt.Sprintf("%s/users", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := []domain.User{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Len(responseBody, 2)
	suite.userUsecase.AssertExpectations(suite.T())

}

func (suite *UserControllerSuite) UpdateUser() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
		Role:     "user",
	}
	suite.userUsecase.On("UpdateUser", user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/users", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Success to update user", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}


func (suite *UserControllerSuite) DeleteUser() {
	user := domain.User{
		Username: "tester1",
		Password: "password",
		Role:     "user",
	}
	suite.userUsecase.On("DeleteUser", user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/users", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Success to delete user", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestGetUserByID() {
	user := domain.User{
		ID:   primitive.NewObjectID(),
		Username: "tester1",
		Password: "password",
		Role:     "user",
	}
	suite.userUsecase.On("GetUserByID", user.ID).Return(user, nil)

	response, err := http.Get(fmt.Sprintf("%s/users/%s", suite.testingServer.URL, user.ID))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.User{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(user, responseBody)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestGetUserByIDInvalidID() {
	response, err := http.Get(fmt.Sprintf("%s/users/invalidID", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *UserControllerSuite) TestGetUserByIDNotFound() {
	user := domain.User{
		ID:   primitive.NewObjectID(),
		Username : "tester1",
		Password : "password",
		Role : "user",
	}
	suite.userUsecase.On("GetUserByID", user.ID).Return(domain.User{}, fmt.Errorf("User not found"))

	response, err := http.Get(fmt.Sprintf("%s/users/%s", suite.testingServer.URL, user.ID))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusNotFound, response.StatusCode)
}

