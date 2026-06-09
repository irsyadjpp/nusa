# Devin AI Assistant Guidelines

This document provides specific guidelines for Devin AI assistants when working on the NUSA Platform project.

## Project Context

**NUSA Platform** is a solo-developed education management system for Indonesian schools implementing Kurikulum Merdeka 2026.

### Solo Developer Context
- **Single Maintainer**: No team, no backup, individual responsibility
- **Time Constraints**: Limited development time (full-time job elsewhere)
- **Quality Standards**: Production-quality code despite solo development
- **Long-term Vision**: Code must be maintainable for years by one person

## Critical Architecture Governance

### Architecture Freeze - MANDATORY READING
**`docs/ARCHITECTURE_FREEZE_V2.md`** is the single source of truth for all architecture decisions.

### ❌ STRICTLY FORBIDDEN
Never suggest or implement:
- **CQRS** (Command Query Responsibility Segregation)
- **Event Sourcing**
- **Event Bus** (RabbitMQ only for AI workflows, not domain events)
- **Command Bus** or **Query Bus**
- **Read Models** or **Projections**
- **Microservices** architecture
- **New bounded contexts** outside Kurikulum Merdeka domain
- **Complex design patterns** without explicit justification

### ✅ REQUIRED ARCHITECTURE
Must always follow:
- **Layered Architecture**: Handler → Service → Repository → PostgreSQL
- **DDD Lite**: Aggregates, bounded contexts, domain logic
- **Modular Monolith**: Single codebase, modular organization
- **Standard Patterns**: Conventional Go and React patterns

## Domain Knowledge: Kurikulum Merdeka

### Educational Domain Flow
```
CP (Capaian Pembelajaran) 
  ↓
TP (Tujuan Pembelajaran with embedded KKTP)
  ↓
ATP (Alur Tujuan Pembelajaran)
  ↓
Modul Ajar (Teaching Modules)
  ↓
Assessment
  ↓
Evidence (Student Work)
  ↓
Evaluation (Teacher Assessment with revision tracking)
  ↓
Achievement (Runtime calculation, no persistence)
  ↓
Narrative Report
```

### Critical Domain Concepts

**KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)**:
- Embedded as Value Object in TP domain
- Contains: mastery thresholds, performance indicators, minimum requirements
- NOT a separate table or aggregate
- Part of TP domain model

**Version Management**:
- TP: VersionNo increments on updates
- Assessment: TPVersionNo + SuccessCriteriaSnapshot for consistency
- Evaluation: RevisionNo increments on each teacher change
- All versioned entities need historical tracking

**Achievement Service**:
- Runtime calculation ONLY (NO database persistence)
- Calculated from Evidence + Evaluation data
- Services: student achievement, competency progress, class achievement
- Always calculate on-demand, never cache results

## Code Structure Standards

### Backend (Go) - Strict Layer Separation

```
backend/internal/
├── application/    # Application services (use case orchestration)
├── domain/        # Domain models, value objects, invariants
├── handler/       # HTTP handlers (request/response ONLY)
├── repository/    # Data access (SQL ONLY)
├── dto/          # Data transfer objects
└── database/     # Database connections
```

**CRITICAL RULE**: Never skip layers. Handler never calls Repository directly. Business logic only in Domain and Application layers.

### Frontend (TypeScript/React) - State Management

```
frontend/src/
├── shared/
│   ├── services/  # API clients, TanStack Query hooks
│   └── store/     # Zustand stores
└── features/      # Feature-specific components
```

**CRITICAL RULE**: TanStack Query for server state, Zustand for client state.

## Code Pattern Examples

### Backend - Correct Layer Implementation

#### Handler Layer (HTTP Only)
```go
// ✅ CORRECT: HTTP request/response only
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
    // Never validate business rules here
    // Never calculate domain logic here
    // Never call repository directly
}
```

