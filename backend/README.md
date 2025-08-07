# byFood Case - Backend

This repository contains the backend implementation for the byFood case study. It's a RESTful API built using Go (Fiber), GORM (MySQL), and Swagger for API documentation.

---

## Project Structure

```
backend/
├── cmd/app                 # Application entrypoint
├── internal/
│   ├── dto                 # Request/response DTOs
│   ├── handler             # HTTP handlers
│   ├── model               # DB models
│   ├── repository          # Data access layer
│   ├── service             # Business logic
│   ├── routes              # Route definitions
│   └── mocks               # Auto-generated mocks
├── pkg/                    # Shared packages (logger, validator, config, middleware)
├── docs/                   # Swagger documentation
├── Dockerfile              # Docker setup
├── go.mod / go.sum         # Go dependencies
```

---

## Getting Started

### Prerequisites

- Go `v1.23.11` or higher
- MySQL
- Swagger UI
- (Optional) Docker + Docker Compose

### Setup

```bash
git clone https://github.com/yusufbulac/byfood-case.git
cd backend
go mod tidy
```

### Run Locally

```bash
go run cmd/app/main.go
```

Server starts at: `http://localhost:8080`

---

## API Documentation

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Docs generated via `swag init` (`docs/` directory)

---

## Running Tests

```bash
go test ./...
go test -cover ./...
```

- BookService and UrlService are covered with unit tests.
- `coverage: 91.9%` in service layer.

---

## Example API Usage

### Create Book

```http
POST /api/v1/books
Content-Type: application/json

{
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "year": 2008
}
```

### Process URL

```http
POST /api/v1/url/transform
Content-Type: application/json

{
  "url": "http://example.com/path?query=1",
  "operation": "canonical"
}
```

---

## Project Notes

- Follows layered architecture (handler → service → repository)
- Unit tests written with Testify + Mockery
- API documented via Swagger annotations

---

## Repo Layout Recommendation

Since this repo contains **multiple projects**, the best practice is:

- One `README.md` per subproject (`/backend`, `/frontend` etc.)
- One root `README.md` to explain the overall case and link to each project

**For this backend**, this README should be saved at: `backend/README.md`
