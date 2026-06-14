-- Sprint 4: Add CURRICULUM_ADMIN Role (Corrected for actual schema)
-- Purpose: Add CURRICULUM_ADMIN role to support curriculum governance features
-- Migration Number: 000011 (Corrected)
-- Risk Level: LOW
-- Notes: Adapted for actual schema where users table has role_id instead of role

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
-- Add CURRICULUM_ADMIN permissions
-- ============================================================================

-- Academic Year permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'academic_year',
    'CREATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'CREATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'academic_year',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'UPDATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'academic_year',
    'ACTIVATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'ACTIVATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'academic_year',
    'ARCHIVE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'ARCHIVE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'academic_year',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'READ'
);

-- Semester permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'semester',
    'CREATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'semester' AND action = 'CREATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'semester',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'semester' AND action = 'UPDATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'semester',
    'DELETE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'semester' AND action = 'DELETE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'semester',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'semester' AND action = 'READ'
);

-- Subject Category permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'subject_category',
    'CREATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'subject_category' AND action = 'CREATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'subject_category',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'subject_category' AND action = 'UPDATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'subject_category',
    'DELETE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'subject_category' AND action = 'DELETE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'subject_category',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'subject_category' AND action = 'READ'
);

-- Graduate Profile permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'graduate_profile_dimension',
    'CREATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'graduate_profile_dimension' AND action = 'CREATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'graduate_profile_dimension',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'graduate_profile_dimension' AND action = 'UPDATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'graduate_profile_dimension',
    'DELETE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'graduate_profile_dimension' AND action = 'DELETE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'graduate_profile_dimension',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'graduate_profile_dimension' AND action = 'READ'
);

-- CP Alignment permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'cp_alignment',
    'CREATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'cp_alignment' AND action = 'CREATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'cp_alignment',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'cp_alignment' AND action = 'UPDATE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'cp_alignment',
    'DELETE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'cp_alignment' AND action = 'DELETE'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'cp_alignment',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'cp_alignment' AND action = 'READ'
);

-- System Config permissions
INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'system_config',
    'READ',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'system_config' AND action = 'READ'
);

INSERT INTO permissions (id, resource, action, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'system_config',
    'UPDATE',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'system_config' AND action = 'UPDATE'
);

-- ============================================================================
-- Assign permissions to CURRICULUM_ADMIN role
-- ============================================================================

-- Get the CURRICULUM_ADMIN role ID
DO $$
DECLARE
    curriculum_admin_role_id UUID;
BEGIN
    SELECT id INTO curriculum_admin_role_id FROM roles WHERE name = 'CURRICULUM_ADMIN';
    
    IF curriculum_admin_role_id IS NOT NULL THEN
    THEN
        -- Assign curriculum governance permissions
        INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, p.id, NOW(), NOW()
        FROM permissions p
        WHERE p.resource IN ('subject_category', 'graduate_profile_dimension', 'cp_alignment')
          AND p.action IN ('CREATE', 'UPDATE', 'DELETE', 'READ')
        ON CONFLICT DO NOTHING;
        
        -- Assign system config read/update permissions
        INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, p.id, NOW(), NOW()
        FROM permissions p
        WHERE p.resource = 'system_config'
          AND p.action IN ('READ', 'UPDATE')
        ON CONFLICT DO NOTHING;
    END IF;
END $$;

-- ============================================================================
-- Assign additional academic year read permissions to CURRICULUM_ADMIN
-- ============================================================================

DO $$
DECLARE
    curriculum_admin_role_id UUID;
BEGIN
    SELECT id INTO curriculum_admin_role_id FROM roles WHERE name = 'CURRICULUM_ADMIN';
    
    IF curriculum_admin_role_id IS NOT NULL THEN
    THEN
        -- Assign academic year read permission
        INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, p.id, NOW(), NOW()
        FROM permissions p
        WHERE p.resource = 'academic_year' AND p.action = 'READ'
        ON CONFLICT DO NOTHING;
        
        -- Assign semester read permission
        INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, p.id, NOW(), NOW()
        FROM permissions p
        WHERE p.resource = 'semester' AND p.action = 'READ'
        ON CONFLICT DO NOTHING;
    END IF;
END $$;

-- ============================================================================
-- END OF MIGRATION
-- ============================================================================
