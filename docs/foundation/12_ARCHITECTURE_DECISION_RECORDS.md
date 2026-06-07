# 12_ARCHITECTURE_DECISION_RECORDS.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Validated against Foundation Architecture (00A, 00B, 00C, 01, 02, 03, 04, 05, 06, 07, 08, 09)

**Purpose**: Document all frozen architecture decisions for NUSA MVP Wave 1 using Architecture Decision Records (ADRs). This document serves as the official architecture governance record, providing context, rationale, consequences, and trade-offs for each architectural decision.

---

# SECTION 1 — ADR Overview

## ADR Format

Each Architecture Decision Record follows the standard format:

- **Status**: Current state of the decision (Accepted, Proposed, Deprecated, Superseded)
- **Context**: Background and problem statement
- **Decision**: The architectural decision made
- **Consequences**: Positive and negative impacts of the decision
- **Alternatives Considered**: Other options evaluated
- **Trade-offs**: Key trade-offs and compromises

## ADR Index

| ADR | Title | Status | Date |
|-----|-------|--------|------|
| ADR-001 | Modular Monolith Architecture | Accepted | June 2026 |
| ADR-002 | Go + Gin Backend | Accepted | June 2026 |
| ADR-003 | React Frontend | Accepted | June 2026 |
| ADR-004 | PostgreSQL Database | Accepted | June 2026 |
| ADR-005 | RabbitMQ Message Queue | Accepted | June 2026 |
| ADR-006 | Custom JWT Authentication | Accepted | June 2026 |
| ADR-007 | REST API Style | Accepted | June 2026 |
| ADR-008 | Docker Compose Deployment | Accepted | June 2026 |
| ADR-009 | AI Orchestration Module | Accepted | June 2026 |
| ADR-010 | External LLM Integration | Accepted | June 2026 |

---

# SECTION 2 — Architecture Decision Records

## ADR-001: Modular Monolith Architecture

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires an architecture that:
- Supports the curriculum-to-report workflow (CP → TP → ATP → Modul Ajar → Assessment → Rubric → Evidence → Narrative Report)
- Enables AI-assisted content generation with human-in-the-loop governance
- Can be delivered within 20 days by 5 developers
- Provides clear module boundaries aligned with capabilities
- Allows for future evolution to more distributed architectures if needed

The platform must balance:
- Development speed for MVP delivery
- Operational simplicity for initial deployment
- Architectural quality for long-term maintainability
- Future scalability for platform expansion

### Decision

**Adopt Modular Monolith Architecture**

NUSA will implement a Modular Monolith architecture where:
- All modules (Curriculum, Learning Planning, Assessment, Reporting, Authentication, AI Orchestration) are deployed as a single application
- Module boundaries are strictly enforced through code organization and dependency rules
- Each module has its own domain, application, and infrastructure layers
- Modules communicate through well-defined interfaces (application services, domain events)
- The application is deployed as a single deployable unit

### Consequences

**Positive**:
- Faster development speed for MVP (no distributed system complexity)
- Simpler deployment and operations (single deployable unit)
- Easier debugging and monitoring (single process)
- Lower infrastructure costs (no need for multiple servers)
- Clear module boundaries enable future extraction to microservices if needed
- Transactional consistency across modules (single database)
- Reduced network latency (in-process communication)

**Negative**:
- Limited scalability (cannot scale individual modules independently)
- Single point of failure (entire application goes down if one module fails)
- Technology coupling (all modules must use the same technology stack)
- Deployment coupling (all modules must be deployed together)
- Memory constraints (all modules share the same memory space)

### Alternatives Considered

**Microservices Architecture**:
- Pros: Independent scaling, fault isolation, technology diversity
- Cons: Too complex for 20-day MVP, operational overhead, distributed system complexity, network latency
- Rejected: MVP timeline too aggressive, operational complexity not justified for initial scope

**Monolithic Architecture**:
- Pros: Simplest architecture, fastest development
- Cons: No module boundaries, difficult to maintain long-term, no clear separation of concerns
- Rejected: Does not provide architectural quality needed for long-term platform evolution

