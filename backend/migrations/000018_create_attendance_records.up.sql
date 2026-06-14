-- Create attendance_records table for student attendance tracking
CREATE TABLE attendance_records (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  date DATE NOT NULL,
  status VARCHAR(50) NOT NULL CHECK (status IN ('PRESENT', 'ABSENT', 'LATE', 'EXCUSED')),
  notes TEXT,
  recorded_by UUID NOT NULL REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  UNIQUE(class_id, student_id, date)
);

-- Indexes for attendance_records table
CREATE INDEX idx_attendance_records_class_id ON attendance_records(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_student_id ON attendance_records(student_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_date ON attendance_records(date) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_status ON attendance_records(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_class_date ON attendance_records(class_id, date) WHERE deleted_at IS NULL;
CREATE INDEX idx_attendance_records_student_date ON attendance_records(student_id, date) WHERE deleted_at IS NULL;
