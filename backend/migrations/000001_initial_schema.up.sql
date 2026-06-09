-- Migration 000001_initial_schema
-- Purpose: Create initial database schema for MVP
-- Based on: Database Schema Freeze v1
-- Affected Tables: All 25 tables
-- Risk Level: HIGH (initial schema creation)
-- Description: Creates all 25 tables with their initial schema definitions. This is the foundation migration that establishes the complete database structure for the NUSA Platform MVP.

-- Enable UUID extension for UUID v7 support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create UUID v7 generation function for PostgreSQL 18+ compatibility
CREATE OR REPLACE FUNCTION gen_uuid_v7() RETURNS UUID AS $$
BEGIN
    -- PostgreSQL 18+ has native uuid_generate_v7()
    -- Fall back to uuid_generate_v4() for compatibility
    BEGIN
        RETURN uuid_generate_v7();
    EXCEPTION
        WHEN undefined_function THEN
            RETURN uuid_generate_v4();
    END;
END;
$$ LANGUAGE plpgsql;

-- ============================================================================
-- IDENTITY & ACCESS CONTEXT
-- ============================================================================

-- Schools table
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL,
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_schools_code ON schools(code);
CREATE INDEX idx_schools_status ON schools(status);

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('TEACHER', 'SCHOOL_ADMIN', 'SYSTEM_ADMIN')),
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_users_school_id FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE RESTRICT,
    CONSTRAINT uk_users_username UNIQUE (username),
    CONSTRAINT uk_users_email UNIQUE (email)
);

CREATE INDEX idx_users_school_id ON users(school_id);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_status ON users(status);

-- Roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT uk_roles_name UNIQUE (name)
);

CREATE INDEX idx_roles_name ON roles(name);

-- Permissions table
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_permissions_resource_action ON permissions(resource, action);

-- Role Permissions junction table
CREATE TABLE role_permissions (
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    granted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    granted_by UUID,
    PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permissions_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_permissions_permission_id FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_permissions_granted_by FOREIGN KEY (granted_by) REFERENCES users(id)
);

CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- User Roles junction table
CREATE TABLE user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    assigned_by UUID,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_roles_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_assigned_by FOREIGN KEY (assigned_by) REFERENCES users(id)
);

CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);

-- Students table
CREATE TABLE students (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    student_number VARCHAR(50) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    grade_level VARCHAR(50),
    class_name VARCHAR(50),
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE', 'GRADUATED')),
    enrolled_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    graduated_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_students_school_id FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE RESTRICT,
    CONSTRAINT uk_students_school_student UNIQUE (school_id, student_number)
);

CREATE INDEX idx_students_school_id ON students(school_id);
CREATE INDEX idx_students_student_number ON students(student_number);
CREATE INDEX idx_students_status ON students(status);

