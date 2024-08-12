package repository_test

import (
	"context"
	"testing"

	"task_manager_testing/Infrastructure/database"
	repository "task_manager_testing/Repository"
	"task_manager_testing/domain"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite
	repository repository.UserRepository
	collection *mongo.Collection
	client     *mongo.Client
}

func (suite *UserRepositorySuite) SetupTest() {
	client, err := database.ConnectToMongoDB("mongodb+srv://nathnaeldes:12345678n@cluster0.w8bpdtf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		suite.T().Fatal("Failed to connect to MongoDB:", err)
	}
	suite.client = client
	db := "taskdb"
	collectionName := "userstest"
	suite.collection = client.Database(db).Collection(collectionName)
	repository := repository.NewUserRepository(client, db, collectionName)
	suite.repository = *repository
}

func (suite *UserRepositorySuite) TestAddUser() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	}
	err := suite.repository.RegisterUser(user.Username, user.Password, user.Role)
	suite.NoError(err)
}

func (suite *UserRepositorySuite) TestGetAllUsers() {
	users, err := suite.repository.GetAllUsers()
	suite.NoError(err)
	suite.Empty(users)
}

func (suite *UserRepositorySuite) TestGetUserById() {
	id := primitive.NewObjectID()
	_, err := suite.collection.InsertOne(context.TODO(), domain.User{
		ID:       id,
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	})
	suite.NoError(err)

	user, err := suite.repository.GetUserById(id)
	suite.NoError(err)
	suite.NotEmpty(user)
	suite.Equal("tester1", user.Username)
	suite.Equal("user", user.Role)
}

func (suite *UserRepositorySuite) TestUpdateUser() {
	id := primitive.NewObjectID()
	_, err := suite.collection.InsertOne(context.TODO(), domain.User{
		ID:       id,
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	})
	suite.NoError(err)

	updatedUser := domain.User{
		ID:       id,
		Username: "tester2",
		Password: "87654321",
		Role:     "admin",
	}

	err = suite.repository.UpdateUser(id, updatedUser)
	suite.NoError(err)

	userFromDb := domain.User{}
	err = suite.collection.FindOne(context.TODO(), primitive.M{"_id": id}).Decode(&userFromDb)
	suite.NoError(err)
	suite.Equal("tester2", userFromDb.Username)
	suite.Equal("87654321", userFromDb.Password)
	suite.Equal("admin", userFromDb.Role)
}

func (suite *UserRepositorySuite) TestDeleteUser() {
	id := primitive.NewObjectID()
	_, err := suite.collection.InsertOne(context.TODO(), domain.User{
		ID:       id,
		Username: "tester1",
		Password: "12345678",
		Role:     "user",
	})
	suite.NoError(err)

	err = suite.repository.DeleteUser(id)
	suite.NoError(err)

	// Verify the user is deleted
	user, err := suite.repository.GetUserById(id)
	suite.Error(err)
	suite.Empty(user)
}

func (suite *UserRepositorySuite) TearDownTest() {
	err := suite.client.Database("taskdb").Collection("userstest").Drop(context.TODO())
	if err != nil {
		suite.T().Fatal("Failed to drop collection:", err)
	}
}

func (suite *UserRepositorySuite) TearDownSuite() {
	err := suite.client.Disconnect(context.TODO())
	if err != nil {
		suite.T().Fatal("Failed to disconnect from MongoDB:", err)
	}
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
