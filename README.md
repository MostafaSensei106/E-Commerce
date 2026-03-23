<h1 align="center">E-Commerce</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/E-Commerce/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="E-Commerce Banner">
</p>

<p align="center">
  <strong>E-Commerce backend built in Go.</strong><br>
  Fast. Scalable. Robust. Built with Clean Architecture.
</p>

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#technologies">Technologies</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

---

## About

Welcome to **E-Commerce** — backend built with Go, leveraging the power of `chi` for routing and `pgx` for high-performance PostgreSQL interaction. 
E-Commerce empowers developers with a clean, modular architecture that separates business logic from data access. Whether you're building a small shop or a large-scale marketplace, it provides the solid foundation you need for product management, order processing, and more.

---

## Features

### 🌟 Core Functionality

- **High-Performance Engine**: Optimized for high throughput and low latency.
- **Product Management**: Full CRUD operations for managing a diverse product catalog.
- **Order Processing**: Efficient handling of customer orders with database transactions.
- **Clean Architecture**: Decoupled layers (Adapters, Services, Handlers) for maximum maintainability.
- **PostgreSQL Power**: Native performance with `pgx` and type-safe queries via `sqlc`.

### 🛠️ Advanced Capabilities

- **Modular Design**: Easily extendable components for adding new features like payments or shipping.
- **Logging**: Structured logging for better observability and debugging.
- **Docker Ready**: Fully containerized for seamless deployment and development.
- **Health Checks**: Built-in monitoring endpoint for system status.

### 🛡️ Security & Reliability

- **Graceful Error Handling**: Robust recovery from panics and clear error reporting.
- **Input Validation**: Ensures data integrity before processing.
- **Database Migrations**: Versioned schema changes for reliable deployments.

---

## Installation

### ⚠️ Prerequisites

- **Go 1.24+**: For compiling the application.
- **PostgreSQL 16+**: The primary database.
- **Docker & Docker Compose**: For containerized environments.

---

## 📦 Easy Setup (Docker)

Ensure you have Docker and Docker Compose installed.

```bash
# Clone the repository
git clone https://github.com/MostafaSensei106/E-Commerce.git
cd E-Commerce

# Build and start the services
docker-compose up --build
```

---

## 🏗️ Build from Source

Ensure you have `Go`, `git`, and `PostgreSQL` installed first.

```bash
# Clone the repository
git clone --depth 1 https://github.com/MostafaSensei106/E-Commerce.git
cd E-Commerce

# Download dependencies
go mod download

# Build the binary
go build -o e-commerce .
```

---

## 🚀 Quick Start

```bash
# Start the API server
./e-commerce
```

### 📋 API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/health` | Check service health |
| `GET` | `/products` | List all products |
| `POST` | `/products` | Create a new product |
| `GET` | `/products/{id}` | Get product details |
| `POST` | `/orders` | Place a new order |
| `GET` | `/orders` | List all orders |

---

## Configuration

The application uses environment variables for configuration. You can set them in a `.env` file or directly in your shell.

### 🧾 Example Environment Variables:

```bash
SENSEI_E_COMMERCE_DB_DSN="host=localhost user=root password=root dbname=ecommerce sslmode=disable"
PORT=":8080"
```

---

## Technologies

| Technology          | Description                                                                                                 |
| ------------------- | ----------------------------------------------------------------------------------------------------------- |
| 🧠 **Golang**       | [go.dev](https://go.dev) — The core language powering the backend: fast and efficient                      |
| 🚀 **Chi**          | [go-chi/chi](https://github.com/go-chi/chi) — Lightweight, idiomatic, and composable router for Go          |
| 🗄️ **PostgreSQL**   | [postgresql.org](https://www.postgresql.org/) — World's most advanced open source relational database       |
| ⚡ **pgx**           | [jackc/pgx](https://github.com/jackc/pgx) — Pure Go driver and toolkit for PostgreSQL                      |
| 🛠️ **SQLC**         | [sqlc.dev](https://sqlc.dev) — Generate type-safe code from SQL                                            |
| 🐳 **Docker**       | [docker.com](https://www.docker.com/) — Containerization for consistent environments                       |

---

## Contributing

Contributions are welcome! Here’s how to get started:

1.  Fork the repository.
2.  Create a new branch:
    `git checkout -b feature/YourFeature`
3.  Commit your changes:
    `git commit -m "Add amazing feature"`
4.  Push to your branch:
    `git push origin feature/YourFeature`
5.  Open a pull request.

---

## License

This project is licensed under the **MIT License**.
See the [LICENSE](LICENSE) file for full details.

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>