**Serverless Architecture**:
- Pros: Auto-scaling, pay-per-use, no server management
- Cons: Cold start latency, vendor lock-in, complexity for stateful applications
- Rejected: Not suitable for curriculum-to-report workflow with stateful operations

### Trade-offs

**Development Speed vs. Architectural Quality**:
- Chose Modular Monolith to balance development speed with architectural quality
- Module boundaries provide architectural quality without distributed system complexity

**Simplicity vs. Scalability**:
- Chose simplicity for MVP delivery
- Scalability can be addressed in future waves by extracting modules to microservices

**Coupling vs. Consistency**:
- Chose coupling for transactional consistency
- Eventual consistency can be introduced in future waves if needed

---

## ADR-002: Go + Gin Backend

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires a backend technology that:
- Supports the curriculum-to-report workflow with AI orchestration
- Can be developed quickly by 5 developers in 20 days
- Provides good performance for AI agent orchestration
- Has strong typing and concurrency support
- Has good ecosystem for web APIs, database access, and message queues
- Is maintainable for long-term platform evolution

The backend must handle:
- REST API endpoints for all modules
- AI agent orchestration and prompt management
- Database operations with PostgreSQL
- Message queue operations with RabbitMQ
- JWT authentication and authorization
- Domain event publishing and consumption

### Decision

**Use Go + Gin for Backend**

NUSA will use Go programming language with Gin web framework for the backend:
- Go for the language (performance, concurrency, strong typing, simplicity)
- Gin for the web framework (fast, minimal, good middleware support)

### Consequences

**Positive**:
- Excellent performance for AI orchestration (compiled language, efficient runtime)
- Strong concurrency support (goroutines) for parallel AI agent execution
- Strong typing reduces runtime errors
- Simple deployment (single binary)
- Good ecosystem for web APIs, database access, message queues
- Fast compilation and build times
- Good tooling for testing and profiling
- Large talent pool for Go developers

**Negative**:
- Steeper learning curve for developers unfamiliar with Go
- Less expressive than dynamically typed languages
- More verbose than some modern languages
- Ecosystem smaller than JavaScript/Python
- Error handling can be verbose (explicit error checking)

### Alternatives Considered

**Node.js + Express**:
- Pros: Large ecosystem, JavaScript everywhere, fast development
- Cons: Single-threaded event loop limits concurrency, dynamic typing increases runtime errors, performance lower than Go
- Rejected: Concurrency limitations for AI orchestration, performance concerns

**Python + FastAPI**:
- Pros: Excellent AI/ML ecosystem, fast development, type hints available
- Cons: Performance lower than Go, Global Interpreter Lock limits concurrency
- Rejected: Performance and concurrency limitations for AI orchestration

**Java + Spring Boot**:
- Pros: Enterprise-grade, large ecosystem, strong typing
- Cons: Heavy framework, slower startup, more verbose than Go
- Rejected: Complexity and verbosity not justified for MVP timeline

**Rust + Actix**:
- Pros: Excellent performance, memory safety, strong typing
- Cons: Steep learning curve, smaller ecosystem, longer development time
- Rejected: Learning curve too steep for 20-day MVP

### Trade-offs

**Performance vs. Development Speed**:
- Chose Go for performance without sacrificing too much development speed
- Go's simplicity offsets performance optimization time

**Concurrency vs. Ecosystem**:
- Chose Go's concurrency over larger ecosystems of other languages
- Go's ecosystem is sufficient for MVP requirements

**Learning Curve vs. Long-term Benefits**:
- Chose Go despite learning curve for long-term performance and maintainability benefits
- Team can learn Go during development

---

## ADR-003: React Frontend

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires a frontend technology that:
- Provides a modern, responsive user interface for teachers and administrators
- Supports the curriculum-to-report workflow with complex forms and data visualization
- Can be developed quickly by 5 developers in 20 days
- Has strong ecosystem for UI components, state management, and API integration
- Is maintainable for long-term platform evolution
- Provides good developer experience

