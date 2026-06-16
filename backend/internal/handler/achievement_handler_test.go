package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
)

// MockAchievementService is a mock for testing
type MockAchievementService struct {
	calculateStudentAchievementFunc func(ctx context.Context, studentID string, studentName string, tpID string) (*domain.Achievement, error)
	calculateCompetencyProgressFunc func(ctx context.Context, studentID string, studentName string, subjectID string, subjectName string, phaseID string, phaseName string) (*domain.CompetencyProgress, error)
	generateClassAchievementFunc    func(ctx context.Context, classID string, className string, subjectID string, subjectName string) (*domain.ClassAchievement, error)
	generateAchievementSummaryFunc  func(ctx context.Context, studentID string, studentName string, classID string, className string, reportPeriod interface{}) (*domain.AchievementSummary, error)
}

func (m *MockAchievementService) CalculateStudentAchievement(ctx context.Context, studentID string, studentName string, tpID string) (*domain.Achievement, error) {
	if m.calculateStudentAchievementFunc != nil {
		return m.calculateStudentAchievementFunc(ctx, studentID, studentName, tpID)
	}
	return &domain.Achievement{
		StudentID:   studentID,
		StudentName: studentName,
		TPID:        tpID,
		TPTitle:     "Test TP Title",
	}, nil
}

func (m *MockAchievementService) CalculateCompetencyProgress(ctx context.Context, studentID string, studentName string, subjectID string, subjectName string, phaseID string, phaseName string) (*domain.CompetencyProgress, error) {
	if m.calculateCompetencyProgressFunc != nil {
		return m.calculateCompetencyProgressFunc(ctx, studentID, studentName, subjectID, subjectName, phaseID, phaseName)
	}
	return &domain.CompetencyProgress{
		StudentID:   studentID,
		StudentName: studentName,
		SubjectID:   subjectID,
		SubjectName: subjectName,
		PhaseID:     phaseID,
		PhaseName:   phaseName,
	}, nil
}

func (m *MockAchievementService) GenerateClassAchievement(ctx context.Context, classID string, className string, subjectID string, subjectName string) (*domain.ClassAchievement, error) {
	if m.generateClassAchievementFunc != nil {
		return m.generateClassAchievementFunc(ctx, classID, className, subjectID, subjectName)
	}
	return &domain.ClassAchievement{
		ClassID:     classID,
		ClassName:   className,
		SubjectID:   subjectID,
		SubjectName: subjectName,
	}, nil
}

func (m *MockAchievementService) GenerateAchievementSummary(ctx context.Context, studentID string, studentName string, classID string, className string, reportPeriod interface{}) (*domain.AchievementSummary, error) {
	if m.generateAchievementSummaryFunc != nil {
		return m.generateAchievementSummaryFunc(ctx, studentID, studentName, classID, className, reportPeriod)
	}
	return &domain.AchievementSummary{
		StudentID:   studentID,
		StudentName: studentName,
		ClassID:     classID,
		ClassName:   className,
	}, nil
}

func setupAchievementTestRouter(service IAchievementService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := NewAchievementHandler(service)

	router.GET("/students/:id/achievement", handler.GetStudentAchievement)
	router.GET("/students/:id/progress", handler.GetStudentProgress)
	router.GET("/classes/:id/achievement", handler.GetClassAchievement)
	router.GET("/reports/:id/achievement-summary", handler.GetReportAchievementSummary)

	return router
}

