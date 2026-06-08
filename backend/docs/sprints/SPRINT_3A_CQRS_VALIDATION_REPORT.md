# Sprint 3A CQRS Validation Report

## Executive Summary

This report documents the CQRS (Command Query Responsibility Segregation) validation for the Achievement domain in Sprint 3A. The current implementation uses a runtime calculation approach without persistence. This report designs a future projection-based Achievement system that would improve performance, scalability, and enable advanced analytics while maintaining the current architecture's benefits.

## Document Date
2024-06-08

## Current Achievement Implementation

### Architecture
- **Pattern:** Runtime Calculation (On-Demand)
- **Persistence:** None (Achievement is calculated at request time)
- **Data Source:** Evaluations, Assessments, TPs
- **Calculation Service:** `AchievementService` in domain layer

### Current Flow
```
Request → AchievementHandler → AchievementService → Repository Queries → Domain Calculation → Response
```

### Advantages
1. **Data Consistency:** Always reflects latest evaluation data
2. **Simplicity:** No synchronization complexity
3. **Storage Efficiency:** No redundant data storage
4. **Flexibility:** Easy to modify calculation logic

### Disadvantages
1. **Performance:** Calculation overhead on each request
2. **Scalability:** Doesn't scale well with large datasets
3. **Analytics:** Limited historical trend analysis
4. **Caching:** No built-in caching mechanism

---

## Future Projection-Based Achievement Design

### Architecture Overview

The proposed CQRS implementation separates command and query responsibilities for Achievement:

```
Command Side (Write Model):
- Evaluation Events → Event Store → Projection Builders → Read Models

Query Side (Read Model):
- Achievement Projections → Optimized Read Database → Fast Queries
```

### Component Design

#### 1. Event Sourcing for Evaluations

**Purpose:** Capture all evaluation changes as immutable events

**Event Types:**
```go
type EvaluationEvent struct {
    EventID      string
    EventType    string // "EVALUATION_CREATED", "EVALUATION_UPDATED", "EVALUATION_DELETED"
    AggregateID  string // Evidence ID
    Data         EvaluationEventData
    Timestamp    time.Time
    Version      int64
}

type EvaluationEventData struct {
    EvaluationID   string
    StudentID      string
    EvidenceID     string
    AssessmentID   string
    Score          float64
    Feedback       string
    TeacherID      string
    RevisionNo     int
    // ... other evaluation fields
}
```

**Event Store Schema:**
```sql
CREATE TABLE evaluation_events (
    event_id UUID PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    aggregate_id UUID NOT NULL,
    data JSONB NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    version BIGINT NOT NULL,
    INDEX idx_aggregate_id (aggregate_id),
    INDEX idx_timestamp (timestamp)
);
```

---

#### 2. Projection Builders

**Purpose:** Transform events into optimized read models

**Achievement Projection Builder:**
```go
type AchievementProjectionBuilder struct {
    eventStore      EventStore
    readDB          *sql.DB
    logger          *zap.Logger
}

func (b *AchievementProjectionBuilder) Handle(event EvaluationEvent) error {
    switch event.EventType {
    case "EVALUATION_CREATED":
        return b.handleEvaluationCreated(event)
    case "EVALUATION_UPDATED":
        return b.handleEvaluationUpdated(event)
    case "EVALUATION_DELETED":
        return b.handleEvaluationDeleted(event)
    }
}

func (b *AchievementProjectionBuilder) handleEvaluationCreated(event EvaluationEvent) error {
    // 1. Extract evaluation data
    // 2. Recalculate student achievement for the relevant TP
    // 3. Update achievement projection table
    // 4. Update competency progress projection
    // 5. Update class achievement projection
}
```

---

#### 3. Read Models (Projections)

**Achievement Projection:**
```sql
CREATE TABLE achievement_projection (
    student_id UUID NOT NULL,
    tp_id UUID NOT NULL,
    achievement_score DECIMAL(5,2),
    achievement_level VARCHAR(20),
    criteria_met JSONB,
    last_evaluation_id UUID,
    last_updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (student_id, tp_id),
    INDEX idx_tp_id (tp_id),
    INDEX idx_last_updated (last_updated_at)
);
```

**Competency Progress Projection:**
```sql
CREATE TABLE competency_progress_projection (
    student_id UUID NOT NULL,
    subject_id UUID NOT NULL,
    phase_id UUID,
    total_tps INTEGER,
    completed_tps INTEGER,
    overall_progress DECIMAL(5,2),
    progress_items JSONB,
    last_updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (student_id, subject_id, phase_id),
    INDEX idx_subject_phase (subject_id, phase_id)
);
```

**Class Achievement Projection:**
```sql
CREATE TABLE class_achievement_projection (
    class_id UUID NOT NULL,
    subject_id UUID NOT NULL,
    tp_id UUID NOT NULL,
    average_score DECIMAL(5,2),
    student_count INTEGER,
    achievement_distribution JSONB,
    last_updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (class_id, subject_id, tp_id),
    INDEX idx_class_subject (class_id, subject_id)
);
```

