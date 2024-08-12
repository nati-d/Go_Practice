package usecase_test

import (
	"testing"

	// infrastructure "task_manager_testing/Infrastructure"
	infrastructure "task_manager_testing/Infrastructure"
	usecase "task_manager_testing/Usecase"
	"task_manager_testing/domain"
	"task_manager_testing/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecaseSuite defines the suite for user usecase tests
type UserUsecaseSuite struct {
	suite.Suite
	userRepo    *mocks.UserRepository
	userUsecase *usecase.UserUsecase
}

// SetupTest sets up the necessary resources before each test
func (suite *UserUsecaseSuite) SetupTest() {
	suite.userRepo = &mocks.UserRepository{}
	suite.userUsecase = usecase.NewUserUsecase(suite.userRepo)
}



// TestRegisterUser tests the RegisterUser functionality
func (suite *UserUsecaseSuite) TestRegisterUser() {
	// Setup the test case
	user := domain.User{
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}

	// temp := user.Password


	suite.userRepo.On("RegisterUser", user.Username, mock.AnythingOfType("string")  , user.Role).Return(nil)


	// Execute the test case
	err := suite.userUsecase.RegisterUser(user.Username, user.Password, user.Role)

	// Verify the test results
	suite.Require().NoError(err)
	suite.userRepo.AssertExpectations(suite.T())
	
}

//test login
func (suite *UserUsecaseSuite) TestLogin() {
	// Setup the test case
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username : "tester1",
		Password : "12345678",
		Role : "user",
	}

	// Hash the password that will be stored in the mock repository
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	suite.Require().NoError(err)

	// Update the user object with the hashed password
	storedUser := user
	storedUser.Password = hashedPassword

	// Mock the Login method to return the user with the hashed password
	suite.userRepo.On("Login", user.Username, user.Password).Return(storedUser, nil)

	// Call the Login method with the plain text password
	result, err := suite.userUsecase.Login(user.Username, "12345678")

	// Ensure there were no errors during login
	suite.Require().NoError(err)

	// Compare the provided password with the hashed password stored in the mock repository
	err = infrastructure.ComparePasswords(result.Password, "12345678")

	// Check if other fields match as expected
	suite.Require().NoError(err)

	suite.Require().Equal(user.Username, result.Username)
	suite.Require().Equal(user.Role, result.Role)

	suite.userRepo.AssertExpectations(suite.T())

}

// TestGetAllUsers tests the GetAllUsers functionality
func (suite *UserUsecaseSuite) TestGetAllUsers() {
	// Setup the test case
	users := []domain.User{
		{
			ID:       primitive.NewObjectID(),
			Username: "tester1",
			Password : "12345678",
			Role:     "user",
		},

	}

	suite.userRepo.On("GetAllUsers").Return(users, nil)

	// Execute the test case
	result, err := suite.userUsecase.GetAllUsers()

	// Verify the test results
	suite.Require().NoError(err)

	suite.Require().Equal(users, result)

	suite.userRepo.AssertExpectations(suite.T())

}


// TestGetUserById tests the GetUserById functionality
func (suite *UserUsecaseSuite) TestGetUserById() {
	// Setup the test case
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username : "tester1",
		Password : "12345678", // Plain text password
		Role : "user",
	}

	// Hash the password before setting it in the mock repository
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	suite.Require().NoError(err)

	// Update the user object with the hashed password
	storedUser := user
	storedUser.Password = hashedPassword

	// Mock the GetUserById method to return the user with the hashed password
	suite.userRepo.On("GetUserById", user.ID).Return(storedUser, nil)

	// Call the GetUserById usecase method
	result, err := suite.userUsecase.GetUserById(user.ID)
	suite.Require().NoError(err)

	// Compare the provided password with the hashed password stored in the mock repository
	err = infrastructure.ComparePasswords(result.Password, "12345678")

	// Verify the test results
	suite.Require().NoError(err)

	suite.userRepo.AssertExpectations(suite.T())

}


func (suite *UserUsecaseSuite) TestUpdateUser() {
	// Setup the test case
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username : "tester1",
		Password : "12345678",
		Role : "user",
	}

	// Hash the password before setting it in the mock repository
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	suite.Require().NoError(err)

	// Update the user object with the hashed password
	storedUser := user
	storedUser.Password = hashedPassword

	// Mock the UpdateUser method to return the user with the hashed password
	suite.userRepo.On("UpdateUser", user.ID, user).Return(nil)

	// Call the UpdateUser usecase method

	err = suite.userUsecase.UpdateUser(user.ID, user)
	// Verify the test results
	suite.Require().NoError(err)

	suite.userRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestDeleteUser() {
	// Setup the test case
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username : "tester1",
		Password : "12345678",
		Role : "user",
	}

	// Mock the DeleteUser method to return the user with the hashed password
	suite.userRepo.On("DeleteUser", user.ID).Return(nil)

	// Call the DeleteUser usecase method
	err := suite.userUsecase.DeleteUser(user.ID)

	// Verify the test results
	suite.Require().NoError(err)

	suite.userRepo.AssertExpectations(suite.T())

}

// TearDownTest clears resources after each test
func (suite *UserUsecaseSuite) TearDownTest() {
	// Reset the mock expectations
	suite.userRepo.AssertExpectations(suite.T())
}


// Run the test suite
func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
