-- Kurikulum Merdeka 2026 Seeder
-- This migration seeds the database with standard Kurikulum Merdeka curriculum data
-- It is idempotent - can be run multiple times without creating duplicates

-- Helper function to check if data exists before inserting
DO $$ 
BEGIN
    -- Only insert Subjects if table is empty
    IF NOT EXISTS (SELECT 1 FROM curriculum_subjects LIMIT 1) THEN
        -- Core Kurikulum Merdeka Subjects
        INSERT INTO curriculum_subjects (id, code, name, name_en, description, is_active, created_at, updated_at) VALUES
        ('550e8400-e29b-41d4-a716-446655440000', 'PAI', 'Pendidikan Agama dan Budi Pekerti', 'Islamic Education and Character', 'Mata pelajaran agama dan pembentukan karakter', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440001', 'PPKn', 'Pendidikan Pancasila', 'Civic Education', 'Mata pelajaran pendidikan Pancasila dan kewarganegaraan', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440002', 'B_Indo', 'Bahasa Indonesia', 'Indonesian Language', 'Mata pelajaran bahasa Indonesia', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440003', 'MTK', 'Matematika', 'Mathematics', 'Mata pelajaran matematika', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440004', 'IPAS', 'IPAS', 'Science and Social Studies', 'Ilmu Pengetahuan Alam dan Ilmu Pengetahuan Sosial', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440005', 'S_Bud', 'Seni Budaya', 'Arts and Culture', 'Mata pelajaran seni budaya', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440006', 'PJOK', 'PJOK', 'Physical Education', 'Pendidikan Jasmani, Olahraga, dan Kesehatan', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440007', 'B_Ing', 'Bahasa Inggris', 'English Language', 'Mata pelajaran bahasa Inggris', true, NOW(), NOW()),
        ('550e8400-e29b-41d4-a716-446655440008', 'B_Daerah', 'Bahasa Daerah', 'Local Language', 'Mata pelajaran bahasa daerah', true, NOW(), NOW());
        
        RAISE NOTICE 'Inserted 9 Kurikulum Merdeka subjects';
    END IF;

    -- Only insert Phases if table is empty
    IF NOT EXISTS (SELECT 1 FROM curriculum_phases LIMIT 1) THEN
        INSERT INTO curriculum_phases (id, code, name, name_en, description, level, grade_range, is_active, created_at, updated_at) VALUES
        ('660e8400-e29b-41d4-a716-446655440000', 'Fase A', 'Fase A - SD Kelas 1-2', 'Phase A - Grades 1-2', 'Fase fondasi untuk kelas 1-2 SD', '1-2', 'Kelas 1-2', true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440001', 'Fase B', 'Fase B - SD Kelas 3-4', 'Phase B - Grades 3-4', 'Fase konseptual untuk kelas 3-4 SD', '3-4', 'Kelas 3-4', true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440002', 'Fase C', 'Fase C - SD Kelas 5-6', 'Phase C - Grades 5-6', 'Fase operasional untuk kelas 5-6 SD', '5-6', 'Kelas 5-6', true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440003', 'Fase D', 'Fase D - SMP Kelas 7-8', 'Phase D - Grades 7-8', 'Fase hipotesis untuk kelas 7-8 SMP', '7-8', 'Kelas 7-8', true, NOW(), NOW()),
        ('660e8400-e29b-41d4-a716-446655440004', 'Fase E', 'Fase E - SMP Kelas 9', 'Phase E - Grade 9', 'Fase aplikatif untuk kelas 9 SMP', '9', 'Kelas 9', true, NOW(), NOW());
        
        RAISE NOTICE 'Inserted 5 Kurikulum Merdeka phases';
    END IF;

    -- Only insert sample Elements if table is empty (one per subject-phase pair as example)
    IF NOT EXISTS (SELECT 1 FROM curriculum_elements LIMIT 1) THEN
        -- Example elements for IPAS in Phase C
        INSERT INTO curriculum_elements (id, subject_id, phase_id, code, name, name_en, description, is_active, created_at, updated_at) VALUES
        ('770e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002', 'IPAS_IPA_C', 'IPAS - Ilmu Pengetahuan Alam', 'Science', 'Elemen IPAs untuk Ilmu Pengetahuan Alam di Fase C', true, NOW(), NOW()),
        ('770e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002', 'IPAS_IPS_C', 'IPAS - Ilmu Pengetahuan Sosial', 'Social Studies', 'Elemen IPAS untuk Ilmu Pengetahuan Sosial di Fase C', true, NOW(), NOW());
        
        RAISE NOTICE 'Inserted sample curriculum elements';
    END IF;

    -- Only insert sample Subelements if table is empty
    IF NOT EXISTS (SELECT 1 FROM curriculum_subelements LIMIT 1) THEN
        INSERT INTO curriculum_subelements (id, element_id, code, name, name_en, description, is_active, created_at, updated_at) VALUES
        ('880e8400-e29b-41d4-a716-446655440000', '770e8400-e29b-41d4-a716-446655440000', 'IPA_C_Benda', 'Struktur dan Sifat Benda', 'Structure and Properties of Matter', 'Memahami struktur dan sifat benda di sekitar', true, NOW(), NOW()),
        ('880e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440000', 'IPA_C_Energi', 'Energi dan Perubahannya', 'Energy and Its Changes', 'Memahami berbagai bentuk energi dan perubahannya', true, NOW(), NOW());
        
        RAISE NOTICE 'Inserted sample curriculum subelements';
    END IF;

    RAISE NOTICE 'Kurikulum Merdeka seeder completed successfully';
END $$;
