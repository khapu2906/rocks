

# 🥃  Rocks Framework

**Rocks** is a modular, serverless-ready backend framework built with Go and Gin. It provides a solid foundation for building scalable, maintainable web applications using modern development practices.

> 📌 *Note: Replace this with your repository URL and update with a project-specific description if needed.*

---

## 📘 Introduction

Rocks is designed with a clean modular architecture that simplifies development, testing, and scaling. It emphasizes code separation, maintainability, and ease of extension for backend systems.

---

## 🏗️ Architecture Overview

The project is structured into well-defined components:

```
.
├── cmd/main.go           # Entry point of the application
├── internal/user         # User management module
├── pkg/core              # Core utilities and middleware
├── pkg/database          # Database setup and migration logic
├── config                # Configuration management
├── migrations            # SQL migration scripts
├── stubs                 # Definatin file to generate
```

---

## 🧰 Prerequisites

Ensure you have the following installed:

* **Go**: [https://go.dev/doc/install](https://go.dev/doc/install)
* **Docker**: [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/)
* **Docker Compose**: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

---

## ⚙️ Setup

1. **Clone the repository:**

   ```bash
   git clone git@github.com:khapu2906/rocks.git
   cd rocks
   ```

2. **Install dependencies:**

   ```bash
   make setup
   ```

3. **Create and configure the environment file:**

   ```bash
   cp .env.example .env
   ```

   Update the `.env` file:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_NAME=rocks_db
   DATABASE_URL=postgres://postgres:secret@localhost:5432/rocks_db?sslmode=disable
   # Add other required environment variables
   ```

---

## 🚀 Running the Application

1. **Install the migration tool:**

   ```bash
   make migrate-install
   ```

2. **Start the database:**

   ```bash
   docker-compose up -d db
   ```

3. **Run database migrations:**

   ```bash
   make migrate-up
   ```

4. **Start the app:**

   ```bash
   make run
   ```

---

## 🛠️ Building the Application

Compile the Go code:

```bash
make build
```

---

## 🐳 Docker Support

**Build Docker image:**

```bash
docker build -t rocks-app .
```

**Run the container:**

```bash
docker run -p 8080:8080 rocks-app
```

---

## 🧱 Module Generator

Generate a new module:

```bash
make generate-module name=[module_name]
```

This scaffolds a new module with the necessary boilerplate.

---

## 🧩 Key Features

* ✅ **Modular Architecture** — easy to extend, test, and maintain.
* ✅ **User Management** — includes registration and authentication.
* ✅ **Database Migration Support** — integrated with `migrate` CLI.
* ✅ **Environment-based Configuration** — flexible for all environments.
* ✅ **Docker Ready** — easy to run and deploy via Docker.

---

## 🧯 Troubleshooting

* **Database connection errors:**

  * Ensure Docker is running and `db` service is up.
  * Double-check credentials in `.env`.
* **Migration issues:**

  * Confirm database is available before running `make migrate-up`.
* **Startup errors:**

  * Review logs with `docker-compose logs app`.
  * Ensure all `.env` values are correctly set.

