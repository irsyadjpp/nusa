package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/application"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
)

// MockAcademicYearApplicationService is a mock for testing
type MockAcademicYearApplicationService struct {
	createAcademicYearFunc   func(ctx context.Context, cmd *application.CreateAcademicYearCommand) (*application.CreateAcademicYearResponse, error)
	updateAcademicYearFunc   func(ctx context.Context, cmd *application.UpdateAcademicYearCommand) (*application.UpdateAcademicYearResponse, error)
	activateAcademicYearFunc func(ctx context.Context, cmd *application.ActivateAcademicYearCommand) (*application.ActivateAcademicYearResponse, error)
	archiveAcademicYearFunc  func(ctx context.Context, cmd *application.ArchiveAcademicYearCommand) (*application.ArchiveAcademicYearResponse, error)
	getAcademicYearFunc      func(ctx context.Context, cmd *application.GetAcademicYearCommand) (*application.GetAcademicYearResponse, error)
	listAcademicYearsFunc    func(ctx context.Context, cmd *application.ListAcademicYearsCommand) (*application.ListAcademicYearsResponse, error)
}

func (m *MockAcademicYearApplicationService) CreateAcademicYear(ctx context.Context, cmd *application.CreateAcademicYearCommand) (*application.CreateAcademicYearResponse, error) {
	if m.createAcademicYearFunc != nil {
		return m.createAcademicYearFunc(ctx, cmd)
	}
	return &application.CreateAcademicYearResponse{
		AcademicYearID: "test-id",
		Status:         domain.AcademicYearStatusDraft,
	}, nil
}

func (m *MockAcademicYearApplicationService) UpdateAcademicYear(ctx context.Context, cmd *application.UpdateAcademicYearCommand) (*application.UpdateAcademicYearResponse, error) {
	if m.updateAcademicYearFunc != nil {
		return m.updateAcademicYearFunc(ctx, cmd)
	}
	return &application.UpdateAcademicYearResponse{
		AcademicYearID: cmd.AcademicYearID,
		Status:         domain.AcademicYearStatusActive,
	}, nil
}

func (m *MockAcademicYearApplicationService) ActivateAcademicYear(ctx context.Context, cmd *application.ActivateAcademicYearCommand) (*application.ActivateAcademicYearResponse, error) {
	if m.activateAcademicYearFunc != nil {
		return m.activateAcademicYearFunc(ctx, cmd)
	}
	return &application.ActivateAcademicYearResponse{
		AcademicYearID: cmd.AcademicYearID,
		Status:         domain.AcademicYearStatusActive,
	}, nil
}

func (m *MockAcademicYearApplicationService) ArchiveAcademicYear(ctx context.Context, cmd *application.ArchiveAcademicYearCommand) (*application.ArchiveAcademicYearResponse, error) {
	if m.archiveAcademicYearFunc != nil {
		return m.archiveAcademicYearFunc(ctx, cmd)
	}
	return &application.ArchiveAcademicYearResponse{
		AcademicYearID: cmd.AcademicYearID,
		Status:         domain.AcademicYearStatusArchived,
	}, nil
}

func (m *MockAcademicYearApplicationService) GetAcademicYear(ctx context.Context, cmd *application.GetAcademicYearCommand) (*application.GetAcademicYearResponse, error) {
	if m.getAcademicYearFunc != nil {
		return m.getAcademicYearFunc(ctx, cmd)
	}
	return &application.GetAcademicYearResponse{
		AcademicYear: &domain.AcademicYear{
			ID:     cmd.AcademicYearID,
			Name:   "2024-2025",
			Status: domain.AcademicYearStatusActive,
		},
		Semesters: []*domain.Semester{},
	}, nil
}

func (m *MockAcademicYearApplicationService) ListAcademicYears(ctx context.Context, cmd *application.ListAcademicYearsCommand) (*application.ListAcademicYearsResponse, error) {
	if m.listAcademicYearsFunc != nil {
		return m.listAcademicYearsFunc(ctx, cmd)
	}
	return &application.ListAcademicYearsResponse{
		AcademicYears: []*domain.AcademicYear{},
	}, nil
}

func NewAcademicYearHandlerWithInterface(service IAcademicYearApplicationService) *AcademicYearHandler {
	return &AcademicYearHandler{
		academicYearService: service,
	}
}

func setupAcademicYearTestRouter(service IAcademicYearApplicationService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := NewAcademicYearHandlerWithInterface(service)

	// Add mock auth middleware
	router.Use(func(c *gin.Context) {
		schoolID := "test-school-id"
		authCtx := &middleware.AuthContext{
			UserID:      "test-user-id",
			SchoolID:    &schoolID,
			Role:        "SCHOOL_ADMIN",
			Permissions: []string{"academic_year:CREATE", "academic_year:READ", "academic_year:UPDATE"},
		}
		c.Set(middleware.AuthContextKey, authCtx)
		c.Next()
	})

	router.POST("/academic-years", handler.CreateAcademicYear)
	router.PUT("/academic-years/:id", handler.UpdateAcademicYear)
	router.POST("/academic-years/:id/activate", handler.ActivateAcademicYear)
	router.POST("/academic-years/:id/archive", handler.ArchiveAcademicYear)
	router.GET("/academic-years/:id", handler.GetAcademicYear)
	router.GET("/academic-years", handler.ListAcademicYears)

	return router
}

