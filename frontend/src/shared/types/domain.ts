/**
 * Domain Types for NUSA Platform
 * Shared type definitions for all domain entities
 * This file provides comprehensive TypeScript types for the Kurikulum Merdeka domain
 */

// ============================================================================
// TP (Teaching Plan) Types
// ============================================================================

export interface KKTPCriteria {
  mastery_thresholds: string[];
  performance_indicators: string[];
  minimum_requirements: string[];
  minimum_mastery_level?: number; // Convenience property matching component expectations
}

export interface TP {
  id: string;
  tp_set_id: string;
  sequence_number: number;
  cp_id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  user_id: string;
  status: TPStatus;
  title: string;
  learning_objectives: LearningObjectives;
  time_allocation: TimeAllocation;
  prerequisites: Prerequisites;
  estimated_weeks: number;
  success_criteria: KKTPCriteria;
  version_no: number;
  is_current_version: boolean;
  parent_version_id?: string;
  created_at: string;
  updated_at: string;
}

export type TPStatus = 'DRAFT' | 'UNDER_REVIEW' | 'APPROVED' | 'REJECTED' | 'ARCHIVED';

export interface LearningObjectives {
  main_objective: string;
  supporting_objectives: string[];
}

export interface TimeAllocation {
  total_hours: number;
  per_week_hours: number;
  hours_per_week: number; // Convenience property matching component expectations
  hours_per_session?: number; // Convenience property matching component expectations
  breakdown: Record<string, number>;
}

export interface Prerequisites {
  required_tps: string[];
  required_skills: string[];
  notes?: string;
}

export interface TPSet {
  id: string;
  cp_id: string;
  version_no: number;
  status: TPStatus;
  generation_source: 'MANUAL' | 'AI_GENERATED';
  generation_reason?: string;
  generated_by: string;
  ai_generation_id?: string;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
  // Convenience properties matching component expectations
  name?: string;
}

// ============================================================================
// Assessment Types
// ============================================================================

export interface Assessment {
  id: string;
  tp_id: string;
  tp_title?: string;
  tp_version_no: number;
  success_criteria_snapshot: KKTPCriteria;
  user_id: string;
  user_name?: string;
  assessment_type: AssessmentType;
  status: AssessmentStatus;
  assessment_items: AssessmentItems;
  answer_key: AnswerKey;
  scoring_guidelines: ScoringGuidelines;
  ai_confidence_score?: number;
  ai_generated_at?: string;
  ai_agent_version?: string;
  version_no: number;
  is_current_version: boolean;
  parent_version_id?: string;
  created_at: string;
  updated_at: string;
  approved_at?: string;
  approved_by?: string;
}

export type AssessmentType = 'FORMATIVE' | 'SUMMATIVE';

export type AssessmentStatus = 'DRAFT' | 'UNDER_REVIEW' | 'APPROVED' | 'REJECTED' | 'ARCHIVED';

export interface AssessmentItems {
  questions: AssessmentQuestion[];
  total_score: number;
  duration_minutes: number;
}

export interface AssessmentQuestion {
  id: string;
  question_text: string;
  question_type: 'MULTIPLE_CHOICE' | 'ESSAY' | 'SHORT_ANSWER' | 'PRACTICAL';
  points: number;
  options?: string[];
  correct_answer?: string;
  rubric_criteria?: string[];
}

export interface AnswerKey {
  version: string;
  answers: Record<string, string | string[]>;
  notes?: Record<string, string>;
}

export interface ScoringGuidelines {
  version: string;
  rubric: ScoringRubric[];
  grading_scale: GradingScale[];
}

export interface ScoringRubric {
  criteria: string;
  levels: {
    level: string;
    score_range: [number, number];
    description: string;
  }[];
}

export interface GradingScale {
  grade: string;
  min_score: number;
  max_score: number;
  description: string;
}

// ============================================================================
// Achievement Types
// ============================================================================

