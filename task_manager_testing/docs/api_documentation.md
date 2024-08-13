
# Task Manager Project - Testing Guide

## Overview

This document outlines the testing process for the Task Manager project. The project is built using Go, and follows Clean Architecture principles. The tests are written using the `testify` package along with `httptest` for HTTP testing and custom mocks for isolating dependencies using go `mockery`.

## Prerequisites

Ensure that you have the following tools and dependencies installed:

- Go (version 1.21+)
- `testify` package
- `httptest` package
- `mockery` package

## Project Structure

The project follows the Clean Architecture pattern, with the following main folders:

- **Delivery**: Contains the controllers for handling HTTP requests.
- **Domain**: Contains the domain models and interfaces.
- **Usecase**: Contains the business logic.
- **Repositories**: Contains the database interactions.
- **Mocks**: Contains the generated mocks for testing purposes.

## Setup

To set up the project for testing, follow these steps:

1. **Install dependencies:**

   ```sh
   go get github.com/stretchr/testify
   go get go.mongodb.org/mongo-driver/bson/primitive
   ```

2. **Generate Mocks (if applicable):**

   If you're using `mockery` to generate mocks:

   ```sh
   go install github.com/vektra/mockery/v2/.../
   mockery --name=TaskUsecase --dir=domain mockery --name=UserUsecase --dir=domain mockery --name=TaskInterface --dir=domain mockery --name=UserInterface--dir=domain       --output=mocks --outpkg=mocks
   ```

3. **Run Tests:**

   You can run all the tests in the project using the `go test` command:

   ```sh
   go test ./... -v
   ```

   This command runs all the test files recursively and outputs verbose results.

## Testing Process

The project is tested at multiple levels, focusing on both unit and integration tests. Below are some of the key tests implemented:

### 1. Unit Tests

Unit tests are written to ensure that individual components behave as expected. Mocks are used to isolate dependencies and test each unit in isolation.

**Example: Testing `AddTask` Functionality**

- **File:** `task_controller_test.go`
- **Description:** Tests the `AddTask` function in the `TaskController`.

```go
func (suite *TaskControllerSuite) TestAddTask() {
    task := domain.Task{
        ID: primitive.NewObjectID(),
        Title: "Task 1",
        Description: "Description 1",
        DueDate: primitive.NewDateTimeFromTime(time.Now()),
        Status: "Completed",
        CreatedBy: primitive.NewObjectID(),
    }

    suite.taskUsecase.On("AddTask", task).Return(nil)

    taskJson, _ := json.Marshal(task)

    req, err := http.NewRequest("POST", fmt.Sprintf("%s/task", suite.testingServer.URL), bytes.NewBuffer(taskJson))
    suite.NoError(err)

    resp, err := http.DefaultClient.Do(req)
    suite.NoError(err)
    suite.Equal(http.StatusOK, resp.StatusCode)

    // Check response body
    var response map[string]string
    err = json.NewDecoder(resp.Body).Decode(&response)
    suite.NoError(err)
    suite.Equal("Task added successfully!", response["message"])

    suite.taskUsecase.AssertExpectations(suite.T())
}
```

### 2. Integration Tests

Integration tests are written to ensure that multiple components work together as expected.

**Example: Testing `GetAllTasks` Functionality**

- **File:** `task_controller_test.go`
- **Description:** Tests the `GetAllTasks` function in the `TaskController`.

```go
func (suite *TaskControllerSuite) TestGetAllTasks() {
    tasks := []domain.Task{
        {
            ID: primitive.NewObjectID(),
            Title: "Task 1",
            Description: "Description 1",
            DueDate: primitive.NewDateTimeFromTime(time.Now()),
            Status: "Completed",
            CreatedBy: primitive.NewObjectID(),
        },
        {
            ID: primitive.NewObjectID(),
            Title: "Task 2",
            Description: "Description 2",
            DueDate: primitive.NewDateTimeFromTime(time.Now()),
            Status: "Pending",
            CreatedBy: primitive.NewObjectID(),
        },
    }

    suite.taskUsecase.On("GetAllTasks").Return(tasks, nil)

    req, err := http.NewRequest("GET", fmt.Sprintf("%s/tasks", suite.testingServer.URL), nil)
    suite.NoError(err)

    resp, err := http.DefaultClient.Do(req)
    suite.NoError(err)
    suite.Equal(http.StatusOK, resp.StatusCode)

    var returnedTasks []domain.Task
    err = json.NewDecoder(resp.Body).Decode(&returnedTasks)
    suite.NoError(err)
    suite.Equal(len(tasks), len(returnedTasks))

    suite.taskUsecase.AssertExpectations(suite.T())
}
```

## Running the Tests

To run the tests, simply execute the following command in the root of the project:

```sh
go test ./... -v
```

This will run all the tests across the project, providing detailed output for each test case.

## Coverage

You can also check the test coverage by running:

```sh
go test ./... -cover
```

This command will provide a coverage percentage, showing how much of your code is covered by tests.

## Conclusion

This guide provides an overview of how testing is implemented in the Task Manager project. Following the Clean Architecture pattern ensures that the codebase is modular, maintainable, and testable. The use of mocks and the `testify` package helps in writing effective and isolated tests, ensuring that each component behaves as expected.

