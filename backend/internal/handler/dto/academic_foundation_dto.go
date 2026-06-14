package dto

import "time"

// ==================== Academic Year DTOs ====================

// CreateAcademicYearRequest represents the request body for creating an academic year
type CreateAcademicYearRequest struct {
	SchoolID    string    `json:"school_id" binding:"required"`
	Name        string    `json:"name" binding:"required,min=1,max=100"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Description string    `json:"description,omitempty,max=500"`
}

// AcademicYearResponse represents the response body for academic year operations
type AcademicYearResponse struct {
	ID        string `json:"id"`
	SchoolID  string `json:"school_id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UpdateAcademicYearRequest represents the request body for updating an academic year
type UpdateAcademicYearRequest struct {
	Name      *string    `json:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// AcademicYearWithSemestersResponse represents an academic year with its semesters
type AcademicYearWithSemestersResponse struct {
	ID        string             `json:"id"`
	SchoolID  string             `json:"school_id"`
	Name      string             `json:"name"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Status    string             `json:"status"`
	CreatedBy string             `json:"created_by"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Semesters []SemesterResponse `json:"semesters"`
}

// ListAcademicYearsResponse represents the response body for listing academic years
type ListAcademicYearsResponse struct {
	AcademicYears []AcademicYearResponse `json:"academic_years"`
	Total         int                    `json:"total"`
}

// ==================== Semester DTOs ====================

// CreateSemesterRequest represents the request body for creating a semester
type CreateSemesterRequest struct {
	AcademicYearID string    `json:"academic_year_id" binding:"required"`
	Type           string    `json:"type" binding:"required,oneof=GANJIL GENAP"`
	Name           string    `json:"name" binding:"required,min=1,max=100"`
	StartDate      time.Time `json:"start_date" binding:"required"`
	EndDate        time.Time `json:"end_date" binding:"required"`
	SequenceNumber int       `json:"sequence_number" binding:"required,min=1,max=2"`
}

// SemesterResponse represents the response body for semester operations
type SemesterResponse struct {
	ID             string `json:"id"`
	AcademicYearID string `json:"academic_year_id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Status         string `json:"status"`
	SequenceNumber int    `json:"sequence_number"`
	CreatedBy      string `json:"created_by"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// UpdateSemesterRequest represents the request body for updating a semester
type UpdateSemesterRequest struct {
	Name      *string    `json:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Status    *string    `json:"status,omitempty" binding:"omitempty,oneof=ACTIVE INACTIVE"`
}

// ListSemestersResponse represents the response body for listing semesters
type ListSemestersResponse struct {
	Semesters []SemesterResponse `json:"semesters"`
	Total     int                `json:"total"`
}

// ==================== Subject Category DTOs ====================

// CreateSubjectCategoryRequest represents the request body for creating a subject category
type CreateSubjectCategoryRequest struct {
	Code        string  `json:"code" binding:"required,min=1,max=20"`
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	IsMandatory bool    `json:"is_mandatory"`
}

// SubjectCategoryResponse represents the response body for subject category operations
type SubjectCategoryResponse struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	IsMandatory bool    `json:"is_mandatory"`
	IsActive    bool    `json:"is_active"`
	CreatedBy   string  `json:"created_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UpdateSubjectCategoryRequest represents the request body for updating a subject category
type UpdateSubjectCategoryRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,max=100"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	IsMandatory *bool   `json:"is_mandatory,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// ListSubjectCategoriesResponse represents the response body for listing subject categories
type ListSubjectCategoriesResponse struct {
	SubjectCategories []SubjectCategoryResponse `json:"subject_categories"`
	Total             int                       `json:"total"`
}

// ==================== Graduate Profile Dimension DTOs ====================

// CreateGraduateProfileDimensionRequest represents the request body for creating a graduate profile dimension
type CreateGraduateProfileDimensionRequest struct {
	Code           string  `json:"code" binding:"required,min=1,max=20"`
	Name           string  `json:"name" binding:"required,min=1,max=100"`
	Description    *string `json:"description,omitempty" binding:"omitempty,max=1000"`
	SequenceNumber int     `json:"sequence_number" binding:"required,min=1,max=6"`
}

// GraduateProfileDimensionResponse represents the response body for graduate profile dimension operations
type GraduateProfileDimensionResponse struct {
	ID             string  `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	SequenceNumber int     `json:"sequence_number"`
	IsActive       bool    `json:"is_active"`
	CreatedBy      string  `json:"created_by"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// UpdateGraduateProfileDimensionRequest represents the request body for updating a graduate profile dimension
type UpdateGraduateProfileDimensionRequest struct {
	Name           *string `json:"name,omitempty" binding:"omitempty,max=100"`
	Description    *string `json:"description,omitempty" binding:"omitempty,max=1000"`
	SequenceNumber *int    `json:"sequence_number,omitempty" binding:"omitempty,min=1,max=6"`
	IsActive       *bool   `json:"is_active,omitempty"`
}

// ListGraduateProfileDimensionsResponse represents the response body for listing graduate profile dimensions
type ListGraduateProfileDimensionsResponse struct {
	GraduateProfileDimensions []GraduateProfileDimensionResponse `json:"graduate_profile_dimensions"`
	Total                     int                                `json:"total"`
}

// ==================== CP Alignment DTOs ====================

// CreateCPAlignmentRequest represents the request body for creating a CP alignment
type CreateCPAlignmentRequest struct {
	CurriculumSubjectID        string  `json:"curriculum_subject_id" binding:"required"`
	GraduateProfileDimensionID string  `json:"graduate_profile_dimension_id" binding:"required"`
	AlignmentDescription       *string `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
}

// CPAlignmentResponse represents the response body for CP alignment operations
type CPAlignmentResponse struct {
	ID                         string  `json:"id"`
	CurriculumSubjectID        string  `json:"curriculum_subject_id"`
	GraduateProfileDimensionID string  `json:"graduate_profile_dimension_id"`
	AlignmentDescription       *string `json:"alignment_description,omitempty"`
	IsActive                   bool    `json:"is_active"`
	CreatedBy                  string  `json:"created_by"`
	CreatedAt                  string  `json:"created_at"`
	UpdatedAt                  string  `json:"updated_at"`
}

// CreateCPAlignmentBulkRequest represents the request body for creating multiple CP alignments
type CreateCPAlignmentBulkRequest struct {
	CurriculumSubjectID  string   `json:"curriculum_subject_id" binding:"required"`
	AlignmentIDs         []string `json:"alignment_ids" binding:"required"`
	AlignmentDescription *string  `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
}

// UpdateCPAlignmentRequest represents the request body for updating a CP alignment
type UpdateCPAlignmentRequest struct {
	AlignmentDescription *string `json:"alignment_description,omitempty" binding:"omitempty,max=500"`
	IsActive             *bool   `json:"is_active,omitempty"`
}

// ListCPAlignmentsResponse represents the response body for listing CP alignments
type ListCPAlignmentsResponse struct {
	CPAlignments []CPAlignmentResponse `json:"cp_alignments"`
	Total        int                   `json:"total"`
}

// CPAlignmentReportResponse represents the CP alignment report response
type CPAlignmentReportResponse struct {
	GraduateProfileDimensionID   string  `json:"graduate_profile_dimension_id"`
	GraduateProfileDimensionName string  `json:"graduate_profile_dimension_name"`
	TotalCPCount                 int     `json:"total_cp_count"`
	AlignedCPCount               int     `json:"aligned_cp_count"`
	CoveragePercentage           float64 `json:"coverage_percentage"`
	MeetsThreshold               bool    `json:"meets_threshold"`
}

// GenerateCPAlignmentReportResponse represents the response for generating CP alignment report
type GenerateCPAlignmentReportResponse struct {
	Reports []CPAlignmentReportResponse `json:"reports"`
}

// ==================== System Configuration DTOs ====================

// CreateSystemConfigurationRequest represents the request body for creating a system configuration
type CreateSystemConfigurationRequest struct {
	Key         string  `json:"key" binding:"required,min=1,max=100"`
	Value       string  `json:"value" binding:"required"`
	ValueType   string  `json:"value_type" binding:"required,oneof=string number boolean json"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	Category    string  `json:"category" binding:"required,min=1,max=50"`
	IsSystem    bool    `json:"is_system"`
}

// SystemConfigurationResponse represents the response body for system configuration operations
type SystemConfigurationResponse struct {
	ID          string  `json:"id"`
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	ValueType   string  `json:"value_type"`
	Description *string `json:"description,omitempty"`
	Category    string  `json:"category"`
	IsSystem    bool    `json:"is_system"`
	IsActive    bool    `json:"is_active"`
	CreatedBy   string  `json:"created_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UpdateSystemConfigurationRequest represents the request body for updating a system configuration
type UpdateSystemConfigurationRequest struct {
	Value       *string `json:"value,omitempty"`
	ValueType   *string `json:"value_type,omitempty" binding:"omitempty,oneof=string number boolean json"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	Category    *string `json:"category,omitempty" binding:"omitempty,max=50"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// ListSystemConfigurationsResponse represents the response body for listing system configurations
type ListSystemConfigurationsResponse struct {
	SystemConfigurations []SystemConfigurationResponse `json:"system_configurations"`
	Total                int                           `json:"total"`
}