export interface StudentAchievement {
  student_id: string;
  student_name: string;
  subject_id: string;
  subject_name: string;
  tp_id: string;
  tp_title: string;
  mastery_level: MasteryLevel;
  competency_progress: CompetencyProgress[];
  calculated_at: string;
  // Convenience properties matching component expectations
  total_achievements?: number;
  average_mastery?: number;
  achievements?: any[];
}

export type MasteryLevel = 'EMERGING' | 'DEVELOPING' | 'PROFICIENT' | 'ADVANCED';

export interface CompetencyProgress {
  competency_id: string;
  competency_name: string;
  current_level: MasteryLevel;
  target_level: MasteryLevel;
  progress_percentage: number;
  evidence_count: number;
  last_assessment_date: string;
}

export interface ClassAchievement {
  class_id: string;
  class_name: string;
  subject_id: string;
  subject_name: string;
  total_students: number;
  mastery_distribution: Record<MasteryLevel, number>;
  average_mastery_level: MasteryLevel;
  average_mastery?: number; // Convenience property matching component expectations
  competency_achievements?: any[]; // Convenience property matching component expectations
  areas_for_improvement?: any[]; // Convenience property matching component expectations
  top_performers?: any[]; // Convenience property matching component expectations
  tp_completion_rate: number;
  calculated_at: string;
  // Convenience properties matching component expectations
  overall_class_score?: number;
  student_achievements?: any[];
}

// ============================================================================
// Evidence Types
// ============================================================================

export interface Evidence {
  id: string;
  student_id: string;
  student_name?: string;
  title?: string;
  assessment_id: string;
  assessment_title?: string;
  evidence_type: EvidenceType;
  file_url: string;
  file_metadata: FileMetadata;
  submission_date: string;
  status: EvidenceStatus;
  evaluation_count: number;
  latest_evaluation_id?: string;
  teacher_notes?: string;
  created_at: string;
  updated_at: string;
}

export type EvidenceType = 'DOCUMENT' | 'IMAGE' | 'VIDEO' | 'AUDIO' | 'PROJECT' | 'PRESENTATION';

export type EvidenceStatus = 'SUBMITTED' | 'UNDER_REVIEW' | 'APPROVED' | 'REJECTED' | 'NEEDS_REVISION';

export interface FileMetadata {
  filename: string;
  file_size: number;
  file_format: string;
  mime_type: string;
  dimensions?: { width: number; height: number };
  duration_seconds?: number;
}

// ============================================================================
// Evaluation Types
// ============================================================================

export interface Evaluation {
  id: string;
  evidence_id: string;
  student_id: string;
  student_name?: string;
  teacher_id: string;
  teacher_name?: string;
  revision_no: number;
  performance_scores: PerformanceScores;
  performance_level: MasteryLevel;
  teacher_feedback: string;
  evaluation_date: string;
  is_current_revision: boolean;
  parent_revision_id?: string;
  created_at: string;
  updated_at: string;
}

export interface PerformanceScores {
  total_score: number;
  max_score: number;
  percentage: number;
  criteria_scores: Record<string, number>;
}

export interface EvaluationFeedbackHistory {
  id: string;
  evaluation_id: string;
  revision_no: number;
  teacher_feedback: string;
  feedback_date: string;
  teacher_id: string;
  teacher_name?: string;
}

// ============================================================================
// ATP (Alur Tujuan Pembelajaran) Types
// ============================================================================

export interface ATP {
  id: string;
  atp_set_id: string;
  sequence_number: number;
  tp_id: string;
  tp_title?: string;
  week: number;
  week_number: number; // Convenience property matching component expectations
  learning_activities: LearningActivities;
  assessment_methods: string[];
  time_allocation: TimeAllocation;
  estimated_hours?: number; // Convenience property matching component expectations
  status: TPStatus;
  created_at: string;
  updated_at: string;
}

export interface LearningActivities {
  opening: string[];
  core_activities: string[];
  closing: string[];
}