#### Service Layer (Business Logic Orchestration)
```go
// ✅ CORRECT: Orchestrate domain and repository
func (s *TPService) CreateTP(req CreateTPRequest) (*TP, error) {
    // 1. Validate request DTO
    if err := req.Validate(); err != nil {
        return nil, err
    }
    
    // 2. Create domain object
    tp := NewTP(req)
    
    // 3. Apply domain validation (business rules)
    if err := tp.Validate(); err != nil {
        return nil, err
    }
    
    // 4. Persist via repository
    if err := s.tpRepo.Save(tp); err != nil {
        return nil, err
    }
    
    return tp, nil
}
```

#### Repository Layer (Data Access Only)
```go
// ✅ CORRECT: SQL operations only
func (r *TPRepository) Save(tp *TP) error {
    query := `INSERT INTO tp (id, cp_id, kktp, version_no, created_at) 
              VALUES ($1, $2, $3, $4, $5) 
              RETURNING id, created_at`
    return r.db.QueryRow(query, tp.ID, tp.CPID, tp.KKTP, tp.VersionNo, tp.CreatedAt).Scan(&tp.ID, &tp.CreatedAt)
}

// ❌ WRONG: Business logic in repository
func (r *TPRepository) Save(tp *TP) error {
    // Never validate business rules here
    // Never calculate domain logic here
    // Never apply domain invariants here
}
```

#### Domain Layer (Business Rules)
```go
// ✅ CORRECT: Domain invariants and behavior
type TP struct {
    ID        string
    CPID      string
    KKTP      KKTPCriteria // Value Object
    VersionNo int
    CreatedAt time.Time
}

func (tp *TP) Validate() error {
    // Domain invariants
    if tp.CPID == "" {
        return errors.New("CP ID is required")
    }
    
    if tp.VersionNo < 1 {
        return errors.New("Version number must be >= 1")
    }
    
    // Validate embedded value object
    if err := tp.KKTP.Validate(); err != nil {
        return fmt.Errorf("KKTP validation failed: %w", err)
    }
    
    return nil
}
```

### Frontend - State Management

#### Server State (TanStack Query)
```typescript
// ✅ CORRECT: Use TanStack Query
const { data: tps, isLoading, error } = useQuery({
  queryKey: ['tps'],
  queryFn: () => tpService.getTPs(),
  staleTime: 5 * 60 * 1000, // 5 minutes
  refetchOnWindowFocus: false
})

// ❌ WRONG: Manual state management
const [tps, setTPs] = useState([])
useEffect(() => {
  fetchTPs().then(data => setTPs(data))
}, [])
```

#### Client State (Zustand)
```typescript
// ✅ CORRECT: Zustand store
interface TPStore {
  selectedTP: TP | null
  isEditing: boolean
  setSelectedTP: (tp: TP) => void
  setIsEditing: (editing: boolean) => void
}

const useTPStore = create<TPStore>((set) => ({
  selectedTP: null,
  isEditing: false,
  setSelectedTP: (tp) => set({ selectedTP: tp }),
  setIsEditing: (editing) => set({ isEditing: editing })
}))
```

## Solo Developer Implications

### What This Means for Devin

**Prioritize**:
1. **Simplicity**: Simple solutions over clever ones
2. **Maintainability**: Code easy to understand and modify
3. **Self-documenting**: Clear code reduces need for external docs
4. **Testing**: Comprehensive tests reduce debugging burden
5. **Pattern Consistency**: Follow existing patterns exactly

**Avoid**:
1. **Unnecessary Complexity**: Each layer of complexity is maintenance cost
2. **New Dependencies**: Each new dependency requires ongoing maintenance
3. **Clever Code**: Clever code is hard to understand and modify
4. **Team Assumptions**: Don't assume code review or backup developer
5. **Breaking Changes**: Hard to undo, need careful consideration

### Example Decision Framework

**Scenario**: User asks for better performance on TP listing

❌ **Bad Devin Response**:
"I'll implement Redis caching with pub/sub invalidation and also add a read model projection for faster queries."

