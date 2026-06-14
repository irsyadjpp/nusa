-- Sprint 4 Academic Foundation Migration
-- Based on: Sprint 4 Requirement Package (Refactored v2)
-- Based on: sprint4-implementation-plan-refactored.md
-- Migration Number: 000010
-- Purpose: Add academic year, semester, subject category, graduate profile dimension infrastructure
-- Tables Created: 5 (academic_years, semesters, subject_categories, graduate_profile_dimensions, cp_alignments, system_configuration)
-- Tables Extended: 2 (curriculum_subjects, cp)
-- Risk Level: MEDIUM (new tables, existing table extensions)
-- Notes: Students table exists but is NOT modified (out of scope for Sprint 4)

-- ============================================================================
-- NEW TABLES
-- ============================================================================

-- Table: system_configuration (NEW)
-- Priority: Create first as it's infrastructure for other features
CREATE TABLE IF NOT EXISTS system_configuration (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    key VARCHAR(100) NOT NULL UNIQUE,
    value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_system_configuration_key ON system_configuration(key);

COMMENT ON TABLE system_configuration IS 'System-level configuration values for configurable parameters like CP alignment threshold';

-- Seed default CP alignment threshold
INSERT INTO system_configuration (id, key, value, description, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'cp_alignment_threshold',
    '60.0',
    'CP alignment threshold percentage (default: 60%)',
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM system_configuration WHERE key = 'cp_alignment_threshold'
);

-- ============================================================================
-- Table: academic_years (NEW)
-- Purpose: Store academic year configuration with simplified workflow (DRAFT → ACTIVE → ARCHIVED)
CREATE TABLE IF NOT EXISTS academic_years (
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
-- Purpose: Store semester configuration within academic years
CREATE TABLE IF NOT EXISTS semesters (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    academic_year_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    sequence INTEGER NOT NULL CHECK (sequence IN (1, 2)),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'DRAFT' CHECK (status IN ('DRAFT', 'ACTIVE', 'INACTIVE')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_semesters_academic_year_id FOREIGN KEY (academic_year_id) REFERENCES academic_years(id) ON DELETE CASCADE,
    CONSTRAINT uq_semesters_academic_year_sequence UNIQUE (academic_year_id, sequence),
    CONSTRAINT chk_semesters_dates CHECK (start_date < end_date)
);

CREATE INDEX idx_semesters_academic_year_id ON semesters(academic_year_id);
CREATE INDEX idx_semesters_sequence ON semesters(academic_year_id, sequence);
CREATE INDEX idx_semesters_dates ON semesters(academic_year_id, start_date, end_date);
CREATE INDEX idx_semesters_status ON semesters(status);

COMMENT ON TABLE semesters IS 'Semester configuration within academic years (Ganjil/Genap)';

-- ============================================================================
-- Table: subject_categories (NEW)
-- Purpose: Store subject category definitions (Intrakurikuler, Kokurikuler, Ekstrakurikuler)
CREATE TABLE IF NOT EXISTS subject_categories (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(100),
    description TEXT NOT NULL,
    guidelines TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_subject_categories_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_subject_categories_code ON subject_categories(code);
CREATE INDEX idx_subject_categories_name ON subject_categories(name);
CREATE INDEX idx_subject_categories_status ON subject_categories(status);

COMMENT ON TABLE subject_categories IS 'Subject categories following Kurikulum Merdeka standards';

-- Seed subject categories
INSERT INTO subject_categories (id, code, name, name_en, description, guidelines, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'INTRAKURIKULER',
    'Intrakurikuler',
    'Intracurricular',
    'Mata pelajaran utama kurikulum',
    'Mata pelajaran inti yang wajib diikuti semua siswa sesuai fase',
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'INTRAKURIKULER'
);

INSERT INTO subject_categories (id, code, name, name_en, description, guidelines, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'KOKURIKULER',
    'Kokurikuler',
    'Cocurricular',
    'Kegiatan pengembangan diri',
    'Kegiatan untuk mengembangkan potensi dan bakat siswa',
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'KOKURIKULER'
);

INSERT INTO subject_categories (id, code, name, name_en, description, guidelines, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'EKSTRAKURIKULER',
    'Ekstrakurikuler',
    'Extracurricular',
    'Kegiatan tambahan di luar kurikulum',
    'Kegiatan pilihan untuk pengembangan minat dan bakat',
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM subject_categories WHERE code = 'EKSTRAKURIKULER'
);

-- ============================================================================
-- Table: graduate_profile_dimensions (NEW)
-- Purpose: Store graduate profile dimension definitions (8 dimensions dari Profil Lulusan)
CREATE TABLE IF NOT EXISTS graduate_profile_dimensions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    name_en VARCHAR(100),
    description TEXT NOT NULL,
    indicators JSONB NOT NULL,
    weight DECIMAL(5,4) NOT NULL CHECK (weight > 0 AND weight <= 1.0),
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_graduate_profile_dimensions_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE INDEX idx_graduate_profile_dimensions_code ON graduate_profile_dimensions(code);
CREATE INDEX idx_graduate_profile_dimensions_name ON graduate_profile_dimensions(name);
CREATE INDEX idx_graduate_profile_dimensions_status ON graduate_profile_dimensions(status);
CREATE INDEX idx_graduate_profile_dimensions_indicators ON graduate_profile_dimensions USING GIN (indicators);

COMMENT ON TABLE graduate_profile_dimensions IS '8-dimensional Profil Lulusan framework following Permendikdasmen No. 10 Tahun 2025';

-- Seed graduate profile dimensions (8 dimensions)
INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'KEIMANAN_KETAKWAAN',
    'Keimanan & Ketakwaan',
    'Faith and Piety',
    'Dimensi keimanan dan ketakwaan kepada Tuhan Yang Maha Esa',
    '["Berakhlak mulia", "Menjaga kebersihan hati", "Melaksanakan ibadah"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'KEIMANAN_KETAKWAAN'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'KEWARGAAN',
    'Kewargaan',
    'Citizenship',
    'Dimensi kesadaran berbangsa dan bernegara',
    '["Cinta tanah air", "Menghargai keberagaman", "Taat aturan"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'KEWARGAAN'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'BERAKHLAK_MULIA',
    'Berakhlak Mulia',
    'Noble Character',
    'Dimensi pembentukan karakter mulia',
    '["Jujur", "Disiplin", "Tanggung jawab"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'BERAKHLAK_MULIA'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'BERANI_BERTANGGUNG_JAWAB',
    'Berani Bertanggung Jawab',
    'Courageous and Responsible',
    'Dimensi keberanian dan tanggung jawab',
    '["Berani mengambil keputusan", "Pertanggung jawaban atas tindakan"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'BERANI_BERTANGGUNG_JAWAB'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'PEDULI',
    'Peduli',
    'Caring',
    'Dimensi kepedulian terhadap sesama',
    '["Empati", "Saling membantu", "Toleransi"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'PEDULI'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'GOTONG_ROYONG',
    'Gotong Royong',
    'Collaboration',
    'Dimensi kerja sama dan gotong royong',
    '["Kerja tim", "Solidaritas", "Kolaborasi"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'GOTONG_ROYONG'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'MANDIRI',
    'Mandiri',
    'Independent',
    'Dimensi kemandirian dan otonomi',
    '["Berpikir kritis", "Mengambil inisiatif", "Mandiri belajar"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'MANDIRI'
);

INSERT INTO graduate_profile_dimensions (id, code, name, name_en, description, indicators, weight, status, created_by, created_at, updated_at)
SELECT 
    gen_uuid_v7(),
    'KREATIF',
    'Kreatif',
    'Creative',
    'Dimensi kreativitas dan inovasi',
    '["Berpikir kreatif", "Inovasi", "Pemecahan masalah"]'::jsonb,
    0.125,
    'ACTIVE',
    (SELECT id FROM users WHERE role = 'SYSTEM_ADMIN' LIMIT 1),
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM graduate_profile_dimensions WHERE code = 'KREATIF'
);

-- ============================================================================
-- Table: cp_alignments (NEW)
-- Purpose: Store CP to graduate profile dimension alignments with strength scoring
CREATE TABLE IF NOT EXISTS cp_alignments (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    cp_id UUID NOT NULL,
    dimension_id UUID NOT NULL,
    alignment_strength VARCHAR(20) NOT NULL CHECK (alignment_strength IN ('STRONG', 'MEDIUM', 'WEAK')),
    rationale TEXT,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_cp_alignments_cp_id FOREIGN KEY (cp_id) REFERENCES cp(id) ON DELETE CASCADE,
    CONSTRAINT fk_cp_alignments_dimension_id FOREIGN KEY (dimension_id) REFERENCES graduate_profile_dimensions(id) ON DELETE CASCADE,
    CONSTRAINT fk_cp_alignments_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT uq_cp_alignments_cp_dimension UNIQUE (cp_id, dimension_id)
);

CREATE INDEX idx_cp_alignments_cp_id ON cp_alignments(cp_id);
CREATE INDEX idx_cp_alignments_dimension_id ON cp_alignments(dimension_id);
CREATE INDEX idx_cp_alignments_strength ON cp_alignments(alignment_strength);

COMMENT ON TABLE cp_alignments IS 'CP alignment to graduate profile dimensions with configurable threshold';

-- ============================================================================
-- TABLE EXTENSIONS
-- ============================================================================

-- Extend: curriculum_subjects
-- Purpose: Add subject categorization to existing subjects table
-- Note: Students table exists but is NOT modified (out of scope for Sprint 4)
ALTER TABLE curriculum_subjects
ADD COLUMN IF NOT EXISTS subject_category_id UUID;

ALTER TABLE curriculum_subjects
ADD CONSTRAINT fk_curriculum_subjects_category_id 
FOREIGN KEY (subject_category_id) REFERENCES subject_categories(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_curriculum_subjects_category_id ON curriculum_subjects(subject_category_id);

COMMENT ON COLUMN curriculum_subjects.subject_category_id IS 'Subject category (Intrakurikuler/Kokurikuler/Ekstrakurikuler)';

-- Categorize existing subjects as Intrakurikuler by default
UPDATE curriculum_subjects
SET subject_category_id = (SELECT id FROM subject_categories WHERE code = 'INTRAKURIKULER' LIMIT 1)
WHERE subject_category_id IS NULL;

-- ============================================================================
-- Extend: cp
-- Purpose: Add academic year scoping to existing CP table
-- Note: This extends the curriculum module to support temporal scoping
-- Note: Students table exists but is NOT modified (out of scope for Sprint 4)

ALTER TABLE cp
ADD COLUMN IF NOT EXISTS academic_year_id UUID;

ALTER TABLE cp
ADD CONSTRAINT fk_cp_academic_year_id 
FOREIGN KEY (academic_year_id) REFERENCES academic_years(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_cp_academic_year_id ON cp(academic_year_id);

ALTER TABLE cp
ADD COLUMN IF NOT EXISTS semester_id UUID;

ALTER TABLE cp
ADD CONSTRAINT fk_cp_semester_id 
FOREIGN KEY (semester_id) REFERENCES semesters(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_cp_semester_id ON cp(semester_id);
CREATE INDEX IF NOT EXISTS idx_cp_academic_semester ON cp(academic_year_id, semester_id);

COMMENT ON COLUMN cp.academic_year_id IS 'Academic year scoping for curriculum planning';
COMMENT ON COLUMN cp.semester_id IS 'Semester scoping for curriculum planning';

-- ============================================================================
-- END OF MIGRATION
-- ============================================================================
