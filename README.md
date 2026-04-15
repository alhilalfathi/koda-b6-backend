# ☕ Koda B6 — Coffee Shop Backend

A RESTful API backend for a **Coffee Shop E-Commerce** application, built with **Go (Golang)**. This service handles product management, user authentication, and order processing.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Framework | Gin-Gonic |
| Database | PostgreSQL (via `pgx/v5`) |
| Cache | Redis (`go-redis/v9`) |
| Auth | JWT (`golang-jwt/jwt v5`) |
| Password Hashing | Argon2 (`matthewhartstonge/argon2`) |
| API Docs | Swagger (`swaggo/swag`) |
| Config | godotenv |
| Container | Docker |

---

## Project Structure

```
koda-b6-backend/
├── cmd/                  # Application entry point
│   └── main.go
├── internal/             # Core application logic
│   ├── handler/          # HTTP handlers
│   ├── repository/       # Database queries
│   ├── service/          # Business logic
│   └── middleware/       # Auth & other middleware
├── migrations/           # Database migration files
├── uploads/
│   └── products/         # Uploaded product images
├── docs/                 # Swagger generated docs
├── .github/workflows/    # CI/CD pipelines
├── Dockerfile
├── go.mod
├── go.sum
├── rest.http             # HTTP request examples
└── .gitignore
```

---

## Prerequisites

Make sure the following are installed on your machine:

- [Go](https://go.dev/dl/) >= 1.21
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/) *(optional, for containerized setup)*

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/alhilalfathi/koda-b6-backend.git
cd koda-b6-backend
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory based on the following template:

```env
# Server
APP_PORT=8080

# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=koda_coffee

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRED=24h
```

### 3. Run Database Migrations

```bash
# Apply all migrations
psql -U postgres -d koda_coffee -f migrations/<migration_file>.sql
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Application

```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`.

---

## Running with Docker

### Build & Run

```bash
docker build -t koda-b6-backend .
docker run -p 8080:8080 --env-file .env koda-b6-backend
```

---

## API Documentation

Once the server is running, Swagger UI is accessible at:

```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger docs after changes:

```bash
swag init -g cmd/main.go
```

---

## Authentication

This API uses **JWT (JSON Web Token)** for authentication. Include the token in the `Authorization` header:

```
Authorization: Bearer <your_token>
```

---

## Key Dependencies

| Package | Purpose |
|---|---|
| `gin-gonic/gin` | HTTP web framework |
| `golang-jwt/jwt/v5` | JWT authentication |
| `jackc/pgx/v5` | PostgreSQL driver |
| `joho/godotenv` | `.env` file loader |
| `redis/go-redis/v9` | Redis client |
| `matthewhartstonge/argon2` | Password hashing |
| `swaggo/swag` | Swagger API docs generator |
| `google/uuid` | UUID generation |

---

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m 'feat: add some feature'`
4. Push to the branch: `git push origin feature/your-feature`
5. Open a Pull Request

---

## License

This project is open source and available under the [MIT License](LICENSE).

---

