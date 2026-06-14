-- Sprint 4: Rollback CURRICULUM_ADMIN Role Addition
-- Purpose: Remove CURRICULUM_ADMIN role (rollback for 000011)
-- Migration Number: 000011
-- Risk Level: LOW

-- ============================================================================
-- ROLLBACK INSTRUCTIONS
-- Note: This rollback will remove the CURRICULUM_ADMIN role
-- Any users with CURRICULUM_ADMIN role should be reassigned before rollback
-- ============================================================================

-- Drop existing CHECK constraint
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_check;

-- Recreate CHECK constraint without CURRICULUM_ADMIN
ALTER TABLE users 
ADD CONSTRAINT users_role_check 
CHECK (role IN ('TEACHER', 'SCHOOL_ADMIN', 'SYSTEM_ADMIN'));

-- Remove CURRICULUM_ADMIN role from roles table
DELETE FROM roles WHERE name = 'CURRICULUM_ADMIN';

-- ============================================================================
-- END OF ROLLBACK MIGRATION
-- ============================================================================