The frontend must handle:
- Complex forms for teaching artifact creation and editing
- AI-generated content display and editing
- Document viewing and export
- User authentication and authorization
- Real-time updates for AI generation status
- Responsive design for various screen sizes

### Decision

**Use React + TypeScript for Frontend**

NUSA will use React with TypeScript for the frontend:
- React for the UI framework (component-based, large ecosystem, good performance)
- TypeScript for type safety (reduces runtime errors, improves maintainability)
- Vite for build tool (fast development server, optimized production builds)

### Consequences

**Positive**:
- Component-based architecture aligns with modular backend
- Large ecosystem of UI components and libraries
- Strong typing with TypeScript reduces runtime errors
- Excellent developer experience with hot module replacement
- Good performance with virtual DOM
- Large talent pool for React developers
- Good state management options (Redux, Zustand, Context API)
- Strong community support and documentation

**Negative**:
- Steeper learning curve for developers unfamiliar with React
- Bundle size can be large if not optimized
- State management can be complex for large applications
- JSX syntax can be unfamiliar to some developers
- Frequent updates to React ecosystem can cause maintenance burden

### Alternatives Considered

**Vue.js + TypeScript**:
- Pros: Simpler learning curve, good performance, strong typing
- Cons: Smaller ecosystem than React, less talent pool
- Rejected: Smaller ecosystem and talent pool

**Angular + TypeScript**:
- Pros: Enterprise-grade, comprehensive framework, strong typing built-in
- Cons: Steeper learning curve, more opinionated, larger bundle size
- Rejected: Complexity and bundle size not justified for MVP

**Svelte + TypeScript**:
- Pros: Smaller bundle size, simpler syntax, good performance
- Cons: Smaller ecosystem, less mature than React
- Rejected: Smaller ecosystem and maturity concerns

**Vanilla JavaScript**:
- Pros: No framework overhead, fastest performance
- Cons: No component architecture, harder to maintain, no type safety
- Rejected: Lack of component architecture and type safety

### Trade-offs

**Ecosystem Size vs. Bundle Size**:
- Chose React's large ecosystem over smaller bundle size of alternatives
- Bundle size can be optimized with code splitting and lazy loading

**Learning Curve vs. Talent Pool**:
- Chose React despite learning curve for larger talent pool
- Team can learn React during development

**Maturity vs. Innovation**:
- Chose React's maturity over newer frameworks
- Maturity reduces risk for MVP delivery

---

## ADR-004: PostgreSQL Database

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires a database that:
- Stores all curriculum, teaching, assessment, and reporting data
- Supports complex queries for reporting and analytics
- Provides transactional consistency across modules
- Integrates well with Go backend
- Can be deployed with Docker Compose
- Is reliable and well-maintained
- Supports future scalability needs

The database must handle:
- User accounts and authentication data
- National curriculum data
- Teaching plans (TP, ATP, Modul Ajar)
- Assessments and rubrics
- Evidence and evaluations
- Narrative reports
- Audit trails and change history

### Decision

**Use PostgreSQL as the Primary Database**

NUSA will use PostgreSQL as the single database for MVP Wave 1:
- Single database for all modules (Modular Monolith pattern)
- PostgreSQL for relational data with ACID transactions
- JSONB support for flexible schema where needed
- Full-text search capabilities for curriculum data
- Strong data integrity with foreign key constraints

### Consequences

**Positive**:
- ACID transactions ensure data consistency across modules
- Strong data integrity with foreign key constraints
- JSONB support provides flexibility for evolving schemas
- Excellent query performance with proper indexing
- Mature and well-maintained database
- Good integration with Go (pgx, sqlx libraries)
- Strong community support and documentation
- Supports complex queries for reporting and analytics
- Full-text search capabilities

**Negative**:
- Single database limits independent scaling of modules
- Database becomes bottleneck if not properly optimized
- Schema migrations can be complex for large datasets
- Vertical scaling required for performance (cannot scale horizontally easily)
- Single point of failure for data (requires backup strategy)

### Alternatives Considered

**MySQL**:
- Pros: Popular, good performance, widely used
- Cons: Less advanced features than PostgreSQL, weaker JSON support
- Rejected: PostgreSQL provides more advanced features and better JSON support

