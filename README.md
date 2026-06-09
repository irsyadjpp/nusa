# NUSA - Education Operating System

![Supported Go Versions](https://img.shields.io/badge/Go-1.21%2B-lightgrey.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
![React](https://img.shields.io/badge/React-18%2B-blue.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5%2B-blue.svg)
![Node](https://img.shields.io/badge/Node-20%2B-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-blue.svg)
![Podman](https://img.shields.io/badge/Podman-4%2B-informational.svg)
[![Development Status](https://img.shields.io/badge/status-Active--Development-yellow.svg)](https://github.com/sdibonerate85/nusa)
![Architecture](https://img.shields.io/badge/architecture-Modular--Monolith-orange.svg)
![REST API](https://img.shields.io/badge/API-REST-green.svg)
![DDD](https://img.shields.io/badge/domain-DDD--Lite-blue.svg)
![Education](https://img.shields.io/badge/domain-Education-blue.svg)
![Indonesia](https://img.shields.io/badge/country-Indonesia-red.svg)

A comprehensive education management system designed for Indonesian schools implementing Kurikulum Merdeka 2026. Built as a personal project with production-quality architecture, focusing on curriculum planning, learning design, assessment tools, and administrative workflows.

## 🎯 Overview

NUSA (National Unified System Administration) is a modern full-stack web application that streamlines educational processes for teachers, administrators, and students in Indonesia. This is a solo-developed project following Domain-Driven Design (DDD) Lite principles with a modular monolith architecture, specifically designed to support the Kurikulum Merdeka implementation.

**Current Status**: Active development - Sprint 3B-3C completed, Sprint 4 roadmap defined

## ✨ Features

### 📚 Curriculum Management
- **Curriculum Plan (CP)** - Create and manage curriculum plans
- **Teaching Plan (TP)** - Detailed teaching plans with embedded KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)
- **Annual Teaching Plan (ATP)** - Yearly teaching schedules and planning
- **TP Set Management** - Group and version teaching plans

### 🎓 Learning Design
- **Modul Ajar** - Teaching modules and learning materials
- **Resource Management** - Educational content and resources
- **AI-Assisted Content Generation** - AI-powered lesson plan creation

### 📊 Assessment & Evaluation
- **Assessment Tools** - Create and manage various assessment types
- **Rubrics** - Define and use assessment rubrics
- **Evidence Management** - Track student evidence and evaluations
- **Evaluation Revision Tracking** - Complete history of evaluation changes
- **Narrative Reports** - Generate narrative student reports
- **Achievement Dashboard** - Real-time competency progress tracking

### 🤖 AI Integration
- **AI Orchestration** - Structured AI workflows for content generation
- **Prompt Versioning** - Versioned AI prompts with approval checkpoints
- **Human-in-the-Loop** - AI assistance with teacher oversight

### 👥 User Management
- **Role-based Access Control** - Admin, Teacher, Principal roles
- **School Management** - Multi-school support with proper isolation
- **User Authentication** - Secure JWT-based authentication
- **Resource Authorization** - Multi-level permission system

### 🔧 Administrative Features
- **Settings Management** - System configuration
- **Approval Workflows** - Content approval processes
- **Version Management** - Unified versioning across all entities
- **Reporting** - Comprehensive reporting tools

## 🏗️ Architecture

### Design Philosophy
- **Architecture**: Modular Monolith with DDD Lite
- **Layer Pattern**: Handler → Service → Repository → PostgreSQL
- **No Complex Patterns**: No CQRS, Event Sourcing, or Event Bus (kept simple for maintainability)
- **Single Contributor**: Optimized for solo development with proper governance

### Tech Stack

#### Backend
- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL 18.4
- **ORM**: sqlx with pgxpool
- **Authentication**: JWT with refresh tokens
- **Message Queue**: RabbitMQ (for AI workflows)
- **Cache**: Redis
- **Vector Database**: Qdrant (for AI features)
- **Object Storage**: MinIO
- **Container**: Podman/Docker

#### Frontend
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite 7.3.5
- **UI Library**: Material-UI (MUI)
- **Styling**: Tailwind CSS
- **State Management**: TanStack Query + Zustand
- **Routing**: React Router v6
- **Form Handling**: Formik
- **HTTP Client**: Axios

### Project Structure

```
nusa/
├── backend/                 # Go backend application
│   ├── cmd/
│   │   └── api/
│   │       └── main.go     # Application entry point
│   ├── internal/           # Private application code
│   │   ├── application/   # Application services (use cases)
│   │   ├── domain/        # Domain models and value objects
│   │   ├── handler/       # HTTP handlers
│   │   ├── repository/    # Data access layer
│   │   ├── dto/           # Data transfer objects
│   │   └── database/      # Database connections
│   ├── migrations/        # Database migrations
│   └── docs/              # Backend documentation
├── frontend/              # React frontend application
│   ├── src/
│   │   ├── shared/        # Shared utilities and services
│   │   │   ├── services/  # API clients and state management
│   │   │   └── store/     # Zustand stores
│   │   ├── features/      # Feature-specific components
│   │   │   ├── tp/       # TP Workspace
│   │   │   ├── atp/      # ATP Workspace
│   │   │   ├── modul-ajar/ # Modul Ajar Workspace
│   │   │   ├── assessment/ # Assessment Designer
│   │   │   ├── evidence/  # Evidence Workspace
│   │   │   ├── evaluation/ # Evaluation Workspace
│   │   │   ├── achievement/ # Achievement Dashboard
│   │   │   └── report/    # Narrative Report Builder
│   │   ├── App.tsx        # Root component
│   │   └── main.tsx       # Entry point
│   ├── package.json
│   └── vite.config.ts
├── docs/                  # Project documentation
│   ├── centralized/      # Centralized documentation
│   └── *.md              # Various architecture docs
├── ai-runtime/           # AI integration service
├── podman-compose.yml    # Container orchestration
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- **Podman** (or Docker) for container management
- **Go** 1.21+ for backend development
- **Node.js** 20+ and npm for frontend development
- **PostgreSQL** client tools (for database management)

### Quick Start with Podman

1. **Clone the repository**
```bash
git clone https://github.com/sdibonerate85/nusa.git
cd nusa
```

2. **Start all infrastructure services**
```bash
podman-compose -f podman-compose.yml up -d
```

This starts:
- PostgreSQL (port 5432)
- RabbitMQ (port 5672)
- Redis (port 6379)
- Qdrant (port 6333)
- MinIO (port 9000)
- Backend API (port 8081)
- Frontend (port 3001)

3. **Run database migrations**
```bash
# Apply migrations
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /docker-entrypoint-initdb.d/000001_initial_schema.up.sql
```

4. **Access the application**
- Frontend: http://localhost:3001
- Backend API: http://localhost:8081
- MinIO Console: http://localhost:9001

### Default Credentials

- **Email**: admin@nusa.local
- **Password**: admin123
- **Role**: SYSTEM_ADMIN

## 🔧 Configuration

### Backend Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Application environment | `development` |
| `SERVER_PORT` | Server port | `:8080` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_NAME` | Database name | `nusa_db` |
| `DB_USER` | Database user | `nusa_user` |
| `DB_PASSWORD` | Database password | `nusa_password` |
| `JWT_SECRET` | JWT signing secret | `your-secret-key` |
| `JWT_EXPIRATION` | Token expiration | `24h` |
| `RABBITMQ_HOST` | RabbitMQ host | `rabbitmq` |
| `RABBITMQ_PORT` | RabbitMQ port | `5672` |
| `REDIS_HOST` | Redis host | `redis` |
| `REDIS_PORT` | Redis port | `6379` |
| `MINIO_ENDPOINT` | MinIO endpoint | `http://minio:9000` |

### Frontend Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `VITE_API_BASE_URL` | Backend API URL | `http://localhost:8081` |

## 📦 API Documentation

The backend provides RESTful APIs for all features. OpenAPI specification is available at:

- **OpenAPI Spec**: `/backend/docs/api/openapi.yaml`
- **API Endpoints**: `http://localhost:8081/api/v1/`

### Main API Endpoints

- **Authentication**: `/api/v1/public/auth/*`
- **Teaching Plans (TP)**: `/api/v1/tp/*`
- **Assessment**: `/api/v1/assessment/*`
- **Evidence**: `/api/v1/evidence/*`
- **Evaluation**: `/api/v1/evaluation/*`
- **Achievement**: `/api/v1/achievement/*`
- **Reports**: `/api/v1/reports/*`

## 🛠️ Development

### Backend Development

```bash
cd backend

# Run tests
go test ./... -v

# Build locally
go build -o bin/api cmd/api/main.go

# Run locally
./bin/api

# Format code
go fmt ./...

# Vet code
go vet ./...
```

### Frontend Development

```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Type check
npm run type-check

# Lint
npm run lint
```

### Database Migrations

```bash
# Apply migration
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /path/to/migration.up.sql

# Rollback migration
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /path/to/migration.down.sql
```

## 🐳 Container Management

### Start all services
```bash
podman-compose -f podman-compose.yml up -d
```

### Stop all services
```bash
podman-compose -f podman-compose.yml down
```

### View logs
```bash
# Backend logs
podman logs nusa-backend

# Database logs
podman logs nusa-postgres

# All services
podman-compose logs
```

### Restart services
```bash
podman-compose restart
```

## 🧪 Testing

### Backend Testing
```bash
cd backend
go test ./... -v
```

### Frontend Testing
```bash
cd frontend
npm test
```

## � Project Status

### Completed Sprints
- **Sprint 3A**: Domain implementation, backend API, database migration ✅
- **Sprint 3B**: Frontend implementation, workspace components ✅
- **Sprint 3C**: UAT validation, integration testing ✅

### Current Sprint
- **Sprint 3.5**: Resource authorization, unified versioning, evidence storage 🚧

### Planned
- **Sprint 4**: AI Copilot, Analytics, Student Progress Tracking 📋

See [CHANGELOG.md](CHANGELOG.md) for detailed history.

## �🔒 Security

- **Authentication**: JWT-based with refresh tokens
- **Authorization**: Role-based access control (RBAC) with school-level isolation
- **Password Security**: bcrypt hashing
- **CORS**: Configurable CORS policies
- **SQL Injection**: Parameterized queries via sqlx
- **XSS Protection**: Input sanitization
- **Multi-tenancy**: School-level data isolation

See [SECURITY.md](SECURITY.md) for security policies.

## 📝 Deployment

### Production Considerations

1. Change default passwords and secrets
2. Enable SSL/TLS for database connections
3. Configure proper CORS origins
4. Set up monitoring and logging
5. Configure backup strategy
6. Use environment-specific configurations
7. Enable rate limiting
8. Set up proper error handling

### Build for Production

```bash
# Backend
cd backend
CGO_ENABLED=0 GOOS=linux go build -o bin/api cmd/api/main.go

# Frontend
cd frontend
npm run build
```

## 🤝 Contributing

This is currently a solo-developed project, but contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## � Author

**Single Contributor Project**
- **Developer**: [Your Name]
- **Contact]: Via GitHub Issues
- **Location**: Indonesia

Built with passion for improving education technology in Indonesia. 🇮🇩

## 📞 Support

For support, questions, or discussions:
- **GitHub Issues**: [Create an issue](https://github.com/sdibonerate85/nusa/issues)
- **GitHub Discussions**: [Start a discussion](https://github.com/sdibonerate85/nusa/discussions)
- **Documentation**: See [docs/](docs/) folder

## 🙏 Acknowledgments

- **Kementerian Pendidikan, Kebudayaan, Riset, dan Teknologi (Kemendikbudristek)** for Kurikulum Merdeka
- **Open Source Community** for the amazing tools and libraries
- **AI Assistants** (Claude, Devin, etc.) that helped in development

## 🔗 Resources

- [Kurikulum Merdeka Documentation](https://kurikulum.kemdikbud.go.id/)
- [Gin Framework](https://gin-gonic.com/)
- [React Documentation](https://react.dev/)
- [Material-UI](https://mui.com/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Domain-Driven Design](https://martinfowler.com/tags/domain%20driven%20design.html)

---

**NUSA** - Empowering Education through Technology 🎓

*Built with ❤️ for Indonesian Education*
