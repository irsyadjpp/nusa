-- Initial schema for NUSA Education Operating System
-- This migration creates the foundation tables for the system
-- Based on 14_DATABASE_SCHEMA.md
-- PostgreSQL 18+ with UUID v7 support
-- Enable UUID extension for UUID v7 support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Create UUID v7 generation function for PostgreSQL 18+ compatibility
CREATE OR REPLACE FUNCTION gen_uuid_v7() RETURNS UUID AS $$ BEGIN -- PostgreSQL 18+ has native uuid_generate_v7()
    -- Fall back to uuid_generate_v4() for compatibility
    BEGIN RETURN uuid_generate_v7();
EXCEPTION
WHEN undefined_function THEN RETURN uuid_generate_v4();
END;
END;
$$ LANGUAGE plpgsql;
-- Schools table
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_by UUID
);
-- Roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- Permissions table
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_permission UNIQUE (role_id, resource, action)
);
-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role_id UUID NOT NULL REFERENCES roles(id),
    school_id UUID REFERENCES schools(id) ON DELETE
    SET NULL,
        is_active BOOLEAN NOT NULL DEFAULT true,
        failed_login_attempts INTEGER NOT NULL DEFAULT 0,
        locked_until TIMESTAMP WITH TIME ZONE,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        created_by UUID REFERENCES users(id),
        updated_by UUID REFERENCES users(id),
        CONSTRAINT chk_users_email_format CHECK (
            email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
        ),
        CONSTRAINT chk_users_failed_login_attempts CHECK (failed_login_attempts >= 0)
);
-- Refresh tokens table
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    ip_address INET,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by UUID,
    CONSTRAINT chk_refresh_tokens_expires_at CHECK (expires_at > created_at)
);
-- AI Generation Logs table
CREATE TABLE ai_generation_logs (
    id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    school_id UUID REFERENCES schools(id) ON DELETE
    SET NULL,
        artifact_type VARCHAR(50) NOT NULL,
        artifact_id UUID,
        provider VARCHAR(50) NOT NULL,
        model VARCHAR(100),
        tokens_used INTEGER,
        estimated_cost DECIMAL(10, 4),
        response_time_ms INTEGER,
        status VARCHAR(20) NOT NULL,
        error_message TEXT,
        prompt_snapshot TEXT,
        response_snapshot TEXT,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- Indexes for schools
CREATE INDEX idx_schools_code ON schools(code);
CREATE INDEX idx_schools_is_active ON schools(is_active);
-- Indexes for roles
CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_is_active ON roles(is_active);
-- Indexes for permissions
CREATE INDEX idx_permissions_role_id ON permissions(role_id);
CREATE INDEX idx_permissions_resource ON permissions(resource);
-- Indexes for users
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_school_id ON users(school_id);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_created_at ON users(created_at);
-- Indexes for refresh tokens
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
-- Indexes for ai generation logs
CREATE INDEX idx_ai_generation_logs_user_id ON ai_generation_logs(user_id);
CREATE INDEX idx_ai_generation_logs_school_id ON ai_generation_logs(school_id);
CREATE INDEX idx_ai_generation_logs_artifact ON ai_generation_logs(artifact_type, artifact_id);
CREATE INDEX idx_ai_generation_logs_created_at ON ai_generation_logs(created_at);