**Why Bad**: 
- Introduces forbidden patterns (read models)
- Adds operational complexity (Redis)
- Solo developer can't maintain this complexity
- Violates architecture freeze

✅ **Good Devin Response**:
"For TP listing performance, let's:
1. First, check if database indexes exist on query fields
2. Review current query execution plan
3. Consider TanStack Query caching (staleTime configuration)
4. If still slow, add pagination to reduce data transfer
5. Only consider major changes after measuring actual bottleneck

Simple optimizations first. Complex caching adds maintenance burden for solo developer."

## Important Documentation Files

### Must Read (In Priority Order)
1. **`docs/ARCHITECTURE_FREEZE_V2.md`** - Architecture specification (READ FIRST)
2. **`docs/REPOSITORY_MODIFICATION_MAP.md`** - AI coding agent governance
3. **`docs/DOMAIN_INVARIANT_CATALOG.md`** - All domain invariants
4. **`docs/DATABASE_SCHEMA_FREEZE_V1.md`** - Database schema specification
5. **`docs/SPRINT_3.5_EXECUTION_SEQUENCE.md`** - Current sprint plan

### Reference Documentation
- **`docs/EVENT_STORMING_ARCHITECTURE_REVIEW.md`** - Domain events and policies
- **`CHANGELOG.md`** - Project history and changes
- **`CONTRIBUTING.md`** - Contribution guidelines

## Before Making Code Changes

### Pre-Change Checklist
- [ ] Read relevant sections of ARCHITECTURE_FREEZE_V2.md
- [ ] Searched codebase for similar existing patterns
- [ ] Considered solo developer maintenance burden
- [ ] Planned documentation updates (CHANGELOG.md, inline comments)
- [ ] Thought about testing approach and coverage
- [ ] Verified no forbidden patterns are being introduced

### Stop Signs (Red Flags)
🚩 User requests CQRS implementation → STOP, explain architecture freeze
🚩 Need for Event Sourcing → STOP, explain not allowed
🚩 Suggestion for microservices → STOP, explain modular monolith
🚩 New bounded context proposal → STOP, discuss with maintainer
🚩 Complex pattern suggestion → STOP, consider simpler alternative
🚩 "Ignore architecture for this" → STOP, architecture freeze is mandatory

## Database Migration Rules

### Migration File Structure
```sql
-- 000001_add_tp_set_aggregate.up.sql
CREATE TABLE tp_set (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    school_id UUID NOT NULL REFERENCES schools(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_tp_set_school_id ON tp_set(school_id);

-- 000001_add_tp_set_aggregate.down.sql
DROP INDEX IF EXISTS idx_tp_set_school_id;
DROP TABLE IF EXISTS tp_set;
```

### Migration Process
1. Always create both `.up.sql` and `.down.sql` files
2. Number sequentially: 000001, 000002, 000003
3. Test on fresh database before committing
4. Update DATABASE_SCHEMA_FREEZE_V1.md for significant changes
5. Include indexes for performance optimization

## Testing Requirements

### Backend Testing
```bash
cd backend
go test ./... -v
```
- **Domain tests**: Business rules and invariants
- **Service tests**: Use case orchestration
- **Repository tests**: Data access operations

### Frontend Testing
```bash
cd frontend
npm test
```
- **Component tests**: UI components
- **Hook tests**: Custom hooks
- **Integration tests**: User flows

## Devin-Specific Working Patterns

### When Devin Should Read Files
- Always read ARCHITECTURE_FREEZE_V2.md before architectural changes
- Read existing implementation files before adding similar code
- Read domain files before modifying business logic
- Read migration files before database changes

### When Devin Should Execute Commands
- Run tests after code changes
- Run linting/formatting tools
- Build to check for compilation errors
- Run migrations to verify they work

### When Devin Should Ask for Clarification
- Architecture compliance is unclear
- Domain logic is ambiguous
- Multiple implementation approaches exist
- Change scope is large
- Breaking changes might be introduced

