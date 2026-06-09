# Sprint 3A Migration Impact Report

**Date:** 2026-06-07
**Migration Version:** 000003_add_success_criteria_and_refactor_assessment
**Status:** Ready for Review

---

## Executive Summary

This migration implements the domain model changes required for Sprint 3A, including:
- Embedding SuccessCriteria (KKTP) into TP
- Replacing Assessment's ModulAjar reference with TP reference
- Adding revision history to Evaluation
- Creating optimal indexes for new query patterns

**Risk Level:** MEDIUM
**Estimated Downtime:** 5-10 minutes
**Rollback Capability:** YES

---

## Migration Components

### Migration A: Add success_criteria to TP table

**Changes:**
- Add `success_criteria JSONB` column to `tp` table
- Add GIN index for JSONB queries

**Impact:**
- **Tables Affected:** 1 (tp)
- **Rows Affected:** 0 (new column, no data migration)
- **Breaking Changes:** NO
- **Application Impact:** LOW - backward compatible (nullable)

**Validation:**
- Column added successfully
- Index created successfully
- Existing data unaffected

---

### Migration B: Assessment table changes

**Changes:**
- Add `tp_id UUID` (references tp)
- Add `tp_version_no INTEGER`
- Add `success_criteria_snapshot JSONB`
- Data migration from `modul_ajar_id` to `tp_id` via ATP chain
- Drop `modul_ajar_id` column
- Drop and recreate indexes

**Impact:**
- **Tables Affected:** 1 (assessments)
- **Rows Affected:** ALL existing assessments (data migration)
- **Breaking Changes:** YES - application must be updated to use tp_id
- **Application Impact:** HIGH - requires coordinated deployment

**Data Migration Strategy:**
1. Add new columns (nullable initially)
2. Migrate data: assessments → modul_ajar → atp → tp
3. Validate all rows have tp_id
4. Make tp_id NOT NULL
5. Drop modul_ajar_id column
6. Update indexes

**Risk Mitigation:**
- Data migration is reversible (down migration restores modul_ajar_id)
- Validation checks ensure no data loss
- Application must be deployed after migration completes

---

### Migration C: Evaluation table updates

**Changes:**
- Add `teacher_feedback TEXT`
- Add `revision_no INTEGER`
- Add `created_at TIMESTAMP`
- Add `updated_at TIMESTAMP`
- Backfill timestamps for existing records

**Impact:**
- **Tables Affected:** 1 (evaluations)
- **Rows Affected:** ALL existing evaluations (timestamp backfill)
- **Breaking Changes:** NO - new columns are nullable or have defaults
- **Application Impact:** LOW - backward compatible

**Validation:**
- All existing records get created_at = evaluated_at
- All existing records get updated_at = evaluated_at
- New records will have proper timestamps

---

### Migration D: Create optimal indexes

**Changes:**
- Add GIN index on tp.success_criteria
- Add composite index on tp(tp_set_id, sequence_number)
- Add index on assessments.tp_id
- Add composite index on assessments(tp_id, tp_version_no)
- Add GIN index on assessments.success_criteria_snapshot
- Add index on evaluations(evidence_id, revision_no)
- Add partial index on evaluations.teacher_feedback
- Add indexes on evaluations.created_at, evaluations.updated_at
- Add composite index on evaluations(student_id, evidence_id, revision_no DESC)

**Impact:**
- **Tables Affected:** 3 (tp, assessments, evaluations)
- **Rows Affected:** 0 (index creation only)
- **Breaking Changes:** NO
- **Application Impact:** NONE - performance improvement only

**Performance Impact:**
- **Positive:** Faster queries on new access patterns
- **Negative:** Minimal write overhead for index maintenance
- **Storage:** Estimated 5-10% increase in index size

---

## Data Risk Analysis

### High Risk Areas

1. **Assessment Data Migration (Migration B)**
   - **Risk:** Data loss if ATP/TP chain is broken
   - **Mitigation:** Validation checks before dropping modul_ajar_id
   - **Fallback:** Down migration restores original structure
   - **Monitoring:** Log records that fail migration

2. **Application Compatibility**
   - **Risk:** Application deployed before migration completes
   - **Mitigation:** Coordinated deployment with application team
   - **Fallback:** Feature flags to toggle new behavior
   - **Monitoring:** Application error logs

### Medium Risk Areas

1. **Index Creation**
   - **Risk:** Long-running index creation on large tables
   - **Mitigation:** Run during low-traffic period
   - **Fallback:** Drop indexes if performance issues
   - **Monitoring:** Query performance metrics

2. **Timestamp Backfill**
   - **Risk:** Inaccurate timestamps for historical data
   - **Mitigation:** Use evaluated_at as source of truth
   - **Fallback:** Manual correction if needed
   - **Monitoring:** Data quality checks

### Low Risk Areas

1. **New Column Addition**
   - **Risk:** Schema changes fail
   - **Mitigation:** Transaction rollback
   - **Fallback:** Down migration
   - **Monitoring:** Migration logs

---

## Migration Plan

### Pre-Migration Checklist

