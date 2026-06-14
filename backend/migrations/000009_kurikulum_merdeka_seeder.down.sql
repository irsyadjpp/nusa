-- Rollback Kurikulum Merdeka Seeder
-- Removes the seeded data while preserving any user-created data

DO $$
BEGIN
    -- Remove seeded subelements (check by ID prefix)
    DELETE FROM curriculum_subelements 
    WHERE id LIKE '880e8400-e29b-41d4-a716-446655440%';
    
    -- Remove seeded elements (check by ID prefix)
    DELETE FROM curriculum_elements 
    WHERE id LIKE '770e8400-e29b-41d4-a716-446655440%';
    
    -- Remove seeded phases (check by ID prefix)
    DELETE FROM curriculum_phases 
    WHERE id LIKE '660e8400-e29b-41d4-a716-446655440%';
    
    -- Remove seeded subjects (check by ID prefix)
    DELETE FROM curriculum_subjects 
    WHERE id LIKE '550e8400-e29b-41d4-a716-446655440%';
    
    RAISE NOTICE 'Kurikulum Merdeka seeder rollback completed';
END $$;
