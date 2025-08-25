# Gorm Chi Template

A boilerplate Go project using [GORM](https://gorm.io/) for ORM and [Chi](https://github.com/go-chi/chi) for HTTP routing. This template provides a clean structure for building RESTful APIs with authentication, middleware, and Docker support.

## Features

- Modular folder structure for controllers, routes, middlewares, and utilities
- JWT-based authentication middleware
- User controller and routes as examples
- Dockerized setup with PostgreSQL integration
- Environment variable configuration via `.env`
- Ready-to-use response utility functions

## Project Structure

```
gorm-chi-template/
├── cmd/
│   └── main/
│       └── main.go           # Application entry point
├── controllers/
│   └── user.controller.go    # User-related handlers
├── middlewares/
│   └── auth.go               # Authentication middleware
├── routes/
│   ├── main.go               # Route initialization
│   └── user.routes.go        # User routes
├── utils/
│   ├── auth.go               # Auth utility functions
│   └── response.go           # Response formatting utilities
├── Dockerfile                # Docker image build instructions
├── docker-compose.yml        # Multi-container orchestration
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies checksum
├── .env.example              # Example environment variables
└── README.md                 # Project documentation
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18+
- [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/install/)

### Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/suryanshvermaa/Go_Projects.git
   cd Go_Projects/gorm-chi-template
   ```

2. **Configure environment variables:**
   - Copy `.env.example` to `.env` and update values as needed.

3. **Run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

   This will start the Go API and a PostgreSQL database.

4. **Access the API:**
   - The server runs on `http://localhost:3000` by default.

## Environment Variables

See `.env.example` for all required variables:

- `PORT`: API server port
- `AUTH_SECRET_KEY`: JWT secret key
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Database connection details

## Usage

- Extend controllers and routes for new resources.
- Add middlewares in `middlewares/` as needed.
- Use utility functions from `utils/` for common tasks.

## License

MIT

---

For questions or contributions, feel free to open an issue or pull