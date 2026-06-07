-- Education Domain Tables Migration
-- Based on 14_DATABASE_SCHEMA.md
-- This migration creates all education domain tables for Curriculum, Learning Planning, Assessment, and Reporting modules
-- ============================================
-- CURRICULUM MODULE TABLES
-- ============================================
-- Table: curriculum_subjects
CREATE TABLE curriculum_subjects (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  code VARCHAR(50) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description TEXT,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_curriculum_subjects_code ON curriculum_subjects(code);
CREATE INDEX idx_curriculum_subjects_is_active ON curriculum_subjects(is_active);
-- Table: curriculum_phases
CREATE TABLE curriculum_phases (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  code VARCHAR(50) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description TEXT,
  grade_level_start INTEGER,
  grade_level_end INTEGER,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_curriculum_phases_code ON curriculum_phases(code);
CREATE INDEX idx_curriculum_phases_is_active ON curriculum_phases(is_active);
-- Table: curriculum_elements
CREATE TABLE curriculum_elements (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  subject_id UUID NOT NULL REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
  phase_id UUID NOT NULL REFERENCES curriculum_phases(id) ON DELETE CASCADE,
  code VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description TEXT,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_curriculum_elements_subject_phase_code UNIQUE (subject_id, phase_id, code)
);
CREATE INDEX idx_curriculum_elements_subject_id ON curriculum_elements(subject_id);
CREATE INDEX idx_curriculum_elements_phase_id ON curriculum_elements(phase_id);
CREATE INDEX idx_curriculum_elements_code ON curriculum_elements(code);
CREATE INDEX idx_curriculum_elements_is_active ON curriculum_elements(is_active);
-- Table: curriculum_subelements
CREATE TABLE curriculum_subelements (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  element_id UUID NOT NULL REFERENCES curriculum_elements(id) ON DELETE CASCADE,
  code VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description TEXT,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_curriculum_subelements_element_code UNIQUE (element_id, code)
);
CREATE INDEX idx_curriculum_subelements_element_id ON curriculum_subelements(element_id);
CREATE INDEX idx_curriculum_subelements_code ON curriculum_subelements(code);
CREATE INDEX idx_curriculum_subelements_is_active ON curriculum_subelements(is_active);
-- Table: cp (Curriculum Plans)
CREATE TABLE cp (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  subject_id UUID NOT NULL REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
  phase_id UUID NOT NULL REFERENCES curriculum_phases(id) ON DELETE CASCADE,
  element_id UUID NOT NULL REFERENCES curriculum_elements(id) ON DELETE CASCADE,
  subelement_id UUID NOT NULL REFERENCES curriculum_subelements(id) ON DELETE CASCADE,
  code VARCHAR(50) NOT NULL,
  description TEXT NOT NULL,
  competency_code VARCHAR(50),
  learning_objectives JSONB NOT NULL,
  competency_standards JSONB NOT NULL,
  time_allocation_hours INTEGER NOT NULL,
  hours_per_week INTEGER NOT NULL,
  version VARCHAR(20) NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT true,
  imported_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  imported_by UUID REFERENCES users(id),
  CONSTRAINT uq_cp_hierarchy_code UNIQUE (
    subject_id,
    phase_id,
    element_id,
    subelement_id,
    code
  ),
  CONSTRAINT chk_cp_time_allocation_hours CHECK (time_allocation_hours > 0),
  CONSTRAINT chk_cp_hours_per_week CHECK (
    hours_per_week > 0
    AND hours_per_week <= 40
  )
);
CREATE INDEX idx_cp_subject_id ON cp(subject_id);
CREATE INDEX idx_cp_phase_id ON cp(phase_id);
CREATE INDEX idx_cp_element_id ON cp(element_id);
CREATE INDEX idx_cp_subelement_id ON cp(subelement_id);
CREATE INDEX idx_cp_code ON cp(code);
CREATE INDEX idx_cp_version ON cp(version);
CREATE INDEX idx_cp_is_active ON cp(is_active);
-- ============================================
-- LEARNING PLANNING MODULE TABLES
-- ============================================
-- Table: tp_sets (Teaching Plan Sets)
CREATE TABLE tp_sets (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  cp_id UUID NOT NULL REFERENCES cp(id) ON DELETE CASCADE,
  version_no INTEGER NOT NULL DEFAULT 1,
  status VARCHAR(20) NOT NULL,
  generation_source VARCHAR(50) NOT NULL,
  generation_reason TEXT,
  generated_by UUID NOT NULL REFERENCES users(id),
  ai_generation_id UUID REFERENCES ai_generation_logs(id),
  approved_by UUID REFERENCES users(id),
  approved_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_tp_sets_cp_version UNIQUE (cp_id, version_no),
  CONSTRAINT chk_tp_sets_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  ),
  CONSTRAINT chk_tp_sets_generation_source CHECK (generation_source IN ('AI_GENERATED', 'MANUAL'))
);
CREATE INDEX idx_tp_sets_cp_id ON tp_sets(cp_id);
CREATE INDEX idx_tp_sets_version_no ON tp_sets(version_no);
CREATE INDEX idx_tp_sets_status ON tp_sets(status);
CREATE INDEX idx_tp_sets_generated_by ON tp_sets(generated_by);
CREATE INDEX idx_tp_sets_approved_by ON tp_sets(approved_by);
CREATE INDEX idx_tp_sets_ai_generation_id ON tp_sets(ai_generation_id);
-- Table: tp (Teaching Plan Items)
CREATE TABLE tp (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  tp_set_id UUID NOT NULL REFERENCES tp_sets(id) ON DELETE CASCADE,
  sequence_number INTEGER NOT NULL,
  cp_id UUID NOT NULL REFERENCES cp(id) ON DELETE CASCADE,
  subject_id UUID NOT NULL REFERENCES curriculum_subjects(id) ON DELETE CASCADE,
  phase_id UUID NOT NULL REFERENCES curriculum_phases(id) ON DELETE CASCADE,
  element_id UUID NOT NULL REFERENCES curriculum_elements(id) ON DELETE CASCADE,
  subelement_id UUID NOT NULL REFERENCES curriculum_subelements(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  status VARCHAR(20) NOT NULL,
  title VARCHAR(255),
  learning_objectives JSONB NOT NULL,
  time_allocation JSONB NOT NULL,
  prerequisites JSONB,
  estimated_weeks INTEGER,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_tp_sequence_number CHECK (sequence_number > 0),
  CONSTRAINT chk_tp_estimated_weeks CHECK (
    estimated_weeks IS NULL
    OR estimated_weeks > 0
  ),
  CONSTRAINT chk_tp_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  )
);
CREATE INDEX idx_tp_tp_set_id ON tp(tp_set_id);
CREATE INDEX idx_tp_sequence_number ON tp(tp_set_id, sequence_number);
CREATE INDEX idx_tp_cp_id ON tp(cp_id);
CREATE INDEX idx_tp_subject_id ON tp(subject_id);
CREATE INDEX idx_tp_phase_id ON tp(phase_id);
CREATE INDEX idx_tp_element_id ON tp(element_id);
CREATE INDEX idx_tp_subelement_id ON tp(subelement_id);
CREATE INDEX idx_tp_user_id ON tp(user_id);
CREATE INDEX idx_tp_status ON tp(status);
CREATE INDEX idx_tp_created_at ON tp(created_at);
-- Table: atp_sets (Annual Teaching Plan Sets)
CREATE TABLE atp_sets (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  tp_set_id UUID NOT NULL REFERENCES tp_sets(id) ON DELETE CASCADE,
  version_no INTEGER NOT NULL DEFAULT 1,
  status VARCHAR(20) NOT NULL,
  generation_source VARCHAR(50) NOT NULL,
  generation_reason TEXT,
  generated_by UUID NOT NULL REFERENCES users(id),
  ai_generation_id UUID REFERENCES ai_generation_logs(id),
  approved_by UUID REFERENCES users(id),
  approved_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_atp_sets_tp_version UNIQUE (tp_set_id, version_no),
  CONSTRAINT chk_atp_sets_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  )
);
CREATE INDEX idx_atp_sets_tp_set_id ON atp_sets(tp_set_id);
CREATE INDEX idx_atp_sets_version_no ON atp_sets(version_no);
CREATE INDEX idx_atp_sets_status ON atp_sets(status);
CREATE INDEX idx_atp_sets_generated_by ON atp_sets(generated_by);
CREATE INDEX idx_atp_sets_approved_by ON atp_sets(approved_by);
CREATE INDEX idx_atp_sets_ai_generation_id ON atp_sets(ai_generation_id);
-- Table: atp (Annual Teaching Plan Items)
CREATE TABLE atp (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  atp_set_id UUID NOT NULL REFERENCES atp_sets(id) ON DELETE CASCADE,
  tp_id UUID NOT NULL REFERENCES tp(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  status VARCHAR(20) NOT NULL,
  academic_calendar JSONB NOT NULL,
  class_schedule JSONB NOT NULL,
  weekly_sequence JSONB NOT NULL,
  assessment_schedule JSONB,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_atp_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  )
);
CREATE INDEX idx_atp_atp_set_id ON atp(atp_set_id);
CREATE INDEX idx_atp_tp_id ON atp(tp_id);
CREATE INDEX idx_atp_user_id ON atp(user_id);
CREATE INDEX idx_atp_status ON atp(status);
CREATE INDEX idx_atp_created_at ON atp(created_at);
-- Table: modul_ajar_sets (Modul Ajar Sets)
CREATE TABLE modul_ajar_sets (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  atp_set_id UUID NOT NULL REFERENCES atp_sets(id) ON DELETE CASCADE,
  version_no INTEGER NOT NULL DEFAULT 1,
  status VARCHAR(20) NOT NULL,
  generation_source VARCHAR(50) NOT NULL,
  generation_reason TEXT,
  generated_by UUID NOT NULL REFERENCES users(id),
  ai_generation_id UUID REFERENCES ai_generation_logs(id),
  approved_by UUID REFERENCES users(id),
  approved_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_modul_ajar_sets_atp_version UNIQUE (atp_set_id, version_no),
  CONSTRAINT chk_modul_ajar_sets_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  )
);
CREATE INDEX idx_modul_ajar_sets_atp_set_id ON modul_ajar_sets(atp_set_id);
CREATE INDEX idx_modul_ajar_sets_version_no ON modul_ajar_sets(version_no);
CREATE INDEX idx_modul_ajar_sets_status ON modul_ajar_sets(status);
CREATE INDEX idx_modul_ajar_sets_generated_by ON modul_ajar_sets(generated_by);
CREATE INDEX idx_modul_ajar_sets_approved_by ON modul_ajar_sets(approved_by);
CREATE INDEX idx_modul_ajar_sets_ai_generation_id ON modul_ajar_sets(ai_generation_id);
-- Table: modul_ajar (Modul Ajar Items)
CREATE TABLE modul_ajar (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  modul_ajar_set_id UUID NOT NULL REFERENCES modul_ajar_sets(id) ON DELETE CASCADE,
  atp_id UUID NOT NULL REFERENCES atp(id) ON DELETE CASCADE,
  week INTEGER NOT NULL,
  topic JSONB NOT NULL,
  resources JSONB,
  class_characteristics JSONB,
  learning_activities JSONB NOT NULL,
  resource_requirements JSONB,
  assessment_methods JSONB,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_modul_ajar_week CHECK (week >= 1),
  CONSTRAINT chk_modul_ajar_status CHECK (
    status IN (
      'DRAFT',
      'UNDER_REVIEW',
      'APPROVED',
      'REJECTED',
      'ARCHIVED'
    )
  )
);
CREATE INDEX idx_modul_ajar_modul_ajar_set_id ON modul_ajar(modul_ajar_set_id);
CREATE INDEX idx_modul_ajar_atp_id ON modul_ajar(atp_id);
CREATE INDEX idx_modul_ajar_week ON modul_ajar(week);
CREATE INDEX idx_modul_ajar_status ON modul_ajar(status);
CREATE INDEX idx_modul_ajar_created_at ON modul_ajar(created_at);
-- ============================================
-- ASSESSMENT MODULE TABLES
-- ============================================
-- Table: assessments
CREATE TABLE assessments (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  modul_ajar_id UUID NOT NULL REFERENCES modul_ajar(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  assessment_type VARCHAR(20) NOT NULL,
  status VARCHAR(20) NOT NULL,
  assessment_items JSONB NOT NULL,
  answer_key JSONB NOT NULL,
  scoring_guidelines JSONB NOT NULL,
  ai_confidence_score DECIMAL(3, 2),
  ai_generated_at TIMESTAMP WITH TIME ZONE,
  ai_agent_version VARCHAR(20),
  version_no INTEGER NOT NULL DEFAULT 1,
  is_current_version BOOLEAN NOT NULL DEFAULT true,
  parent_version_id UUID REFERENCES assessments(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  approved_at TIMESTAMP WITH TIME ZONE,
  approved_by UUID REFERENCES users(id),
  CONSTRAINT chk_assessments_assessment_type CHECK (assessment_type IN ('FORMATIVE', 'SUMMATIVE')),
  CONSTRAINT chk_assessments_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED')),
  CONSTRAINT chk_assessments_ai_confidence_score CHECK (
    ai_confidence_score IS NULL
    OR (
      ai_confidence_score >= 0.00
      AND ai_confidence_score <= 1.00
    )
  )
);
CREATE INDEX idx_assessments_modul_ajar_id ON assessments(modul_ajar_id);
CREATE INDEX idx_assessments_user_id ON assessments(user_id);
CREATE INDEX idx_assessments_assessment_type ON assessments(assessment_type);
CREATE INDEX idx_assessments_status ON assessments(status);
CREATE INDEX idx_assessments_created_at ON assessments(created_at);
CREATE INDEX idx_assessments_approved_at ON assessments(approved_at);
-- Table: rubrics
CREATE TABLE rubrics (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  rubric_type VARCHAR(20) NOT NULL,
  status VARCHAR(20) NOT NULL,
  performance_criteria JSONB NOT NULL,
  performance_levels JSONB NOT NULL,
  scoring_guidelines JSONB NOT NULL,
  ai_confidence_score DECIMAL(3, 2),
  ai_generated_at TIMESTAMP WITH TIME ZONE,
  ai_agent_version VARCHAR(20),
  version_no INTEGER NOT NULL DEFAULT 1,
  is_current_version BOOLEAN NOT NULL DEFAULT true,
  parent_version_id UUID REFERENCES rubrics(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  approved_at TIMESTAMP WITH TIME ZONE,
  approved_by UUID REFERENCES users(id),
  CONSTRAINT chk_rubrics_rubric_type CHECK (rubric_type IN ('ANALYTIC', 'HOLISTIC')),
  CONSTRAINT chk_rubrics_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED')),
  CONSTRAINT chk_rubrics_ai_confidence_score CHECK (
    ai_confidence_score IS NULL
    OR (
      ai_confidence_score >= 0.00
      AND ai_confidence_score <= 1.00
    )
  )
);
CREATE INDEX idx_rubrics_assessment_id ON rubrics(assessment_id);
CREATE INDEX idx_rubrics_user_id ON rubrics(user_id);
CREATE INDEX idx_rubrics_rubric_type ON rubrics(rubric_type);
CREATE INDEX idx_rubrics_status ON rubrics(status);
CREATE INDEX idx_rubrics_created_at ON rubrics(created_at);
CREATE INDEX idx_rubrics_approved_at ON rubrics(approved_at);
-- Table: evidences
CREATE TABLE evidences (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  student_id VARCHAR(50) NOT NULL,
  assessment_id UUID NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  evidence_type VARCHAR(50) NOT NULL,
  status VARCHAR(20) NOT NULL,
  evidence_data JSONB NOT NULL,
  teacher_notes TEXT,
  rubric_id UUID REFERENCES rubrics(id),
  linked_criteria JSONB,
  evaluation_notes TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_evidences_evidence_type CHECK (
    evidence_type IN (
      'STUDENT_WORK',
      'ASSESSMENT_RESULT',
      'OBSERVATION'
    )
  ),
  CONSTRAINT chk_evidences_status CHECK (status IN ('COLLECTED', 'LINKED', 'EVALUATED'))
);
CREATE INDEX idx_evidences_student_id ON evidences(student_id);
CREATE INDEX idx_evidences_assessment_id ON evidences(assessment_id);
CREATE INDEX idx_evidences_user_id ON evidences(user_id);
CREATE INDEX idx_evidences_evidence_type ON evidences(evidence_type);
CREATE INDEX idx_evidences_status ON evidences(status);
CREATE INDEX idx_evidences_rubric_id ON evidences(rubric_id);
CREATE INDEX idx_evidences_created_at ON evidences(created_at);
-- Table: evaluations
CREATE TABLE evaluations (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  student_id VARCHAR(50) NOT NULL,
  rubric_id UUID NOT NULL REFERENCES rubrics(id) ON DELETE CASCADE,
  evidence_id UUID NOT NULL REFERENCES evidences(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id),
  performance_scores JSONB NOT NULL,
  total_score INTEGER NOT NULL,
  max_score INTEGER NOT NULL,
  performance_level VARCHAR(20) NOT NULL,
  evaluated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT chk_evaluations_total_score CHECK (total_score >= 0),
  CONSTRAINT chk_evaluations_max_score CHECK (max_score > 0),
  CONSTRAINT chk_evaluations_total_score_max CHECK (total_score <= max_score),
  CONSTRAINT chk_evaluations_performance_level CHECK (
    performance_level IN (
      'EXCELLENT',
      'PROFICIENT',
      'DEVELOPING',
      'BEGINNING'
    )
  )
);
CREATE INDEX idx_evaluations_student_id ON evaluations(student_id);
CREATE INDEX idx_evaluations_rubric_id ON evaluations(rubric_id);
CREATE INDEX idx_evaluations_evidence_id ON evaluations(evidence_id);
CREATE INDEX idx_evaluations_user_id ON evaluations(user_id);
CREATE INDEX idx_evaluations_performance_level ON evaluations(performance_level);
CREATE INDEX idx_evaluations_evaluated_at ON evaluations(evaluated_at);
-- ============================================
-- REPORTING MODULE TABLES
-- ============================================
-- Table: narrative_reports
CREATE TABLE narrative_reports (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  student_id VARCHAR(50) NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id),
  status VARCHAR(20) NOT NULL,
  report_period JSONB NOT NULL,
  language VARCHAR(20) NOT NULL,
  content JSONB NOT NULL,
  ai_confidence_score DECIMAL(3, 2),
  ai_generated_at TIMESTAMP WITH TIME ZONE,
  ai_agent_version VARCHAR(20),
  version_no INTEGER NOT NULL DEFAULT 1,
  is_current_version BOOLEAN NOT NULL DEFAULT true,
  parent_version_id UUID REFERENCES narrative_reports(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  approved_at TIMESTAMP WITH TIME ZONE,
  approved_by UUID REFERENCES users(id),
  CONSTRAINT chk_narrative_reports_status CHECK (status IN ('DRAFT', 'APPROVED', 'REJECTED')),
  CONSTRAINT chk_narrative_reports_language CHECK (language IN ('INDONESIAN', 'ENGLISH')),
  CONSTRAINT chk_narrative_reports_ai_confidence_score CHECK (
    ai_confidence_score IS NULL
    OR (
      ai_confidence_score >= 0.00
      AND ai_confidence_score <= 1.00
    )
  )
);
CREATE INDEX idx_narrative_reports_student_id ON narrative_reports(student_id);
CREATE INDEX idx_narrative_reports_user_id ON narrative_reports(user_id);
CREATE INDEX idx_narrative_reports_status ON narrative_reports(status);
CREATE INDEX idx_narrative_reports_language ON narrative_reports(language);
CREATE INDEX idx_narrative_reports_created_at ON narrative_reports(created_at);
CREATE INDEX idx_narrative_reports_approved_at ON narrative_reports(approved_at);
-- ============================================
-- ADMINISTRATION MODULE TABLES
-- ============================================
-- Table: audit_logs
CREATE TABLE audit_logs (
  id UUID PRIMARY KEY DEFAULT gen_uuid_v7(),
  user_id UUID REFERENCES users(id),
  action VARCHAR(50) NOT NULL,
  entity_type VARCHAR(50) NOT NULL,
  entity_id UUID NOT NULL,
  changes JSONB,
  ip_address INET,
  user_agent TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_entity_type ON audit_logs(entity_type);
CREATE INDEX idx_audit_logs_entity_id ON audit_logs(entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);
CREATE INDEX idx_audit_logs_user_id_created_at ON audit_logs(user_id, created_at);