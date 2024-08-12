package usecase_test

import (
	usecase "task_manager_testing/Usecase"
	"task_manager_testing/domain"
	"task_manager_testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseSuite struct {
	suite.Suite
	userUsecase *usecase.UserUsecase
	mockRepo    *mocks.UserRepository
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
	suite.userUsecase = usecase.NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseSuite) TestRegisterUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}

	suite.mockRepo.On("RegisterUser", user.Username, user.Password, user.Role).Return(nil)

	err := suite.userUsecase.RegisterUser(user.Username, user.Password, user.Role)

	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLogin() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678", // Plain text password to be provided during login
		Role:     "user",
	}

	// Hash the password that will be stored in the mock repository
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	assert.NoError(suite.T(), err)

	// Update the user object with the hashed password
	storedUser := user
	storedUser.Password = string(hashedPassword)

	// Mock the Login method to return the user with the hashed password
	suite.mockRepo.On("Login", user.Username, user.Password).Return(storedUser, nil)

	// Call the Login method with the plain text password
	result, err := suite.userUsecase.Login(user.Username, "12345678")

	// Ensure there were no errors during login
	assert.NoError(suite.T(), err)

	// Compare the provided password with the hashed password stored in the mock repository
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte("12345678"))
	assert.NoError(suite.T(), err, "Password should match the hashed password")

	// Check if other fields match as expected
	assert.Equal(suite.T(), user.Username, result.Username)
	assert.Equal(suite.T(), user.Role, result.Role)
	suite.mockRepo.AssertExpectations(suite.T())
}


func (suite *UserUsecaseSuite) TestGetAllUsers() {
	users := []domain.User{
		{
			ID:       primitive.NewObjectID(),
			Username: "tester1",
			Password: "12345678",
			Role:     "user",
		},
	}

	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := suite.userUsecase.GetAllUsers()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), users, result)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetUserById() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678", // Plain text password
		Role:     "user",
	}

	// Hash the password before setting it in the mock repository
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	assert.NoError(suite.T(), err)

	// Update the user object with the hashed password
	storedUser := user
	storedUser.Password = string(hashedPassword)

	// Mock the GetUserById method to return the user with the hashed password
	suite.mockRepo.On("GetUserById", user.ID).Return(storedUser, nil)

	// Call the GetUserById usecase method
	result, err := suite.userUsecase.GetUserById(user.ID)
	assert.NoError(suite.T(), err)

	// Compare the provided password with the hashed password stored in the mock repository
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte("12345678"))
	assert.NoError(suite.T(), err, "Password should match the hashed password")

	// Check if other fields match as expected
	assert.Equal(suite.T(), storedUser.Username, result.Username)
	assert.Equal(suite.T(), storedUser.Role, result.Role)

	suite.mockRepo.AssertExpectations(suite.T())
}


func (suite *UserUsecaseSuite) TestUpdateUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}

	suite.mockRepo.On("UpdateUser", user.ID, user).Return(nil)

	err := suite.userUsecase.UpdateUser(user.ID, user)

	assert.NoError(suite.T(), err)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestDeleteUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}

	suite.mockRepo.On("DeleteUser", user.ID).Return(nil)

	err := suite.userUsecase.DeleteUser(user.ID)

	assert.NoError(suite.T(), err)

	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