## Communication Guidelines

### When Providing Code
- Explain the architectural decision
- Reference existing patterns in codebase
- Highlight maintenance implications
- Suggest testing strategy
- Consider simpler alternatives

### When Explaining Trade-offs
- Be explicit about pros and cons
- Consider solo developer context
- Reference architecture freeze
- Suggest incremental approach

### When Uncertain
- Ask questions rather than assume
- Provide multiple approaches with trade-offs
- Reference relevant documentation
- Be explicit about uncertainty

## Quality Standards

### Code Quality
- Clean, readable, self-explanatory
- Proper error handling at all layers
- No hardcoded values (use environment variables)
- Comprehensive comments for complex logic
- Follow language idioms (Go conventions, React best practices)

### Architecture Quality
- Strict adherence to layered architecture
- Business logic only in appropriate layers
- No forbidden patterns (CQRS, Event Sourcing, etc.)
- Respect existing domain boundaries
- Follow DDD Lite constraints

### Documentation Quality
- Update CHANGELOG.md for significant changes
- Add inline comments for complex logic
- Update relevant architecture documentation
- Keep docs in sync with code changes

## Common Mistakes to Avoid

### Architecture Violations
- ❌ Suggesting CQRS/Event Sourcing patterns
- ❌ Business logic in handlers
- ❌ Direct repository access from handlers
- ❌ New bounded contexts without discussion
- ❌ Breaking layered architecture

### Code Quality Issues
- ❌ Over-engineering simple problems
- ❌ Ignoring existing code patterns
- ❌ Insufficient testing
- ❌ Inadequate documentation
- ❌ Breaking existing functionality

### Solo Developer Blind Spots
- ❌ Assuming team development context
- ❌ Introducing high-maintenance solutions
- ❌ Not considering learning curve
- ❌ Under-documenting complex logic
- ❌ Adding unnecessary dependencies

## Success Criteria

Devin's assistance is successful if:
- ✅ Code follows ARCHITECTURE_FREEZE_V2.md exactly
- ✅ Solution is maintainable by solo developer
- ✅ Tests are comprehensive and passing
- ✅ Documentation is updated and accurate
- ✅ No forbidden patterns introduced
- ✅ Existing functionality is preserved
- ✅ Domain logic is correct for Kurikulum Merdeka
- ✅ Code is self-explanatory with good comments

## Example Devin Responses

### ❌ Bad Response
User: "I need better performance for TP queries."
Devin: "I'll implement CQRS with separate read models and use Redis caching. This will significantly improve performance."

**Why Bad**: 
- Introduces forbidden patterns (CQRS, read models)
- Adds operational complexity (Redis)
- Violates architecture freeze
- Not maintainable by solo developer

### ✅ Good Response
User: "I need better performance for TP queries."
Devin: "Let's improve TP query performance step by step:

1. **First, let's analyze the current bottleneck**:
   - Check if indexes exist on query fields
   - Review query execution plan
   - Measure current query performance

2. **Simple optimizations**:
   - Add database indexes if missing
   - Optimize the SQL query
   - Implement pagination to reduce data transfer

3. **Frontend caching**:
   - TanStack Query already configured with caching
   - Increase staleTime for TP data (e.g., 10 minutes)
   - Consider selective refetching

4. **If still slow after optimizations**:
   - Discuss with maintainer before major changes
   - Consider simpler caching strategies
   - Avoid complex solutions that add maintenance burden

Let me start by checking the current TP query implementation and database indexes."

**Why Good**: 
- Respects architecture freeze
- Suggests simple optimizations first
- Considers solo developer context
- Step-by-step approach
- Acknowledges maintenance burden

---

**Remember for Devin**: This is a solo-developed project. Your assistance should enable one person to maintain production-quality code long-term. Simplicity, maintainability, and strict adherence to architecture governance are more important than cleverness or theoretical "best practices." Always prioritize the solo developer's ability to understand, modify, and maintain the code.
