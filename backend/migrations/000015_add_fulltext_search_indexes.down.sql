-- Remove full-text search indexes
-- This migration removes the GIN indexes added in the up migration

DROP INDEX IF EXISTS idx_messages_content_search;
DROP INDEX IF EXISTS idx_announcements_content_search;
DROP INDEX IF EXISTS idx_assessments_title_search;
DROP INDEX IF EXISTS idx_tps_title_search;
DROP INDEX IF EXISTS idx_cps_text_search;
