# 09_REPOSITORY_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 16)

**Purpose**: Define the official repository structure for NUSA MVP Wave 1 implementation, serving as the single source of truth for repository strategy, folder structure, module boundaries, naming conventions, ownership rules, dependency rules, import rules, and CI/CD configuration. This document is the blueprint for all implementation work.

---

# Repository Principles

## Single Source of Truth

All code for NUSA MVP Wave 1 resides in this repository. No external repositories for MVP components. Shared contracts defined once, used everywhere.

## Module Boundaries

Backend modules follow capability boundaries. Frontend modules align with backend modules. Shared contracts enforce module boundaries. No direct cross-module dependencies allowed.

## Ownership Clarity

Each module has clear ownership. Ownership follows domain ownership. Cross-module changes require coordination.

## Dependency Management

Backend dependencies managed via Go modules. Frontend dependencies managed via npm. Shared dependencies versioned together. No circular dependencies allowed.

## MVP Scope Enforcement

Repository structure must support MVP Wave 1 only. Future domains (Competency Graph, Digital Twin, Lifelong Learning Record) must not introduce repository structure requirements.

---

# Monorepo Structure

## Root Structure

```
nusa/
├── backend/
├── frontend/
├── ai-runtime/
├── database/
│   ├── migrations/
│   ├── seed/
│   └── schema/
├── docs/
│   ├── foundation/
│   │   ├── 00A-16/
│   ├── adr/
│   ├── api/
│   ├── database/
│   └── prompt-specification/
├── prompts/
│   ├── tp/
│   ├── atp/
│   ├── modul-ajar/
│   ├── assessment/
│   ├── rubric/
│   └── report/
├── scripts/
├── deploy/
├── adr/
├── .github/
├── .gitignore
├── .env.example
├── README.md
└── LICENSE
```

## Directory Ownership

| Directory | Owner | Responsibility |
|----------|-------|----------------|
| backend/ | Backend Team | Go + Gin backend application |
| frontend/ | Frontend Team | React + TypeScript application |
| ai-runtime/ | AI Engineer | Python + LangGraph AI runtime |
| database/ | Backend Team + DevOps | PostgreSQL migrations and schema |
| docs/ | Architecture Team | Architecture documentation |
| prompts/ | AI Engineer | AI prompt specifications |
| scripts/ | DevOps | Development and deployment scripts |
| deploy/ | DevOps | Docker and deployment configurations |
| adr/ | Architecture Team | Architecture Decision Records |
| .github/ | DevOps | CI/CD workflows |

---

# Backend Structure

## Technology Stack

- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Message Broker**: RabbitMQ
- **Authentication**: JWT (Custom)

