# Future CQRS Expansion Plan for Achievement System

**Date:** 2026-06-07
**Status:** Design Document (Not for Implementation)
**Purpose:** Design projection-based Achievement system for future scalability

---

## Executive Summary

The current Achievement system uses runtime calculation via AchievementService. This design document outlines a future CQRS (Command Query Responsibility Segregation) pattern implementation that would replace runtime calculations with pre-computed projections for improved performance and scalability at scale.

**Current State:** Runtime calculation (on-demand)
**Future State:** Projection-based (event-sourced, pre-computed)
**Migration Strategy:** Gradual rollout with feature flags

---

## Current Architecture

### Achievement Service (Runtime Calculation)
- **CalculateStudentAchievement()** - Calculates on-demand from evaluations
- **CalculateCompetencyProgress()** - Aggregates across TPs on-demand
- **GenerateAchievementSummary()** - Summarizes achievements on-demand
- **GenerateClassAchievement()** - Aggregates class data on-demand

### Advantages
- Simple implementation
- Always up-to-date
- No data duplication
- Easy to maintain

### Disadvantages
- Performance degrades with data volume
- Complex queries become slow
- No caching layer
- High database load during peak times
- Difficult to optimize for specific read patterns

---

## Future CQRS Architecture

### Command Side (Write Model)
Handles all write operations and emits events.

#### Commands
- `CreateEvaluationCommand` - Create new evaluation
- `UpdateEvaluationCommand` - Update existing evaluation
- `DeleteEvaluationCommand` - Delete evaluation
- `CreateAssessmentCommand` - Create new assessment
- `UpdateAssessmentCommand` - Update assessment

#### Events
- `EvaluationCreatedEvent` - Emitted when evaluation is created
- `EvaluationUpdatedEvent` - Emitted when evaluation is updated
- `EvaluationDeletedEvent` - Emitted when evaluation is deleted
- `AssessmentCreatedEvent` - Emitted when assessment is created
- `AssessmentUpdatedEvent` - Emitted when assessment is updated

#### Event Store
- Append-only log of all domain events
- Immutable event history
- Supports event replay for projection rebuilding

---

### Query Side (Read Model)
Handles all read operations using pre-computed projections.

#### Projections

##### 1. StudentProgressProjection
**Purpose:** Track student progress across all TPs and competencies

**Schema:**
```sql
CREATE TABLE student_progress_projection (
    id UUID PRIMARY KEY,
    student_id UUID NOT NULL,
    tp_id UUID NOT NULL,
    tp_version_no INTEGER NOT NULL,
    competency_code VARCHAR(50),
    overall_score DECIMAL(5,2),
    performance_level VARCHAR(20),
    mastery_status VARCHAR(20),
    evidence_count INTEGER DEFAULT 0,
    evaluation_count INTEGER DEFAULT 0,
    last_evaluation_at TIMESTAMP,
    last_calculated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(student_id, tp_id, tp_version_no)
);

CREATE INDEX idx_student_progress_student ON student_progress_projection(student_id);
CREATE INDEX idx_student_progress_tp ON student_progress_projection(tp_id);
CREATE INDEX idx_student_progress_competency ON student_progress_projection(competency_code);
```

**Update Triggers:**
- On `EvaluationCreatedEvent` - Recalculate student progress for affected TP
- On `EvaluationUpdatedEvent` - Recalculate student progress for affected TP
- On `EvaluationDeletedEvent` - Recalculate student progress for affected TP

##### 2. CompetencyMatrixProjection
**Purpose:** Track competency mastery across students in a class

**Schema:**
```sql
CREATE TABLE competency_matrix_projection (
    id UUID PRIMARY KEY,
    class_id UUID NOT NULL,
    subject_id UUID NOT NULL,
    phase_id UUID NOT NULL,
    competency_code VARCHAR(50),
    total_students INTEGER DEFAULT 0,
    mastered_count INTEGER DEFAULT 0,
    in_progress_count INTEGER DEFAULT 0,
    not_started_count INTEGER DEFAULT 0,
    exceeding_count INTEGER DEFAULT 0,
    average_score DECIMAL(5,2),
    last_calculated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(class_id, subject_id, phase_id, competency_code)
);

CREATE INDEX idx_competency_matrix_class ON competency_matrix_projection(class_id);
CREATE INDEX idx_competency_matrix_subject ON competency_matrix_projection(subject_id);
CREATE INDEX idx_competency_matrix_phase ON competency_matrix_projection(phase_id);
```

**Update Triggers:**
- On `StudentProgressProjection` update - Recalculate competency matrix for affected class
- Batch update job - Recalculate all matrices nightly

##### 3. SchoolAnalyticsProjection
**Purpose:** Aggregate analytics for school-wide reporting

**Schema:**
```sql
CREATE TABLE school_analytics_projection (
    id UUID PRIMARY KEY,
    school_id UUID NOT NULL,
    subject_id UUID,
    phase_id UUID,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    total_students INTEGER DEFAULT 0,
    total_evaluations INTEGER DEFAULT 0,
    average_achievement DECIMAL(5,2),
    mastery_rate DECIMAL(5,2),
    top_performing_classes JSONB,
    areas_for_improvement JSONB,
    last_calculated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(school_id, subject_id, phase_id, period_start, period_end)
);

CREATE INDEX idx_school_analytics_school ON school_analytics_projection(school_id);
CREATE INDEX idx_school_analytics_period ON school_analytics_projection(period_start, period_end);
```