- [ ] Backup database (full dump)
- [ ] Review migration scripts with DBA
- [ ] Test migration on staging environment
- [ ] Verify application compatibility
- [ ] Schedule maintenance window
- [ ] Notify stakeholders
- [ ] Prepare rollback plan
- [ ] Set up monitoring/alerts

### Migration Execution Steps

1. **Preparation (5 minutes)**
   - Take database backup
   - Verify backup integrity
   - Stop application writes (read-only mode)

2. **Migration A - TP success_criteria (1 minute)**
   - Add column
   - Create index
   - Verify success

3. **Migration B - Assessment changes (5-10 minutes)**
   - Add new columns
   - Run data migration
   - Validate data integrity
   - Make tp_id NOT NULL
   - Drop modul_ajar_id
   - Update indexes
   - Verify success

4. **Migration C - Evaluation updates (2 minutes)**
   - Add new columns
   - Backfill timestamps
   - Verify success

5. **Migration D - Indexes (3-5 minutes)**
   - Create all indexes
   - Verify index creation
   - Test query performance

6. **Post-Migration (5 minutes)**
   - Run validation queries
   - Verify application connectivity
   - Enable application writes
   - Monitor for errors

### Post-Migration Validation

- [ ] All migrations applied successfully
- [ ] No orphaned records (assessments without tp_id)
- [ ] All indexes created and valid
- [ ] Application connects successfully
- [ ] Sample queries return expected results
- [ ] No performance degradation
- [ ] Error logs clean

---

## Rollback Plan

### Rollback Triggers

- Migration failure at any step
- Data validation failures
- Application errors post-migration
- Performance degradation
- Stakeholder request

### Rollback Steps

1. **Stop Application**
   - Put application in maintenance mode
   - Stop all writes

2. **Execute Down Migration**
   - Run `000003_add_success_criteria_and_refactor_assessment.down.sql`
   - Verify each step completes
   - Check for errors

3. **Restore from Backup** (if down migration fails)
   - Stop database
   - Restore from pre-migration backup
   - Start database
   - Verify data integrity

4. **Restart Application**
   - Deploy previous application version
   - Verify application health
   - Enable writes

5. **Post-Rollback Validation**
   - Verify data integrity
   - Run application tests
   - Monitor error logs
   - Document incident

### Rollback Validation

- [ ] Down migration completes successfully
- [ ] Original schema restored
- [ ] Data integrity verified
- [ ] Application connects successfully
- [ ] No data loss
- [ ] Performance baseline restored

---

## Migration Checklist

### Before Migration

- [ ] Database backup completed
- [ ] Staging environment tested
- [ ] Migration scripts reviewed
- [ ] Rollback plan documented
- [ ] Stakeholders notified
- [ ] Maintenance window scheduled
- [ ] Monitoring configured
- [ ] Application team ready

### During Migration

- [ ] Backup verified
- [ ] Application in read-only mode
- [ ] Migration A executed successfully
- [ ] Migration B executed successfully
- [ ] Migration C executed successfully
- [ ] Migration D executed successfully
- [ ] Validation queries passed
- [ ] No errors in logs

### After Migration

- [ ] Application writes enabled
- [ ] Application health checks pass
- [ ] Sample queries tested
- [ ] Performance metrics normal
- [ ] Error logs clean
- [ ] Stakeholders notified of completion
- [ ] Documentation updated
- [ ] Incident report created (if issues)

---

## Estimated Timelines

| Phase | Duration | Buffer | Total |
|-------|----------|--------|-------|
| Preparation | 30 min | 15 min | 45 min |
| Migration A | 1 min | 5 min | 6 min |
| Migration B | 10 min | 10 min | 20 min |
| Migration C | 2 min | 5 min | 7 min |
| Migration D | 5 min | 10 min | 15 min |
| Validation | 5 min | 10 min | 15 min |
| **Total** | **53 min** | **55 min** | **108 min** |

---

## Contacts

- **Database Administrator:** [TBD]
- **Application Lead:** [TBD]
- **DevOps Engineer:** [TBD]
- **Project Manager:** [TBD]

---

## Appendix: Validation Queries

```sql
-- Verify all assessments have tp_id
SELECT COUNT(*) FROM assessments WHERE tp_id IS NULL;

-- Verify tp_id references exist
SELECT COUNT(*) FROM assessments a 
LEFT JOIN tp ON a.tp_id = tp.id 
WHERE tp.id IS NULL;

-- Verify evaluation timestamps
SELECT COUNT(*) FROM evaluations WHERE created_at IS NULL OR updated_at IS NULL;

-- Check index creation
SELECT indexname FROM pg_indexes WHERE tablename IN ('tp', 'assessments', 'evaluations')
AND indexname LIKE 'idx_%';

-- Verify data migration counts
SELECT 
    (SELECT COUNT(*) FROM assessments) as total_assessments,
    (SELECT COUNT(*) FROM assessments WHERE tp_id IS NOT NULL) as with_tp_id,
    (SELECT COUNT(*) FROM assessments WHERE modul_ajar_id IS NULL) as without_modul_ajar;
```
