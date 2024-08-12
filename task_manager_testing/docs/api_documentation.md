# Task Manager Using Go Integrated With MongoDB

## Overview

This project is a task management API built using Go, providing endpoints to create, retrieve, update, and delete tasks. The API is designed to be RESTful, making it easy to integrate with various client applications.

## API Documentation

For detailed API documentation, please refer to the [Postman Documentation](https://documenter.getpostman.com/view/23906890/2sA3rwKtdG).

This documentation includes:
- **Endpoint URLs**: The paths for each API endpoint.
- **HTTP Methods**: The request methods (GET, POST, PUT, PATCH, DELETE) supported by each endpoint.
- **Request Parameters**: The required and optional parameters for each request.
- **Response Formats**: The structure of the response data returned by the API.
- **Examples**: Sample requests and responses for each endpoint.

## Roles and Authentication

This API supports three roles: user, admin, and root user. JWT authentication and authorization are implemented to ensure secure access to the API endpoints.

### Role Descriptions

- **User**: Can only edit and delete their tasks and update their profiles but cannot change others' data or passwords.
- **Admin**: Can edit users' profiles (but not those of other admins or the root user) and manage tasks of other users. Admins cannot change others' passwords.
- **Root User**: Has the highest level of access, including the ability to manage all data and tasks related to other admins and users but cannot edit their passwords. The root user can also add, edit, update, and delete their tasks and add users and admins. There is only one root user.

### Authentication and Authorization

- **JWT Authentication**: Ensures secure access to the API endpoints. Users must authenticate using a valid JSON Web Token.
- **Authorization**: Based on the role assigned, users can access specific endpoints and perform actions allowed by their role.

## Getting Started

### Prerequisites

- Go (version 1.22 or later)
- Gin (Web Framework for Go)
- MongoDB (version 4.0 or later)


