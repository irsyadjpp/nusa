-- Add full-text search indexes for text fields
-- This migration adds GIN indexes for full-text search on text fields

-- Add full-text search for CP description
CREATE INDEX idx_cps_text_search ON cps USING gin(to_tsvector('english', description)) WHERE deleted_at IS NULL;

-- Add full-text search for TP title
CREATE INDEX idx_tps_title_search ON tps USING gin(to_tsvector('english', title)) WHERE deleted_at IS NULL;

-- Add full-text search for Assessment title
CREATE INDEX idx_assessments_title_search ON assessments USING gin(to_tsvector('english', title)) WHERE deleted_at IS NULL;

-- Add full-text search for Announcement content
CREATE INDEX idx_announcements_content_search ON announcements USING gin(to_tsvector('english', content)) WHERE deleted_at IS NULL;

-- Add full-text search for Message content
CREATE INDEX idx_messages_content_search ON messages USING gin(to_tsvector('english', content)) WHERE deleted_at IS NULL;
