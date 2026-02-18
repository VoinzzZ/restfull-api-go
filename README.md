# 🚀 Go RESTful API

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/pulls)
[![Maintained](https://img.shields.io/badge/Maintained-yes-green.svg)](https://github.com/)
[![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat&logo=mysql&logoColor=white)](https://dev.mysql.com/downloads/)

A RESTful API for user management built with **Go (Golang)** using **Clean Architecture** and **MySQL** as the database.

---

## 🛠️ Tech Stack

| Technology | Description |
|---|---|
| **Go** | Backend language |
| **net/http** | HTTP server (native) |
| **MySQL** | Database |
| **godotenv** | Load environment variables |
| **bcrypt** | Password hashing |

---

## 📁 Project Structure

```
restfull-api-go/
├── cmd/api/main.go                          # Entry point
├── config/config.go                         # Configuration
├── internal/
│   ├── domain/user.go                       # Entity & interface
│   ├── repository/user_repository.go        # Database layer
│   ├── usecase/user_usecase.go              # Business logic
│   └── delivery/http/
│       ├── handler/user_handler.go          # HTTP handlers
│       ├── middleware/logger.go             # Logger middleware
│       └── router/
│           ├── router.go                    # Main router
│           └── user_router.go              # User routes
├── pkg/
│   ├── database/mysql.go                    # DB connection
│   └── helper/response.go                  # Response helper
├── .env.example                             # Environment template
├── Makefile                                 # Build scripts
└── go.mod
```

---

## ⚙️ Requirements

Make sure the following tools are installed on your machine:

| Tool | Version | Link |
|---|---|---|
| **Go** | >= 1.22 | [golang.org/dl](https://golang.org/dl/) |
| **MySQL** | >= 8.0 | [mysql.com](https://dev.mysql.com/downloads/) |
| **Make** | any | Available by default on Linux/macOS |
| **Git** | any | [git-scm.com](https://git-scm.com/) |

Check your Go version:
```bash
go version
# Expected output: go version go1.22.x ...
```

---

## 🚀 Quick Start

### 1. Clone the repository
```bash
git clone https://github.com/VoinzzZ/restfull-api-go.git
cd restfull-api-go
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Setup environment
```bash
# Copy the environment template
cp .env.example .env

# Edit .env with your database configuration
nano .env   # or use your preferred editor
```

Fill in your `.env` file:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=go_api_db
PORT=8080
```

### 4. Setup database
Log in to MySQL and run:
```sql
CREATE DATABASE go_api_db;

USE go_api_db;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 5. Run the server
```bash
make run
```

If successful, the terminal will display:
```
2026/02/19 01:12:00 Database connected successfully
2026/02/19 01:12:00 Server starting on http://localhost:8080
```

---

## 📡 API Endpoints

Base URL: `http://localhost:8080`

### Users

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/v1/users` | Create a new user |
| `GET` | `/api/v1/users` | Get all users |
| `GET` | `/api/v1/users/{id}` | Get user by ID |
| `PUT` | `/api/v1/users/{id}` | Update user |
| `DELETE` | `/api/v1/users/{id}` | Delete user |

---

## 📋 API Reference

### Create User
**POST** `/api/v1/users`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201 Created):**
```json
{
  "status": 201,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-02-19T01:12:00Z",
    "updated_at": "2026-02-19T01:12:00Z"
  }
}
```

---

### Get All Users
**GET** `/api/v1/users`

**Response (200 OK):**
```json
{
  "status": 200,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2026-02-19T01:12:00Z",
      "updated_at": "2026-02-19T01:12:00Z"
    }
  ]
}
```

---

### Get User by ID
**GET** `/api/v1/users/{id}`

**Response (200 OK):**
```json
{
  "status": 200,
  "message": "User found",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-02-19T01:12:00Z",
    "updated_at": "2026-02-19T01:12:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "status": 404,
  "message": "user not found"
}
```

---

### Update User
**PUT** `/api/v1/users/{id}`

**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.new@example.com"
}
```

**Response (200 OK):**
```json
{
  "status": 200,
  "message": "User updated successfully"
}
```

---

### Delete User
**DELETE** `/api/v1/users/{id}`

**Response (200 OK):**
```json
{
  "status": 200,
  "message": "User deleted successfully"
}
```

---

## 🧪 Testing with cURL

```bash
# Create user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'

# Get all users
curl http://localhost:8080/api/v1/users

# Get user by ID
curl http://localhost:8080/api/v1/users/1

# Update user
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.new@example.com"}'

# Delete user
curl -X DELETE http://localhost:8080/api/v1/users/1
```

---

## 📦 Makefile Commands

```bash
make run    # Run the server (development)
make build  # Compile to binary
make start  # Run the compiled binary
make clean  # Remove binary
```

---

## 🏗️ Architecture

This project uses **Clean Architecture** with the following flow:

```
Request → Router → Middleware → Handler → Usecase → Repository → Database
                                                                      ↓
Response ← Router ← Middleware ← Handler ← Usecase ← Repository ← Database
```

| Layer | Package | Responsibility |
|---|---|---|
| **Domain** | `internal/domain` | Entity & interface contract |
| **Repository** | `internal/repository` | Database operations |
| **Usecase** | `internal/usecase` | Business logic |
| **Handler** | `internal/delivery/http/handler` | HTTP request/response |
| **Middleware** | `internal/delivery/http/middleware` | Cross-cutting concerns |
| **Router** | `internal/delivery/http/router` | Route definitions |

---

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