**MongoDB**:
- Pros: Flexible schema, good for document storage, horizontal scaling
- Cons: No ACID transactions across collections, weaker data integrity
- Rejected: Lack of ACID transactions and data integrity not suitable for educational data

**SQLite**:
- Pros: Simplest database, no server required
- Cons: Not suitable for production, limited concurrency, no network access
- Rejected: Not suitable for production application with multiple users

**CockroachDB**:
- Pros: Distributed SQL, horizontal scaling, ACID transactions
- Cons: More complex to deploy, smaller ecosystem, less mature
- Rejected: Complexity not justified for MVP single-database requirement

### Trade-offs

**Consistency vs. Flexibility**:
- Chose PostgreSQL's ACID transactions over MongoDB's flexibility
- Data consistency is critical for educational artifacts

**Maturity vs. Innovation**:
- Chose PostgreSQL's maturity over newer distributed databases
- Maturity reduces risk for MVP delivery

**Single Database vs. Distributed Database**:
- Chose single database for MVP simplicity
- Can migrate to distributed database in future waves if needed

---

## ADR-005: RabbitMQ Message Queue

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires a message queue for:
- Asynchronous communication between modules
- Domain event publishing and consumption
- AI agent orchestration and coordination
- Decoupling of module dependencies
- Event-driven architecture foundation

The message queue must handle:
- Curriculum events (CP imported, TP created, TP updated)
- Learning planning events (ATP created, Modul Ajar created)
- Assessment events (Assessment created, Rubric created)
- Reporting events (Narrative Report created)
- AI orchestration events (Agent started, Agent completed, Agent failed)

### Decision

**Use RabbitMQ as the Message Queue**

NUSA will use RabbitMQ for message queuing:
- RabbitMQ for message broker (mature, reliable, feature-rich)
- AMQP protocol for message communication
- Exchange and queue architecture for event routing
- Dead letter queue for error handling
- Message persistence for reliability

### Consequences

**Positive**:
- Mature and reliable message broker
- Flexible routing with exchanges and queues
- Good integration with Go (amqp091-go library)
- Supports message persistence for reliability
- Dead letter queue for error handling
- Good monitoring and management tools
- Supports multiple messaging patterns (pub/sub, RPC, routing)
- Strong community support and documentation

**Negative**:
- Additional infrastructure component to manage
- Adds operational complexity
- Requires proper configuration for production
- Message ordering not guaranteed across queues
- Learning curve for developers unfamiliar with message queues

### Alternatives Considered

**Kafka**:
- Pros: High throughput, good for event streaming, strong durability
- Cons: More complex to deploy and operate, overkill for MVP requirements
- Rejected: Complexity not justified for MVP event volume

**Redis Streams**:
- Pros: Simple, fast, built-in to Redis
- Cons: Less mature than RabbitMQ, fewer features
- Rejected: RabbitMQ provides more features and better reliability

**AWS SQS**:
- Pros: Managed service, no operational overhead
- Cons: Vendor lock-in, additional cost, not suitable for Docker Compose deployment
- Rejected: Vendor lock-in and cost not justified for MVP

**In-memory events (no message queue)**:
- Pros: Simplest approach, no additional infrastructure
- Cons: No persistence, no reliability, no decoupling
- Rejected: Lack of persistence and reliability not suitable for production

### Trade-offs

**Maturity vs. Simplicity**:
- Chose RabbitMQ's maturity over simpler alternatives
- Maturity reduces risk for production reliability

**Feature Richness vs. Complexity**:
- Chose RabbitMQ's features over simpler message queues
- Features needed for proper event-driven architecture

**Self-hosted vs. Managed**:
- Chose self-hosted RabbitMQ for Docker Compose deployment
- Managed services add cost and vendor lock-in

---

## ADR-006: Custom JWT Authentication

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires an authentication strategy that:
- Provides secure user authentication for teachers and administrators
- Supports role-based access control (RBAC)
- Integrates well with Go backend and React frontend
- Can be deployed with Docker Compose
- Does not require external identity providers
- Provides good performance for API calls
- Supports token refresh for seamless user experience

