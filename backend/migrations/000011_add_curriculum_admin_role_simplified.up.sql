-- Sprint 4: Add CURRICULUM_ADMIN Role and Permissions (Simplified)
-- Purpose: Add CURRICULUM_ADMIN role and permissions for curriculum governance
-- Migration Number: 000011 (Simplified)
-- Risk Level: LOW

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
-- Add new Sprint 4 permissions
-- ============================================================================

-- Academic Year permissions
INSERT INTO permissions (id, role_id, resource, action, created_at)
SELECT 
    gen_uuid_v7(),
    (SELECT id FROM roles WHERE name = 'SCHOOL_ADMIN'),
    'academic_year',
    'CREATE',
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'CREATE'
);

INSERT INTO permissions (id, role_id, resource, action, created_at)
SELECT 
    gen_uuid_v7(),
    (SELECT id FROM roles WHERE name = 'SCHOOL_ADMIN'),
    'academic_year',
    'ACTIVATE',
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'ACTIVATE'
);

INSERT INTO permissions (id, role_id, resource, action, created_at)
SELECT 
    gen_uuid_v7(),
    (SELECT id FROM roles WHERE name = 'SCHOOL_ADMIN'),
    'academic_year',
    'ARCHIVE',
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'academic_year' AND action = 'ARCHIVE'
);

INSERT INTO permissions (id, role_id, resource, action, created_at)
SELECT 
    gen_uuid_v7(),
    (SELECT id FROM roles WHERE name = 'SYSTEM_ADMIN'),
    'system_config',
    'READ',
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'system_config' AND action = 'READ'
);

INSERT INTO permissions (id, role_id, resource, action, created_at)
SELECT 
    gen_uuid_v7(),
    (SELECT id FROM roles WHERE name = 'SYSTEM_ADMIN'),
    'system_config',
    'UPDATE',
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM permissions WHERE resource = 'system_config' AND action = 'UPDATE'
);

-- ============================================================================
-- Assign CURRICULUM_ADMIN permissions
-- ============================================================================

-- Get CURRICULUM_ADMIN role ID and assign permissions
DO $$
DECLARE
    curriculum_admin_role_id UUID;
BEGIN
    SELECT id INTO curriculum_admin_role_id FROM roles WHERE name = 'CURRICULUM_ADMIN';
    
    IF curriculum_admin_role_id IS NOT NULL THEN
    THEN
        -- Subject Category permissions
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'subject_category', 'CREATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'subject_category' AND action = 'CREATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'subject_category', 'UPDATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'subject_category' AND action = 'UPDATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'subject_category', 'DELETE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'subject_category' AND action = 'DELETE');
        
        -- Graduate Profile permissions
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'graduate_profile_dimension', 'CREATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'graduate_profile_dimension' AND action = 'CREATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'graduate_profile_dimension', 'UPDATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'graduate_profile_dimension' AND action = 'UPDATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'graduate_profile_dimension', 'DELETE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'graduate_profile_dimension' AND action = 'DELETE');
        
        -- CP Alignment permissions
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'cp_alignment', 'CREATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'cp_alignment' AND action = 'CREATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'cp_alignment', 'UPDATE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'cp_alignment' AND action = 'UPDATE');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'cp_alignment', 'DELETE', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'cp_alignment' AND action = 'DELETE');
        
        -- System config read permission
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'system_config', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'system_config' AND action = 'READ');
        
        -- Academic year read permission
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'academic_year', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'academic_year' AND action = 'READ');
        
        -- Semester read permission
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), curriculum_admin_role_id, 'semester', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = curriculum_admin_role_id AND resource = 'semester' AND action = 'READ');
    END IF;
END $$;

-- ============================================================================
-- Add Teacher read permissions for new Sprint 4 resources
-- ============================================================================

DO $$
DECLARE
    teacher_role_id UUID;
BEGIN
    SELECT id INTO teacher_role_id FROM roles WHERE name = 'TEACHER';
    
    IF teacher_role_id IS NOT NULL THEN
    THEN
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), teacher_role_id, 'academic_year', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = teacher_role_id AND resource = 'academic_year' AND action = 'READ');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), teacher_role_id, 'subject_category', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = teacher_role_id AND resource = 'subject_category' AND action = 'READ');
        
        INSERT INTO permissions (id, role_id, resource, action, created_at)
        SELECT gen_uuid_v7(), teacher_role_id, 'graduate_profile_dimension', 'READ', NOW()
        WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = teacher_role_id AND resource = 'graduate_profile_dimension' AND action = 'READ');
    END IF;
END $$;

-- ============================================================================
-- END OF MIGRATION
-- ============================================================================
