# NUSA Backend API

Backend foundation for NUSA Education Operating System MVP.

## Prerequisites

- Go 1.25+
- Docker and Docker Compose
- PostgreSQL 15+
- RabbitMQ 3.12+

## Quick Start

### Using Docker Compose (Recommended)

```bash
# Start all services (PostgreSQL, RabbitMQ, Backend)
make docker-up

# Run database migrations
make migrate-up

# View logs
docker-compose logs -f backend
```

The API will be available at `http://localhost:8080`

### Local Development

```bash
# Install dependencies
make install-deps

# Set up environment variables (copy .env.example to .env)
cp .env.example .env
# Edit .env with your configuration

# Start PostgreSQL and RabbitMQ
make docker-up

# Run database migrations
make migrate-up

# Run the application
make run
```

## Available Endpoints

- `GET /health` - Health check endpoint
- `GET /ready` - Readiness check endpoint
- `GET /version` - Version information endpoint

## Makefile Commands

```bash
make build          # Build the application
make run            # Run the application
make test           # Run tests
make lint           # Run linter
make migrate-up     # Run database migrations up
make migrate-down   # Run database migrations down
make migrate-create # Create a new migration (usage: make migrate-create NAME=migration_name)
make generate       # Generate code from sqlc
make clean          # Clean build artifacts
make install-deps   # Install dependencies
make docker-up      # Start Docker Compose services
make docker-down    # Stop Docker Compose services
```

## Configuration

Configuration is loaded from environment variables. See `.env.example` for all available options.

Required environment variables:
- `DB_PASSWORD` - Database password
- `JWT_SECRET` - JWT secret key (must be set in production)

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── bootstrap/            # Dependency injection and app initialization
│   ├── config/               # Configuration management
│   ├── db/                   # Database layer (pgx)
│   ├── error/                # Error handling and response helpers
│   ├── logger/               # Structured logging (zap)
│   ├── middleware/           # HTTP middleware
│   └── server/               # HTTP server (Gin)
├── migrations/              # Database migrations
├── tests/                   # Unit and integration tests
├── Dockerfile               # Docker image build
├── docker-compose.yml       # Docker Compose configuration
├── Makefile                 # Build and run commands
└── go.mod                  # Go dependencies
```

## Testing

```bash
# Run all tests
make test

# Run unit tests only
go test -v ./tests/unit/...

# Run integration tests (requires running database)
go test -v ./tests/integration/...
```

## Database Migrations

```bash
# Run migrations up
make migrate-up

# Run migrations down
make migrate-down

# Create a new migration
make migrate-create NAME=add_new_table
```