The authentication must handle:
- User login with email and password
- JWT access token generation
- JWT refresh token generation
- Token validation on API calls
- Role-based access control
- Session management

### Decision

**Use Custom JWT Authentication**

NUSA will implement custom JWT authentication:
- Custom JWT implementation in Go backend
- JWT access tokens (24-hour expiration)
- JWT refresh tokens (7-day expiration)
- Role-based access control (RBAC)
- Token storage in frontend (localStorage)
- No external identity provider (Keycloak, Auth0, etc.)

### Consequences

**Positive**:
- No external dependency on identity providers
- Full control over authentication logic
- Good performance (no external API calls for validation)
- Simple deployment (no additional infrastructure)
- Cost-effective (no subscription fees)
- Flexible for custom requirements
- Good integration with Go and React

**Negative**:
- Must implement and maintain authentication logic
- Must implement security best practices (token signing, validation, revocation)
- No built-in user management UI
- Must implement password hashing and validation
- Must implement token refresh logic
- Must implement role-based access control
- Security responsibility on development team

### Alternatives Considered

**Keycloak**:
- Pros: Enterprise-grade identity provider, built-in user management, SSO support
- Cons: Additional infrastructure complexity, overkill for MVP, operational overhead
- Rejected: Complexity not justified for MVP requirements

**Auth0**:
- Pros: Managed service, enterprise features, good documentation
- Cons: Vendor lock-in, subscription cost, not suitable for Docker Compose deployment
- Rejected: Vendor lock-in and cost not justified for MVP

**OAuth2 + OpenID Connect**:
- Pros: Industry standard, wide support, good security
- Cons: Requires external identity provider, adds complexity
- Rejected: External identity provider not needed for MVP

**Session-based Authentication**:
- Pros: Simple to implement, server-controlled sessions
- Cons: Not suitable for stateless APIs, requires session storage
- Rejected: JWT better suited for stateless REST API

### Trade-offs

**Control vs. Convenience**:
- Chose custom JWT for control over authentication logic
- Convenience of managed services not worth vendor lock-in

**Security Responsibility vs. External Dependency**:
- Chose to own security responsibility rather than depend on external providers
- Team can implement security best practices

**Simplicity vs. Enterprise Features**:
- Chose simple custom JWT over enterprise identity providers
- Enterprise features not needed for MVP

---

## ADR-007: REST API Style

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires an API style that:
- Provides clear communication between frontend and backend
- Supports all curriculum-to-report workflow operations
- Integrates well with Go backend and React frontend
- Is well-understood by developers
- Provides good tooling and documentation support
- Supports future API evolution

The API must handle:
- CRUD operations for all entities (CP, TP, ATP, Modul Ajar, Assessment, Rubric, Evidence, Narrative Report)
- AI generation operations (generate TP, generate ATP, generate Modul Ajar, generate Assessment, generate Rubric, generate Narrative Report)
- Authentication and authorization
- File uploads and downloads
- Search and filtering operations

### Decision

**Use REST API Style**

NUSA will implement REST API for all backend-frontend communication:
- RESTful API design principles
- HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Resource-based URLs
- JSON request/response format
- HTTP status codes for error handling
- API versioning (v1)
- OpenAPI/Swagger documentation

### Consequences

**Positive**:
- Well-understood and widely adopted
- Good tooling support (Postman, Swagger UI)
- Easy to document with OpenAPI/Swagger
- Stateless and scalable
- Good integration with Go (Gin) and React (axios, fetch)
- Cacheable with HTTP caching
- Supports multiple clients (web, mobile, API)
- Industry standard with large talent pool

**Negative**:
- Can lead to over-fetching or under-fetching data
- Multiple round trips for related data
- No built-in real-time communication
- URL design can be subjective
- Versioning can be complex

### Alternatives Considered

**GraphQL**:
- Pros: Flexible queries, no over-fetching/under-fetching, single endpoint
- Cons: More complex to implement, steeper learning curve, overkill for MVP
- Rejected: Complexity not justified for MVP requirements

