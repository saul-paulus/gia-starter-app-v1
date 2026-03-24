# **GIA Starter App - Clean Architecture**

This project is a backend application built using the Go (Golang) programming language with the Gin Gonic framework, following the **Clean Architecture** (or Hexagonal Architecture) pattern. This structure is designed to decouple business logic from technical details (such as frameworks, databases, etc.) to enhance scalability and ease of testing.

### 🚀 Tech Stack & Usage

This project leverages modern Go tools and libraries to ensure performance, maintainability, and a great developer experience:

- **[Gin Gonic](https://gin-gonic.com/)** (Web Framework): Handles all HTTP routing, request parsing, and middleware. It's the core of the `delivery` layer.
- **[GORM](https://gorm.io/)** (Database ORM): Manages database interactions using an Object-Relational Mapping approach. Used in the `infrastructure/persistence` layer.
- **[sql-migrate](https://github.com/rubenv/sql-migrate)** (Migrations): Version controls the database schema. It tracks which SQL scripts have been applied by their filenames.
- **[Viper](https://github.com/spf13/viper)** (Configuration): A complete configuration solution. It loads settings from `configs/config.yaml` and handles environment variable overrides.
- **[Uber Zap](https://github.com/uber-go/zap)** (Logging): A blazing fast, structured logger. Configured to output logs to both the console and files in `storage/logs`.
- **[Swagger](https://github.com/swaggo/swag)** (Documentation): Automatically generates API documentation from code annotations, accessible via a web UI.
- **[Air](https://github.com/cosmtrek/air)** (Hot Reloading): Watches for file changes during development and automatically rebuilds/restarts the application.
- **[Go Playground Validator](https://github.com/go-playground/validator)** (Validation): Ensures that incoming request data (JSON) is valid and meets business requirements before processing.
- **Dependency Injection (Modular Pattern)**: Each module handles its own dependency injection in `module.go`, and all are aggregated in `internal/bootstrap/bootstrap.go`.

## Folder Structure

```text
gia-starter-app-V1/
│
├── cmd/                # Application entry point
│   └── api/
│       └── main.go     # Minimal entry point (calls bootstrap)
│
├── internal/           # Private application code
│   ├── bootstrap/      # Centralized App & Module registration
│   │   └── bootstrap.go
│   │
│   ├── modules/        # Feature-based Modules (Self-contained)
│   │   └── user/       # Example Module: User
│   │       ├── domain/        # Entities & Repository Interfaces
│   │       ├── usecase/       # Business Logic
│   │       ├── interface/     # Adapters (HTTPHandlers, Repo Impls)
│   │       │   ├── http/      # Handlers & Module Router
│   │       │   └── repository/# GORM Repository Implementation
│   │       └── module.go      # Module setup & DI
│   │
│   ├── shared/         # Shared Kernel (Cross-cutting concerns)
│   │   ├── domain/model/  # Base Models
│   │   ├── response/      # API Response standardization
│   │   ├── middleware/    # Global HTTP Middleware
│   │   ├── errors/        # Custom error definitions
│   │   └── util/          # Utility functions
│   │
│   ├── delivery/       # Global Transport Layer
│   │   └── http/
│   │       └── router.go   # Main router (aggregates module routes)
│   │
│   └── infrastructure/ # Global technical implementations
│       ├── database/    # Database connections (Postgres)
│       ├── config/      # Configuration loader (Viper)
│       └── logger/      # Structured logging (Zap)
│
├── configs/            # Configuration files (YAML)
├── migrations/         # Database migration scripts
├── pkg/                # Public shared libraries
├── go.mod              # Module definition
└── README.md           # Documentation
```

## Architecture

This application follows the **Modular Clean Architecture**. This approach ensures that business logic is isolated and that modules remain decoupled.

1.  **Shared Kernel**: All cross-cutting concerns (Response, Errors, Middleware, Base Models) are placed in `internal/shared`.
2.  **Module Registration**: Each module has a `module.go` that defines its `Init` function, handling dependency injection internally.
3.  **Central Bootstrap**: `internal/bootstrap/bootstrap.go` is the orchestrator that initializes the database, all modules, and the main router.
4.  **Decoupled Routing**: Each module defines its own routes in `interface/http/router.go`, which are then aggregated by the main router in `delivery/http/router.go`.

## Setup & Usage

### 📚 API Documentation

This project uses **Swagger** (via `swaggo`).

- **Swagger UI**: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

To regenerate:
```bash
swag init -g cmd/api/main.go
```

### 🛠️ Running the Application

The application features **Resilient Configuration Loading**. It will automatically search for `configs/config.yaml` and `.env` in the current directory and parent directories.

1.  Prepare `configs/config.yaml` and `.env`.
2.  Install dependencies: `go mod tidy`
3.  Run from **any folder** (Project Root Recommended):
    ```bash
    go run cmd/api/main.go
    # OR using Air
    air
    ```

### 🏗️ Adding a New Module (e.g., `Product`)

1.  **Migration**: `make migrate-new name=create_products_table`.
2.  **Domain**: Create `modules/product/domain/entity` and repository interfaces.
3.  **UseCase**: Implement logic in `modules/product/usecase`.
4.  **Interface**:
    - Implement Repository in `modules/product/interface/repository`.
    - Implement Handlers and `RegisterRoutes` in `modules/product/interface/http`.
5.  **Module Init**: Create `modules/product/module.go` with an `Init` function.
6.  **Bootstrap**: Register the module in `internal/bootstrap/bootstrap.go`.
7.  **Routing**: Update `internal/delivery/http/router.go` to call `productHttp.RegisterRoutes`.

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
