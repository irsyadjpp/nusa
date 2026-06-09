-- Rollback migration 000001_initial_schema
-- Purpose: Drop all tables created in initial schema migration
-- Based on: Database Schema Freeze v1
-- Rollback Strategy: DROP all tables in reverse order of creation
-- Risk Level: HIGH (complete schema removal)

-- ============================================================================
-- REPORTING CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS narrative_report_versions CASCADE;
DROP TABLE IF EXISTS narrative_reports CASCADE;

-- ============================================================================
-- ACHIEVEMENT CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS achievement_snapshots CASCADE;
DROP TABLE IF EXISTS achievement_criteria CASCADE;
DROP TABLE IF EXISTS achievements CASCADE;

-- ============================================================================
-- EVIDENCE CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS evidences CASCADE;

-- ============================================================================
-- ASSESSMENT CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS evaluation_feedback_history CASCADE;
DROP TABLE IF EXISTS evaluations CASCADE;
DROP TABLE IF EXISTS scoring_guidelines CASCADE;
DROP TABLE IF EXISTS answer_keys CASCADE;
DROP TABLE IF EXISTS assessment_items CASCADE;
DROP TABLE IF EXISTS assessment_versions CASCADE;
DROP TABLE IF EXISTS assessments CASCADE;

-- ============================================================================
-- LEARNING PLANNING CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS modul_ajar_items CASCADE;
DROP TABLE IF EXISTS modul_ajar_set_versions CASCADE;
DROP TABLE IF EXISTS modul_ajar_sets CASCADE;
DROP TABLE IF EXISTS atp_items CASCADE;
DROP TABLE IF EXISTS atp_set_versions CASCADE;
DROP TABLE IF EXISTS atp_sets CASCADE;
DROP TABLE IF EXISTS tp_set_versions CASCADE;
DROP TABLE IF EXISTS tp_sets CASCADE;
DROP TABLE IF EXISTS cp CASCADE;

-- ============================================================================
-- IDENTITY & ACCESS CONTEXT (reverse order)
-- ============================================================================

DROP TABLE IF EXISTS permission_changes CASCADE;
DROP TABLE IF EXISTS user_roles CASCADE;
DROP TABLE IF EXISTS role_permissions CASCADE;
DROP TABLE IF EXISTS permissions CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS students CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS schools CASCADE;

-- ============================================================================
-- CLEANUP
-- ============================================================================

-- Drop UUID v7 generation function
DROP FUNCTION IF EXISTS gen_uuid_v7();

-- Note: We keep the uuid-ossp extension as it may be used by other parts of the system
