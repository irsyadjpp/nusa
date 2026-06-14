-- Create assignments table for assignment management
CREATE TABLE assignments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  due_date TIMESTAMP NOT NULL,
  max_score INTEGER DEFAULT 100,
  status VARCHAR(50) DEFAULT 'ASSIGNED' CHECK (status IN ('ASSIGNED', 'IN_PROGRESS', 'SUBMITTED', 'GRADED', 'CANCELLED')),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes for assignments table
CREATE INDEX idx_assignments_class_id ON assignments(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_assessment_id ON assignments(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_due_date ON assignments(due_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_status ON assignments(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_assignments_class_due ON assignments(class_id, due_date) WHERE deleted_at IS NULL;
