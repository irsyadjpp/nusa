# NUSA - Education Operating System

![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)
[![License](https://img.shields.io/badge/license-Proprietary-red.svg)](LICENSE)
![React](https://img.shields.io/badge/React-18%2B-blue.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5%2B-blue.svg)
![Node](https://img.shields.io/badge/Node-18%2B-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-blue.svg)
![Podman](https://img.shields.io/badge/Podman-4%2B-informational.svg)
[![Development Status](https://img.shields.io/badge/status-Development-yellow.svg)](https://github.com/{username}/nusa)
![Microservices](https://img.shields.io/badge/architecture-Microservices-orange.svg)
![REST API](https://img.shields.io/badge/API-REST-green.svg)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3%2B-orange.svg)
![Redis](https://img.shields.io/badge/Redis-7%2B-red.svg)
![Education](https://img.shields.io/badge/domain-Education-blue.svg)
![Indonesia](https://img.shields.io/badge/country-Indonesia-red.svg)

A comprehensive education management system designed for Indonesian schools, featuring curriculum planning, learning design, assessment tools, and administrative workflows.

## 🎯 Overview

NUSA (National Unified System Administration) is a modern full-stack web application that streamlines educational processes for teachers, administrators, and students in Indonesia. Built with a focus on the Kurikulum Merdeka implementation, it provides tools for curriculum planning, teaching material development, assessment management, and school administration.

## ✨ Features

### 📚 Curriculum Management
- **Curriculum Plan (CP)** - Create and manage curriculum plans
- **Teaching Plan (TP)** - Detailed teaching plans and lesson preparation
- **Annual Teaching Plan (ATP)** - Yearly teaching schedules and planning

### 🎓 Learning Design
- **Modul Ajar** - Teaching modules and learning materials
- **Resource Management** - Educational content and resources

### 📊 Assessment & Evaluation
- **Assessment Tools** - Create and manage various assessment types
- **Rubrics** - Define and use assessment rubrics
- **Narrative Reports** - Generate narrative student reports
- **Evaluation Tracking** - Track and manage evaluation revisions

### 👥 User Management
- **Role-based Access Control** - Admin, Teacher, Principal roles
- **School Management** - Multi-school support
- **User Authentication** - Secure JWT-based authentication

### 🔧 Administrative Features
- **Settings Management** - System configuration
- **Approval Workflows** - Content approval processes
- **Reporting** - Comprehensive reporting tools

## 🏗️ Architecture

### Tech Stack

#### Backend
- **Language**: Go 1.25
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL 18
- **ORM**: sqlx with pgxpool
- **Authentication**: JWT with refresh tokens
- **Message Queue**: RabbitMQ
- **Cache**: Redis
- **Vector Database**: Qdrant (for AI features)
- **Object Storage**: MinIO
- **Container**: Podman

#### Frontend
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite 7
- **UI Library**: Material-UI (MUI)
- **Styling**: Tailwind CSS
- **State Management**: React Context
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
│   │   ├── auth/          # Authentication logic
│   │   ├── config/        # Configuration management
│   │   ├── db/            # Database connections
│   │   ├── domain/        # Domain models
│   │   ├── logger/        # Logging utilities
│   │   ├── middleware/    # HTTP middleware
│   │   ├── repository/    # Data access layer
│   │   ├── router/        # Route definitions
│   │   ├── server/        # Server setup
│   │   └── service/       # Business logic
│   ├── modules/           # Feature modules
│   │   ├── auth/          # Authentication handlers
│   │   ├── users/         # User management
│   │   ├── curriculum/    # Curriculum features
│   │   ├── assessment/    # Assessment features
│   │   └── reporting/     # Reporting features
│   ├── migrations/        # Database migrations
│   └── services/          # External service integrations
├── frontend/              # React frontend application
│   ├── src/
│   │   ├── api/           # API client
│   │   ├── components/    # Reusable components
│   │   ├── features/      # Feature modules
│   │   ├── pages/         # Page components
│   │   ├── theme/         # Theme configuration
│   │   ├── types/         # TypeScript types
│   │   ├── App.tsx        # Root component
│   │   ├── config.ts      # App configuration
│   │   ├── main.tsx       # Entry point
│   │   ├── menu-items.tsx # Navigation menu
│   │   └── routes.tsx     # Route definitions
│   ├── package.json
│   └── vite.config.ts
├── podman-compose.yml     # Container orchestration
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- **Podman** (or Docker) for container management
- **Go** 1.25+ for backend development
- **Node.js** 18+ and npm for frontend development
- **PostgreSQL** client tools (for database management)

### Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd nusa
```

2. **Start infrastructure services**
```bash
podman-compose -f podman-compose.yml up -d
```

This starts:
- PostgreSQL (port 5432)
- RabbitMQ (port 5672)
- Redis (port 6379)
- Qdrant (port 6333)
- MinIO (port 9000)

3. **Build and start backend**
```bash
cd backend
podman build -t nusa-backend .
podman run -d --name nusa-backend --network nusa_network \
  -e APP_ENV=development \
  -e SERVER_PORT=:8080 \
  -e DB_HOST=10.89.0.2 \
  -e DB_PORT=5432 \
  -e DB_NAME=nusa_db \
  -e DB_USER=nusa_user \
  -e DB_PASSWORD=nusa_password \
  -e DB_SSLMODE=disable \
  -e JWT_SECRET=your-secret-key-change-in-production \
  -e JWT_EXPIRATION=24h \
  -e RABBITMQ_HOST=rabbitmq \
  -e RABBITMQ_PORT=5672 \
  -e RABBITMQ_USER=nusa_user \
  -e RABBITMQ_PASSWORD=nusa_password \
  -e RABBITMQ_QUEUE=ai_generation \
  -e AI_PRIMARY_PROVIDER=openai \
  -e AI_OPENAI_KEY=your-ai-api-key \
  -e REDIS_HOST=redis \
  -e REDIS_PORT=6379 \
  -e REDIS_PASSWORD= \
  -e MINIO_ENDPOINT=http://minio:9000 \
  -e MINIO_ACCESS_KEY=admin \
  -e MINIO_SECRET_KEY=admin123 \
  -e MINIO_BUCKET=nusa-documents \
  -e MINIO_REGION=us-east-1 \
  -e MINIO_SECURE=false \
  -p 8081:8080 \
  nusa-backend
```

4. **Run database migrations**
```bash
# Copy migration files to PostgreSQL container
podman cp backend/migrations/000001_init_schema.up.sql nusa-postgres:/tmp/migration.sql
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /tmp/migration.sql

# Run subsequent migrations in order
podman cp backend/migrations/000002_add_education_domain_tables.up.sql nusa-postgres:/tmp/migration2.sql
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /tmp/migration2.sql

# Continue with remaining migrations...
```

5. **Install frontend dependencies**
```bash
cd frontend
npm install
```

6. **Configure frontend environment**
```bash
# Create .env file
echo "VITE_API_BASE_URL=http://localhost:8081" > .env
```

7. **Start frontend development server**
```bash
npm run dev
```

The frontend will be available at `http://localhost:3001`

### Default Credentials

After setup, create an admin user in the database:

```sql
-- First create admin role
INSERT INTO roles (name, description) VALUES ('admin', 'Administrator role');

-- Create admin user (password: admin123)
INSERT INTO users (email, password_hash, name, role_id) 
VALUES ('admin@nusa.id', '$2a$10$jjqseoG0NY6SzveYfGVQ4ejcQ7kkzJfEi63TtsT1eEEO13rpTBLwu', 'Admin Nusa', 
  (SELECT id FROM roles WHERE name = 'admin'));
```

**Default Login:**
- Email: `admin@nusa.id`
- Password: `admin123`

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
| `DB_SSLMODE` | SSL mode | `disable` |
| `JWT_SECRET` | JWT signing secret | - |
| `JWT_EXPIRATION` | Token expiration | `24h` |
| `RABBITMQ_HOST` | RabbitMQ host | `rabbitmq` |
| `RABBITMQ_PORT` | RabbitMQ port | `5672` |
| `RABBITMQ_USER` | RabbitMQ user | `nusa_user` |
| `RABBITMQ_PASSWORD` | RabbitMQ password | `nusa_password` |
| `REDIS_HOST` | Redis host | `redis` |
| `REDIS_PORT` | Redis port | `6379` |
| `MINIO_ENDPOINT` | MinIO endpoint | `http://minio:9000` |
| `MINIO_ACCESS_KEY` | MinIO access key | `admin` |
| `MINIO_SECRET_KEY` | MinIO secret key | `admin123` |

### Frontend Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `VITE_API_BASE_URL` | Backend API URL | `http://localhost:8082` |

## 📦 API Documentation

The backend provides RESTful APIs for all features. After starting the backend, API documentation is typically available at:

- Swagger UI: `http://localhost:8081/swagger/index.html` (if configured)
- API endpoints: `http://localhost:8081/api/v1/`

### Main API Endpoints

- **Authentication**: `/api/v1/public/auth/*`
- **Users**: `/api/v1/users/*`
- **Curriculum**: `/api/v1/curriculum/*`
- **Assessment**: `/api/v1/assessment/*`
- **Reporting**: `/api/v1/reporting/*`

## 🛠️ Development

### Backend Development

```bash
cd backend

# Run tests
go test ./...

# Build locally
go build -o bin/api cmd/api/main.go

# Run locally
./bin/api
```

### Frontend Development

```bash
cd frontend

# Run development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

### Database Migrations

```bash
# Create new migration
# (Manually create files in backend/migrations/)

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

# Frontend logs (if containerized)
podman logs nusa-frontend
```

### Restart services
```bash
podman restart nusa-backend
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

## 🔒 Security

- **Authentication**: JWT-based with refresh tokens
- **Authorization**: Role-based access control (RBAC)
- **Password Security**: bcrypt hashing
- **CORS**: Configurable CORS policies
- **SQL Injection**: Parameterized queries
- **XSS Protection**: Input sanitization

## 📝 Deployment

### Production Considerations

1. **Change default passwords and secrets**
2. **Enable SSL/TLS for database connections**
3. **Configure proper CORS origins**
4. **Set up monitoring and logging**
5. **Configure backup strategy**
6. **Use environment-specific configurations**
7. **Enable rate limiting**
8. **Set up proper error handling**

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

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is proprietary software. All rights reserved.

## 👥 Team

- **Development Team**: NUSA Development Team
- **Project Management**: Education Technology Team

## 📞 Support

For support, please contact:
- Email: support@nusa.id
- Documentation: [Internal Documentation Portal]

## 🔗 Resources

- [Kurikulum Merdeka Documentation](https://kurikulum.kemdikbud.go.id/)
- [Gin Framework](https://gin-gonic.com/)
- [React Documentation](https://react.dev/)
- [Material-UI](https://mui.com/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

---

**NUSA** - Empowering Education through Technology
