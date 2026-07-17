# Production-Ready Authentication System in Go (Phase 1)

This project contains a production-ready, clean-architecture backend for a user authentication system written in Go, using the Gin framework, GORM, and MySQL.

---

## Technical Stack
- **Language**: Go 1.24+
- **HTTP Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL
- **Hashing**: bcrypt
- **Validation**: go-playground/validator/v10

---

## Folder Structure
```text
project/
│
├── cmd/
│      └── server/
│              main.go          # Application Entrypoint / DI bootstrap
│
├── internal/
│      ├── config/
│      │      database.go       # GORM MySQL Connection Pool Setup
│      │      env.go            # Environmental Config Parser
│      │
│      ├── models/
│      │      user.go           # GORM User Database Model
│      │
│      ├── dto/
│      │      register_request.go   # Registration Validation Rules
│      │      register_response.go  # Sanitized Output DTO
│      │
│      ├── repository/
│      │      interfaces.go         # Database interfaces
│      │      user_repository.go    # GORM MySQL implementation
│      │
│      ├── service/
│      │      interfaces.go         # Business workflow interfaces
│      │      auth_service.go       # Registration business workflow
│      │
│      ├── handler/
│      │      auth_handler.go       # HTTP handler mapping HTTP <-> Service
│      │
│      ├── routes/
│      │      routes.go             # Gin Router setups
│      │
│      └── utils/
│             password.go           # Bcrypt Hashing helpers
│             response.go           # Standard JSON response wrappers
│
├── .env                        # Configuration file (Local/Docker)
├── schema.sql                  # MySQL schema definition
├── postman_collection.json     # Postman testing file
├── go.mod                      # Go Module definitions
└── README.md                   # Project documentation
```

---

## Setup Instructions

### 1. Database Setup
1. Start your local MySQL server (port `3306`).
2. Run the `schema.sql` script to create the database `bridge_ai` and the `users` table:
   ```sql
   SOURCE schema.sql;
   ```

### 2. Environment Setup
Configure your environment parameters in the `.env` file at the root directory:
```env
PORT=8080
ENV=development
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=bridge_ai
```

### 3. Run the Server
1. Download dependencies:
   ```bash
   go mod download
   ```
2. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

---

## API Testing Guide

### Endpoint
* **URL**: `http://localhost:8080/api/v1/auth/register`
* **Method**: `POST`
* **Content-Type**: `application/json`

### Success Case (201 Created)
**Request Body**:
```json
{
    "username": "vinith",
    "email": "abc@gmail.com",
    "phone_number": "9876543210",
    "address": "Tamil Nadu",
    "password": "Password@123"
}
```
**Response Body**:
```json
{
    "success": true,
    "message": "User registered successfully",
    "data": {
        "id": 1,
        "username": "vinith",
        "email": "abc@gmail.com",
        "phone_number": "9876543210",
        "address": "Tamil Nadu"
    }
}
```

### Duplicate Error Case (409 Conflict)
**Response Body** (if username, email, or phone number already exists):
```json
{
    "success": false,
    "message": "username already exists"
}
```

### Validation Error Case (400 Bad Request)
**Response Body** (if request lacks required fields or is invalid):
```json
{
    "success": false,
    "message": "Validation failed",
    "data": {
        "Email": "email",
        "Password": "min"
    }
}
```
