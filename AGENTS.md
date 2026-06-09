# AI Agent Guidelines

This document provides comprehensive guidelines for AI agents working on the NUSA Platform project.

## Project Overview

**NUSA Platform** is a solo-developed education management system designed for Indonesian schools implementing Kurikulum Merdeka 2026.

### Project Identity
- **Maintainer**: Solo developer with full-time job elsewhere
- **Scale**: Production-quality code, solo-maintained
- **Domain**: Education technology for Kurikulum Merdeka
- **Architecture**: Modular Monolith with DDD Lite
- **Philosophy**: Maintainability over cleverness, simplicity over complexity

## Critical Architecture Rules

### ❌ STRICTLY FORBIDDEN
Do not suggest or implement:
- **CQRS** (Command Query Responsibility Segregation)
- **Event Sourcing**
- **Event Bus** (except RabbitMQ for AI workflows only)
- **Command Bus** or **Query Bus**
- **Read Models** or **Projections**
- **Microservices** architecture
- **New domains** outside Kurikulum Merdeka scope
- **Complex design patterns** without explicit justification

### ✅ REQUIRED PATTERNS
Must follow:
- **Layered Architecture**: Handler → Service → Repository → PostgreSQL
- **DDD Lite**: Aggregates, bounded contexts, domain logic
- **Modular Monolith**: Single codebase, modular organization
- **Standard Go**: Conventional Go patterns
- **Standard React**: idiomatic React/TypeScript patterns

## Domain Knowledge: Kurikulum Merdeka

### Educational Domain Chain
```
CP (Capaian Pembelajaran) → TP (Tujuan Pembelajaran) → ATP (Alur Tujuan Pembelajaran) → Modul Ajar → Assessment → Evidence → Evaluation → Achievement → Narrative Report
```

### Key Domain Concepts

**TP (Tujuan Pembelajaran)**:
- Learning objectives
- Contains embedded KKTP as Value Object
- Has version tracking
- Links to CP (Capaian Pembelajaran)

**KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)**:
- Mastery thresholds
- Performance indicators
- Minimum requirements
- Embedded in TP, not separate entity

**Assessment**:
- References TP via TPID and TPVersionNo
- Contains SuccessCriteriaSnapshot for historical consistency
- Links to Evidence for student work

**Evidence & Evaluation**:
- Evidence: Student work artifacts
- Evaluation: Teacher assessment with revision tracking
- RevisionNo increments for each change
- Teacher feedback preserved across revisions

**Achievement**:
- Runtime calculation only (no persistence)
- Services for student achievement, competency progress, class achievement
- Calculated from Evidence and Evaluation data

## File Structure & Standards

### Backend Structure (Go)
```
backend/internal/
├── application/     # Application services (use cases)
├── domain/         # Domain models, value objects, invariants
├── handler/        # HTTP handlers (request/response only)
├── repository/     # Data access (database only)
├── dto/           # Data transfer objects
└── database/      # Database connections and migrations
```

### Frontend Structure (TypeScript/React)
```
frontend/src/
├── shared/
│   ├── services/   # API clients, TanStack Query hooks
│   └── store/      # Zustand stores
└── features/
    ├── tp/        # TP Workspace components
    ├── assessment/# Assessment Designer components
    └── ...
```

## Layered Architecture Rules

### Handler Layer (HTTP)
```go
// ✅ CORRECT: HTTP only
func (h *TPHandler) CreateTP(c *gin.Context) {
    var req CreateTPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    tp, err := h.tpService.CreateTP(req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(201, tp)
}

// ❌ WRONG: Business logic in handler
func (h *TPHandler) CreateTP(c *gin.Context) {
    // Don't validate business rules here
    // Don't calculate domain logic here
}
```