**gRPC**:
- Pros: High performance, strong typing, good for microservices
- Cons: Not suitable for web browsers, more complex, less tooling
- Rejected: Not suitable for web frontend integration

**WebSocket**:
- Pros: Real-time communication, bi-directional
- Cons: Not suitable for CRUD operations, more complex state management
- Rejected: Not suitable for primary API style, can add for specific features later

**SOAP**:
- Pros: Enterprise-grade, strong typing, WS-Security
- Cons: Complex, verbose, outdated, poor browser support
- Rejected: Outdated and overly complex for modern web applications

### Trade-offs

**Simplicity vs. Flexibility**:
- Chose REST's simplicity over GraphQL's flexibility
- GraphQL's flexibility not needed for MVP requirements

**Tooling vs. Performance**:
- Chose REST's tooling over gRPC's performance
- REST performance sufficient for MVP requirements

**Industry Standard vs. Innovation**:
- Chose REST's industry standard over newer API styles
- Industry standard reduces risk and improves developer onboarding

---

## ADR-008: Docker Compose Deployment

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires a deployment strategy that:
- Can be deployed quickly for MVP delivery
- Supports local development for 5 developers
- Integrates all components (backend, frontend, PostgreSQL, RabbitMQ)
- Is simple to operate and maintain
- Does not require complex infrastructure
- Supports future migration to production deployment
- Provides good developer experience

The deployment must handle:
- Backend service (Go + Gin)
- Frontend service (React + TypeScript)
- PostgreSQL database
- RabbitMQ message queue
- Nginx reverse proxy (optional)
- Volume mounts for data persistence
- Environment variable configuration

### Decision

**Use Docker Compose for Deployment**

NUSA will use Docker Compose for deployment:
- Docker Compose for orchestration
- Individual services defined in docker-compose.yml
- Volume mounts for data persistence
- Environment variables for configuration
- Single command to start all services
- Development and production configurations

### Consequences

**Positive**:
- Simple deployment with single command
- Consistent environment across development and production
- Easy to set up for local development
- Good integration with Go and React Dockerfiles
- Volume mounts for data persistence
- Environment variable configuration
- Good developer experience
- Low operational complexity
- No infrastructure as code complexity

**Negative**:
- Not suitable for large-scale production deployments
- Limited scalability (single host)
- No built-in load balancing
- No built-in service discovery
- Manual scaling required
- No built-in monitoring and alerting
- Limited high availability

### Alternatives Considered

**Kubernetes**:
- Pros: Scalable, production-grade, auto-scaling, service discovery
- Cons: Too complex for MVP, steep learning curve, operational overhead
- Rejected: Complexity not justified for MVP timeline and scale

**AWS ECS**:
- Pros: Managed service, scalable, production-grade
- Cons: Vendor lock-in, additional cost, complexity
- Rejected: Vendor lock-in and cost not justified for MVP

**Heroku**:
- Pros: Managed service, simple deployment, good developer experience
- Cons: Vendor lock-in, cost, not suitable for Docker Compose workflow
- Rejected: Vendor lock-in and cost not justified for MVP

**Manual Deployment**:
- Pros: No Docker overhead
- Cons: Inconsistent environments, difficult to reproduce, error-prone
- Rejected: Inconsistent environments increase risk for MVP delivery

### Trade-offs

**Simplicity vs. Scalability**:
- Chose Docker Compose's simplicity over Kubernetes' scalability
- Scalability can be addressed in future waves if needed

**Control vs. Managed Services**:
- Chose self-hosted Docker Compose over managed services
- Managed services add cost and vendor lock-in

**Development Speed vs. Production Readiness**:
- Chose development speed for MVP delivery
- Production-grade deployment can be added in future waves

---

## ADR-009: AI Orchestration Module

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires an AI orchestration strategy that:
- Coordinates 6 AI agents (TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report)
- Manages prompt construction and optimization
- Provides workflow orchestration for the curriculum-to-report pipeline
- Handles AI service integration
- Implements human-in-the-loop governance
- Provides graceful degradation when AI fails
- Logs AI agent executions for auditability