**Update Triggers:**
- Batch update job - Recalculate nightly
- On-demand refresh - Manual trigger for reports

---

## Projection Update Strategies

### 1. Event-Driven Updates (Real-time)
- **Trigger:** Domain events
- **Latency:** < 100ms
- **Use Case:** Student progress, individual achievements
- **Implementation:** Event handlers update projections immediately

### 2. Batch Updates (Scheduled)
- **Trigger:** Cron jobs
- **Latency:** Hourly/Daily
- **Use Case:** Competency matrices, school analytics
- **Implementation:** Scheduled jobs recalculate aggregates

### 3. On-Demand Refresh
- **Trigger:** Manual API call
- **Latency:** Variable (depends on data volume)
- **Use Case:** Report generation, data corrections
- **Implementation:** API endpoint triggers recalculation

---

## Migration Strategy

### Phase 1: Dual-Write (6-8 weeks)
- Keep existing AchievementService
- Add event emission to write operations
- Implement projection update handlers
- Populate projections from existing data
- Feature flag to switch read source

### Phase 2: Read-Through (4-6 weeks)
- Enable feature flag for read operations
- Route reads to projections
- Keep AchievementService as fallback
- Monitor performance and data consistency
- Fix any discrepancies

### Phase 3: Cutover (2-4 weeks)
- Remove AchievementService from read path
- Deprecate runtime calculation endpoints
- Update all clients to use projection endpoints
- Remove feature flags

### Phase 4: Cleanup (2 weeks)
- Remove AchievementService
- Clean up unused code
- Update documentation
- Archive event store retention policy

---

## Benefits of CQRS Approach

### Performance
- **Read Performance:** 10-100x faster (pre-computed data)
- **Write Performance:** Unchanged (event emission is fast)
- **Scalability:** Independent scaling of read/write
- **Caching:** Projections can be cached easily

### Maintainability
- **Separation of Concerns:** Read/write logic separated
- **Event Sourcing:** Complete audit trail
- **Replayability:** Projections can be rebuilt from events
- **Testing:** Easier to test read/write independently

### Flexibility
- **Multiple Projections:** Different views for different use cases
- **Event Replay:** Fix bugs by replaying events
- **Schema Evolution:** Easy to add new projections
- **Analytics:** Rich event data for analysis

---

## Risks and Mitigations

### Risk 1: Eventual Consistency
- **Description:** Projections may be slightly stale
- **Mitigation:** Real-time updates for critical paths, SLA monitoring
- **Acceptable Latency:** < 1 second for student progress, < 1 hour for analytics

### Risk 2: Data Inconsistency
- **Description:** Projections may diverge from source of truth
- **Mitigation:** Reconciliation jobs, data validation checks, manual refresh capability
- **Detection:** Daily consistency checks between projections and source

### Risk 3: Complexity
- **Description:** CQRS adds architectural complexity
- **Mitigation:** Clear documentation, team training, gradual rollout
- **Team Impact:** Requires CQRS expertise, event-driven architecture knowledge

### Risk 4: Event Store Growth
- **Description:** Event store grows indefinitely
- **Mitigation:** Retention policy, event archiving, compaction strategies
- **Retention:** 7 years for audit, 1 year for replay

---

## Implementation Timeline

| Phase | Duration | Deliverables |
|-------|----------|--------------|
| Phase 1: Design | 2 weeks | Architecture document, schema design |
| Phase 2: Event Infrastructure | 4 weeks | Event store, event handlers, event emission |
| Phase 3: Projection Implementation | 6 weeks | All 3 projections, update handlers |
| Phase 4: Migration - Dual Write | 6-8 weeks | Dual-write implementation, data population |
| Phase 5: Migration - Read Through | 4-6 weeks | Feature flags, monitoring, validation |
| Phase 6: Migration - Cutover | 2-4 weeks | Full cutover, client updates |
| Phase 7: Cleanup | 2 weeks | Code cleanup, documentation |
| **Total** | **26-32 weeks** | **Full CQRS implementation** |

---

## Success Criteria

### Performance
- Read latency < 100ms for student progress
- Read latency < 500ms for competency matrix
- Read latency < 1s for school analytics
- Write latency unchanged (< 50ms)

### Reliability
- 99.9% projection update success rate
- < 1 minute data freshness for critical paths
- < 1 hour data freshness for analytics
- Zero data loss events

### Maintainability
- Clear event schema documentation
- Automated projection rebuild capability
- Monitoring and alerting for projection health
- Runbooks for common scenarios

---

## Conclusion

The CQRS pattern with projection-based Achievement system provides significant performance and scalability benefits for the NUSA platform. The migration should be approached gradually with careful monitoring and validation. The current runtime calculation system should remain in place until the CQRS system is fully validated and proven.

**Recommendation:** Implement CQRS when the platform reaches 10,000+ students or when read performance becomes a bottleneck. The current runtime calculation system is sufficient for the current scale.

**Next Steps:**
1. Monitor current system performance
2. Define performance thresholds for CQRS migration
3. Build proof-of-concept for one projection
4. Validate benefits before full implementation