## Backend Folder Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
├── internal/
│   ├── config/
│   │   ├── config.go                  # Configuration loading
│   │   ├── database.go                # Database configuration
│   │   ├── rabbitmq.go                # RabbitMQ configuration
│   │   └── jwt.go                     # JWT configuration
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── curriculum/
│   │   │   ├── planning/
│   │   │   ├── assessment/
│   │   │   ├── reporting/
│   │   │   ├── administration/
│   │   │   └── ai/
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   ├── logging.go
│   │   │   └── error.go
│   │   └── routes/
│   │       └── routes.go
│   ├── modules/
│   │   ├── curriculum/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   ├── planning/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   ├── assessment/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   ├── reporting/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   ├── administration/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   ├── ai/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   ├── repositories/
│   │   │   │   └── services/
│   │   │   └── application/
│   │   │       └── usecases/
│   │   └── shared/
│   │       ├── domain/
│   │       │   ├── entities/
│   │       │   │   ├── user.go
│   │       │   │   └── audit.go
│   │       │   └── valueobjects/
│   │       └── infrastructure/
│   │           ├── persistence/
│   │           ├── messaging/
│   │           └── logging/
│   └── pkg/
│       ├── jwt/
│       ├── middleware/
│       └── utils/
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── go.mod
├── go.sum
└── Dockerfile
```

## Backend Module Ownership

| Module | Owner | Responsibility |
|--------|-------|----------------|
| curriculum/ | Backend Team | Graduate Profile (CP) and Tujuan Pembelajaran (TP) |
| planning/ | Backend Team | Alur Tujuan Pembelajaran (ATP) and Modul Ajar |
| assessment/ | Backend Team | Assessment, Rubric, and Evidence |
| reporting/ | Backend Team | Narrative Report |
| administration/ | Backend Team | Authentication, School, and User management |
| ai/ | AI Engineer | AI agent orchestration and integration |
| shared/ | Backend Team | Shared domain entities and infrastructure |

---

# Frontend Structure

## Technology Stack

- **Framework**: React
- **Language**: TypeScript

## Frontend Folder Structure

```
frontend/
├── src/
│   ├── pages/
│   │   ├── auth/
│   │   │   └── LoginPage.tsx
│   │   ├── curriculum/
│   │   │   ├── CPListPage.tsx
│   │   │   ├── CPDetailPage.tsx
│   │   │   ├── TPListPage.tsx
│   │   │   └── TPDetailPage.tsx
│   │   ├── planning/
│   │   │   ├── ATPListPage.tsx
│   │   │   ├── ATPDetailPage.tsx
│   │   │   ├── ModulAjarListPage.tsx
│   │   │   └── ModulAjarDetailPage.tsx
│   │   ├── assessment/
│   │   │   ├── AssessmentListPage.tsx
│   │   │   ├── AssessmentDetailPage.tsx
│   │   │   ├── RubricListPage.tsx
│   │   │   ├── RubricDetailPage.tsx
│   │   │   ├── EvidenceListPage.tsx
│   │   │   └── EvidenceDetailPage.tsx
│   │   ├── reporting/
│   │   │   ├── NarrativeReportListPage.tsx
│   │   │   └── NarrativeReportDetailPage.tsx
│   │   ├── administration/
│   │   │   ├── SchoolListPage.tsx
│   │   │   ├── SchoolDetailPage.tsx
│   │   │   ├── UserListPage.tsx
│   │   │   └── UserDetailPage.tsx
│   │   ├── DashboardPage.tsx
│   │   └── NotFoundPage.tsx
│   ├── components/
│   │   ├── common/
│   │   │   ├── Button/
│   │   │   ├── Input/
│   │   │   ├── Modal/
│   │   │   ├── Table/
│   │   │   └── Card/
│   │   └── layout/
│   │       ├── Header/
│   │       ├── Sidebar/
│   │       └── Layout/
│   ├── features/
│   │   ├── curriculum/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   └── hooks/
│   │   ├── planning/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   └── hooks/
│   │   ├── assessment/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   └── hooks/
│   │   ├── reporting/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   └── hooks/
│   │   ├── administration/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   └── hooks/
│   │   └── ai/
│   │       ├── components/
│   │       ├── services/
│   │       └── hooks/
│   ├── services/
│   │   ├── api/
│   │   │   ├── apiClient.ts
│   │   │   └── apiConfig.ts
│   │   └── auth/
│   │       └── authService.ts
│   ├── hooks/
│   │   ├── useAuth.ts
│   │   ├── useApi.ts
│   │   └── useModule.ts
│   ├── contexts/
│   │   ├── AuthContext.tsx
│   │   └── ThemeContext.tsx
│   ├── layouts/
│   │   ├── MainLayout.tsx
│   │   └── AuthLayout.tsx
│   ├── types/
│   │   ├── common.types.ts
│   │   ├── api.types.ts
│   │   └── domain.types.ts
│   ├── utils/
│   │   ├── validation.ts
│   │   ├── formatting.ts
│   │   └── dateHelpers.ts
│   ├── App.tsx
│   └── main.tsx
├── public/
│   ├── index.html
│   ├── favicon.ico
│   └── assets/
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── package.json
├── tsconfig.json
├── vite.config.ts
├── tailwind.config.js
├── .eslintrc.js
└── Dockerfile
```

## Frontend Feature Ownership

| Feature | Owner | Responsibility |
|---------|-------|----------------|
| curriculum/ | Frontend Team | CP and TP UI components and pages |
| planning/ | Frontend Team | ATP and Modul Ajar UI components and pages |
| assessment/ | Frontend Team | Assessment, Rubric, Evidence UI components and pages |
| reporting/ | Frontend Team | Narrative Report UI components and pages |
| administration/ | Frontend Team | Authentication, School, User UI components and pages |
| ai/ | Frontend Team + AI Engineer | AI generation UI components and pages |

---

# Database Structure

## Database Folder Structure

```
database/
├── migrations/
│   ├── 000001_init_schema.up.sql
│   ├── 000001_init_schema.down.sql
│   ├── 000002_users.up.sql
│   ├── 000002_users.down.sql
│   ├── 000003_curriculum.up.sql
│   ├── 000003_curriculum.down.sql
│   ├── 000004_planning.up.sql
│   ├── 000004_planning.down.sql
│   ├── 000005_assessment.up.sql
│   ├── 000005_assessment.down.sql
│   ├── 000006_reporting.up.sql
│   ├── 000006_reporting.down.sql
│   ├── 000007_ai_audit.up.sql
│   ├── 000007_ai_audit.down.sql
│   └── 000008_workflow_history.up.sql
│       └── 000008_workflow_history.down.sql
├── seed/
│   └── seed_data.sql
└── schema/
    └── schema_diagram.sql
