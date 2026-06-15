package domain

import (
	"testing"
	"time"
)

// TestClass_Validate tests validation of class entities
func TestClass_Validate(t *testing.T) {
	tests := []struct {
		name                  string
		class                 Class
		wantError             bool
		expectedErrorContains string
	}{
		{
			name: "Valid class",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid class with room",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				Room:           makeStringPtr("Room 101"),
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid class - minimum students (1)",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    1,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid class - maximum students (100)",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    100,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid class - inactive",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       false,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid class - with deleted_at",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				DeletedAt:      func() *time.Time { t := time.Now(); return &t }(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Invalid - empty ID",
			class: Class{
				ID:             "",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required",
		},
		{
			name: "Invalid - empty school ID",
			class: Class{
				ID:             "class-1",
				SchoolID:       "",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "school_id is required",
		},
		{
			name: "Invalid - empty academic year ID",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "academic_year_id is required",
		},
		{
			name: "Invalid - empty semester ID",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "semester_id is required",
		},
		{
			name: "Invalid - empty subject ID",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "subject_id is required",
		},
		{
			name: "Invalid - empty teacher ID",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "teacher_id is required",
		},
		{
			name: "Invalid - empty name",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "name is required",
		},
		{
			name: "Invalid - empty grade level",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "grade_level is required",
		},
		{
			name: "Invalid - max students less than 1",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    0,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "max_students must be between 1 and 100",
		},
		{
			name: "Invalid - max students greater than 100",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    101,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "max_students must be between 1 and 100",
		},
		{
			name: "Invalid - negative max students",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    -5,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "max_students must be between 1 and 100",
		},
		{
			name: "Invalid - all required fields missing",
			class: Class{
				ID:             "",
				SchoolID:       "",
				AcademicYearID: "",
				SemesterID:     "",
				SubjectID:      "",
				TeacherID:      "",
				Name:           "",
				GradeLevel:     "",
				MaxStudents:    0,
				IsActive:       true,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.class.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErrorContains != "" && !contains(err.Error(), tc.expectedErrorContains) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErrorContains, err.Error())
				}
			}
		})
	}
}