-- Permission Changes audit table
CREATE TABLE permission_changes (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    role_id UUID NOT NULL,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    change_type VARCHAR(20) NOT NULL CHECK (change_type IN ('GRANT', 'REVOKE')),
    changed_by UUID NOT NULL,
    changed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    reason TEXT,
    CONSTRAINT fk_permission_changes_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE RESTRICT,
    CONSTRAINT fk_permission_changes_changed_by FOREIGN KEY (changed_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_permission_changes_role_id ON permission_changes(role_id);
CREATE INDEX idx_permission_changes_changed_at ON permission_changes(changed_at);

-- ============================================================================
-- LEARNING PLANNING CONTEXT
-- ============================================================================

-- CP (Capaian Pembelajaran) table
CREATE TABLE cp (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL,
    text TEXT NOT NULL,
    subject VARCHAR(100),
    grade_level VARCHAR(50),
    phase VARCHAR(50),
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT uk_cp_code UNIQUE (code)
);

CREATE INDEX idx_cp_code ON cp(code);
CREATE INDEX idx_cp_subject ON cp(subject);
CREATE INDEX idx_cp_grade_level ON cp(grade_level);
CREATE INDEX idx_cp_status ON cp(status);

-- TP Sets table
CREATE TABLE tp_sets (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    cp_id UUID NOT NULL,
    generated_by UUID NOT NULL,
    learning_objectives TEXT,
    time_allocation INTEGER,
    success_criteria TEXT,
    generation_reason TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    approved_by_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tp_sets_cp_id FOREIGN KEY (cp_id) REFERENCES cp(id) ON DELETE RESTRICT,
    CONSTRAINT fk_tp_sets_generated_by FOREIGN KEY (generated_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_tp_sets_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_tp_sets_cp_id ON tp_sets(cp_id);
CREATE INDEX idx_tp_sets_generated_by ON tp_sets(generated_by);
CREATE INDEX idx_tp_sets_status ON tp_sets(status);

-- TP Set Versions table
CREATE TABLE tp_set_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    tp_set_id UUID NOT NULL,
    version_no INTEGER NOT NULL,
    is_current_version BOOLEAN NOT NULL DEFAULT false,
    parent_revision_id UUID,
    cp_id UUID NOT NULL,
    learning_objectives TEXT,
    time_allocation INTEGER,
    success_criteria TEXT,
    version_reason TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tp_set_versions_tp_set_id FOREIGN KEY (tp_set_id) REFERENCES tp_sets(id) ON DELETE CASCADE,
    CONSTRAINT fk_tp_set_versions_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES tp_set_versions(id),
    CONSTRAINT fk_tp_set_versions_cp_id FOREIGN KEY (cp_id) REFERENCES cp(id) ON DELETE RESTRICT,
    CONSTRAINT fk_tp_set_versions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_tp_set_versions_tp_version UNIQUE (tp_set_id, version_no)
);

CREATE INDEX idx_tp_set_versions_tp_set_id ON tp_set_versions(tp_set_id);
CREATE INDEX idx_tp_set_versions_version_no ON tp_set_versions(version_no);
CREATE INDEX idx_tp_set_versions_is_current ON tp_set_versions(is_current_version);

-- ATP Sets table
CREATE TABLE atp_sets (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    tp_set_id UUID NOT NULL,
    tp_set_version_no INTEGER NOT NULL,
    generated_by UUID NOT NULL,
    generation_reason TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_atp_sets_tp_set_id FOREIGN KEY (tp_set_id) REFERENCES tp_sets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_atp_sets_generated_by FOREIGN KEY (generated_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_atp_sets_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_atp_sets_tp_set_id ON atp_sets(tp_set_id);
CREATE INDEX idx_atp_sets_generated_by ON atp_sets(generated_by);
CREATE INDEX idx_atp_sets_status ON atp_sets(status);

-- ATP Set Versions table
CREATE TABLE atp_set_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    atp_set_id UUID NOT NULL,
    version_no INTEGER NOT NULL,
    is_current_version BOOLEAN NOT NULL DEFAULT false,
    parent_revision_id UUID,
    tp_set_id UUID NOT NULL,
    tp_set_version_no INTEGER NOT NULL,
    generation_reason TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_atp_set_versions_atp_set_id FOREIGN KEY (atp_set_id) REFERENCES atp_sets(id) ON DELETE CASCADE,
    CONSTRAINT fk_atp_set_versions_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES atp_set_versions(id),
    CONSTRAINT fk_atp_set_versions_tp_set_id FOREIGN KEY (tp_set_id) REFERENCES tp_sets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_atp_set_versions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_atp_set_versions_atp_version UNIQUE (atp_set_id, version_no)
);

CREATE INDEX idx_atp_set_versions_atp_set_id ON atp_set_versions(atp_set_id);
CREATE INDEX idx_atp_set_versions_version_no ON atp_set_versions(version_no);
CREATE INDEX idx_atp_set_versions_is_current ON atp_set_versions(is_current_version);

-- ATP Items table
CREATE TABLE atp_items (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    atp_set_version_id UUID NOT NULL,
    week_number INTEGER NOT NULL,
    cp_code VARCHAR(50) NOT NULL,
    cp_text TEXT,
    sequence_number INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_atp_items_atp_set_version_id FOREIGN KEY (atp_set_version_id) REFERENCES atp_set_versions(id) ON DELETE CASCADE
);

CREATE INDEX idx_atp_items_atp_set_version_id ON atp_items(atp_set_version_id);
CREATE INDEX idx_atp_items_week_number ON atp_items(week_number);

-- Modul Ajar Sets table
CREATE TABLE modul_ajar_sets (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    atp_set_id UUID NOT NULL,
    atp_set_version_no INTEGER NOT NULL,
    generated_by UUID NOT NULL,
    generation_reason TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_modul_ajar_sets_atp_set_id FOREIGN KEY (atp_set_id) REFERENCES atp_sets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_modul_ajar_sets_generated_by FOREIGN KEY (generated_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_modul_ajar_sets_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_modul_ajar_sets_atp_set_id ON modul_ajar_sets(atp_set_id);
CREATE INDEX idx_modul_ajar_sets_generated_by ON modul_ajar_sets(generated_by);
CREATE INDEX idx_modul_ajar_sets_status ON modul_ajar_sets(status);

-- Modul Ajar Set Versions table
CREATE TABLE modul_ajar_set_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    modul_ajar_set_id UUID NOT NULL,
    version_no INTEGER NOT NULL,
    is_current_version BOOLEAN NOT NULL DEFAULT false,
    parent_revision_id UUID,
    atp_set_id UUID NOT NULL,
    atp_set_version_no INTEGER NOT NULL,
    generation_reason TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_modul_ajar_set_versions_modul_ajar_set_id FOREIGN KEY (modul_ajar_set_id) REFERENCES modul_ajar_sets(id) ON DELETE CASCADE,
    CONSTRAINT fk_modul_ajar_set_versions_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES modul_ajar_set_versions(id),
    CONSTRAINT fk_modul_ajar_set_versions_atp_set_id FOREIGN KEY (atp_set_id) REFERENCES atp_sets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_modul_ajar_set_versions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_modul_ajar_set_versions_modul_version UNIQUE (modul_ajar_set_id, version_no)
);

CREATE INDEX idx_modul_ajar_set_versions_modul_ajar_set_id ON modul_ajar_set_versions(modul_ajar_set_id);
CREATE INDEX idx_modul_ajar_set_versions_version_no ON modul_ajar_set_versions(version_no);
CREATE INDEX idx_modul_ajar_set_versions_is_current ON modul_ajar_set_versions(is_current_version);

-- Modul Ajar Items table
CREATE TABLE modul_ajar_items (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    modul_ajar_set_version_id UUID NOT NULL,
    sequence_number INTEGER NOT NULL,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_modul_ajar_items_modul_ajar_set_version_id FOREIGN KEY (modul_ajar_set_version_id) REFERENCES modul_ajar_set_versions(id) ON DELETE CASCADE
);

CREATE INDEX idx_modul_ajar_items_modul_ajar_set_version_id ON modul_ajar_items(modul_ajar_set_version_id);
CREATE INDEX idx_modul_ajar_items_sequence_number ON modul_ajar_items(sequence_number);

-- ============================================================================
-- ASSESSMENT CONTEXT
-- ============================================================================

-- Assessments table
CREATE TABLE assessments (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    tp_id UUID NOT NULL,
    tp_version_no INTEGER NOT NULL,
    user_id UUID NOT NULL,
    success_criteria_snapshot TEXT,
    tp_learning_objectives_snapshot TEXT,
    tp_time_allocation_snapshot INTEGER,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_assessments_tp_id FOREIGN KEY (tp_id) REFERENCES cp(id) ON DELETE RESTRICT,
    CONSTRAINT fk_assessments_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_assessments_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_assessments_tp_id ON assessments(tp_id);
CREATE INDEX idx_assessments_user_id ON assessments(user_id);
CREATE INDEX idx_assessments_status ON assessments(status);

-- Assessment Versions table
CREATE TABLE assessment_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    assessment_id UUID NOT NULL,
    version_no INTEGER NOT NULL,
    is_current_version BOOLEAN NOT NULL DEFAULT false,
    parent_revision_id UUID,
    tp_id UUID NOT NULL,
    tp_version_no INTEGER NOT NULL,
    success_criteria_snapshot TEXT,
    tp_learning_objectives_snapshot TEXT,
    tp_time_allocation_snapshot INTEGER,
    version_reason TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_assessment_versions_assessment_id FOREIGN KEY (assessment_id) REFERENCES assessments(id) ON DELETE CASCADE,
    CONSTRAINT fk_assessment_versions_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES assessment_versions(id),
    CONSTRAINT fk_assessment_versions_tp_id FOREIGN KEY (tp_id) REFERENCES cp(id) ON DELETE RESTRICT,
    CONSTRAINT fk_assessment_versions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_assessment_versions_assessment_version UNIQUE (assessment_id, version_no)
);

CREATE INDEX idx_assessment_versions_assessment_id ON assessment_versions(assessment_id);
CREATE INDEX idx_assessment_versions_version_no ON assessment_versions(version_no);
CREATE INDEX idx_assessment_versions_is_current ON assessment_versions(is_current_version);

-- Assessment Items table
CREATE TABLE assessment_items (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    assessment_version_id UUID NOT NULL,
    sequence_number INTEGER NOT NULL,
    question_text TEXT,
    item_type VARCHAR(50),
    max_score INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_assessment_items_assessment_version_id FOREIGN KEY (assessment_version_id) REFERENCES assessment_versions(id) ON DELETE CASCADE
);

CREATE INDEX idx_assessment_items_assessment_version_id ON assessment_items(assessment_version_id);
CREATE INDEX idx_assessment_items_sequence_number ON assessment_items(sequence_number);

-- Answer Keys table
CREATE TABLE answer_keys (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    assessment_item_id UUID NOT NULL,
    correct_answer TEXT,
    scoring_rubric TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_answer_keys_assessment_item_id FOREIGN KEY (assessment_item_id) REFERENCES assessment_items(id) ON DELETE CASCADE
);

CREATE INDEX idx_answer_keys_assessment_item_id ON answer_keys(assessment_item_id);

-- Scoring Guidelines table
CREATE TABLE scoring_guidelines (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    assessment_version_id UUID NOT NULL,
    guideline_text TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_scoring_guidelines_assessment_version_id FOREIGN KEY (assessment_version_id) REFERENCES assessment_versions(id) ON DELETE CASCADE
);

CREATE INDEX idx_scoring_guidelines_assessment_version_id ON scoring_guidelines(assessment_version_id);

-- Evaluations table
CREATE TABLE evaluations (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    assessment_id UUID NOT NULL,
    student_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    revision_no INTEGER NOT NULL DEFAULT 1,
    is_current_version BOOLEAN NOT NULL DEFAULT true,
    parent_revision_id UUID,
    total_score INTEGER CHECK (total_score >= 0 AND total_score <= 100),
    performance_level VARCHAR(50),
    teacher_feedback TEXT,
    evaluation_criteria_scores JSONB,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_evaluations_assessment_id FOREIGN KEY (assessment_id) REFERENCES assessments(id) ON DELETE RESTRICT,
    CONSTRAINT fk_evaluations_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT,
    CONSTRAINT fk_evaluations_teacher_id FOREIGN KEY (teacher_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_evaluations_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES evaluations(id),
    CONSTRAINT fk_evaluations_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_evaluations_assessment_id ON evaluations(assessment_id);
CREATE INDEX idx_evaluations_student_id ON evaluations(student_id);
CREATE INDEX idx_evaluations_teacher_id ON evaluations(teacher_id);
CREATE INDEX idx_evaluations_revision_no ON evaluations(revision_no);
CREATE INDEX idx_evaluations_is_current ON evaluations(is_current_version);
CREATE INDEX idx_evaluations_status ON evaluations(status);

-- Evaluation Feedback History table
CREATE TABLE evaluation_feedback_history (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    evaluation_id UUID NOT NULL,
    revision_no INTEGER NOT NULL,
    teacher_feedback TEXT NOT NULL,
    changed_by UUID NOT NULL,
    changed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    change_reason TEXT,
    CONSTRAINT fk_evaluation_feedback_history_evaluation_id FOREIGN KEY (evaluation_id) REFERENCES evaluations(id) ON DELETE CASCADE,
    CONSTRAINT fk_evaluation_feedback_history_changed_by FOREIGN KEY (changed_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_evaluation_feedback_history_evaluation_id ON evaluation_feedback_history(evaluation_id);
CREATE INDEX idx_evaluation_feedback_history_revision_no ON evaluation_feedback_history(revision_no);
CREATE INDEX idx_evaluation_feedback_history_changed_at ON evaluation_feedback_history(changed_at);

-- ============================================================================
-- EVIDENCE CONTEXT
-- ============================================================================

-- Evidences table
CREATE TABLE evidences (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    student_id UUID NOT NULL,
    assessment_id UUID NOT NULL,
    evidence_type VARCHAR(50) NOT NULL CHECK (evidence_type IN ('DOCUMENT', 'IMAGE', 'VIDEO', 'AUDIO')),
    file_id UUID NOT NULL,
    storage_key VARCHAR(500) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    file_size_bytes BIGINT NOT NULL CHECK (file_size_bytes > 0 AND file_size_bytes <= 52428800),
    file_hash VARCHAR(64) NOT NULL,
    teacher_notes TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'UPLOADED' CHECK (status IN ('UPLOADED', 'PROCESSING', 'PROCESSED', 'ERROR')),
    uploaded_by UUID NOT NULL,
    uploaded_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_evidences_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT,
    CONSTRAINT fk_evidences_assessment_id FOREIGN KEY (assessment_id) REFERENCES assessments(id) ON DELETE RESTRICT,
    CONSTRAINT fk_evidences_uploaded_by FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_evidences_file_id UNIQUE (file_id)
);

CREATE INDEX idx_evidences_student_id ON evidences(student_id);
CREATE INDEX idx_evidences_assessment_id ON evidences(assessment_id);
CREATE INDEX idx_evidences_file_id ON evidences(file_id);
CREATE INDEX idx_evidences_status ON evidences(status);
CREATE INDEX idx_evidences_is_deleted ON evidences(is_deleted);

-- ============================================================================
-- ACHIEVEMENT CONTEXT
-- ============================================================================

-- Achievements table
CREATE TABLE achievements (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    student_id UUID NOT NULL,
    competency_level VARCHAR(50),
    progress_percentage INTEGER CHECK (progress_percentage >= 0 AND progress_percentage <= 100),
    achievement_date DATE NOT NULL,
    calculation_metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_achievements_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT
);

CREATE INDEX idx_achievements_student_id ON achievements(student_id);
CREATE INDEX idx_achievements_achievement_date ON achievements(achievement_date);

-- Achievement Criteria table
CREATE TABLE achievement_criteria (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    achievement_id UUID NOT NULL,
    criteria_name VARCHAR(255) NOT NULL,
    criteria_value TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_achievement_criteria_achievement_id FOREIGN KEY (achievement_id) REFERENCES achievements(id) ON DELETE CASCADE
);

CREATE INDEX idx_achievement_criteria_achievement_id ON achievement_criteria(achievement_id);

-- Achievement Snapshots table
CREATE TABLE achievement_snapshots (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    student_id UUID NOT NULL,
    snapshot_date DATE NOT NULL,
    competency_level VARCHAR(50),
    progress_percentage INTEGER CHECK (progress_percentage >= 0 AND progress_percentage <= 100),
    snapshot_data JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_achievement_snapshots_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT
);

CREATE INDEX idx_achievement_snapshots_student_id ON achievement_snapshots(student_id);
CREATE INDEX idx_achievement_snapshots_snapshot_date ON achievement_snapshots(snapshot_date);

-- ============================================================================
-- REPORTING CONTEXT
-- ============================================================================

-- Narrative Reports table
CREATE TABLE narrative_reports (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    student_id UUID NOT NULL,
    report_period VARCHAR(50) NOT NULL,
    narrative_content TEXT,
    achievement_data JSONB,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'UNDER_REVIEW', 'APPROVED')),
    generated_by UUID NOT NULL,
    approved_by UUID,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_narrative_reports_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT,
    CONSTRAINT fk_narrative_reports_generated_by FOREIGN KEY (generated_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_narrative_reports_approved_by FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE INDEX idx_narrative_reports_student_id ON narrative_reports(student_id);
CREATE INDEX idx_narrative_reports_report_period ON narrative_reports(report_period);
CREATE INDEX idx_narrative_reports_status ON narrative_reports(status);

-- Narrative Report Versions table
CREATE TABLE narrative_report_versions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    narrative_report_id UUID NOT NULL,
    version_no INTEGER NOT NULL,
    is_current_version BOOLEAN NOT NULL DEFAULT false,
    parent_revision_id UUID,
    student_id UUID NOT NULL,
    report_period VARCHAR(50) NOT NULL,
    narrative_content TEXT,
    achievement_data JSONB,
    version_reason TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_narrative_report_versions_narrative_report_id FOREIGN KEY (narrative_report_id) REFERENCES narrative_reports(id) ON DELETE CASCADE,
    CONSTRAINT fk_narrative_report_versions_parent_revision_id FOREIGN KEY (parent_revision_id) REFERENCES narrative_report_versions(id),
    CONSTRAINT fk_narrative_report_versions_student_id FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT,
    CONSTRAINT fk_narrative_report_versions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uk_narrative_report_versions_report_version UNIQUE (narrative_report_id, version_no)
);

CREATE INDEX idx_narrative_report_versions_narrative_report_id ON narrative_report_versions(narrative_report_id);
CREATE INDEX idx_narrative_report_versions_version_no ON narrative_report_versions(version_no);
CREATE INDEX idx_narrative_report_versions_is_current ON narrative_report_versions(is_current_version);