The AI orchestration must handle:
- TP Generator Agent (CP → TP)
- ATP Generator Agent (TP → ATP)
- Modul Ajar Agent (ATP → Modul Ajar)
- Assessment Agent (Modul Ajar → Assessment)
- Rubric Agent (Assessment → Rubric)
- Narrative Report Agent (Evidence → Narrative Report)

### Decision

**Use Centralized AI Orchestration Module**

NUSA will implement a centralized AI Orchestration Module:
- Single module for all AI agent coordination
- Workflow orchestration for agent execution sequence
- Prompt orchestration for prompt construction and optimization
- Generation pipeline for end-to-end generation process
- Human-in-the-loop enforcement (teacher approval required)
- Graceful degradation (manual mode when AI fails)
- Audit logging for all AI executions

### Consequences

**Positive**:
- Centralized control over AI agent execution
- Consistent prompt management across all agents
- Simplified monitoring and logging
- Easier to implement human-in-the-loop governance
- Graceful degradation to manual mode
- Clear separation of AI logic from business logic
- Easier to test and debug AI orchestration
- Foundation for future AI capabilities

**Negative**:
- Single point of failure for AI orchestration
- Centralized module can become bottleneck
- More complex to implement than distributed agents
- Requires careful design of agent interfaces
- Adds architectural complexity

### Alternatives Considered

**Distributed AI Agents**:
- Pros: Independent agent execution, fault isolation, scalability
- Cons: More complex to orchestrate, harder to enforce governance, operational overhead
- Rejected: Complexity not justified for MVP requirements

**Embedded AI Logic in Each Module**:
- Pros: Simpler architecture, no separate orchestration layer
- Cons: Duplicated AI logic, inconsistent prompt management, harder to govern
- Rejected: Duplicated logic and inconsistent management not acceptable

**External AI Orchestration Service**:
- Pros: Managed service, no operational overhead
- Cons: Vendor lock-in, cost, not suitable for Docker Compose deployment
- Rejected: Vendor lock-in and cost not justified for MVP

**No Orchestration (Direct LLM Calls)**:
- Pros: Simplest approach
- Cons: No prompt management, no governance, no auditability
- Rejected: Lack of governance and auditability not acceptable for educational artifacts

### Trade-offs

**Centralization vs. Distribution**:
- Chose centralized orchestration for control and governance
- Distribution adds complexity not justified for MVP

**Complexity vs. Governance**:
- Chose orchestration complexity for proper governance
- Governance is critical for educational artifacts

**Custom vs. External**:
- Chose custom orchestration for control and flexibility
- External services add cost and vendor lock-in

---

## ADR-010: External LLM Integration

**Status**: Accepted
**Date**: June 2026

### Context

NUSA MVP Wave 1 requires an AI runtime strategy that:
- Provides high-quality AI assistance for content generation
- Supports prompt-based AI interaction
- Integrates well with Go backend
- Can be deployed with Docker Compose
- Does not require custom model training infrastructure
- Provides good performance for real-time generation
- Supports multiple LLM providers for flexibility

The AI integration must handle:
- TP generation from CP
- ATP generation from TP
- Modul Ajar generation from ATP
- Assessment generation from Modul Ajar
- Rubric generation from Assessment
- Narrative Report generation from Evidence

### Decision

**Use External LLM Integration**

NUSA will integrate with external LLM providers:
- External LLM providers (OpenAI, Anthropic, etc.)
- Prompt-based AI interaction
- No custom model training infrastructure
- HTTP API integration with LLM providers
- Provider abstraction layer for flexibility
- Fallback to manual mode when AI fails

### Consequences

**Positive**:
- No need for custom model training infrastructure
- Access to state-of-the-art LLM models
- Faster time to market (no model training)
- Lower infrastructure costs (no GPU servers)
- Flexible provider switching
- Continuous model improvements from providers
- Good performance with provider APIs

**Negative**:
- Dependency on external providers (vendor lock-in risk)
- API costs can be high at scale
- Latency from external API calls
- Limited control over model behavior
- Privacy concerns (data sent to external providers)
- Rate limits from providers
- Provider outages affect system