// TestClass_ToClassResponse tests DTO transformation for Class
func TestClass_ToClassResponse(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name              string
		class             Class
		schoolName        string
		academicYear      string
		semesterName      string
		subjectName       string
		teacherName       string
		currentStudents   int
		expectedID        string
		expectedHasSchoolName bool
		expectedHasAcademicYear bool
		expectedHasSemesterName bool
		expectedHasSubjectName bool
		expectedHasTeacherName bool
		expectedHasRoom      bool
	}{
		{
			name: "Full transformation with all data",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				Room:           makeStringPtr("Room 101"),
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "Sekolah A",
			academicYear:     "2026-2027",
			semesterName:     "Ganjil 1",
			subjectName:      "Mathematics",
			teacherName:       "John Doe",
			currentStudents:  25,
			expectedID:       "class-1",
			expectedHasSchoolName: true,
			expectedHasAcademicYear: true,
			expectedHasSemesterName: true,
			expectedHasSubjectName: true,
			expectedHasTeacherName: true,
			expectedHasRoom: true,
		},
		{
			name: "Transformation with empty optional fields",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "",
			academicYear:     "",
			semesterName:     "",
			subjectName:      "",
			teacherName:       "",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: false,
			expectedHasAcademicYear: false,
			expectedHasSemesterName: false,
			expectedHasSubjectName: false,
			expectedHasTeacherName: false,
			expectedHasRoom: false,
		},
		{
			name: "Transformation with only school name",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "Sekolah A",
			academicYear:     "",
			semesterName:     "",
			subjectName:      "",
			teacherName:       "",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: true,
			expectedHasAcademicYear: false,
			expectedHasSemesterName: false,
			expectedHasSubjectName: false,
			expectedHasTeacherName: false,
			expectedHasRoom: false,
		},
		{
			name: "Transformation with only academic year",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "",
			academicYear:     "2026-2027",
			semesterName:     "",
			subjectName:      "",
			teacherName:       "",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: false,
			expectedHasAcademicYear: true,
			expectedHasSemesterName: false,
			expectedHasSubjectName: false,
			expectedHasTeacherName: false,
			expectedHasRoom: false,
		},
		{
			name: "Transformation with inactive class",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				MaxStudents:    30,
				IsActive:       false,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "",
			academicYear:     "",
			semesterName:     "",
			subjectName:      "",
			teacherName:       "",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: false,
			expectedHasAcademicYear: false,
			expectedHasSemesterName: false,
			expectedHasSubjectName: false,
			expectedHasTeacherName: false,
			expectedHasRoom: false,
		},
		{
			name: "Transformation with no room",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				Room:           nil,
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "Sekolah A",
			academicYear:     "2026-2027",
			semesterName:     "Ganjil 1",
			subjectName:      "Mathematics",
			teacherName:       "John Doe",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: true,
			expectedHasAcademicYear: true,
			expectedHasSemesterName: true,
			expectedHasSubjectName: true,
			expectedHasTeacherName: true,
			expectedHasRoom: false,
		},
		{
			name: "Transformation with empty room string",
			class: Class{
				ID:             "class-1",
				SchoolID:       "school-1",
				AcademicYearID: "academic-year-1",
				SemesterID:     "semester-1",
				SubjectID:      "subject-1",
				TeacherID:      "teacher-1",
				Name:           "Class A",
				GradeLevel:     "10",
				Room:           makeStringPtr(""),
				MaxStudents:    30,
				IsActive:       true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			schoolName:       "",
			academicYear:     "",
			semesterName:     "",
			subjectName:      "",
			teacherName:       "",
			currentStudents: 25,
			expectedID:       "class-1",
			expectedHasSchoolName: false,
			expectedHasAcademicYear: false,
			expectedHasSemesterName: false,
			expectedHasSubjectName: false,
			expectedHasTeacherName: false,
			expectedHasRoom: true, // Empty string is still a string
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := tc.class.ToClassResponse(tc.schoolName, tc.academicYear, tc.semesterName, tc.subjectName, tc.teacherName, tc.currentStudents)

			if response == nil {
				t.Fatal("expected response, got nil")
			}

			if response.ID != tc.expectedID {
				t.Errorf("expected ID %s, got %s", tc.expectedID, response.ID)
			}
			if response.SchoolID != tc.class.SchoolID {
				t.Errorf("expected school ID %s, got %s", tc.class.SchoolID, response.SchoolID)
			}
			if response.AcademicYearID != tc.class.AcademicYearID {
				t.Errorf("expected academic year ID %s, got %s", tc.class.AcademicYearID, response.AcademicYearID)
			}
			if response.SemesterID != tc.class.SemesterID {
				t.Errorf("expected semester ID %s, got %s", tc.class.SemesterID, response.SemesterID)
			}
			if response.SubjectID != tc.class.SubjectID {
				t.Errorf("expected subject ID %s, got %s", tc.class.SubjectID, response.SubjectID)
			}
			if response.TeacherID != tc.class.TeacherID {
				t.Errorf("expected teacher ID %s, got %s", tc.class.TeacherID, response.TeacherID)
			}
			if response.Name != tc.class.Name {
				t.Errorf("expected name %s, got %s", tc.class.Name, response.Name)
			}
			if response.GradeLevel != tc.class.GradeLevel {
				t.Errorf("expected grade level %s, got %s", tc.class.GradeLevel, response.GradeLevel)
			}
			if response.MaxStudents != tc.class.MaxStudents {
				t.Errorf("expected max students %d, got %d", tc.class.MaxStudents, response.MaxStudents)
			}
			if response.IsActive != tc.class.IsActive {
				t.Errorf("expected is_active %v, got %v", tc.class.IsActive, response.IsActive)
			}
			if response.CurrentStudents != tc.currentStudents {
				t.Errorf("expected current students %d, got %d", tc.currentStudents, response.CurrentStudents)
			}
			if !response.CreatedAt.Equal(tc.class.CreatedAt) {
				t.Error("CreatedAt should be preserved")
			}
			if !response.UpdatedAt.Equal(tc.class.UpdatedAt) {
				t.Error("UpdatedAt should be preserved")
			}

			// Check optional fields
			if tc.expectedHasSchoolName && response.SchoolName == nil {
				t.Error("expected SchoolName to be set")
			}
			if !tc.expectedHasSchoolName && response.SchoolName != nil {
				t.Error("expected SchoolName to be nil")
			}
			if tc.expectedHasAcademicYear && response.AcademicYear == nil {
				t.Error("expected AcademicYear to be set")
			}
			if !tc.expectedHasAcademicYear && response.AcademicYear != nil {
				t.Error("expected AcademicYear to be nil")
			}
			if tc.expectedHasSemesterName && response.SemesterName == nil {
				t.Error("expected SemesterName to be set")
			}
			if !tc.expectedHasSemesterName && response.SemesterName != nil {
				t.Error("expected SemesterName to be nil")
			}
			if tc.expectedHasSubjectName && response.SubjectName == nil {
				t.Error("expected SubjectName to be set")
			}
			if !tc.expectedHasSubjectName && response.SubjectName != nil {
				t.Error("expected SubjectName to be nil")
			}
			if tc.expectedHasTeacherName && response.TeacherName == nil {
				t.Error("expected TeacherName to be set")
			}
			if !tc.expectedHasTeacherName && response.TeacherName != nil {
				t.Error("expected TeacherName to be nil")
			}
			if tc.expectedHasRoom && response.Room == nil {
				t.Error("expected Room to be set")
			}
			if !tc.expectedHasRoom && response.Room != nil {
				t.Error("expected Room to be nil")
			}
		})
	}
}

