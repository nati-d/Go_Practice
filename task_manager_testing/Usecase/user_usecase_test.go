package usecase_test

import (
	"fmt"
	// "log"
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

	suite.userRepo.On("RegisterUser", user.Username, mock.AnythingOfType("string"), user.Role).Return(nil).
		Run(func(args mock.Arguments) {
			password := args.Get(1).(string)
			err := infrastructure.ComparePasswords(password, user.Password)
			suite.NoError(err)
			user.Password = password

		})

	// Execute the test case
	err := suite.userUsecase.RegisterUser(user.Username, user.Password, user.Role)

	// Verify the test results
	suite.Require().NoError(err)
	suite.userRepo.AssertExpectations(suite.T())

}

// test login
func (suite *UserUsecaseSuite) TestLogin() {
	testCases := []struct {
		name          string
		user          domain.User
		inputPassword string
		mockReturn    domain.User
		mockError     error
		expectedError bool
	}{
		{
			name: "Valid login",
			user: domain.User{
				Username: "validUser",
				Password: "ValidPassword123",
				Role:     "user",
			},
			inputPassword: "ValidPassword123",
			mockReturn: domain.User{
				Username: "validUser",
				Password: "", // We will set this later after hashing
				Role:     "user",
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name: "Nonexistent user",
			user: domain.User{
				Username: "nonexistentUser",
				Password: "password",
				Role:     "user",
			},
			inputPassword: "password",
			mockReturn:    domain.User{},
			mockError:     fmt.Errorf("user not found"),
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.name != "Nonexistent user" {
				// Hash the password for the mock return user if not "Nonexistent user"
				hashedPassword, err := infrastructure.HashPassword(tc.user.Password)
				suite.Require().NoError(err)
				tc.mockReturn.Password = hashedPassword
			}

			// Mock the Login method
			suite.userRepo.On("Login", tc.user.Username, mock.AnythingOfType("string")).Return(tc.mockReturn, tc.mockError)

			// Execute the test case
			result, err := suite.userUsecase.Login(tc.user.Username, tc.inputPassword)

			// Verify the test results
			if tc.expectedError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(tc.user.Username, result.Username)
				suite.Equal(tc.user.Role, result.Role)
				err := infrastructure.ComparePasswords(result.Password, tc.inputPassword)
				suite.NoError(err)
			}

			suite.userRepo.AssertExpectations(suite.T())
		})
	}
}

// TestGetAllUsers tests the GetAllUsers functionality
func (suite *UserUsecaseSuite) TestGetAllUsers() {
	// Setup the test case
	users := []domain.User{
		{
			ID:       primitive.NewObjectID(),
			Username: "tester1",
			Password: "12345678",
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
	testCases := []struct {
		name          string
		userID        primitive.ObjectID
		storedUser    domain.User
		mockError     error
		expectedError bool
		expectedUser  domain.User
	}{
		{
			name:   "Valid User ID",
			userID: primitive.NewObjectID(),
			storedUser: domain.User{
				ID:       primitive.NewObjectID(),
				Username: "tester1",
				Password: "", // We will hash it later
				Role:     "user",
			},
			mockError:     nil,
			expectedError: false,
			expectedUser: domain.User{
				ID:       primitive.NewObjectID(),
				Username: "tester1",
				Password: "12345678", // Plain text for comparison
				Role:     "user",
			},
		},
		{
			name:          "Non-existent User ID",
			userID:        primitive.NewObjectID(),
			storedUser:    domain.User{},
			mockError:     fmt.Errorf("user not found"),
			expectedError: true,
		},
		{
			name:          "Database Error",
			userID:        primitive.NewObjectID(),
			storedUser:    domain.User{},
			mockError:     fmt.Errorf("database error"),
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// If we are not testing for error scenarios, hash the password
			if !tc.expectedError {
				hashedPassword, err := infrastructure.HashPassword(tc.expectedUser.Password)
				suite.Require().NoError(err)
				tc.storedUser.Password = hashedPassword
			}

			// Mock the GetUserById method to return the user or error
			suite.userRepo.On("GetUserById", tc.userID).Return(tc.storedUser, tc.mockError)

			// Call the GetUserById usecase method
			result, err := suite.userUsecase.GetUserById(tc.userID)

			if tc.expectedError {
				suite.Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Equal(tc.storedUser.ID, result.ID)
				suite.Equal(tc.storedUser.Username, result.Username)
				suite.Equal(tc.storedUser.Role, result.Role)

				// Compare the provided password with the hashed password stored in the mock repository
				err = infrastructure.ComparePasswords(result.Password, tc.expectedUser.Password)
				suite.Require().NoError(err)
			}

			suite.userRepo.AssertExpectations(suite.T())
		})
	}
}

// func (suite *UserUsecaseSuite) TestUpdateUser() {
// 	// Setup the test case
// 	user := domain.User{
// 		ID:       primitive.NewObjectID(),
// 		Username: "tester1",
// 		Password: "12345678",
// 		Role:     "user",
// 	}

// 	// Hash the password before setting it in the mock repository
// 	// hashedPassword, err := infrastructure.HashPassword(user.Password)
// 	// suite.Require().NoError(err)

// 	// // Update the user object with the hashed password
// 	// storedUser := user
// 	// storedUser.Password = hashedPassword

// 	// Mock the UpdateUser method to return the user with the hashed password
// 	// suite.userRepo.On("UpdateUser", user.ID, user).Return(nil)
// 	suite.userRepo.On("UpdateUser", user.ID, user).Return(nil).
// 		Run(func(args mock.Arguments) {
// 			password := args.Get(1).(string)
// 			log.Print(password)
// 			err := infrastructure.ComparePasswords(password, user.Password)
// 			suite.NoError(err)
// 			user.Password = password
// 		})

// 	// Call the UpdateUser usecase method

// 	err := suite.userUsecase.UpdateUser(user.ID, user)
// 	// Verify the test results
// 	suite.Require().NoError(err)

// 	suite.userRepo.AssertExpectations(suite.T())

// }

func (suite *UserUsecaseSuite) TestDeleteUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}

	rootUser, _ := primitive.ObjectIDFromHex("66bb0cfa397c997e09b4afb8")
	testCases := []struct {
		name          string
		userID        primitive.ObjectID
		mockReturnErr error
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "Valid delete",
			userID:        user.ID,
			mockReturnErr: nil,
			expectedError: false,
		},
		{
			name:          "Delete non-existent user",
			userID:        primitive.NewObjectID(),
			mockReturnErr: fmt.Errorf("user not found"),
			expectedError: true,
			expectedMsg:   "user not found",
		},
		{
			name:          "Delete root user",
			userID:        rootUser,
			mockReturnErr: fmt.Errorf("cannot delete root user"),
			expectedError: true,
			expectedMsg:   "cannot delete root user",
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Mock the DeleteUser method to return the error if expected
			suite.userRepo.On("DeleteUser", tc.userID).Return(tc.mockReturnErr)

			// Call the DeleteUser usecase method
			err := suite.userUsecase.DeleteUser(tc.userID)

			// Verify the test results
			if tc.expectedError {
				suite.Error(err)
				suite.EqualError(err, tc.expectedMsg)
			} else {
				suite.NoError(err)
			}

			suite.userRepo.AssertExpectations(suite.T())
		})
	}
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
