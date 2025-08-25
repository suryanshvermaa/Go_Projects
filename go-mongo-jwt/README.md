# Go Mongo JWT API

A RESTful API built with Go, Gin, and MongoDB, featuring JWT-based authentication. This project provides a boilerplate for user management, authentication, and secure API development.

## Features
- User registration and login endpoints
- JWT authentication and role-based access control
- MongoDB integration for persistent storage
- Modular code structure (controllers, models, routes, middlewares, helpers)
- Docker Compose for easy MongoDB setup

## Project Structure
```
.
├── cmd/main/main.go         # Application entry point
├── controllers/             # HTTP handler logic (userController.go)
├── database/                # MongoDB connection logic (connection.go)
├── helpers/                 # Helper functions for authentication (authHelper.go)
├── middlewares/             # (Stub) Middleware for authentication
├── models/                  # Data models (userModel.go)
├── routes/                  # Route definitions (authRouter.go, userRouter.go)
├── utils/                   # Utility functions (response.go)
├── docker-compose.yml       # MongoDB service definition
├── go.mod / go.sum          # Go module dependencies
├── .gitignore               # Git ignore rules
├── mongo-data/              # MongoDB data volume (ignored by git)
```

## Getting Started

### Prerequisites
- Go 1.18+
- Docker (for MongoDB)

### Setup
1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd go-mongo-jwt
   ```
2. **Start MongoDB with Docker Compose:**
   ```sh
   docker-compose up -d
   ```
   This will start MongoDB on port 27000 with default credentials (`root`/`root`).

3. **Configure environment variables:**
   Create a `.env` file in the root directory with the following content:
   ```env
   MONGO_URI=mongodb://root:root@localhost:27000
   PORT=3000
   ```

4. **Install Go dependencies:**
   ```sh
   go mod tidy
   ```

5. **Run the application:**
   ```sh
   go run cmd/main/main.go
   ```
   The server will start on `http://localhost:3000`.

## API Endpoints

### Health Check
- `GET /` — Returns server health status.

### Authentication
- `POST /user/signup` — Register a new user
- `POST /user/login` — Login and receive JWT

### Users
- `GET /user/:user_id` — Get user details (protected, role-based)

## Code Overview

- **cmd/main/main.go**: Initializes Gin, loads environment variables, sets up the root route, and starts the server.
- **controllers/userController.go**: Contains user-related handlers (signup, login, get user). Uses helpers for password hashing and role checks.
- **database/connection.go**: Handles MongoDB connection and exposes collection access.
- **helpers/authHelper.go**: Provides functions for checking user roles and matching user IDs for authorization.
- **models/userModel.go**: Defines the `User` struct with validation tags.
- **routes/authRouter.go & userRouter.go**: Define authentication and user routes.
- **utils/response.go**: Standardizes API JSON responses.
- **middlewares/authMiddleware.go**: (Stub) Placeholder for authentication middleware logic.

## Development Notes
- The project uses Gin for HTTP routing and middleware.
- MongoDB connection details are loaded from `.env`.
- JWT and password logic is referenced but not fully implemented in the provided code.
- The `middlewares/authMiddleware.go` is a stub and should be implemented for full JWT protection.
- The `mongo-data/` directory is used for MongoDB data persistence and is git-ignored.

## License
MIT 