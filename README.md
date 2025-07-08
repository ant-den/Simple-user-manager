# GO_Project

A RESTful API in Go for user management using Gin and GORM.

## Features

- CRUD operations for users (Create, Read, Update, Delete)
- API versioning (`/api/v1`)
- Swagger documentation (`/swagger/index.html`)
- Clean architecture (Repository, Service, Controller layers)
- Uses GORM for database operations

## Quick Start

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/GO_Project.git
cd GO_Project
```

### 2. Configure the database

Create a MySQL database and set the connection parameters in `database/db.go`  
(or use environment variables if implemented).

### 3. Install dependencies

```sh
go mod tidy
```

### 4. Run the application

```sh
go run main.go
```

The server will be available at [http://localhost:8081](http://localhost:8081)

### 5. Open Swagger documentation

Go to: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

## Project Structure

```
GO_Project/
├── controller/      # Controllers (HTTP request handlers)
├── database/        # Database connection
├── docs/            # Swagger documentation
├── models/          # Data models
├── repository/      # Database repositories
├── routes/          # API routes
├── service/         # Business logic
├── main.go          # Entry point
```

## Example Requests

- Get all users:  
  `GET /api/v1/users/`
- Get user by ID:  
  `GET /api/v1/users/{id}`
- Create user:  
  `POST /api/v1/users/`
- Update user:  
  `PUT /api/v1/users/{id}`
- Delete user:  
  `DELETE /api/v1/users/{id}`

## Requirements

- Go 1.18+
- MySQL
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/gin-swagger)
