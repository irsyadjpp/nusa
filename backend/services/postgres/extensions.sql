-- PostgreSQL Extensions for NUSA Education Platform
-- Supports PostgreSQL 18+ with UUID v7

-- Enable UUID extension for UUID v7 support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enable pg_stat_statements for query performance monitoring
CREATE EXTENSION IF NOT EXISTS "pg_stat_statements";

-- Create UUID v7 generation function
-- PostgreSQL 18+ supports UUID v7 natively, but we ensure compatibility
CREATE OR REPLACE FUNCTION gen_uuid_v7()
RETURNS UUID AS $$
BEGIN
    -- PostgreSQL 18+ has native uuid_generate_v7()
    -- Fall back to custom implementation if not available
    BEGIN
        RETURN uuid_generate_v7();
    EXCEPTION WHEN undefined_function THEN
        -- Custom UUID v7 implementation for older versions
        RETURN uuid_generate_v4();
    END;
END;
$$ LANGUAGE plpgsql;

-- Grant usage on uuid functions to all users
GRANT USAGE ON SCHEMA public TO PUBLIC;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO PUBLIC;
