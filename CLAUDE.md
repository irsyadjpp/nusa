# Claude AI Assistant Guidelines

This document provides specific guidelines for Claude (Anthropic) AI assistants when working on the NUSA Platform project.

## Project Context

**NUSA Platform** is a solo-developed education management system for Indonesian schools implementing Kurikulum Merdeka 2026.

### Key Characteristics
- **Solo Developer**: One maintainer with full-time job elsewhere
- **Architecture**: Modular Monolith with DDD Lite
- **Tech Stack**: Go (backend), React/TypeScript (frontend), PostgreSQL
- **Domain**: Kurikulum Merdeka education system
- **Quality**: Production-quality code despite being solo-developed

## Architecture Constraints

### ✅ ALLOWED
- Modular Monolith architecture
- DDD Lite (aggregates, bounded contexts, domain logic)
- Layered architecture: Handler → Service → Repository → PostgreSQL
- TanStack Query + Zustand for frontend state
- Standard Go/React patterns
- Database migrations with up/down SQL

### ❌ FORBIDDEN
- CQRS (Command Query Responsibility Segregation)
- Event Sourcing
- Event Bus or Message Bus (except RabbitMQ for AI workflows)
- Command Bus or Query Bus
- Read Models or Projections
- Microservices architecture
- Complex design patterns without justification
- New domains outside Kurikulum Merdeka scope

## Domain Knowledge

### Kurikulum Merdeka 2026
The education system follows Indonesia's national curriculum:

**Core Entities**:
- **CP** (Capaian Pembelajaran): Learning outcomes/achievements
- **TP** (Tujuan Pembelajaran): Learning objectives with embedded KKTP
- **ATP** (Alur Tujuan Pembelajaran): Annual teaching flow
- **Modul Ajar**: Teaching modules and lesson plans
- **Assessment**: Various assessment types and rubrics
- **Evidence**: Student work evidence and evaluations
- **Achievement**: Competency progress and mastery
- **Narrative Report**: Student progress reports

**Key Concept - KKTP**:
- Kriteria Ketuntasan Tujuan Pembelajaran
- Embedded as Value Object in TP domain
- Contains mastery thresholds, performance indicators, minimum requirements

## Coding Standards

### Backend (Go)
```go
// Handler layer: HTTP only, no business logic
func (h *TPHandler) CreateTP(c *gin.Context) {
    // Parse request
    // Call service
    // Return response
}

// Service layer: Business logic orchestration
func (s *TPService) CreateTP(req CreateTPRequest) (*TP, error) {
    // Validate
    // Coordinate domain logic
    // Call repository
    // Return result
}

// Repository layer: Database access only
func (r *TPRepository) Save(tp *TP) error {
    // SQL operations
    // No business logic
}

// Domain layer: Business rules and invariants
type TP struct {
    // Domain model with methods
    func (tp *TP) Validate() error {
        // Domain invariants
    }
}
```

### Frontend (TypeScript/React)
```typescript
// Use TanStack Query for server state
const { data: tps, isLoading } = useQuery({
  queryKey: ['tps'],
  queryFn: () => tpService.getTPs()
})

// Use Zustand for global client state
const useTPStore = create((set) => ({
  selectedTP: null,
  setSelectedTP: (tp) => set({ selectedTP: tp })
}))

// Component naming: PascalCase
function TPWorkspaceHeader() { }

// Hooks: camelCase with 'use' prefix
function useTPData() { }
```

## Project Structure

```
backend/
├── internal/
│   ├── application/    # Application services (use cases)
│   ├── domain/        # Domain models and value objects
│   ├── handler/       # HTTP handlers
│   ├── repository/    # Data access
│   ├── dto/          # Data transfer objects
│   └── database/     # Database connections
├── migrations/        # Database migrations
└── docs/             # Backend documentation

frontend/
├── src/
│   ├── shared/
│   │   ├── services/  # API clients and TanStack Query hooks
│   │   └── store/     # Zustand stores
│   └── features/      # Feature-specific components
│       ├── tp/        # TP Workspace
│       ├── assessment/# Assessment Designer
│       └── ...
```

## Important Files

### Architecture Documentation
- `docs/ARCHITECTURE_FREEZE_V2.md` - **READ THIS FIRST** - Complete architecture specification
- `docs/DATABASE_SCHEMA_FREEZE_V1.md` - Database schema specification
- `docs/SPRINT_3.5_EXECUTION_SEQUENCE.md` - Current sprint implementation plan
- `docs/REPOSITORY_MODIFICATION_MAP.md` - Governance for AI coding agents

