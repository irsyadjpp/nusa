-- Sprint 4 Academic Foundation Migration (Manual Execution)
-- This migration will be executed step by step to identify any issues

-- ============================================================================
-- Table: system_configurations (NEW)
-- ============================================================================
CREATE TABLE system_configurations (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    key VARCHAR(100) NOT NULL UNIQUE,
    value TEXT NOT NULL,
    value_type VARCHAR(20) NOT NULL CHECK (value_type IN ('string', 'number', 'boolean', 'json')),
    description TEXT,
    category VARCHAR(50) NOT NULL,
    is_system BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_by UUID,
    CONSTRAINT fk_system_configurations_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_system_configurations_key ON system_configurations(key);
CREATE INDEX idx_system_configurations_category ON system_configurations(category);
CREATE INDEX idx_system_configurations_is_active ON system_configurations(is_active);

COMMENT ON TABLE system_configurations IS 'System-level configuration values for configurable parameters like CP alignment threshold';

-- Seed default CP alignment threshold
INSERT INTO system_configurations (key, value, value_type, description, category, is_system, is_active, created_by, created_at, updated_at)
SELECT 
    'cp_alignment_threshold',
    '60.0',
    'number',
    'CP alignment threshold percentage (default: 60%)',
    'CURRICULUM',
    true,
    true,
    (SELECT id FROM users WHERE role_id = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM system_configurations WHERE key = 'cp_alignment_threshold'
);

-- ============================================================================
-- Table: academic_years (NEW)
-- ============================================================================
CREATE TABLE academic_years (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    school_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'ACTIVE', 'ARCHIVED')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_academic_years_school_id FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE RESTRICT,
    CONSTRAINT fk_academic_years_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uq_academic_years_school_name UNIQUE (school_id, name),
    CONSTRAINT chk_academic_years_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_academic_years_school_id ON academic_years(school_id);
CREATE INDEX idx_academic_years_status ON academic_years(status);
CREATE INDEX idx_academic_years_dates ON academic_years(school_id, start_date, end_date);
CREATE INDEX idx_academic_years_school_status ON academic_years(school_id, status) WHERE status = 'ACTIVE';

COMMENT ON TABLE academic_years IS 'Academic year configuration with simplified workflow (no System Admin approval)';
COMMENT ON COLUMN academic_years.status IS 'Workflow status: DRAFT → ACTIVE → ARCHIVED (simplified from v1)';

-- ============================================================================
-- Table: semesters (NEW)
-- ============================================================================
CREATE TABLE semesters (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    academic_year_id UUID NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('GANJIL', 'GENAP')),
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    sequence_number INTEGER NOT NULL CHECK (sequence_number IN (1, 2)),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_semesters_academic_year_id FOREIGN KEY (academic_year_id) REFERENCES academic_years(id) ON DELETE CASCADE,
    CONSTRAINT uq_semesters_academic_year_sequence UNIQUE (academic_year_id, sequence_number),
    CONSTRAINT chk_semesters_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_semesters_academic_year_id ON semesters(academic_year_id);
CREATE INDEX idx_semesters_sequence ON semesters(academic_year_id, sequence_number);
CREATE INDEX idx_semesters_dates ON semesters(academic_year_id, start_date, end_date);
CREATE INDEX idx_semesters_status ON semesters(status);

COMMENT ON TABLE semesters IS 'Semester configuration within academic years (Ganjil/Genap)';

-- ============================================================================
-- Table: subject_categories (NEW)
-- ============================================================================
CREATE TABLE subject_categories (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    name_en VARCHAR(100),
    description TEXT,
    is_mandatory BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_by UUID,
    CONSTRAINT fk_subject_categories_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_subject_categories_code ON subject_categories(code);
CREATE INDEX idx_subject_categories_is_active ON subject_categories(is_active);

COMMENT ON TABLE subject_categories IS 'Subject categories following Kurikulum Merdeka standards';

-- Seed subject categories
INSERT INTO subject_categories (code, name, name_en, description, is_mandatory, is_active, created_by, created_at, updated_at)
SELECT 
    'INTRAKURIKULER',
    'Intrakurikuler',
    'Intracurricular',
    'Mata pelajaran utama kurikulum',
    true,
    true,
    (SELECT id FROM users WHERE role_id = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'INTRAKURIKULER'
);

INSERT INTO subject_categories (code, name, name_en, description, is_mandatory, is_active, created_by, created_at, updated_at)
SELECT 
    'KOKURIKULER',
    'Kokurikuler',
    'Cocurricular',
    'Kegiatan pengembangan diri',
    false,
    true,
    (SELECT id FROM users WHERE role_id = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'KOKURIKULER'
);

INSERT INTO subject_categories (code, name, name_en, description, is_mandatory, is_active, created_by, created_at, updated_at)
SELECT 
    'EKSTRAKURIKULER',
    'Ekstrakurikuler',
    'Extracurricular',
    'Kegiatan tambahan di luar kurikulum',
    false,
    true,
    (SELECT id FROM users WHERE role_id = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'EKSTRAKURIKULER'
);

-- ============================================================================
-- Table: graduate_profile_dimensions (NEW)
-- ============================================================================
CREATE TABLE graduate_profile_dimensions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    name_en VARCHAR(100),
    description TEXT,
    description_en TEXT,
    sequence_number INTEGER NOT NULL CHECK (sequence_number >= 1 AND sequence_number <= 6),
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_by UUID,
    CONSTRAINT fk_graduate_profile_dimensions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_graduate_profile_dimensions_code ON graduate_profile_dimensions(code);
CREATE INDEX idx_graduate_profile_dimensions_sequence_number ON graduate_profile_dimensions(sequence_number);
CREATE INDEX idx_graduate_profile_dimensions_is_active ON graduate_profile_dimensions(is_active);

COMMENT ON TABLE graduate_profile_dimensions IS 'Graduate profile dimensions (6 dimensions dari Profil Lulusan Kurikulum Merdeka)';

-- ============================================================================
-- Table: cp_alignments (NEW)
-- ============================================================================
CREATE TABLE cp_alignments (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    curriculum_subject_id UUID NOT NULL,
    graduate_profile_dimension_id UUID NOT NULL,
    alignment_description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_by UUID,
    CONSTRAINT fk_cp_alignments_curriculum_subject_id FOREIGN KEY (curriculum_subject_id) REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
    CONSTRAINT fk_cp_alignments_graduate_profile_dimension_id FOREIGN KEY (graduate_profile_dimension_id) REFERENCES graduate_profile_dimensions(id) ON DELETE CASCADE,
    CONSTRAINT uq_cp_alignments_combination UNIQUE (curriculum_subject_id, graduate_profile_dimension_id)
);

CREATE INDEX idx_cp_alignments_curriculum_subject_id ON cp_alignments(curriculum_subject_id);
CREATE INDEX idx_cp_alignments_graduate_profile_dimension_id ON cp_alignments(graduate_profile_dimension_id);
CREATE INDEX idx_cp_alignments_is_active ON cp_alignments(is_active);

COMMENT ON TABLE cp_alignments IS 'CP (Capaian Pembelajaran) alignment to graduate profile dimensions';