export interface ATPSet {
  id: string;
  tp_set_id: string;
  version_no: number;
  status: TPStatus;
  generation_source: 'MANUAL' | 'AI_GENERATED';
  generated_by: string;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
  // Convenience properties matching component expectations
  subject_id?: string;
  phase_id?: string;
  grade?: string;
  semester?: string;
  created_by?: string;
}

// ============================================================================
// Modul Ajar Types
// ============================================================================

export interface ModulAjar {
  id: string;
  modul_ajar_set_id: string;
  atp_id: string;
  atp_title?: string;
  week: number;
  session_number: number;
  learning_objectives: string[];
  teaching_materials: TeachingMaterials;
  learning_methods: string[];
  teaching_methods?: string[]; // Convenience property matching component expectations
  assessment_methods: string[];
  time_allocation: TimeAllocation;
  status: TPStatus;
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  title?: string;
  sequence_number?: number;
  learning_activities?: LearningActivities;
  learning_media?: string[];
  learning_resources?: string[];
  tp_id?: string; // Convenience property matching component expectations
}

export interface TeachingMaterials {
  resources: string[];
  media: string[];
  references: string[];
  core_materials: string[]; // Convenience property matching component expectations
  supporting_materials: string[]; // Convenience property matching component expectations
  digital_resources: string[]; // Convenience property matching component expectations
}

export interface ModulAjarSet {
  id: string;
  atp_set_id: string;
  version_no: number;
  status: TPStatus;
  generation_source: 'MANUAL' | 'AI_GENERATED';
  generated_by: string;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
}

// ============================================================================
// Rubric Types
// ============================================================================

export interface Rubric {
  id: string;
  assessment_id: string;
  assessment_title?: string;
  rubric_type: RubricType;
  criteria: RubricCriteria[];
  total_points: number;
  status: AssessmentStatus;
  version_no: number;
  is_current_version: boolean;
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  title?: string;
  description?: string;
}

export type RubricType = 'ANALYTIC' | 'HOLISTIC';

export interface RubricCriteria {
  id: string;
  name: string;
  description: string;
  weight: number;
  levels: RubricLevel[];
}

export interface RubricLevel {
  level: string;
  description: string;
  points: number;
}

// ============================================================================
// Narrative Report Types
// ============================================================================

export interface NarrativeReport {
  id: string;
  student_id: string;
  student_name: string;
  class_id: string;
  class_name: string;
  subject_id: string;
  subject_name: string;
  reporting_period: ReportingPeriod;
  narrative_content: NarrativeContent;
  achievement_summary: StudentAchievement;
  teacher_recommendations: string[];
  parent_feedback?: string;
  status: AssessmentStatus;
  generated_at: string;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  title?: string;
  period?: string;
  period_id?: string;
  content?: NarrativeContent;
  created_by?: string;
  published_by?: string;
  published_at?: string;
}

export interface ReportingPeriod {
  semester: string;
  academic_year: string;
  start_date: string;
  end_date: string;
}

export interface NarrativeContent {
  introduction: string;
  academic_progress: string;
  behavioral_observations: string;
  strengths: string[];
  areas_for_improvement: string[];
  conclusion: string;
}

// ============================================================================
// Curriculum (CP) Types
// ============================================================================

export interface CP {
  id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  cp_code: string;
  cp_text: string;
  phase: string;
  element: string;
  subelement: string;
  status: 'ACTIVE' | 'ARCHIVED';
  version: string;
  effective_date: string;
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  description?: string;
  competency_code?: string;
  learning_objectives?: any;
  competency_standards?: any;
  time_allocation_hours?: number;
  hours_per_week?: number;
  code?: string; // Convenience for cp_code
}

export interface CurriculumSubject {
  id: string;
  name: string;
  code: string;
  level: string;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  description?: string;
  is_active?: boolean;
}

export interface CurriculumPhase {
  id: string;
  subject_id: string;
  name: string;
  grade_level: string;
  order: number;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  code?: string;
  description?: string;
  grade_level_start?: number;
  grade_level_end?: number;
  grade_range?: string;
  level?: string;
  is_active?: boolean;
}

