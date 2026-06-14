-- Create exams table for exam management
CREATE TABLE exams (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  exam_date DATE NOT NULL,
  start_time TIME NOT NULL,
  duration_minutes INTEGER NOT NULL,
  room VARCHAR(100),
  status VARCHAR(50) DEFAULT 'SCHEDULED' CHECK (status IN ('SCHEDULED', 'IN_PROGRESS', 'COMPLETED', 'CANCELLED')),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes for exams table
CREATE INDEX idx_exams_class_id ON exams(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_assessment_id ON exams(assessment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_exam_date ON exams(exam_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_status ON exams(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_exams_class_date ON exams(class_id, exam_date) WHERE deleted_at IS NULL;