func TestAchievementHandler_GetStudentAchievement_Success(t *testing.T) {
	mockService := &MockAchievementService{
		calculateStudentAchievementFunc: func(ctx context.Context, studentID string, studentName string, tpID string) (*domain.Achievement, error) {
			return &domain.Achievement{
				StudentID:   studentID,
				StudentName: studentName,
				TPID:        tpID,
				TPTitle:     "Learning Objective 1",
			}, nil
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/achievement?tp_id=tp-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentAchievement_MissingTPID(t *testing.T) {
	mockService := &MockAchievementService{}
	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/achievement", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentAchievement_ServiceError(t *testing.T) {
	mockService := &MockAchievementService{
		calculateStudentAchievementFunc: func(ctx context.Context, studentID string, studentName string, tpID string) (*domain.Achievement, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/achievement?tp_id=tp-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentProgress_Success(t *testing.T) {
	mockService := &MockAchievementService{
		calculateCompetencyProgressFunc: func(ctx context.Context, studentID string, studentName string, subjectID string, subjectName string, phaseID string, phaseName string) (*domain.CompetencyProgress, error) {
			return &domain.CompetencyProgress{
				StudentID:   studentID,
				StudentName: studentName,
				SubjectID:   subjectID,
				SubjectName: subjectName,
				PhaseID:     phaseID,
				PhaseName:   phaseName,
			}, nil
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/progress?subject_id=subject-123&phase_id=phase-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentProgress_MissingSubjectID(t *testing.T) {
	mockService := &MockAchievementService{}
	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/progress", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentProgress_WithOptionalPhaseID(t *testing.T) {
	mockService := &MockAchievementService{
		calculateCompetencyProgressFunc: func(ctx context.Context, studentID string, studentName string, subjectID string, subjectName string, phaseID string, phaseName string) (*domain.CompetencyProgress, error) {
			return &domain.CompetencyProgress{
				StudentID:   studentID,
				StudentName: studentName,
				SubjectID:   subjectID,
				SubjectName: subjectName,
				PhaseID:     phaseID,
				PhaseName:   phaseName,
			}, nil
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/progress?subject_id=subject-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAchievementHandler_GetStudentProgress_ServiceError(t *testing.T) {
	mockService := &MockAchievementService{
		calculateCompetencyProgressFunc: func(ctx context.Context, studentID string, studentName string, subjectID string, subjectName string, phaseID string, phaseName string) (*domain.CompetencyProgress, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/students/student-123/progress?subject_id=subject-123&phase_id=phase-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestAchievementHandler_GetClassAchievement_Success(t *testing.T) {
	mockService := &MockAchievementService{
		generateClassAchievementFunc: func(ctx context.Context, classID string, className string, subjectID string, subjectName string) (*domain.ClassAchievement, error) {
			return &domain.ClassAchievement{
				ClassID:     classID,
				ClassName:   className,
				SubjectID:   subjectID,
				SubjectName: subjectName,
			}, nil
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/classes/class-123/achievement?subject_id=subject-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAchievementHandler_GetClassAchievement_MissingSubjectID(t *testing.T) {
	mockService := &MockAchievementService{}
	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/classes/class-123/achievement", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAchievementHandler_GetClassAchievement_ServiceError(t *testing.T) {
	mockService := &MockAchievementService{
		generateClassAchievementFunc: func(ctx context.Context, classID string, className string, subjectID string, subjectName string) (*domain.ClassAchievement, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/classes/class-123/achievement?subject_id=subject-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestAchievementHandler_GetReportAchievementSummary_Success(t *testing.T) {
	mockService := &MockAchievementService{
		generateAchievementSummaryFunc: func(ctx context.Context, studentID string, studentName string, classID string, className string, reportPeriod interface{}) (*domain.AchievementSummary, error) {
			return &domain.AchievementSummary{
				StudentID:   studentID,
				StudentName: studentName,
				ClassID:     classID,
				ClassName:   className,
			}, nil
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/reports/report-123/achievement-summary?student_id=student-123&class_id=class-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAchievementHandler_GetReportAchievementSummary_MissingStudentID(t *testing.T) {
	mockService := &MockAchievementService{}
	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/reports/report-123/achievement-summary?class_id=class-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAchievementHandler_GetReportAchievementSummary_MissingClassID(t *testing.T) {
	mockService := &MockAchievementService{}
	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/reports/report-123/achievement-summary?student_id=student-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAchievementHandler_GetReportAchievementSummary_ServiceError(t *testing.T) {
	mockService := &MockAchievementService{
		generateAchievementSummaryFunc: func(ctx context.Context, studentID string, studentName string, classID string, className string, reportPeriod interface{}) (*domain.AchievementSummary, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAchievementTestRouter(mockService)

	req := httptest.NewRequest("GET", "/reports/report-123/achievement-summary?student_id=student-123&class_id=class-123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}
