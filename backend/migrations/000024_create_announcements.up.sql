-- Create announcements table for school announcements
CREATE TABLE announcements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id UUID NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  priority VARCHAR(50) DEFAULT 'NORMAL' CHECK (priority IN ('LOW', 'NORMAL', 'HIGH', 'URGENT')),
  target_audience VARCHAR(50) DEFAULT 'ALL' CHECK (target_audience IN ('ALL', 'TEACHERS', 'STUDENTS', 'PARENTS', 'ADMIN')),
  published_by UUID NOT NULL REFERENCES users(id),
  published_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- Indexes for announcements table
CREATE INDEX idx_announcements_school_id ON announcements(school_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_priority ON announcements(priority) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_target_audience ON announcements(target_audience) WHERE deleted_at IS NULL;
CREATE INDEX idx_announcements_active ON announcements(is_active) WHERE deleted_at IS NULL AND is_active = true;
CREATE INDEX idx_announcements_school_active ON announcements(school_id, is_active) WHERE deleted_at IS NULL AND is_active = true;