### Service Layer (Business Logic)
```go
// ✅ CORRECT: Orchestrate domain and repository
func (s *TPService) CreateTP(req CreateTPRequest) (*TP, error) {
    // Validate request
    if err := req.Validate(); err != nil {
        return nil, err
    }
    
    // Create domain object
    tp := NewTP(req)
    
    // Apply domain logic
    if err := tp.Validate(); err != nil {
        return nil, err
    }
    
    // Persist via repository
    if err := s.tpRepo.Save(tp); err != nil {
        return nil, err
    }
    
    return tp, nil
}
```

### Repository Layer (Data Access)
```go
// ✅ CORRECT: Database operations only
func (r *TPRepository) Save(tp *TP) error {
    query := `INSERT INTO tp (...) VALUES (...) RETURNING id`
    return r.db.QueryRow(query, ...).Scan(&tp.ID)
}

// ❌ WRONG: Business logic in repository
func (r *TPRepository) Save(tp *TP) error {
    // Don't validate business rules here
    // Don't calculate domain logic here
}
```

### Domain Layer (Business Rules)
```go
// ✅ CORRECT: Domain invariants and behavior
type TP struct {
    ID        string
    CPID      string
    KKTP      KKTPCriteria // Value Object
    VersionNo int
}

func (tp *TP) Validate() error {
    if tp.CPID == "" {
        return errors.New("CP ID is required")
    }
    
    if err := tp.KKTP.Validate(); err != nil {
        return err
    }
    
    return nil
}
```

## Frontend Patterns

### TanStack Query for Server State
```typescript
// ✅ CORRECT: Use TanStack Query
const { data: tps, isLoading, error } = useQuery({
  queryKey: ['tps'],
  queryFn: () => tpService.getTPs(),
  staleTime: 5 * 60 * 1000, // 5 minutes
})

// ❌ WRONG: Manual state management
const [tps, setTPs] = useState([])
useEffect(() => {
  fetchTPs().then(data => setTPs(data))
}, [])
```

### Zustand for Client State
```typescript
// ✅ CORRECT: Zustand store
interface TPStore {
  selectedTP: TP | null
  setSelectedTP: (tp: TP) => void
}

const useTPStore = create<TPStore>((set) => ({
  selectedTP: null,
  setSelectedTP: (tp) => set({ selectedTP: tp })
}))
```

## Database Migration Rules

### Migration File Format
```sql
-- 000001_add_tp_set_table.up.sql
CREATE TABLE tp_set (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 000001_add_tp_set_table.down.sql
DROP TABLE IF EXISTS tp_set;
```

### Migration Process
1. Always create both `.up.sql` and `.down.sql` files
2. Test on fresh database before committing
3. Update `DATABASE_SCHEMA_FREEZE_V1.md` for significant changes
4. Number migrations sequentially: 000001, 000002, etc.

## Important Documentation Files

### Must Read (In Order of Priority)
1. **`docs/ARCHITECTURE_FREEZE_V2.md`** - Architecture specification (READ FIRST)
2. **`docs/REPOSITORY_MODIFICATION_MAP.md`** - AI coding agent governance
3. **`docs/DOMAIN_INVARIANT_CATALOG.md`** - All domain invariants
4. **`docs/DATABASE_SCHEMA_FREEZE_V1.md`** - Database schema
5. **`docs/SPRINT_3.5_EXECUTION_SEQUENCE.md`** - Current sprint plan

### Reference Documentation
- **`docs/EVENT_STORMING_ARCHITECTURE_REVIEW.md`** - Domain events and policies
- **`CHANGELOG.md`** - Project history
- **`CONTRIBUTING.md`** - Contribution guidelines

## Coding Standards

### Go Backend
- Use `gofmt` for formatting
- Exported functions: PascalCase
- Internal functions: camelCase
- Constants: UPPER_SNAKE_CASE
- Error handling: Always handle errors explicitly
- No magic numbers: Use constants