### Domain Documentation
- `docs/DOMAIN_INVARIANT_CATALOG.md` - All domain invariants
- `docs/EVENT_STORMING_ARCHITECTURE_REVIEW.md` - Domain events and policies

### Frontend Documentation
- `docs/centralized/` - Centralized documentation files

## Working with This Project

### Before Making Changes
1. **READ ARCHITECTURE_FREEZE_V2.md** - This is the single source of truth
2. Check existing patterns in the codebase
3. Follow layered architecture strictly
4. No CQRS/Event Sourcing/Event Bus
5. Solo developer context - be realistic about complexity

### Backend Changes
- Always follow: Handler → Service → Repository pattern
- Business logic only in Domain and Application layers
- Repository never accessed directly from Handler
- Add tests for new functionality
- Update migration files for schema changes

### Frontend Changes
- Use existing components before creating new ones
- Follow TanStack Query + Zustand patterns
- Teacher-centric UX design (simple, efficient)
- Indonesian language for UI, English for code comments
- TypeScript strict mode

### Database Changes
- All schema changes via migration files
- Both .up.sql and .down.sql required
- Test migrations on fresh database
- Update DATABASE_SCHEMA_FREEZE_V1.md for significant changes

## Solo Developer Context

### What This Means
- No guaranteed response time for PRs/reviews
- Priority on maintainability over cleverness
- Focus on essential features over nice-to-haves
- Documentation is critical for continuity
- Code must be self-explanatory

### Best Practices for AI Assistance
- Provide complete, working solutions
- Include comprehensive comments
- Follow existing patterns exactly
- Don't introduce unnecessary complexity
- Consider the maintenance burden
- Test thoroughly before suggesting changes

### Communication Style
- Be direct and concise
- Explain architectural decisions
- Highlight trade-offs
- Provide alternatives when appropriate
- Reference existing code patterns

## Specific Commands

### Backend Development
```bash
cd backend
go test ./... -v           # Run tests
go fmt ./...               # Format code
go vet ./...               # Vet code
go build -o bin/api cmd/api/main.go  # Build
```

### Frontend Development
```bash
cd frontend
npm install                # Install dependencies
npm run dev               # Dev server
npm run build             # Production build
npm run lint              # Lint
npm run type-check        # TypeScript check
```

### Database
```bash
# Apply migration
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /path/to/migration.up.sql

# Rollback migration
podman exec nusa-postgres psql -U nusa_user -d nusa_db -f /path/to/migration.down.sql
```

## Quality Standards

### Code Quality
- Clean, readable code
- Proper error handling
- No hardcoded values (use environment variables)
- Comprehensive comments for complex logic
- Tests for critical functionality

### Documentation Quality
- Update CHANGELOG.md for significant changes
- Update inline code comments
- Add README for new features if needed
- Keep architecture docs in sync

### Testing Quality
- Unit tests for business logic
- Integration tests for APIs
- Manual testing for UI changes
- Test migrations thoroughly

## Common Pitfalls to Avoid

### Architecture Violations
- ❌ Adding CQRS patterns
- ❌ Implementing Event Sourcing
- ❌ Adding Event/Command buses
- ❌ Business logic in handlers
- ❌ Repository access from handlers
- ❌ New bounded contexts without discussion

### Code Quality Issues
- ❌ Copy-pasting without understanding
- ❌ Ignoring existing patterns
- ❌ Over-engineering simple problems
- ❌ Not testing before suggesting
- ❌ Breaking existing functionality

### Solo Developer Considerations
- ❌ Suggesting complex solutions that are hard to maintain
- ❌ Introducing many new dependencies
- ❌ Not considering the learning curve
- ❌ Ignoring documentation burden
- ❌ Assuming team context (there is no team)

## When to Ask for Clarification

If you're unsure about:
- Architecture compliance - Ask and reference ARCHITECTURE_FREEZE_V2.md
- Domain logic - Ask for clarification on Kurikulum Merdeka concepts
- Implementation approach - Suggest alternatives with trade-offs
- Scope of changes - Confirm if change is appropriate for solo project

## Success Criteria

Your assistance is successful if:
- ✅ Code follows architecture freeze exactly
- ✅ Solution is maintainable by solo developer
- ✅ Tests are included and passing
- ✅ Documentation is updated
- ✅ No unnecessary complexity introduced
- ✅ Existing functionality is not broken
- ✅ Domain logic is correct for Kurikulum Merdeka

---

**Remember**: This is a solo-developed project with strict architecture governance. Quality and maintainability are more important than speed or cleverness. Always prioritize the long-term viability of the codebase.
