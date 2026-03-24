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
- **Dependency Injection (Registry Pattern)**: A custom manual DI container located in `internal/infrastructure/container` to keep `main.go` clean and manageable.

## Folder Structure

```text
gia-starter-app-V1/
│
├── cmd/                # Application entry point
│   └── api/
│       └── main.go     # Main entry point for the API service
│
├── internal/           # Private application code (cannot be imported by other systems)
│   ├── modules/        # Feature-based Modules (Self-contained)
│   │   └── user/       # Example Module: User
│   │       ├── delivery/      # Transport Layer (HTTP handlers)
│   │       ├── domain/        # Business Logic Core (Entities & Repo Interfaces)
│   │       ├── usecase/       # Application Business Rules
│   │       └── infrastructure/ # Module-specific technical details (Persistence)
│   │
│   ├── infrastructure/ # Global technical details (Shared implementations)
│   │   ├── database/    # Global Database connections
│   │   ├── container/   # Dependency Injection Registry (Registry Pattern)
│   │   ├── config/      # Global Application settings
│   │   └── logger/      # Global Logging system
│   │
│   ├── delivery/       # Global Transport Layer (Shared middleware, router)
│   │   └── http/
│   │       ├── middleware/ # Shared HTTP Middleware
│   │       └── router.go   # API route definitions
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

This application follows the **Modular Clean Architecture** (Feature-Oriented) pattern. This approach groups related business logic into independent modules, making it highly scalable and maintainable.

1.  **Modularization**: Code is organized by **feature/domain** (e.g., `user`, `product`). Each module is a self-contained unit containing its own layers.
2.  **Clean Layers inside Modules**:
    - **Domain**: Business entities and repository interfaces.
    - **UseCase**: Application-specific business logic and interactors.
    - **Delivery**: Handles external inputs (HTTP handlers).
    - **Infrastructure**: Module-specific technical details like persistence (Postgres, etc).
3.  **Global Infrastructure**: Shared technical details like database connections, configurations, logging, and common utilities reside outside the modules to be reused across the application.

## Setup & Usage

### Prerequisites

- Go 1.22 or newer
- Database (depending on the implementation in `infrastructure/database`)

### 📚 API Documentation

This project uses **Swagger** (via `swaggo`) to automatically generate and serve interactive API documentation.

- **Swagger UI**: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

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

### 🗄️ Database Migrations

This project uses `sql-migrate` for database schema management, which tracks migrations by their filenames in the database. You can use `Makefile` commands to simplify the process.

#### Basic Commands

| Command                     | Description                                                      |
| :-------------------------- | :--------------------------------------------------------------- |
| `make migrate-up`           | Run all pending migrations.                                      |
| `make migrate-down`         | Undo the last migration.                                         |
| `make migrate-status`       | View the status of migrations and which files have been applied. |
| `make migrate-new name=xxx` | Create a new migration file.                                     |

#### Example: Creating a New Table (`users`)

1.  **Generate migration file**:
    ```bash
    make migrate-new name=create_users_table
    ```
2.  **Define the SQL schema** in the new file under `migrations/`. Use markers for Up and Down:

    ```sql
    -- +migrate Up
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL
    );

    -- +migrate Down
    DROP TABLE users;
    ```

3.  **Apply the migration**:
    ```bash
    make migrate-up
    ```

#### Benefits of Name-Based Tracking

Unlike version-only tools, `sql-migrate` stores the filename in the database (table `gorp_migrations`). This makes it easy to audit which specific migration scripts have been executed.

### 🏗️ Adding a New Module

To add a new functional module (e.g., `Product`), follow these steps to ensure compliance with Clean Architecture:

1.  **Database Migration**:

    ```bash
    make migrate-new name=create_products_table
    ```

    Define your schema in `migrations/xxxx_create_products_table.sql`.

2.  **Domain Layer**:
    - Create the entity in `internal/modules/product/domain/entity/product.go`.
    - Define the repository interface in `internal/modules/product/domain/repository/product_repository.go`.

3.  **UseCase Layer**:
    - Implement business logic in `internal/modules/product/usecase/product_usecase.go`.

4.  **Infrastructure Layer (Persistence)**:
    - Implement the repository in `internal/modules/product/infrastructure/persistence/postgres/product_repository.go`.

5.  **Delivery Layer (HTTP)**:
    - Create the request handler in `internal/modules/product/delivery/http/handler/product_handler.go`.

6.  **Dependency Injection**:
    - Register the new module components in `internal/infrastructure/container/registry.go`.

    ```go
    // internal/infrastructure/container/registry.go
    func NewRegistry(db *gorm.DB) *Registry {
        productRepo := postgres.NewProductRepository(db)
        productUC := usecase.NewProductUseCase(productRepo)
        // ...
    }
    ```

7.  **Routing**:
    - Register the routes in `internal/delivery/http/router.go`.

    ```go
    // internal/delivery/http/router.go
    productHandler := handler.NewProductHandler(reg.ProductUseCase)
    products := v1.Group("/products")
    {
        products.POST("", productHandler.CreateProduct)
        // ...
    }
    ```

8.  **API Documentation**:
    - Add Swagger annotations to your handler methods and regenerate the docs:
    ```bash
    swag init -g cmd/api/main.go
    ```

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