// TestClassEnrollment_Validate tests validation of class enrollment entities
func TestClassEnrollment_Validate(t *testing.T) {
	tests := []struct {
		name                  string
		enrollment           ClassEnrollment
		wantError             bool
		expectedErrorContains string
	}{
		{
			name: "Valid enrollment",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid enrollment - inactive status",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusInactive),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid enrollment - withdrawn status",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusWithdrawn),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid enrollment - completed status",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusCompleted),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid enrollment - with notes",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          makeStringPtr("Student enrolled successfully"),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false,
		},
		{
			name: "Valid enrollment - with zero date",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				EnrollmentDate: time.Time{},
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError: false, // Validation doesn't check date
		},
		{
			name: "Invalid - empty ID",
			enrollment: ClassEnrollment{
				ID:             "",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required",
		},
		{
			name: "Invalid - empty class ID",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "class_id is required",
		},
		{
			name: "Invalid - empty student ID",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "",
				Status:         string(EnrollmentStatusActive),
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "student_id is required",
		},
		{
			name: "Invalid - empty status",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         "",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "status is required",
		},
		{
			name: "Invalid - all required fields missing",
			enrollment: ClassEnrollment{
				ID:             "",
				ClassID:        "",
				StudentID:      "",
				Status:         "",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			wantError:             true,
			expectedErrorContains: "id is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.enrollment.Validate()

			if tc.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.wantError && err != nil {
				if tc.expectedErrorContains != "" && !contains(err.Error(), tc.expectedErrorContains) {
					t.Errorf("expected error to contain %s, got %s", tc.expectedErrorContains, err.Error())
				}
			}
		})
	}
}