```

## Database Ownership

| Directory | Owner | Responsibility |
|-----------|-------|----------------|
| migrations/ | Backend Team + DevOps | Database migration files |
| seed/ | Backend Team | Seed data for development |
| schema/ | Backend Team | Database schema documentation |

## Migration Naming Convention

YYYYMMDDHHMM_<description>

Examples:
- 202606050900_init_schema
- 202606051200_add_status_column

---

# AI Structure

## AI Prompt Folder Structure

```
prompts/
├── tp/
│   ├── prompt_template.txt
│   ├── prompt_version.txt
│   └── examples/
├── atp/
│   ├── prompt_template.txt
│   ├── prompt_version.txt
│   └── examples/
├── modul-ajar/
│   ├── prompt_template.txt
│   ├── prompt_version.txt
│   └── examples/
├── assessment/
│   ├── prompt_template.txt
│   ├── prompt_version.txt
│   └── examples/
├── rubric/
│   ├── prompt_template.txt
│   ├── prompt_version.txt
│   └── examples/
└── report/
    ├── prompt_template.txt
    ├── prompt_version.txt
    └── examples/
```

## AI Prompt Ownership

| Directory | Owner | Responsibility |
|-----------|-------|----------------|
| tp/ | AI Engineer | TP Agent prompt specification |
| atp/ | AI Engineer | ATP Agent prompt specification |
| modul-ajar/ | AI Engineer | Modul Ajar Agent prompt specification |
| assessment/ | AI Engineer | Assessment Agent prompt specification |
| rubric/ | AI Engineer | Rubric Agent prompt specification |
| report/ | AI Engineer | Narrative Report Agent prompt specification |

---

# Documentation Structure

## Documentation Folder Structure

```
docs/
├── foundation/
│   ├── 00A_NATIONAL_EDUCATION_DIRECTION_2045.md
│   ├── 00B_PRODUCT_VISION.md
│   ├── 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md
│   ├── 01_EDUCATION_DOMAIN_MODEL.md
│   ├── 02_CAPABILITY_MODEL.md
│   ├── 03_BUSINESS_PROCESS_ARCHITECTURE.md
│   ├── 04_DATA_ARCHITECTURE.md
│   ├── 05_AI_ARCHITECTURE.md
│   ├── 06_APPLICATION_ARCHITECTURE.md
│   ├── 07_MVP_ARCHITECTURE.md
│   ├── 08_SDLC_ARCHITECTURE.md
│   ├── 09_REPOSITORY_ARCHITECTURE.md
│   ├── 10_PRODUCT_FEATURE_CATALOG.md
│   ├── 11_PRODUCT_BACKLOG.md
│   ├── 12_ARCHITECTURE_DECISION_RECORDS.md
│   ├── 13_API_CONTRACT.md
│   ├── 14_DATABASE_SCHEMA.md
│   ├── 15_AI_PROMPT_SPECIFICATION.md
│   └── 16_ARCHITECTURE_FREEZE.md
├── adr/
│   ├── ADR-001-architecture-style.md
│   ├── ADR-002-backend-technology.md
│   ├── ADR-003-frontend-technology.md
│   ├── ADR-004-database-technology.md
│   ├── ADR-005-messaging-technology.md
│   ├── ADR-006-authentication-strategy.md
│   ├── ADR-007-deployment-strategy.md
│   ├── ADR-008-ai-orchestration.md
│   ├── ADR-009-llm-integration.md
│   └── ADR-010-monorepo-strategy.md
├── api/
│   └── api-documentation.md
├── database/
│   └── migration-guide.md
└── prompt-specification/
    └── prompt-lifecycle.md
