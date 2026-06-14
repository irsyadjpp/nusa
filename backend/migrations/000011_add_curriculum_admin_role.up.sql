-- Sprint 4: Add CURRICULUM_ADMIN Role
-- Purpose: Add CURRICULUM_ADMIN role to support curriculum governance features
-- Migration Number: 000011
-- Risk Level: LOW
-- Notes: This is a follow-up migration to 000010 to add the required CURRICULUM_ADMIN role

-- ============================================================================
-- Update users.role CHECK constraint to include CURRICULUM_ADMIN
-- ============================================================================

-- Drop existing CHECK constraint
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_check;

-- Recreate CHECK constraint with CURRICULUM_ADMIN included
ALTER TABLE users 
ADD CONSTRAINT users_role_check 
CHECK (role IN ('TEACHER', 'SCHOOL_ADMIN', 'SYSTEM_ADMIN', 'CURRICULUM_ADMIN'));

COMMENT ON COLUMN users.role IS 'User role: TEACHER, SCHOOL_ADMIN, SYSTEM_ADMIN, or CURRICULUM_ADMIN';

-- ============================================================================
-- Add CURRICULUM_ADMIN role to roles table
-- ============================================================================

INSERT INTO roles (id, name, description, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'CURRICULUM_ADMIN',
    'Curriculum Administrator with system-wide curriculum governance authority',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM roles WHERE name = 'CURRICULUM_ADMIN'
);

-- ============================================================================
-- END OF MIGRATION
-- ============================================================================
