# Gemini AI Assistant Guidelines

This document provides specific guidelines for Google Gemini AI assistants when working on the NUSA Platform project.

## Project Context

**NUSA Platform** is a solo-developed education management system for Indonesian schools implementing Kurikulum Merdeka 2026.

### Solo Developer Reality
- **One Maintainer**: No team, no backup, no corporate support
- **Full-time Job**: Limited time for development (evenings/weekends)
- **Production Quality**: Despite being solo, code must be production-ready
- **Long-term Vision**: Code must be maintainable for years by one person

## Critical Constraints

### Architecture Freeze
The project is under strict architecture governance defined in:
- **`docs/ARCHITECTURE_FREEZE_V2.md`** - READ THIS FIRST

### Forbidden Patterns (DO NOT SUGGEST)
- ❌ CQRS (Command Query Responsibility Segregation)
- ❌ Event Sourcing
- ❌ Event Bus (except RabbitMQ for AI workflows)
- ❌ Command/Query Bus
- ❌ Read Models/Projections
- ❌ Microservices
- ❌ New bounded contexts outside Kurikulum Merdeka
- ❌ Complex design patterns without explicit approval

### Required Patterns (MUST FOLLOW)
- ✅ Layered Architecture: Handler → Service → Repository → PostgreSQL
- ✅ DDD Lite (aggregates, bounded contexts, domain logic)
- ✅ Modular Monolith
- ✅ Standard Go and React patterns

## Domain Knowledge: Kurikulum Merdeka

### Educational Workflow
```
CP (Learning Outcomes) → TP (Learning Objectives + KKTP) → ATP (Annual Flow) → Modul Ajar (Lesson Plans) → Assessment → Evidence → Evaluation → Achievement → Reports
```

### Key Concepts for Code

**KKTP (Kriteria Ketuntasan Tujuan Pembelajaran)**:
- Embedded as Value Object in TP
- Contains: mastery thresholds, performance indicators, minimum requirements
- NOT a separate entity, part of TP domain

**Version Tracking**:
- TP entities have VersionNo
- Assessment has TPVersionNo + SuccessCriteriaSnapshot
- Evaluation has RevisionNo (increments on each change)
- Historical consistency is critical

**Achievement Service**:
- Runtime calculation only (NO persistence)
- Calculates from Evidence + Evaluation data
- Services: student achievement, competency progress, class achievement

## File Structure Context

### Backend (Go)
```
backend/internal/
├── application/    # Use case orchestration
├── domain/        # Domain models + invariants (business rules)
├── handler/       # HTTP only (request/response)
├── repository/    # Database only (data access)
├── dto/          # Data transfer objects
└── database/     # Database connections
```

**Key Rule**: Never skip layers. Handler never calls Repository directly.

### Frontend (TypeScript/React)
```
frontend/src/
├── shared/
│   ├── services/  # API clients + TanStack Query hooks
│   └── store/     # Zustand stores
└── features/      # Feature-specific components
```

**Key Rule**: Use TanStack Query for server state, Zustand for client state.

## Code Pattern Examples

### Backend - Handler (HTTP Only)
```go
// ✅ CORRECT
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

// ❌ WRONG - Business logic in handler
func (h *TPHandler) CreateTP(c *gin.Context) {
    // Don't validate business rules here
    // Don't calculate anything domain-related
}
```

### Backend - Service (Business Logic)
```go
// ✅ CORRECT
func (s *TPService) CreateTP(req CreateTPRequest) (*TP, error) {
    // 1. Validate request
    if err := req.Validate(); err != nil {
        return nil, err
    }
    
    // 2. Create domain object
    tp := NewTP(req)
    
    // 3. Apply domain validation
    if err := tp.Validate(); err != nil {
        return nil, err
    }
    
    // 4. Persist
    if err := s.tpRepo.Save(tp); err != nil {
        return nil, err
    }
    
    return tp, nil
}
```

### Backend - Repository (Data Only)
```go
// ✅ CORRECT
func (r *TPRepository) Save(tp *TP) error {
    query := `INSERT INTO tp (id, cp_id, kktp, version_no) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(query, tp.ID, tp.CPID, tp.KKTP, tp.VersionNo)
    return err
}

// ❌ WRONG - Business logic in repository
func (r *TPRepository) Save(tp *TP) error {
    // Don't validate
    // Don't calculate
    // Just SQL
}
```

### Frontend - Server State
```typescript
// ✅ CORRECT - TanStack Query
const { data: tps, isLoading } = useQuery({
  queryKey: ['tps'],
  queryFn: () => tpService.getTPs(),
  staleTime: 5 * 60 * 1000
})

