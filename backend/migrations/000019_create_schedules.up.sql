-- Create schedules table for class scheduling
CREATE TABLE schedules (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  day_of_week INTEGER NOT NULL CHECK (day_of_week >= 1 AND day_of_week <= 7),
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  room VARCHAR(100),
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_by UUID REFERENCES users(id),
  deleted_at TIMESTAMP
);

-- Indexes for schedules table
CREATE INDEX idx_schedules_class_id ON schedules(class_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_day_of_week ON schedules(day_of_week) WHERE deleted_at IS NULL;
CREATE INDEX idx_schedules_active ON schedules(is_active) WHERE deleted_at IS NULL AND is_active = true;
CREATE INDEX idx_schedules_room ON schedules(room) WHERE deleted_at IS NULL;
