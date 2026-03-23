# GIA Starter App - Clean Architecture

This project is a backend application built using the Go (Golang) programming language with the Gin Gonic framework, following the **Clean Architecture** (or Hexagonal Architecture) pattern. This structure is designed to decouple business logic from technical details (such as frameworks, databases, etc.) to enhance scalability and ease of testing.

### 🚀 Tech Stack

-   **Web Framework**: [Gin Gonic](https://gin-gonic.com/) (High-performance HTTP web framework)
-   **Configuration**: [Viper](https://github.com/spf13/viper) (Go configuration with fangs)
-   **Logging**: [Uber Zap](https://github.com/uber-go/zap) (Blazing fast, structured, leveled logging)
-   **Hot Reloading**: [Air](https://github.com/cosmtrek/air) (Live reload for Go apps)
-   **API Documentation**: [Swagger](https://github.com/swaggo/swag) (Interactive API documentation)
-   **Database (ORM)**: [GORM](https://gorm.io/) (Fantastic ORM library for Golang)
-   **Validation**: [Go Playground Validator](https://github.com/go-playground/validator)

## Folder Structure

```text
gia-starter-app-V1/
│
├── cmd/                # Application entry point
│   └── api/
│       └── main.go     # Main entry point for the API service
│
├── internal/           # Private application code (cannot be imported by other systems)
│   ├── domain/         # Business Logic Core
│   │   ├── entity/     # Business models/entities
│   │   ├── repository/ # Interfaces for data access (Persistence Layer)
│   │   └── service/    # Domain services
│   │
│   ├── usecase/        # Application Business Rules (Interactors)
│   │
│   ├── delivery/        # Transport Layer (HTTP, gRPC, etc.)
│   │   ├── http/
│   │   │   ├── handler/    # HTTP Request handlers
│   │   │   ├── middleware/ # HTTP Middleware (Auth, Logger, etc.)
│   │   │   └── router.go   # API route definitions
│   │
│   ├── infrastructure/ # Technical details (External implementations)
│   │   ├── database/    # Database connections and drivers
│   │   ├── persistence/ # Concrete implementations of repository interfaces
│   │   ├── config/      # Application settings
│   │   └── logger/      # Logging system
│   │
│   └── shared/         # Code used across various layers
│       ├── util/       # Utility functions
│       ├── constant/   # Global constants
│       └── errors/     # Custom error definitions
│
├── pkg/                # Public libraries that can be reused
│   ├── response/       # Standardization of API responses
│   ├── pagination/     # Helpers for pagination
│   └── validator/      # Custom validation logic
│
├── configs/            # Configuration files (YAML, JSON, etc.)
│   └── config.yaml
│
├── migrations/         # Database migration scripts (e.g., SQL files)
├── scripts/            # Helper scripts (Build, deploy, etc.)
├── test/               # Additional test suites
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## Architecture

This application follows the principle of Separation of Concerns:

1.  **Domain Layer**: Contains business entities and repository interfaces. This layer has no dependencies on other layers.
2.  **UseCase Layer**: Contains application-specific business logic. This layer orchestrates the flow of data to and from entities.
3.  **Delivery Layer**: Handles external inputs (such as HTTP Requests) and returns outputs.
4.  **Infrastructure Layer**: Implements technical details such as databases (PostgreSQL, MySQL), loggers, and other external services.

## Setup & Usage

### Prerequisites

- Go 1.22 or newer
- Database (depending on the implementation in `infrastructure/database`)

### 📚 API Documentation

This project uses **Swagger** (via `swaggo`) to automatically generate and serve interactive API documentation.

-   **Swagger UI**: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

To regenerate the documentation after adding new annotations, run:
```bash
swag init -g cmd/api/main.go
```

### 🛠️ Running the Application

1.  Prepare the configuration in `configs/config.yaml`.
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the application with hot reloading (Recommendation):
    - Install Air: `go install github.com/air-verse/air@latest`
    - Run: `air`
    
4.  Or run normally:
    ```bash
    go run cmd/api/main.go
    ```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