// TestClassEnrollment_ToClassEnrollmentResponse tests DTO transformation for ClassEnrollment
func TestClassEnrollment_ToClassEnrollmentResponse(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name                string
		enrollment          ClassEnrollment
		className           string
		studentName         string
		expectedID          string
		expectedHasClassName  bool
		expectedHasStudentName bool
		expectedHasNotes      bool
	}{
		{
			name: "Full transformation with all data",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          makeStringPtr("Student enrolled successfully"),
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "Class A",
			studentName:         "John Doe",
			expectedID:          "enrollment-1",
			expectedHasClassName:  true,
			expectedHasStudentName: true,
			expectedHasNotes: true,
		},
		{
			name: "Transformation with empty optional fields",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          nil,
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "",
			studentName:         "",
			expectedID:          "enrollment-1",
			expectedHasClassName:  false,
			expectedHasStudentName: false,
			expectedHasNotes: false,
		},
		{
			name: "Transformation with only class name",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          nil,
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "Class A",
			studentName:         "",
			expectedID:          "enrollment-1",
			expectedHasClassName:  true,
			expectedHasStudentName: false,
			expectedHasNotes: false,
		},
		{
			name: "Transformation with only student name",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          nil,
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "",
			studentName:         "John Doe",
			expectedID:          "enrollment-1",
			expectedHasClassName:  false,
			expectedHasStudentName: true,
			expectedHasNotes: false,
		},
		{
			name: "Transformation - inactive status",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusInactive),
				Notes:          nil,
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "",
			studentName:         "",
			expectedID:          "enrollment-1",
			expectedHasClassName:  false,
			expectedHasStudentName: false,
			expectedHasNotes: false,
		},
		{
			name: "Transformation with empty notes string",
			enrollment: ClassEnrollment{
				ID:             "enrollment-1",
				ClassID:        "class-1",
				StudentID:      "student-1",
				Status:         string(EnrollmentStatusActive),
				Notes:          makeStringPtr(""),
				EnrollmentDate: now,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			className:           "",
			studentName:         "",
			expectedID:          "enrollment-1",
			expectedHasClassName:  false,
			expectedHasStudentName: false,
			expectedHasNotes: true, // Empty string is still a string
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			response := tc.enrollment.ToClassEnrollmentResponse(tc.className, tc.studentName)

			if response == nil {
				t.Fatal("expected response, got nil")
			}

			if response.ID != tc.expectedID {
				t.Errorf("expected ID %s, got %s", tc.expectedID, response.ID)
			}
			if response.ClassID != tc.enrollment.ClassID {
				t.Errorf("expected class ID %s, got %s", tc.enrollment.ClassID, response.ClassID)
			}
			if response.StudentID != tc.enrollment.StudentID {
				t.Errorf("expected student ID %s, got %s", tc.enrollment.StudentID, response.StudentID)
			}
			if response.Status != EnrollmentStatus(tc.enrollment.Status) {
				t.Errorf("expected status %s, got %s", EnrollmentStatus(tc.enrollment.Status), response.Status)
			}
			if response.Notes != tc.enrollment.Notes {
				t.Errorf("expected notes %v, got %v", tc.enrollment.Notes, response.Notes)
			}
			if !response.EnrollmentDate.Equal(tc.enrollment.EnrollmentDate) {
				t.Error("EnrollmentDate should be preserved")
			}
			if !response.CreatedAt.Equal(tc.enrollment.CreatedAt) {
				t.Error("CreatedAt should be preserved")
			}
			if !response.UpdatedAt.Equal(tc.enrollment.UpdatedAt) {
				t.Error("UpdatedAt should be preserved")
			}

			// Check optional fields
			if tc.expectedHasClassName && response.ClassName == nil {
				t.Error("expected ClassName to be set")
			}
			if !tc.expectedHasClassName && response.ClassName != nil {
				t.Error("expected ClassName to be nil")
			}
			if tc.expectedHasStudentName && response.StudentName == nil {
				t.Error("expected StudentName to be set")
			}
			if !tc.expectedHasStudentName && response.StudentName != nil {
				t.Error("expected StudentName to be nil")
			}
			if tc.expectedHasNotes && response.Notes == nil {
				t.Error("expected Notes to be set")
			}
			if !tc.expectedHasNotes && response.Notes != nil {
				t.Error("expected Notes to be nil")
			}
		})
	}
}