**Achievement Summary Projection:**
```sql
CREATE TABLE achievement_summary_projection (
    student_id UUID NOT NULL,
    class_id UUID NOT NULL,
    report_period_start DATE NOT NULL,
    report_period_end DATE NOT NULL,
    overall_achievement DECIMAL(5,2),
    subject_breakdown JSONB,
    competency_trends JSONB,
    recommendations JSONB,
    last_updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (student_id, class_id, report_period_start, report_period_end)
);
```

---

#### 4. Query Handlers

**Achievement Query Handler:**
```go
type AchievementQueryHandler struct {
    readDB *sql.DB
}

func (h *AchievementQueryHandler) GetStudentAchievement(
    ctx context.Context,
    studentID string,
    tpID string,
) (*domain.Achievement, error) {
    // Direct query from projection table
    query := `
        SELECT achievement_score, achievement_level, criteria_met, 
               last_evaluation_id, last_updated_at
        FROM achievement_projection
        WHERE student_id = $1 AND tp_id = $2
    `
    // ... execute query and map to domain model
}
```

---

### CQRS Implementation Strategy

#### Phase 1: Event Sourcing Foundation (Future Sprint)
1. Implement Event Store infrastructure
2. Create Evaluation Event types
3. Implement Event Publisher/Subscriber
4. Add event emission to Evaluation service

#### Phase 2: Projection Infrastructure (Future Sprint)
1. Create Projection Builder framework
2. Implement projection tables
3. Create Projection Manager service
4. Add projection rebuild capability

#### Phase 3: Query Optimization (Future Sprint)
1. Implement Query Handlers
2. Update Achievement API to use projections
3. Add projection refresh endpoints
4. Implement projection caching

#### Phase 4: Advanced Analytics (Future Sprint)
1. Implement historical trend queries
2. Add comparative analysis features
3. Implement predictive analytics
4. Create real-time achievement dashboards

---

## Migration Strategy

### Hybrid Approach (Recommended)

**Phase 1: Dual-Write**
- Keep current runtime calculation as fallback
- Add event emission to Evaluation service
- Build projections in background
- Compare results for validation

**Phase 2: Gradual Migration**
- Route read traffic to projections for cached data
- Use runtime calculation for fresh data or cache misses
- Monitor performance and accuracy

**Phase 3: Full CQRS**
- Deprecate runtime calculation
- Use projections for all queries
- Implement projection refresh on demand

### Rollback Plan
- Maintain runtime calculation service as backup
- Add feature flag to switch between implementations
- Keep event store for audit trail even if projections fail

---

## Performance Considerations

### Expected Improvements
- **Query Latency:** 10-100x faster (cached projections vs. runtime calculation)
- **Throughput:** Higher due to reduced database load
- **Scalability:** Better horizontal scaling capability

### Trade-offs
- **Write Latency:** Slightly increased due to event processing
- **Storage:** Additional storage for events and projections
- **Complexity:** Increased system complexity

### Optimization Strategies
1. **Batch Processing:** Process events in batches for efficiency
2. **Async Processing:** Use message queues for event handling
3. **Partitioning:** Partition projections by class/subject
4. **Caching:** Add application-level caching for hot data

---

## Data Consistency

### Eventual Consistency Model
- Projections updated asynchronously after events
- Acceptable delay: < 5 seconds for most use cases
- Critical paths: Force projection refresh on demand

### Consistency Checks
- Compare projection data with runtime calculation
- Scheduled reconciliation jobs
- Alert on significant discrepancies

### Conflict Resolution
- Event versioning prevents conflicts
- Last-write-wins for concurrent updates
- Manual intervention for data anomalies

---

## Security Considerations

### Event Store Security
- Event immutability (append-only)
- Audit trail for all changes
- Role-based access to event replay

### Projection Security
- Read-only access for query handlers
- Separate database credentials for read/write
- Encryption for sensitive achievement data

---

## Monitoring and Observability

### Metrics to Track
- Event processing latency
- Projection freshness (lag time)
- Query performance (p50, p95, p99)
- Projection rebuild success rate

### Alerting
- Event processing backlog
- Projection sync failures
- Query performance degradation
- Data consistency anomalies

---

## Conclusion

The proposed CQRS implementation for Achievement provides a robust foundation for scaling the system while maintaining data integrity and enabling advanced analytics. The hybrid migration approach ensures a smooth transition with minimal risk.

### Recommendations
1. **Implement in phases** to manage complexity
2. **Maintain runtime calculation** as fallback during transition
3. **Invest in monitoring** to ensure system health
4. **Document projection logic** for maintainability

### Next Steps
1. Design detailed event schemas
2. Prototype projection builders
3. Plan database schema changes
4. Define migration timeline

### Implementation Priority
- **High:** Event Sourcing foundation
- **High:** Core projections (achievement, competency progress)
- **Medium:** Advanced analytics projections
- **Low:** Real-time dashboards

---

## Appendix: Comparison Matrix

| Aspect | Current (Runtime) | Future (CQRS) |
|--------|-------------------|---------------|
| Query Performance | Variable (depends on data volume) | Consistent (cached) |
| Data Freshness | Always fresh | Eventual consistency |
| Storage Efficiency | High | Medium (events + projections) |
| Scalability | Limited | High |
| Complexity | Low | High |
| Analytics Capability | Limited | Advanced |
| Implementation Effort | Complete | Significant |
| Maintenance | Simple | Moderate |
