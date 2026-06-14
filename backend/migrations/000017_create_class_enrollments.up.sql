-- Create class_enrollments table for student enrollment in classes
CREATE TABLE class_enrollments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
  status VARCHAR(50) DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'INACTIVE', 'WITHDRAWN', 'COMPLETED')),
  notes TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(class_id, student_id)
);

-- Indexes for class_enrollments table
CREATE INDEX idx_class_enrollments_class_id ON class_enrollments(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_student_id ON class_enrollments(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_status ON class_enrollments(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_class_enrollments_active ON class_enrollments(status) WHERE deleted_at IS NULL AND status = 'ACTIVE';