### TypeScript Frontend
- TypeScript strict mode
- Components: PascalCase
- Hooks: camelCase with `use` prefix
- Functions: camelCase
- Constants: UPPER_SNAKE_CASE
- Variables: camelCase

## Testing Requirements

### Backend Tests
```bash
cd backend
go test ./... -v
```
- Unit tests for domain logic
- Integration tests for services
- Repository tests for data access

### Frontend Tests
```bash
cd frontend
npm test
```
- Component tests for UI
- Hook tests for custom hooks
- Integration tests for flows

## Solo Developer Considerations

### What "Solo Developer" Means
- One person maintaining the codebase
- No team for code review
- Limited time (full-time job elsewhere)
- Documentation is critical
- Code must be self-maintainable

### Implications for AI Assistance
- **Simplicity**: Favor simple solutions over clever ones
- **Documentation**: Comments and docs are essential
- **Patterns**: Follow existing patterns exactly
- **Testing**: Comprehensive tests reduce debugging burden
- **Dependencies**: Minimize new dependencies
- **Learning Curve**: Consider solo developer's time to understand

## Before Making Changes

### Checklist
1. **READ ARCHITECTURE_FREEZE_V2.md** - Non-negotiable
2. Search codebase for existing patterns
3. Consider solo developer maintenance burden
4. Plan for documentation updates
5. Think about long-term viability

### Red Flags
- 🚩 "We could add CQRS for better scalability" → STOP
- 🚩 "Let's implement Event Sourcing for audit trail" → STOP
- 🚩 "Microservices would be better for modularity" → STOP
- 🚩 "Complex design pattern X would solve this" → STOP
- 🚩 "New bounded context Y for this feature" → STOP

## Communication Guidelines

### When Suggesting Changes
- Explain the "why" clearly
- Reference existing patterns
- Consider alternatives with trade-offs
- Highlight maintenance burden
- Suggest testing strategy

### When Uncertain
- Ask questions rather than assume
- Suggest multiple approaches with trade-offs
- Reference relevant documentation
- Be explicit about uncertainties

## Quality Standards

### Code Quality
- Clean, readable, self-explanatory
- Proper error handling
- No hardcoded values
- Comprehensive comments for complex logic
- Follow language idioms

### Architecture Quality
- Strict layered architecture
- No business logic in wrong layer
- No forbidden patterns (CQRS, etc.)
- Follow existing domain boundaries
- Respect DDD Lite constraints

### Documentation Quality
- Update CHANGELOG.md for significant changes
- Add inline comments for complex logic
- Update relevant documentation files
- Keep docs in sync with code

## Common Mistakes to Avoid

### Architecture Violations
- Adding CQRS/Event Sourcing patterns
- Business logic in handlers
- Repository access from handlers
- New domains without approval
- Breaking layered architecture

### Code Quality Issues
- Over-engineering simple problems
- Ignoring existing patterns
- Not testing thoroughly
- Insufficient documentation
- Breaking existing functionality

### Solo Developer Blind Spots
- Assuming team context
- Introducing high-complexity solutions
- Not considering maintenance burden
- Under-documenting complex logic
- Adding unnecessary dependencies

## Success Criteria

Your assistance is successful if:
- ✅ Architecture freeze is followed exactly
- ✅ Code is maintainable by solo developer
- ✅ Tests are comprehensive and passing
- ✅ Documentation is updated and accurate
- ✅ No forbidden patterns introduced
- ✅ Existing functionality preserved
- ✅ Domain logic correct for Kurikulum Merdeka

## Emergency Contacts

If you encounter issues:
1. Check ARCHITECTURE_FREEZE_V2.md first
2. Search for similar patterns in codebase
3. Consider alternatives with trade-offs
4. Ask for clarification if uncertain

---

**Remember**: This is a solo-developed project with strict architecture governance. Your assistance should prioritize maintainability, simplicity, and adherence to established patterns over cleverness or complexity. The goal is production-quality code that one person can maintain long-term.