// ❌ WRONG - Manual state management
const [tps, setTPs] = useState([])
useEffect(() => {
  fetch('/api/tps').then(res => res.json()).then(setTPs)
}, [])
```

## Solo Developer Implications

### What This Means for Your Assistance

**Prioritize**:
1. **Simplicity**: Simple solutions over clever ones
2. **Maintainability**: Code easy to understand and modify
3. **Documentation**: Comments and docs are critical
4. **Testing**: Tests reduce debugging burden
5. **Patterns**: Follow existing patterns exactly

**Avoid**:
1. **Complexity**: Over-engineering creates maintenance burden
2. **New Dependencies**: Each new dependency is maintenance cost
3. **Clever Code**: Clever code is hard to maintain
4. **Assumptions**: Don't assume team context (no team exists)
5. **Breaking Changes**: Hard to undo, careful with changes

### Example Decision Framework

**Scenario**: Adding caching for TP data

❌ **Bad Approach**:
"Suggest implementing Redis caching with complex invalidation strategy using pub/sub"

✅ **Good Approach**:
"TanStack Query has built-in caching with staleTime configuration. For simple cases, set staleTime to 5-10 minutes. For complex caching needs, discuss with maintainer first as it adds operational complexity."

## Important Files to Reference

### Must Read (Priority Order)
1. **`docs/ARCHITECTURE_FREEZE_V2.md`** - Architecture specification
2. **`docs/REPOSITORY_MODIFICATION_MAP.md`** - AI governance
3. **`docs/DOMAIN_INVARIANT_CATALOG.md`** - Domain rules
4. **`docs/DATABASE_SCHEMA_FREEZE_V1.md`** - Database schema

### Reference
- **`docs/EVENT_STORMING_ARCHITECTURE_REVIEW.md`** - Domain events
- **`CHANGELOG.md`** - Project history
- **`CONTRIBUTING.md`** - Contribution guidelines

## Before Suggesting Code

### Checklist
- [ ] Read ARCHITECTURE_FREEZE_V2.md relevant sections
- [ ] Searched codebase for similar patterns
- [ ] Considered solo developer maintenance burden
- [ ] Planned documentation updates
- [ ] Thought about testing approach
- [ ] Verified no forbidden patterns suggested

### Red Flags (Stop and Reconsider)
🚩 "Implement CQRS for better separation"
🚩 "Use Event Sourcing for audit trail"
🚩 "Split into microservices"
🚩 "Add new bounded context X"
🚩 "Complex pattern Y would be better"
🚩 "Ignore architecture freeze for this"

## Testing Requirements

### Backend
```bash
cd backend
go test ./... -v
```
- Domain logic: Unit tests
- Services: Integration tests
- Repositories: Database tests

### Frontend
```bash
cd frontend
npm test
```
- Components: Unit tests
- Hooks: Unit tests
- Flows: Integration tests

## Communication Style

### When Providing Solutions
- Be explicit about trade-offs
- Reference existing patterns in codebase
- Explain maintenance implications
- Suggest testing approach
- Consider alternatives

### When Uncertain
- Ask questions rather than assume
- Provide multiple approaches with pros/cons
- Reference relevant documentation
- Be explicit about what you don't know

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
- No forbidden patterns
- Follow existing domain boundaries
- Respect DDD Lite constraints

### Documentation Quality
- Update CHANGELOG.md for significant changes
- Add inline comments for complex logic
- Update relevant documentation files
- Keep docs in sync with code

## Gemini-Specific Considerations

### Leverage Gemini Strengths
- **Code Generation**: Generate following patterns exactly
- **Code Review**: Check against architecture freeze
- **Documentation**: Help maintain docs in sync
- **Testing**: Generate comprehensive tests
- **Refactoring**: Suggest improvements within constraints

### Avoid Gemini Pitfalls
- Don't suggest patterns outside architecture freeze
- Don't over-engineer for "best practices" that add complexity
- Don't assume team development context
- Don't introduce unnecessary abstractions
- Don't suggest "modern" patterns that violate constraints

## Success Criteria

Your assistance is successful if:
- ✅ Code follows ARCHITECTURE_FREEZE_V2.md exactly
- ✅ Solution is maintainable by solo developer
- ✅ Tests are included and passing
- ✅ Documentation is updated
- ✅ No forbidden patterns introduced
- ✅ Existing functionality preserved
- ✅ Domain logic correct for Kurikulum Merdeka

## Example Responses

### ❌ Bad Response
"For better scalability, I suggest implementing CQRS with separate read and write models. This will improve performance and separation of concerns."

**Why Bad**: Violates architecture freeze, introduces complexity solo developer can't maintain.

### ✅ Good Response
"The current architecture uses layered pattern which is appropriate for this scale. For the performance concern you mentioned, I suggest:
1. First, add database indexes on the slow queries
2. Consider TanStack Query caching (already configured)
3. If still slow, discuss with maintainer before major changes
Performance improvements should be measured before optimizing."

**Why Good**: Respects architecture, suggests simple solutions first, acknowledges solo developer context.

---

**Remember**: This is a solo-developed project. Your assistance should enable one person to maintain production-quality code long-term. Simplicity, maintainability, and adherence to established patterns are more important than cleverness or theoretical "best practices."