export interface CurriculumElement {
  id: string;
  phase_id: string;
  name: string;
  code: string;
  description: string;
  order: number;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  subject_id?: string;
  is_active?: boolean;
}

export interface CurriculumSubelement {
  id: string;
  element_id: string;
  name: string;
  code: string;
  description: string;
  order: number;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
  // Convenience properties matching component expectations
  is_active?: boolean;
}

// ============================================================================
// Academic Foundation Types
// ============================================================================

export interface AcademicYear {
  id: string;
  year: string;
  start_date: string;
  end_date: string;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
}

export interface Semester {
  id: string;
  academic_year_id: string;
  name: string;
  start_date: string;
  end_date: string;
  order: number;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
}

export interface SubjectCategory {
  id: string;
  name: string;
  code: string;
  description: string;
  status: 'ACTIVE' | 'ARCHIVED';
  created_at: string;
  updated_at: string;
}

// ============================================================================
// Common Types
// ============================================================================

export interface PaginationParams {
  limit?: number;
  offset?: number;
  page?: number;
  per_page?: number;
}

export interface FilterParams {
  status?: string;
  subject_id?: string;
  phase_id?: string;
  user_id?: string;
  date_from?: string;
  date_to?: string;
}

export interface ApiErrorResponse {
  error: string;
  message?: string;
  code?: string;
  status?: number;
}

export interface ApiResponse<T> {
  data: T;
  message?: string;
  meta?: {
    total: number;
    page: number;
    per_page: number;
    total_pages: number;
  };
}

// ============================================================================
// Request Types for Create/Update Operations
// ============================================================================

export interface CreateTPRequest {
  tp_set_id: string;
  sequence_number: number;
  cp_id: string;
  subject_id: string;
  phase_id: string;
  element_id: string;
  subelement_id: string;
  title: string;
  learning_objectives: LearningObjectives;
  time_allocation: TimeAllocation;
  prerequisites: Prerequisites;
  estimated_weeks: number;
  success_criteria: KKTPCriteria;
}

export interface UpdateTPRequest {
  title?: string;
  learning_objectives?: LearningObjectives;
  time_allocation?: TimeAllocation;
  prerequisites?: Prerequisites;
  estimated_weeks?: number;
  status?: TPStatus;
  success_criteria?: KKTPCriteria;
}

export interface CreateTPSetRequest {
  cp_id: string;
  subject_id: string;
  phase_id: string;
  generation_source: 'MANUAL' | 'AI_GENERATED';
  generation_reason?: string;
}

export interface CreateAssessmentRequest {
  tp_id: string;
  tp_version_no: number;
  success_criteria_snapshot: KKTPCriteria;
  assessment_type: AssessmentType;
  assessment_items: AssessmentItems;
  answer_key: AnswerKey;
  scoring_guidelines: ScoringGuidelines;
}

export interface UpdateAssessmentRequest {
  assessment_type?: AssessmentType;
  assessment_items?: AssessmentItems;
  answer_key?: AnswerKey;
  scoring_guidelines?: ScoringGuidelines;
  status?: AssessmentStatus;
}

export interface CreateEvidenceRequest {
  student_id: string;
  assessment_id: string;
  evidence_type: EvidenceType;
  file_url: string;
  file_metadata: FileMetadata;
}

export interface CreateEvaluationRequest {
  evidence_id: string;
  performance_scores: PerformanceScores;
  performance_level: MasteryLevel;
  teacher_feedback: string;
}

export interface EvaluationUpdateRequest {
  performance_scores?: PerformanceScores;
  performance_level?: MasteryLevel;
  teacher_feedback?: string;
}

export interface EvidenceUpdateRequest {
  evidence_type?: EvidenceType;
  file_url?: string;
  file_metadata?: FileMetadata;
  status?: EvidenceStatus;
  teacher_notes?: string;
}