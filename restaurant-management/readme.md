# Restaurant Management System

A RESTful API built with Go for managing restaurant operations, including food items, menus, orders, invoices, tables, and users.

## Features

- **Food Management**: Add, update, delete, and list food items.
- **Menu Management**: Organize food items into menus.
- **Order Management**: Create and manage customer orders and order items.
- **Invoice Generation**: Generate invoices for completed orders.
- **Table Management**: Assign and manage tables for orders.
- **User Authentication**: Secure endpoints with authentication middleware.
- **Notes**: Add notes to orders or tables.

## Project Structure

```
controllers/      # Handles HTTP requests and business logic
database/         # Database connection and setup
helpers/          # Utility functions (e.g., token generation)
middlewares/      # Authentication and other middleware
models/           # Data models for all entities
routes/           # API route definitions
main.go           # Application entry point
go.mod            # Go module definition
readme.md         # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.18+
- A running database (e.g., PostgreSQL, MySQL)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/suryanshvermaa/Go_Projects.git
    cd Go_Projects/restaurant-management
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure your database connection in `database/databaseConnection.go`.

### Running the Application

```bash
go run main.go
```

The server will start and expose RESTful endpoints for managing restaurant resources.

## API Endpoints

- `/food` - Manage food items
- `/menu` - Manage menus
- `/order` - Manage orders
- `/order-item` - Manage order items
- `/invoice` - Manage invoices
- `/table` - Manage tables
- `/user` - User registration and authentication
- `/note` - Add notes

Refer to the route files in `routes/` for detailed endpoint definitions.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License.

---

For questions or support, contact