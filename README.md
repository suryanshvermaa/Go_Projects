# 🚀 Go Projects

A comprehensive collection of Go projects demonstrating various web development patterns, database integrations, and modern API development techniques. This repository serves as a learning resource and boilerplate for building scalable Go applications.

## 📋 Table of Contents

- [Overview](#overview)
- [Projects](#projects)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

This repository contains multiple Go projects showcasing different aspects of backend development, from simple web servers to complex microservices with various database integrations and communication protocols.

## 🏗️ Projects

### 🔐 [go-mongo-jwt](./go-mongo-jwt/) - JWT Authentication with MongoDB
A RESTful API built with Go, Gin, and MongoDB featuring JWT-based authentication and role-based access control.

**Features:**
- ✅ User registration and login endpoints
- 🔑 JWT authentication and role-based access control
- 🗄️ MongoDB integration for persistent storage
- 🏗️ Modular code structure (controllers, models, routes, middlewares)
- 🐳 Docker Compose for easy MongoDB setup

**Tech Stack:** Go, Gin, MongoDB, JWT, Docker

---

### 📚 [golang-MySQL](./golang-MySQL/) - Book Management API with MySQL
A book management system using Go, GORM, and MySQL with RESTful API endpoints.

**Features:**
- 📖 CRUD operations for books
- 🗄️ MySQL database integration with GORM
- 🔄 RESTful API endpoints
- 🏗️ Clean architecture with pkg structure

**Tech Stack:** Go, GORM, MySQL, Gorilla Mux

---

### 🔄 [gRPC](./gRPC/) - gRPC Communication Examples
Comprehensive gRPC examples demonstrating all four types of gRPC communication patterns.

**Features:**
- 🎯 Unary RPC
- 📡 Server Streaming RPC
- 📤 Client Streaming RPC
- 🔄 Bidirectional Streaming RPC
- 📝 Protocol Buffers integration

**Tech Stack:** Go, gRPC, Protocol Buffers

---

### 🎯 [graphQl](./graphQl/) - GraphQL API
A GraphQL API implementation using gqlgen for type-safe GraphQL development.

**Features:**
- 🎯 GraphQL schema definition
- 🔧 gqlgen code generation
- 📊 Type-safe resolvers
- 🗄️ Database integration ready

**Tech Stack:** Go, GraphQL, gqlgen

---

### 🐘 [postgres-go-crud](./postgres-go-crud/) - PostgreSQL CRUD Operations
A complete CRUD application with PostgreSQL, JWT authentication, and user management.

**Features:**
- 👤 User management with CRUD operations
- 🔐 JWT authentication and authorization
- 🗄️ PostgreSQL database integration
- 🛡️ Middleware for route protection
- 🔧 Environment configuration

**Tech Stack:** Go, PostgreSQL, JWT, Gorilla Mux, bcrypt

---

### 🌐 [WebServer](./WebServer/) - Simple Web Server
A basic web server serving static files and handling form submissions.

**Features:**
- 📄 Static file serving
- 📝 Form handling
- 🎯 Simple routing
- 🏠 Basic HTML templates

**Tech Stack:** Go, HTML, CSS

---

### 🗄️ [postgres-go](./postgres-go/) - PostgreSQL Integration Template
A template project for PostgreSQL integration with Go.

**Features:**
- 🗄️ PostgreSQL connection setup
- 🏗️ MVC architecture template
- 📝 Route definitions
- 🎯 Controller structure

**Tech Stack:** Go, PostgreSQL

## 🛠️ Technologies Used

### Core Technologies
- **Go** - Primary programming language
- **Gin** - HTTP web framework
- **Gorilla Mux** - HTTP router and URL matcher
- **GORM** - Object-Relational Mapping library

### Databases
- **MongoDB** - NoSQL database
- **MySQL** - Relational database
- **PostgreSQL** - Advanced relational database

### Authentication & Security
- **JWT** - JSON Web Tokens
- **bcrypt** - Password hashing
- **Middleware** - Route protection

### Communication Protocols
- **gRPC** - High-performance RPC framework
- **GraphQL** - Query language for APIs
- **REST** - Representational State Transfer

### Development Tools
- **Docker** - Containerization
- **Protocol Buffers** - Data serialization
- **gqlgen** - GraphQL code generation

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher
- Docker (for containerized databases)
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go_projects
   ```

2. **Navigate to any project:**
   ```bash
   cd <project-name>
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Set up environment variables** (if required):
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Run the application:**
   ```bash
   go run cmd/main/main.go
   # or
   go run main.go
   ```

## 📁 Project Structure

```
go_projects/
├── 🔐 go-mongo-jwt/          # JWT Auth with MongoDB
├── 📚 golang-MySQL/          # Book Management with MySQL
├── 🔄 gRPC/                  # gRPC Communication Examples
├── 🎯 graphQl/              # GraphQL API
├── 🐘 postgres-go-crud/     # PostgreSQL CRUD with JWT
├── 🌐 WebServer/            # Simple Web Server
└── 🗄️ postgres-go/          # PostgreSQL Template
```

## 🎯 Learning Objectives

This repository is designed to help you learn:

- **Web Development** - Building RESTful APIs and web servers
- **Database Integration** - Working with different database types
- **Authentication** - Implementing secure authentication systems
- **Microservices** - Understanding service-to-service communication
- **Modern APIs** - GraphQL and gRPC implementations
- **Best Practices** - Clean architecture and code organization

## 🤝 Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Contribution Guidelines

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go community for excellent documentation and libraries
- Contributors and maintainers of all the open-source libraries used
- The Go team for creating such a powerful and efficient language

---
