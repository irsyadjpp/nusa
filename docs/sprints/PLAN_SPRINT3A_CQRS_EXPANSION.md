# Sprint 3A CQRS Expansion Plan

## Document Date
2024-06-08

## Purpose
This document outlines the detailed expansion plan for implementing CQRS (Command Query Responsibility Segregation) in the NUSA Education Operating System, starting with the Achievement domain and expanding to other domains over time.

## Current State
- Achievement uses runtime calculation (no persistence)
- All other domains use traditional CRUD operations
- No event sourcing or projection infrastructure exists
- Single database for both read and write operations

## Expansion Roadmap

### Phase 1: Foundation Infrastructure (Sprint 4 - 4 weeks)

#### 1.1 Event Store Implementation
**Objective:** Build the core event sourcing infrastructure

**Tasks:**
- Design event store schema
- Implement EventStore interface
- Create EventPublisher/Subscriber
- Add event serialization/deserialization
- Implement event versioning
- Add event replay capability

**Deliverables:**
- `internal/eventstore/event_store.go`
- `internal/eventstore/event_publisher.go`
- `internal/eventstore/event_subscriber.go`
- Migration script for event tables
- Event store unit tests

**Effort:** 2 weeks

---

#### 1.2 Projection Framework
**Objective:** Build the projection building infrastructure

**Tasks:**
- Design projection builder interface
- Implement ProjectionManager
- Create projection rebuild mechanism
- Add projection health monitoring
- Implement projection consistency checks
- Add projection versioning

**Deliverables:**
- `internal/projection/builder.go`
- `internal/projection/manager.go`
- `internal/projection/rebuild.go`
- Projection monitoring metrics
- Projection framework unit tests

**Effort:** 2 weeks

---

### Phase 2: Achievement Domain CQRS (Sprint 5 - 6 weeks)

#### 2.1 Evaluation Event Sourcing
**Objective:** Convert Evaluation to event-sourced aggregate

**Tasks:**
- Define Evaluation event types
- Implement Evaluation aggregate with event sourcing
- Add event emission to Evaluation service
- Create Evaluation event handlers
- Implement event replay for Evaluation
- Add event store migration for existing Evaluations

**Deliverables:**
- `internal/domain/evaluation_events.go`
- Updated `internal/service/evaluation_service.go`
- Evaluation event migration script
- Event replay tool
- Integration tests

**Effort:** 3 weeks

---

#### 2.2 Achievement Projections
**Objective:** Implement Achievement read models

**Tasks:**
- Design achievement projection schemas
- Implement AchievementProjectionBuilder
- Create CompetencyProgressProjectionBuilder
- Create ClassAchievementProjectionBuilder
- Create AchievementSummaryProjectionBuilder
- Implement projection refresh logic
- Add projection API endpoints

**Deliverables:**
- Projection migration scripts
- `internal/projection/achievement_builder.go`
- `internal/projection/competency_progress_builder.go`
- `internal/projection/class_achievement_builder.go`
- `internal/projection/achievement_summary_builder.go`
- Projection API handlers
- Projection unit tests

**Effort:** 3 weeks

---

#### 2.3 Hybrid Implementation
**Objective:** Implement dual-write with fallback to runtime calculation

**Tasks:**
- Update Achievement API to use projections
- Implement projection cache miss fallback
- Add projection refresh endpoint
- Implement projection comparison validation
- Add feature flag for CQRS toggle
- Monitor projection accuracy

**Deliverables:**
- Updated `internal/handler/achievement_handler.go`
- Projection refresh endpoint
- Monitoring dashboards
- Validation scripts
- Performance benchmarks

**Effort:** 2 weeks (overlap with Phase 3)

---

### Phase 3: Assessment Domain CQRS (Sprint 6 - 5 weeks)

#### 3.1 Assessment Event Sourcing
**Objective:** Convert Assessment to event-sourced aggregate

**Tasks:**
- Define Assessment event types
- Implement Assessment aggregate with event sourcing
- Add event emission to Assessment service
- Create Assessment event handlers
- Implement event replay for Assessment
- Add event store migration for existing Assessments

**Deliverables:**
- `internal/domain/assessment_events.go`
- Updated `internal/service/assessment_service.go`
- Assessment event migration script
- Integration tests

**Effort:** 3 weeks

---

#### 3.2 Assessment Projections
**Objective:** Implement Assessment read models