```

## Documentation Ownership

| Directory | Owner | Responsibility |
|-----------|-------|----------------|
| foundation/ | Architecture Team | Foundation architecture documents |
| adr/ | Architecture Team | Architecture Decision Records |
| api/ | Backend Team | API documentation |
| database/ | Backend Team | Database migration guide |
| prompt-specification/ | AI Engineer | AI prompt lifecycle documentation |

---

# Scripts Structure

## Scripts Folder Structure

```
scripts/
├── dev.sh
├── build.sh
├── test.sh
├── deploy.sh
├── migrate.sh
└── seed.sh
```

## Scripts Ownership

| Script | Owner | Responsibility |
|--------|-------|----------------|
| dev.sh | DevOps | Development environment setup |
| build.sh | DevOps | Build automation |
| test.sh | DevOps | Test automation |
| deploy.sh | DevOps | Deployment automation |
| migrate.sh | Backend Team + DevOps | Database migration |
| seed.sh | Backend Team + DevOps | Database seeding |

---

# Deploy Structure

## Deploy Folder Structure

```
deploy/
├── docker/
│   ├── docker-compose.yml
│   ├── docker-compose.dev.yml
│   └── docker-compose.prod.yml
├── nginx/
│   └── nginx.conf
└── kubernetes/
    └── (future - not for MVP)
```

## Deploy Ownership

| Directory | Owner | Responsibility |
|-----------|-------|----------------|
| docker/ | DevOps | Docker Compose configurations |
| nginx/ | DevOps | Nginx configuration |
| kubernetes/ | DevOps | Kubernetes configurations (future) |

---

# ADR Structure

## ADR Folder Structure

```
adr/
├── ADR-001-architecture-style.md
├── ADR-002-backend-technology.md
├── ADR-003-frontend-technology.md
├── ADR-004-database-technology.md
├── ADR-005-messaging-technology.md
├── ADR-006-authentication-strategy.md
├── ADR-007-deployment-strategy.md
├── ADR-008-ai-orchestration.md
├── ADR-009-llm-integration.md
└── ADR-010-monorepo-strategy.md
```

## ADR Ownership

| File | Owner | Responsibility |
|------|-------|----------------|
| All ADRs | Architecture Team | Architecture Decision Records |

---

# Repository Structure Summary

## MVP Wave 1 Scope

This repository structure supports MVP Wave 1 only:

- **Authentication Module**: Custom JWT authentication
- **Curriculum Module**: Graduate Profile (CP) and Tujuan Pembelajaran (TP)
- **Learning Planning Module**: Alur Tujuan Pembelajaran (ATP) and Modul Ajar
- **Assessment Module**: Assessment, Rubric, and Evidence
- **Reporting Module**: Narrative Report
- **Administration Module**: School and User management
- **AI Orchestration Module**: AI agent coordination and integration

## Future Scope Excluded

The following are explicitly excluded from MVP Wave 1 repository structure:

- Competency Graph data structures
- Digital Twin data structures
- Lifelong Learning Record data structures
- Career Guidance features
- Adaptive Learning features
- Student role features
- Parent role features
- Teacher Professional Growth features
- School Improvement features
- Education Analytics features

## Change Control

No repository structure changes may be introduced during MVP Wave 1 implementation without explicit Architecture Freeze Amendment approved by Chief Enterprise Architect.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED

**Repository Structure Status**: FROZEN FOR MVP WAVE 1