func TestCreateAcademicYear(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		createAcademicYearFunc: func(ctx context.Context, cmd *application.CreateAcademicYearCommand) (*application.CreateAcademicYearResponse, error) {
			return &application.CreateAcademicYearResponse{
				AcademicYearID: "test-id",
				Status:         domain.AcademicYearStatusDraft,
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	reqBody := dto.CreateAcademicYearRequest{
		SchoolID:    "school-123",
		Name:        "2024-2025",
		StartDate:   startDate,
		EndDate:     endDate,
		Description: "Academic year 2024-2025",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/academic-years", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var resp dto.AcademicYearResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", resp.ID)
	}
}

func TestCreateAcademicYear_Unauthorized(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	handler := NewAcademicYearHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// No auth middleware
	router.POST("/academic-years", handler.CreateAcademicYear)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	reqBody := dto.CreateAcademicYearRequest{
		SchoolID:  "school-123",
		Name:      "2024-2025",
		StartDate: startDate,
		EndDate:   endDate,
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/academic-years", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestCreateAcademicYear_InvalidRequest(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	reqBody := dto.CreateAcademicYearRequest{
		// Missing required fields
		Name: "2024-2025",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/academic-years", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestCreateAcademicYear_ServiceError(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		createAcademicYearFunc: func(ctx context.Context, cmd *application.CreateAcademicYearCommand) (*application.CreateAcademicYearResponse, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	reqBody := dto.CreateAcademicYearRequest{
		SchoolID:  "school-123",
		Name:      "2024-2025",
		StartDate: startDate,
		EndDate:   endDate,
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/academic-years", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestUpdateAcademicYear(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		updateAcademicYearFunc: func(ctx context.Context, cmd *application.UpdateAcademicYearCommand) (*application.UpdateAcademicYearResponse, error) {
			return &application.UpdateAcademicYearResponse{
				AcademicYearID: cmd.AcademicYearID,
				Status:         domain.AcademicYearStatusActive,
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	newName := "2024-2025 Updated"
	reqBody := dto.UpdateAcademicYearRequest{
		Name: &newName,
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/academic-years/academic-year-123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestUpdateAcademicYear_InvalidID(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	newName := "2024-2025 Updated"
	reqBody := dto.UpdateAcademicYearRequest{
		Name: &newName,
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/academic-years/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Router returns 404 for empty ID in path
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestUpdateAcademicYear_Unauthorized(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	handler := NewAcademicYearHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.PUT("/academic-years/:id", handler.UpdateAcademicYear)

	newName := "2024-2025 Updated"
	reqBody := dto.UpdateAcademicYearRequest{
		Name: &newName,
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/academic-years/academic-year-123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestActivateAcademicYear(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		activateAcademicYearFunc: func(ctx context.Context, cmd *application.ActivateAcademicYearCommand) (*application.ActivateAcademicYearResponse, error) {
			return &application.ActivateAcademicYearResponse{
				AcademicYearID: cmd.AcademicYearID,
				Status:         domain.AcademicYearStatusActive,
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("POST", "/academic-years/academic-year-123/activate", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestActivateAcademicYear_MissingID(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("POST", "/academic-years//activate", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestActivateAcademicYear_ServiceError(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		activateAcademicYearFunc: func(ctx context.Context, cmd *application.ActivateAcademicYearCommand) (*application.ActivateAcademicYearResponse, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("POST", "/academic-years/academic-year-123/activate", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestArchiveAcademicYear(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		archiveAcademicYearFunc: func(ctx context.Context, cmd *application.ArchiveAcademicYearCommand) (*application.ArchiveAcademicYearResponse, error) {
			return &application.ArchiveAcademicYearResponse{
				AcademicYearID: cmd.AcademicYearID,
				Status:         domain.AcademicYearStatusArchived,
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("POST", "/academic-years/academic-year-123/archive", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestArchiveAcademicYear_MissingID(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("POST", "/academic-years//archive", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestGetAcademicYear(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		getAcademicYearFunc: func(ctx context.Context, cmd *application.GetAcademicYearCommand) (*application.GetAcademicYearResponse, error) {
			return &application.GetAcademicYearResponse{
				AcademicYear: &domain.AcademicYear{
					ID:     cmd.AcademicYearID,
					Name:   "2024-2025",
					Status: domain.AcademicYearStatusActive,
				},
				Semesters: []*domain.Semester{},
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years/academic-year-123", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetAcademicYear_InvalidID(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Router returns 301 (redirect) for empty ID in path
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status 301, got %d", w.Code)
	}
}

func TestGetAcademicYear_ServiceError(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		getAcademicYearFunc: func(ctx context.Context, cmd *application.GetAcademicYearCommand) (*application.GetAcademicYearResponse, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years/academic-year-123", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestListAcademicYears(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		listAcademicYearsFunc: func(ctx context.Context, cmd *application.ListAcademicYearsCommand) (*application.ListAcademicYearsResponse, error) {
			return &application.ListAcademicYearsResponse{
				AcademicYears: []*domain.AcademicYear{
					{
						ID:     "academic-year-1",
						Name:   "2024-2025",
						Status: domain.AcademicYearStatusActive,
					},
				},
			}, nil
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years?school_id=school-123", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp dto.ListAcademicYearsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(resp.AcademicYears) != 1 {
		t.Errorf("Expected 1 academic year, got %d", len(resp.AcademicYears))
	}
}

func TestListAcademicYears_Unauthorized(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	handler := NewAcademicYearHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/academic-years", handler.ListAcademicYears)

	req, _ := http.NewRequest("GET", "/academic-years", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestListAcademicYears_MissingSchoolID(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{}
	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestListAcademicYears_ServiceError(t *testing.T) {
	mockService := &MockAcademicYearApplicationService{
		listAcademicYearsFunc: func(ctx context.Context, cmd *application.ListAcademicYearsCommand) (*application.ListAcademicYearsResponse, error) {
			return nil, errors.New("service error")
		},
	}

	router := setupAcademicYearTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/academic-years?school_id=school-123", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}
