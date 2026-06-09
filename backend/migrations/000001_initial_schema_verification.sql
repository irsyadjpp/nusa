-- Verification queries for migration 000001_initial_schema
-- Purpose: Validate that the migration was applied correctly
-- Based on: Database Schema Freeze v1

-- ============================================================================
-- TABLE EXISTENCE VERIFICATION
-- ============================================================================

-- Verify all 25 tables exist
SELECT 
    table_name,
    table_type
FROM information_schema.tables
WHERE table_schema = 'public'
AND table_name IN (
    'schools',
    'users',
    'roles',
    'permissions',
    'role_permissions',
    'user_roles',
    'students',
    'permission_changes',
    'cp',
    'tp_sets',
    'tp_set_versions',
    'atp_sets',
    'atp_set_versions',
    'atp_items',
    'modul_ajar_sets',
    'modul_ajar_set_versions',
    'modul_ajar_items',
    'assessments',
    'assessment_versions',
    'assessment_items',
    'answer_keys',
    'scoring_guidelines',
    'evaluations',
    'evaluation_feedback_history',
    'evidences',
    'achievements',
    'achievement_criteria',
    'achievement_snapshots',
    'narrative_reports',
    'narrative_report_versions'
)
ORDER BY table_name;

-- ============================================================================
-- COLUMN VERIFICATION
-- ============================================================================

-- Verify schools table columns
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = 'schools'
ORDER BY ordinal_position;

-- Verify users table columns
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = 'users'
ORDER BY ordinal_position;

-- Verify tp_sets table columns
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = 'tp_sets'
ORDER BY ordinal_position;

-- ============================================================================
-- CONSTRAINT VERIFICATION
-- ============================================================================

-- Verify primary key constraints
SELECT 
    tc.table_name,
    tc.constraint_name,
    tc.constraint_type
FROM information_schema.table_constraints tc
WHERE tc.table_schema = 'public'
AND tc.constraint_type = 'PRIMARY KEY'
ORDER BY tc.table_name;

-- Verify foreign key constraints
SELECT 
    tc.table_name,
    tc.constraint_name,
    kcu.column_name,
    ccu.table_name AS foreign_table_name,
    ccu.column_name AS foreign_column_name,
    rc.delete_rule
FROM information_schema.table_constraints tc
JOIN information_schema.key_column_usage kcu
    ON tc.constraint_name = kcu.constraint_name
JOIN information_schema.constraint_column_usage ccu
    ON ccu.constraint_name = tc.constraint_name
JOIN information_schema.referential_constraints rc
    ON tc.constraint_name = rc.constraint_name
WHERE tc.table_schema = 'public'
AND tc.constraint_type = 'FOREIGN KEY'
ORDER BY tc.table_name, tc.constraint_name;

-- Verify unique constraints
SELECT 
    tc.table_name,
    tc.constraint_name,
    kcu.column_name
FROM information_schema.table_constraints tc
JOIN information_schema.key_column_usage kcu
    ON tc.constraint_name = kcu.constraint_name
WHERE tc.table_schema = 'public'
AND tc.constraint_type = 'UNIQUE'
ORDER BY tc.table_name, tc.constraint_name;

-- Verify check constraints
SELECT 
    tc.table_name,
    tc.constraint_name,
    cc.check_clause
FROM information_schema.table_constraints tc
JOIN information_schema.check_constraints cc
    ON tc.constraint_name = cc.constraint_name
WHERE tc.table_schema = 'public'
AND tc.constraint_type = 'CHECK'
ORDER BY tc.table_name, tc.constraint_name;

-- ============================================================================
-- INDEX VERIFICATION
-- ============================================================================

-- Verify all indexes
SELECT 
    schemaname,
    tablename,
    indexname,
    indexdef
FROM pg_indexes
WHERE schemaname = 'public'
ORDER BY tablename, indexname;

-- ============================================================================
-- DATA TYPE VERIFICATION
-- ============================================================================

-- Verify UUID primary keys
SELECT 
    table_name,
    column_name,
    data_type,
    column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND column_name = 'id'
AND data_type = 'uuid'
AND column_default LIKE '%gen_uuid_v7%'
ORDER BY table_name;

-- Verify timestamp columns
SELECT 
    table_name,
    column_name,
    data_type,
    is_nullable,
    column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND data_type IN ('timestamp with time zone', 'timestamp without time zone')
ORDER BY table_name, column_name;