**Tasks:**
- Design assessment projection schemas
- Implement AssessmentProjectionBuilder
- Create AssessmentStatusProjectionBuilder
- Create AssessmentAnalyticsProjectionBuilder
- Implement projection refresh logic
- Add projection API endpoints

**Deliverables:**
- Projection migration scripts
- `internal/projection/assessment_builder.go`
- Projection API handlers
- Unit tests

**Effort:** 2 weeks

---

### Phase 4: Evidence Domain CQRS (Sprint 7 - 5 weeks)

#### 4.1 Evidence Event Sourcing
**Objective:** Convert Evidence to event-sourced aggregate

**Tasks:**
- Define Evidence event types
- Implement Evidence aggregate with event sourcing
- Add event emission to Evidence service
- Create Evidence event handlers
- Implement event replay for Evidence
- Add event store migration for existing Evidence

**Deliverables:**
- `internal/domain/evidence_events.go`
- Updated `internal/service/evidence_service.go`
- Evidence event migration script
- Integration tests

**Effort:** 3 weeks

---

#### 4.2 Evidence Projections
**Objective:** Implement Evidence read models

**Tasks:**
- Design evidence projection schemas
- Implement EvidenceProjectionBuilder
- Create EvidenceStatusProjectionBuilder
- Create EvidenceAnalyticsProjectionBuilder
- Implement projection refresh logic
- Add projection API endpoints

**Deliverables:**
- Projection migration scripts
- `internal/projection/evidence_builder.go`
- Projection API handlers
- Unit tests

**Effort:** 2 weeks

---

### Phase 5: TP Domain CQRS (Sprint 8 - 4 weeks)

#### 5.1 TP Event Sourcing
**Objective:** Convert TP to event-sourced aggregate

**Tasks:**
- Define TP event types
- Implement TP aggregate with event sourcing
- Add event emission to TP service
- Create TP event handlers
- Implement event replay for TP
- Add event store migration for existing TPs

**Deliverables:**
- `internal/domain/tp_events.go`
- Updated `internal/service/tp_service.go`
- TP event migration script
- Integration tests

**Effort:** 2 weeks

---

#### 5.2 TP Projections
**Objective:** Implement TP read models

**Tasks:**
- Design TP projection schemas
- Implement TPProjectionBuilder
- Create TPVersionProjectionBuilder
- Create TPAnalyticsProjectionBuilder
- Implement projection refresh logic
- Add projection API endpoints

**Deliverables:**
- Projection migration scripts
- `internal/projection/tp_builder.go`
- Projection API handlers
- Unit tests

**Effort:** 2 weeks

---

### Phase 6: Advanced Features (Sprint 9 - 4 weeks)

#### 6.1 Real-Time Projections
**Objective:** Implement real-time projection updates

**Tasks:**
- Implement WebSocket support for projection updates
- Add projection change notifications
- Implement projection subscription mechanism
- Create real-time dashboard components
- Add projection streaming API

**Deliverables:**
- WebSocket handlers
- Real-time projection updates
- Frontend real-time components
- Streaming API documentation

**Effort:** 2 weeks

---

#### 6.2 Advanced Analytics
**Objective:** Implement advanced analytics on projections

**Tasks:**
- Implement trend analysis queries
- Add comparative analysis features
- Create predictive analytics models
- Implement anomaly detection
- Add custom report generation

**Deliverables:**
- Analytics service
- Analytics API endpoints
- Analytics dashboard
- Custom report builder

**Effort:** 2 weeks

---

## Database Schema Changes

### Event Store Tables
```sql
-- Generic event store
CREATE TABLE events (
    event_id UUID PRIMARY KEY,
    event_type VARCHAR(100) NOT NULL,
    aggregate_type VARCHAR(100) NOT NULL,
    aggregate_id UUID NOT NULL,
    data JSONB NOT NULL,
    metadata JSONB,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    version BIGINT NOT NULL,
    INDEX idx_aggregate (aggregate_type, aggregate_id),
    INDEX idx_timestamp (timestamp),
    INDEX idx_event_type (event_type)
);

-- Event snapshots for faster replay
CREATE TABLE event_snapshots (
    aggregate_type VARCHAR(100) NOT NULL,
    aggregate_id UUID NOT NULL,
    version BIGINT NOT NULL,
    data JSONB NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (aggregate_type, aggregate_id, version)
);
```

### Projection Tables
```sql
-- Achievement projections (already designed in CQRS Validation Report)
-- Assessment projections
-- Evidence projections
-- TP projections
-- Additional analytics projections
```

---

## API Changes

