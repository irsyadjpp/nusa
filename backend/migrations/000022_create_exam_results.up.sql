-- Create exam_results table for exam result tracking
CREATE TABLE exam_results (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  exam_id UUID NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  score DECIMAL(5,2),
  grade VARCHAR(10),
  remarks TEXT,
  graded_at TIMESTAMP,
  graded_by UUID REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(exam_id, student_id)
);

-- Indexes for exam_results table
CREATE INDEX idx_exam_results_exam_id ON exam_results(exam_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_student_id ON exam_results(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_exam_results_grade ON exam_results(grade) WHERE deleted_at IS NULL;