### Alternatives Considered

**Self-hosted LLM (Llama, Mistral)**:
- Pros: No external dependency, data privacy, no API costs
- Cons: Requires GPU infrastructure, model maintenance, higher operational complexity
- Rejected: GPU infrastructure and maintenance not justified for MVP timeline

**Custom Model Training**:
- Pros: Full control over model behavior, no external dependency
- Cons: Requires ML expertise, training infrastructure, long development time
- Rejected: ML expertise and training infrastructure not available for MVP

**Hybrid Approach (Self-hosted + External)**:
- Pros: Best of both worlds, fallback options
- Cons: Most complex, highest operational overhead
- Rejected: Complexity not justified for MVP requirements

**No AI (Manual Only)**:
- Pros: Simplest approach, no AI complexity
- Cons: No AI assistance, defeats MVP purpose
- Rejected: AI assistance is core value proposition of MVP

### Trade-offs

**External Dependency vs. Infrastructure**:
- Chose external LLM providers to avoid GPU infrastructure
- Infrastructure complexity not justified for MVP timeline

**Cost vs. Control**:
- Chose API costs over model training infrastructure costs
- API costs are predictable and scalable

**Privacy vs. Capability**:
- Chose external providers for state-of-the-art capabilities
- Privacy concerns mitigated by data anonymization and provider selection

---

# SECTION 3 — ADR Governance

## ADR Lifecycle

### ADR States

- **Proposed**: ADR is proposed for discussion
- **Accepted**: ADR is approved and implemented
- **Deprecated**: ADR is no longer recommended but still in use
- **Superseded**: ADR is replaced by a new ADR

### ADR Modification Process

1. **Propose Change**: Create new ADR or modify existing ADR
2. **Review**: Architecture Review Board reviews the change
3. **Approve**: Chief Enterprise Architect approves the change
4. **Document**: Update ADR with new status and rationale
5. **Communicate**: Communicate change to development team

### ADR Rejection Criteria

An ADR may be rejected if:
- It contradicts frozen architecture decisions
- It introduces scope beyond MVP Wave 1
- It adds complexity without clear benefit
- It violates architectural principles
- It is not aligned with strategic objectives

## ADR Compliance

### Mandatory Compliance

All development must comply with:
- Accepted ADRs (cannot be violated without formal amendment)
- Frozen architecture decisions (require Architecture Freeze Amendment)
- Architectural principles (cannot be violated)

### ADR Violation Process

If an ADR must be violated:
1. Document the violation and rationale
2. Submit Architecture Freeze Amendment request
3. Chief Enterprise Architect reviews and approves
4. ADR is updated or superseded
5. Development proceeds with approved amendment

---

# SECTION 4 — Conclusion

## ADR Summary

This Architecture Decision Records (12) document provides the official governance record for NUSA MVP Wave 1:

### ADR Coverage
- **Total ADRs**: 10
- **Status**: All Accepted
- **Scope**: All frozen architecture decisions

### ADR Categories
- **Architecture Pattern**: ADR-001 (Modular Monolith)
- **Backend Technology**: ADR-002 (Go + Gin)
- **Frontend Technology**: ADR-003 (React)
- **Database**: ADR-004 (PostgreSQL)
- **Message Queue**: ADR-005 (RabbitMQ)
- **Authentication**: ADR-006 (Custom JWT)
- **API Style**: ADR-007 (REST API)
- **Deployment**: ADR-008 (Docker Compose)
- **AI Architecture**: ADR-009 (AI Orchestration Module)
- **AI Runtime**: ADR-010 (External LLM Integration)

### Governance Readiness
This ADR document is:
- ✅ Complete with all frozen architecture decisions
- ✅ Aligned with foundation architecture documents
- ✅ Ready for architecture governance
- ✅ Ready for implementation compliance

The Architecture Decision Records are officially approved for NUSA MVP Wave 1 implementation.

---

**Document Status**: FOUNDATION DOCUMENT - LOCKED
