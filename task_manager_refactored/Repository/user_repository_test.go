package repository_test

import (
	"task_manager_refactored/Infrastructure/database"
	repository "task_manager_refactored/Repository"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserRepositorySuite struct {
	suite.Suite
	repository repository.UserRepository
	collection *mongo.Collection
	client     *mongo.Client
}

func (suite *UserRepositorySuite) SetupTest() {
	client, _ := database.ConnectToMongoDB("")
	suite.client = client
	db := "taskdb"
	repository := repository.NewUserRepository(client, db, "userstest")
	suite.repository = *repository
	
}

