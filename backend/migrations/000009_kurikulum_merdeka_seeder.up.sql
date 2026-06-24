-- Kurikulum Merdeka 2026 Seeder (Updated for Current Schema)
-- This migration seeds the database with standard Kurikulum Merdeka curriculum data
-- Updated to match current schema without name_en columns and with proper phase structure
-- It is idempotent - can be run multiple times without creating duplicates

-- Helper function to check if data exists before inserting
DO $$
BEGIN
    -- Only insert Subjects if table has less than 9 subjects
    IF (SELECT COUNT(*) FROM curriculum_subjects) < 9 THEN
        -- Clear existing subjects if any (to allow fresh seeding)
        DELETE FROM curriculum_subjects;

        -- Core Kurikulum Merdeka Subjects
        INSERT INTO curriculum_subjects (id, code, name, description, is_active, created_at, updated_at) VALUES
        ('550e8400-e29b-41d4-a716-446655440000', 'PAI', 'Pendidikan Agama dan Budi Pekerti', 'Mata pelajaran agama dan pembentukan karakter', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440001', 'PPKn', 'Pendidikan Pancasila', 'Mata pelajaran pendidikan Pancasila dan kewarganegaraan', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440002', 'B_Indo', 'Bahasa Indonesia', 'Mata pelajaran bahasa Indonesia', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440003', 'MTK', 'Matematika', 'Mata pelajaran matematika', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440004', 'IPAS', 'IPAS', 'Ilmu Pengetahuan Alam dan Ilmu Pengetahuan Sosial', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440005', 'S_Bud', 'Seni Budaya', 'Mata pelajaran seni budaya', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440006', 'PJOK', 'PJOK', 'Pendidikan Jasmani, Olahraga, dan Kesehatan', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440007', 'B_Ing', 'Bahasa Inggris', 'Mata pelajaran bahasa Inggris', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440008', 'B_Daerah', 'Bahasa Daerah', 'Mata pelajaran bahasa daerah', true, NOW(), NOW());

        RAISE NOTICE 'Inserted 9 Kurikulum Merdeka subjects';
    END IF;

    -- Only insert Phases if table is empty
    IF (SELECT COUNT(*) FROM curriculum_phases) = 0 THEN
        -- Core Kurikulum Merdeka Phases with proper grade_level_start/end
        INSERT INTO curriculum_phases (id, code, name, description, grade_level_start, grade_level_end, is_active, created_at, updated_at) VALUES
        ('660e8400-e29b-41d4-a716-446655440000', 'Fase A', 'Fase A - SD Kelas 1-2', 'Fase fondasi untuk kelas 1-2 SD', 1, 2, true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440001', 'Fase B', 'Fase B - SD Kelas 3-4', 'Fase konseptual untuk kelas 3-4 SD', 3, 4, true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440002', 'Fase C', 'Fase C - SD Kelas 5-6', 'Fase operasional untuk kelas 5-6 SD', 5, 6, true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440003', 'Fase D', 'Fase D - SMP Kelas 7-8', 'Fase hipotesis untuk kelas 7-8 SMP', 7, 8, true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440004', 'Fase E', 'Fase E - SMP Kelas 9', 'Fase aplikatif untuk kelas 9 SMP', 9, 9, true, NOW(), NOW());

        RAISE NOTICE 'Inserted 5 Kurikulum Merdeka phases';
    END IF;

    -- Only insert sample Elements if table is empty
    IF (SELECT COUNT(*) FROM curriculum_elements) = 0 THEN
        -- Example elements for IPAS in Phase C
        INSERT INTO curriculum_elements (id, subject_id, phase_id, code, name, description, is_active, created_at, updated_at) VALUES
        ('770e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002', 'IPAS_IPA_C', 'IPAS - Ilmu Pengetahuan Alam', 'Elemen IPAs untuk Ilmu Pengetahuan Alam di Fase C', true, NOW(), NOW()),
        ('770e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002', 'IPAS_IPS_C', 'IPAS - Ilmu Pengetahuan Sosial', 'Elemen IPAS untuk Ilmu Pengetahuan Sosial di Fase C', true, NOW(), NOW());

        RAISE NOTICE 'Inserted sample curriculum elements';
    END IF;

    -- Only insert sample Subelements if table is empty
    IF (SELECT COUNT(*) FROM curriculum_subelements) = 0 THEN
        INSERT INTO curriculum_subelements (id, element_id, code, name, description, is_active, created_at, updated_at) VALUES
        ('880e8400-e29b-41d4-a716-446655440000', '770e8400-e29b-41d4-a716-446655440000', 'IPA_C_Benda', 'Struktur dan Sifat Benda', 'Memahami struktur dan sifat benda di sekitar', true, NOW(), NOW()),
        ('880e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440000', 'IPA_C_Energi', 'Energi dan Perubahannya', 'Memahami berbagai bentuk energi dan perubahannya', true, NOW(), NOW());

        RAISE NOTICE 'Inserted sample curriculum subelements';
    END IF;

    RAISE NOTICE 'Kurikulum Merdeka seeder completed successfully';
END $$;
