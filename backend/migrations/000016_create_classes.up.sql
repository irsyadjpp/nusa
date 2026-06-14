-- Create classes table for class management
CREATE TABLE classes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id UUID NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  academic_year_id UUID NOT NULL REFERENCES academic_years(id) ON DELETE CASCADE,
  semester_id UUID NOT NULL REFERENCES semesters(id) ON DELETE CASCADE,
  subject_id UUID NOT NULL REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
  teacher_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  grade_level VARCHAR(50) NOT NULL,
  room VARCHAR(100),
  max_students INTEGER DEFAULT 40,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes for classes table
CREATE INDEX idx_classes_school_id ON classes(school_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_academic_year_id ON classes(academic_year_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_semester_id ON classes(semester_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_subject_id ON classes(subject_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_teacher_id ON classes(teacher_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_school_semester ON classes(school_id, semester_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_classes_active ON classes(is_active) WHERE deleted_at IS NULL AND is_active = true;