### New Endpoints
- `POST /api/v1/projections/{projection_type}/refresh` - Force projection refresh
- `GET /api/v1/projections/{projection_type}/status` - Get projection status
- `POST /api/v1/projections/{projection_type}/rebuild` - Rebuild projection from events
- `GET /api/v1/events/{aggregate_type}/{aggregate_id}` - Get event stream
- `POST /api/v1/events/{aggregate_type}/{aggregate_id}/replay` - Replay events

### Modified Endpoints
- Achievement endpoints will use projections by default
- Assessment endpoints will use projections by default
- Evidence endpoints will use projections by default
- TP endpoints will use projections by default

---

## Monitoring and Observability

### Metrics to Track
- Event processing rate (events/second)
- Event processing latency (p50, p95, p99)
- Projection freshness (lag time in seconds)
- Projection rebuild duration
- Projection consistency score
- Query performance (with vs without projections)

### Dashboards to Create
- Event processing dashboard
- Projection health dashboard
- Query performance dashboard
- Data consistency dashboard

### Alerts to Configure
- Event processing backlog (> 1000 events)
- Projection lag (> 30 seconds)
- Projection rebuild failure
- Data consistency anomaly (> 5% discrepancy)
- Query performance degradation (> 2x slower)

---

## Risk Mitigation

### Technical Risks
1. **Event Store Performance**
   - Mitigation: Implement event partitioning, use efficient serialization
   - Fallback: Batch event processing, async event handling

2. **Projection Consistency**
   - Mitigation: Implement consistency checks, automated reconciliation
   - Fallback: Manual reconciliation tools, runtime calculation fallback

3. **Migration Complexity**
   - Mitigation: Phased migration, dual-write period
   - Fallback: Rollback procedures, feature flags

4. **Increased Complexity**
   - Mitigation: Comprehensive documentation, training
   - Fallback: Simplified monitoring, automated tooling

### Operational Risks
1. **Increased Storage Requirements**
   - Mitigation: Event retention policies, compression
   - Fallback: Archive old events, optimize storage

2. **Learning Curve**
   - Mitigation: Training sessions, documentation
   - Fallback: Expert support, gradual adoption

---

## Success Criteria

### Phase 1 Success Criteria
- Event store handles 1000 events/second
- Projection framework supports 10 projection types
- Event replay completes within 5 minutes for 10,000 events

### Phase 2 Success Criteria
- Achievement projections achieve 95% accuracy vs runtime calculation
- Query latency reduced by 80% for achievement endpoints
- Projection freshness < 5 seconds

### Phase 3-5 Success Criteria
- All domain projections achieve 95% accuracy
- Overall query latency reduced by 70%
- System maintains 99.9% uptime during migration

### Phase 6 Success Criteria
- Real-time updates delivered within 1 second
- Advanced analytics queries complete within 10 seconds
- Custom reports generated within 30 seconds

---

## Resource Requirements

### Development Team
- 2 Backend Developers (full-time)
- 1 DevOps Engineer (part-time)
- 1 Database Administrator (part-time)

### Infrastructure
- Additional database storage for events (estimated 2x current)
- Additional database storage for projections (estimated 0.5x current)
- Message queue for event processing (RabbitMQ/Kafka)
- Monitoring infrastructure (Prometheus/Grafana)

### Timeline
- Total Duration: 24 weeks (6 sprints)
- Phase 1: 4 weeks
- Phase 2: 6 weeks
- Phase 3: 5 weeks
- Phase 4: 5 weeks
- Phase 5: 4 weeks
- Phase 6: 4 weeks

---

## Conclusion

This CQRS expansion plan provides a structured approach to implementing event sourcing and projections across the NUSA Education Operating System. The phased approach minimizes risk while delivering incremental value. The plan prioritizes the Achievement domain (highest impact) and expands to other domains systematically.

### Key Recommendations
1. **Start with Achievement domain** - highest impact, lowest risk
2. **Maintain dual-write period** - ensures smooth transition
3. **Invest in monitoring** - critical for operational success
4. **Document everything** - essential for long-term maintainability
5. **Train the team** - CQRS requires mindset shift

### Next Steps
1. Review and approve this expansion plan
2. Allocate resources for Phase 1
3. Set up development environment for event sourcing
4. Begin Phase 1 implementation

### Expected Outcomes
- **Performance:** 70-80% improvement in query latency
- **Scalability:** 10x increase in throughput capacity
- **Analytics:** Advanced analytics capabilities
- **Reliability:** Improved data consistency and auditability