-- ============================================================================
-- ENUM VERIFICATION
-- ============================================================================

-- Verify status check constraints
SELECT 
    tc.table_name,
    tc.constraint_name,
    cc.check_clause
FROM information_schema.table_constraints tc
JOIN information_schema.check_constraints cc
    ON tc.constraint_name = cc.constraint_name
WHERE tc.table_schema = 'public'
AND tc.constraint_type = 'CHECK'
AND tc.constraint_name LIKE '%status%'
ORDER BY tc.table_name;

-- ============================================================================
-- EXTENSION VERIFICATION
-- ============================================================================

-- Verify uuid-ossp extension is installed
SELECT 
    extname,
    extversion
FROM pg_extension
WHERE extname = 'uuid-ossp';

-- Verify gen_uuid_v7 function exists
SELECT 
    proname,
    prosrc
FROM pg_proc
WHERE proname = 'gen_uuid_v7';

-- ============================================================================
-- RELATIONSHIP VERIFICATION
-- ============================================================================

-- Verify school isolation (users.school_id → schools.id)
SELECT 
    COUNT(*) AS user_school_relationships
FROM users u
JOIN schools s ON u.school_id = s.id;

-- Verify learning planning relationships
SELECT 
    COUNT(*) AS tp_set_cp_relationships
FROM tp_sets t
JOIN cp c ON t.cp_id = c.id;

SELECT 
    COUNT(*) AS tp_set_user_relationships
FROM tp_sets t
JOIN users u ON t.generated_by = u.id;

-- Verify assessment relationships
SELECT 
    COUNT(*) AS assessment_cp_relationships
FROM assessments a
JOIN cp c ON a.tp_id = c.id;

SELECT 
    COUNT(*) AS assessment_user_relationships
FROM assessments a
JOIN users u ON a.user_id = u.id;

-- Verify evaluation relationships
SELECT 
    COUNT(*) AS evaluation_assessment_relationships
FROM evaluations e
JOIN assessments a ON e.assessment_id = a.id;

SELECT 
    COUNT(*) AS evaluation_student_relationships
FROM evaluations e
JOIN students s ON e.student_id = s.id;

SELECT 
    COUNT(*) AS evaluation_teacher_relationships
FROM evaluations e
JOIN users u ON e.teacher_id = u.id;

-- ============================================================================
-- VERSION TABLE VERIFICATION
-- ============================================================================

-- Verify version tables have is_current_version column
SELECT 
    table_name,
    column_name,
    data_type
FROM information_schema.columns
WHERE table_schema = 'public'
AND column_name = 'is_current_version'
AND table_name IN (
    'tp_set_versions',
    'atp_set_versions',
    'modul_ajar_set_versions',
    'assessment_versions',
    'narrative_report_versions'
)
ORDER BY table_name;

-- Verify version tables have version_no column
SELECT 
    table_name,
    column_name,
    data_type
FROM information_schema.columns
WHERE table_schema = 'public'
AND column_name = 'version_no'
AND table_name IN (
    'tp_set_versions',
    'atp_set_versions',
    'modul_ajar_set_versions',
    'assessment_versions',
    'narrative_report_versions'
)
ORDER BY table_name;

-- ============================================================================
-- SOFT DELETE VERIFICATION
-- ============================================================================

-- Verify evidences table has soft delete columns
SELECT 
    column_name,
    data_type,
    is_nullable,
    column_default
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = 'evidences'
AND column_name IN ('is_deleted', 'deleted_at')
ORDER BY column_name;

-- ============================================================================
-- AUDIT FIELD VERIFICATION
-- ============================================================================

-- Verify all tables have created_at and updated_at
SELECT 
    table_name,
    COUNT(*) AS audit_field_count
FROM information_schema.columns
WHERE table_schema = 'public'
AND column_name IN ('created_at', 'updated_at')
GROUP BY table_name
HAVING COUNT(*) = 2
ORDER BY table_name;

-- ============================================================================
-- SUMMARY VERIFICATION
-- ============================================================================

-- Count total tables
SELECT COUNT(*) AS total_tables
FROM information_schema.tables
WHERE table_schema = 'public'
AND table_type = 'BASE TABLE';

-- Count total indexes
SELECT COUNT(*) AS total_indexes
FROM pg_indexes
WHERE schemaname = 'public';

-- Count total constraints
SELECT COUNT(*) AS total_constraints
FROM information_schema.table_constraints
WHERE table_schema = 'public';
