# Task Management API

A backend REST API for a task management system, built in Go with [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/), and PostgreSQL. It provides user registration/authentication, JWT-based access and refresh tokens, role-based access control (RBAC), and CRUD operations for tasks.

> go udacity project #3

## Features

- **User accounts** — registration, login, and profile retrieval
- **JWT authentication** — short-lived access tokens plus rotating refresh tokens
- **Role-based access control** — roles (`admin`, `user`) map to fine-grained permissions (e.g. `task:create`, `user:view`)
- **Task CRUD** — create, read, update (partial), delete tasks scoped to the owning user
- **Middleware stack** — auth, RBAC, rate limiting, security headers, centralized error handling, access-denied logging
- **Swagger** — auto-generated via `swagger`, served at `/swagger/*any`
- **Dockerized** — multi-stage Dockerfile + `docker-compose.yml` for API, Postgres, and DB migrations

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Web framework | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) (`gorm.io/driver/postgres`) |
| Database | PostgreSQL 16 |
| Auth | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) (HS256), `bcrypt` for password hashing |
| Migrations | [golang-migrate](https://github.com/golang-migrate/migrate) |
| API docs | [swaggo/swag](https://github.com/swaggo/swag) + `gin-swagger` |
| IDs | `github.com/gofrs/uuid` |

## Data Model

- **User** — username, email, password (bcrypt hash), many-to-many with `Role`
- **Role** — e.g. `admin`, `user`; many-to-many with `Permission`
- **Permission** — `resource` + `action` pairs (e.g. `task` + `create`)
- **Task** — title, description, status (`pending` / `in_progress` / `completed`), owned by a `User`
- **Token** — stores active refresh tokens per user, with expiry, to support rotation

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose (recommended for local development)
- PostgreSQL 16 (if running outside Docker)

### Run with Docker Compose

```bash
docker compose up --build
```

This starts:
1. `postgres-db` — PostgreSQL database
2. `database-migrations` — runs all pending migrations against the database
3. `backend` — the API, available at `http://localhost:8080`

> **Note:** The `docker-compose.yml` in this repo contains placeholder credentials (`POSTGRES_PASSWORD`, `JWT_SECRET`) for local development only. Replace them with secrets from a secure store before deploying anywhere real.

#### Configuring secrets with `.env`
 
Create a `.env` file in the project root (same directory as `docker-compose.yml`). 
 
```
# .env
POSTGRES_USER=taskmanager
POSTGRES_PASSWORD=change_me_locally
POSTGRES_DB=taskmanager
JWT_SECRET=change_me_locally
JWT_EXPIRATION=3600s
```

### Run locally (without Docker)

```bash
# Install dependencies
go mod download

# Apply migrations (requires golang-migrate CLI)
migrate -path database-migrations/migrations \
  -database "postgres://<user>:<password>@localhost:5432/<db>?sslmode=disable" up

# Run the API
go run main.go
```

### Environment Variables

| Variable | Description | Default |
|---|---|---|
| `DB_HOST` | Postgres host | — |
| `DB_PORT` | Postgres port | — |
| `DB_USER` | Postgres user | — |
| `DB_PASSWORD` | Postgres password | — |
| `DB_NAME` | Postgres database name | — |
| `JWT_SECRET` | Secret used to sign JWTs | `""` |
| `JWT_EXPIRATION` | Access/refresh token lifetime  | `1h` |

## API Overview

Interactive docs available at `http://localhost:8080/swagger/index.html` once the server is running.

## Authentication & Authorization

1. On login/registration, the server issues a short-lived **access token** (JWT, HS256) embedding the user's ID, roles, and permissions, plus a long-lived **refresh token** (UUID) persisted in the `tokens` table.
2. Protected routes are guarded by `RequireAuth`, which parses and validates the JWT and populates the request context with `userID`, `userRoles`, and `userPermissions`.
3. `RequireRole` / `RequirePermission` middleware then authorize the request against those context values.
4. `POST /auth/refresh` rotates tokens: it validates the presented refresh token, deletes it, and issues a new access/refresh pair.

## Database Migrations

Located in `database-migrations/migrations`, managed with `golang-migrate`. Each migration has an `up` and `down` file:

1. `users` table
2. `tokens` table (refresh tokens)
3. `roles` table
4. `permissions` table
5. `user_roles` (join table)
6. `role_permissions` (join table)
7. Seed data — default `admin`/`user` roles, permissions, and sample users
8. `tasks` table (with seed sample tasks)

## Security Notes

- Passwords are hashed with `bcrypt` before storage.
- `SecurityHeaders` middleware sets `X-Content-Type-Options`, `X-Frame-Options`, `X-XSS-Protection`, HSTS, and `Referrer-Policy` headers on every response.
- `RateLimiter` middleware throttles requests per client IP (token bucket, 1 req/sec with a burst of 3).