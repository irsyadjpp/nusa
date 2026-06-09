package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/application"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
)

// MockTPSetApplicationService is a mock for testing
type MockTPSetApplicationService struct {
	createTPSetFunc  func(ctx context.Context, cmd *application.CreateTPSetCommand) (*application.CreateTPSetResponse, error)
	listTPSetsFunc   func(ctx context.Context, query *application.ListTPSetsQuery) (*application.ListTPSetsResponse, error)
	getTPSetFunc     func(ctx context.Context, query *application.GetTPSetQuery) (*application.GetTPSetResponse, error)
	approveTPSetFunc func(ctx context.Context, cmd *application.ApproveTPSetCommand) (*application.ApproveTPSetResponse, error)
}

func (m *MockTPSetApplicationService) CreateTPSet(ctx context.Context, cmd *application.CreateTPSetCommand) (*application.CreateTPSetResponse, error) {
	if m.createTPSetFunc != nil {
		return m.createTPSetFunc(ctx, cmd)
	}
	return &application.CreateTPSetResponse{TPSetID: "test-id", Status: domain.WorkflowStatusDraft}, nil
}

func (m *MockTPSetApplicationService) ListTPSets(ctx context.Context, query *application.ListTPSetsQuery) (*application.ListTPSetsResponse, error) {
	if m.listTPSetsFunc != nil {
		return m.listTPSetsFunc(ctx, query)
	}
	return &application.ListTPSetsResponse{TPSets: []*domain.TPSet{}, Total: 0, Page: 1, PageSize: 20}, nil
}

func (m *MockTPSetApplicationService) GetTPSet(ctx context.Context, query *application.GetTPSetQuery) (*application.GetTPSetResponse, error) {
	if m.getTPSetFunc != nil {
		return m.getTPSetFunc(ctx, query)
	}
	return &application.GetTPSetResponse{TPSet: &domain.TPSet{}}, nil
}

func (m *MockTPSetApplicationService) ApproveTPSet(ctx context.Context, cmd *application.ApproveTPSetCommand) (*application.ApproveTPSetResponse, error) {
	if m.approveTPSetFunc != nil {
		return m.approveTPSetFunc(ctx, cmd)
	}
	return &application.ApproveTPSetResponse{TPSetID: "test-id", Status: domain.WorkflowStatusApproved, ApprovedBy: "approver-id", ApprovedAt: "2026-06-09T00:00:00Z"}, nil
}

func NewTPSetHandlerWithInterface(service ITPSetApplicationService) *TPSetHandler {
	return &TPSetHandler{
		tpSetApplicationService: service,
	}
}

func setupTestRouter(service ITPSetApplicationService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := NewTPSetHandlerWithInterface(service)

	// Add mock auth middleware
	router.Use(func(c *gin.Context) {
		authCtx := &middleware.AuthContext{
			UserID:      "test-user-id",
			SchoolID:    stringPtr("test-school-id"),
			Role:        "TEACHER",
			Permissions: []string{"tp_set:CREATE", "tp_set:READ", "tp_set:UPDATE"},
		}
		c.Set(middleware.AuthContextKey, authCtx)
		c.Next()
	})

	router.POST("/tp-sets", handler.CreateTPSet)
	router.GET("/tp-sets", handler.ListTPSets)
	router.GET("/tp-sets/:id", handler.GetTPSet)
	router.POST("/tp-sets/:id/approve", handler.ApproveTPSet)
	router.POST("/tps", handler.CreateTP)
	router.GET("/tps", handler.ListTPs)
	router.GET("/tps/:id", handler.GetTP)

	return router
}

func stringPtr(s string) *string {
	return &s
}

func TestCreateTPSet(t *testing.T) {
	mockService := &MockTPSetApplicationService{
		createTPSetFunc: func(ctx context.Context, cmd *application.CreateTPSetCommand) (*application.CreateTPSetResponse, error) {
			return &application.CreateTPSetResponse{TPSetID: "test-id", Status: domain.WorkflowStatusDraft}, nil
		},
		getTPSetFunc: func(ctx context.Context, query *application.GetTPSetQuery) (*application.GetTPSetResponse, error) {
			return &application.GetTPSetResponse{
				TPSet: &domain.TPSet{
					ID:               "test-id",
					CPID:             "cp-id",
					VersionNo:        1,
					Status:           domain.WorkflowStatusDraft,
					GenerationSource: domain.GenerationSourceManual,
					GeneratedBy:      "test-user-id",
				},
			}, nil
		},
	}

	router := setupTestRouter(mockService)

	reqBody := dto.CreateTPSetRequest{
		CPID:             "cp-id",
		VersionNo:        1,
		GenerationSource: "MANUAL",
		GenerationReason: "Test reason",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tp-sets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var resp dto.TPSetResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", resp.ID)
	}
}

func TestCreateTPSet_Unauthorized(t *testing.T) {
	mockService := &MockTPSetApplicationService{}
	handler := NewTPSetHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// No auth middleware
	router.POST("/tp-sets", handler.CreateTPSet)

	reqBody := dto.CreateTPSetRequest{
		CPID:             "cp-id",
		VersionNo:        1,
		GenerationSource: "MANUAL",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tp-sets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestListTPSets(t *testing.T) {
	mockService := &MockTPSetApplicationService{
		listTPSetsFunc: func(ctx context.Context, query *application.ListTPSetsQuery) (*application.ListTPSetsResponse, error) {
			return &application.ListTPSetsResponse{
				TPSets: []*domain.TPSet{
					{
						ID:               "test-id",
						CPID:             "cp-id",
						VersionNo:        1,
						Status:           domain.WorkflowStatusDraft,
						GenerationSource: domain.GenerationSourceManual,
						GeneratedBy:      "test-user-id",
					},
				},
				Total:    1,
				Page:     1,
				PageSize: 20,
			}, nil
		},
	}

	router := setupTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/tp-sets", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp dto.ListTPSetsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.Total != 1 {
		t.Errorf("Expected Total 1, got %d", resp.Total)
	}
}

func TestGetTPSet(t *testing.T) {
	mockService := &MockTPSetApplicationService{
		getTPSetFunc: func(ctx context.Context, query *application.GetTPSetQuery) (*application.GetTPSetResponse, error) {
			return &application.GetTPSetResponse{
				TPSet: &domain.TPSet{
					ID:               "test-id",
					CPID:             "cp-id",
					VersionNo:        1,
					Status:           domain.WorkflowStatusDraft,
					GenerationSource: domain.GenerationSourceManual,
					GeneratedBy:      "test-user-id",
				},
			}, nil
		},
	}

	router := setupTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/tp-sets/test-id", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp dto.TPSetResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", resp.ID)
	}
}

func TestApproveTPSet(t *testing.T) {
	mockService := &MockTPSetApplicationService{
		approveTPSetFunc: func(ctx context.Context, cmd *application.ApproveTPSetCommand) (*application.ApproveTPSetResponse, error) {
			return &application.ApproveTPSetResponse{
				TPSetID:    "test-id",
				Status:     domain.WorkflowStatusApproved,
				ApprovedBy: "approver-id",
				ApprovedAt: "2026-06-09T00:00:00Z",
			}, nil
		},
	}

	handler := NewTPSetHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add mock auth middleware with SCHOOL_ADMIN role
	router.Use(func(c *gin.Context) {
		authCtx := &middleware.AuthContext{
			UserID:      "test-user-id",
			SchoolID:    stringPtr("test-school-id"),
			Role:        "SCHOOL_ADMIN",
			Permissions: []string{"tp_set:APPROVE"},
		}
		c.Set(middleware.AuthContextKey, authCtx)
		c.Next()
	})

	router.POST("/tp-sets/:id/approve", handler.ApproveTPSet)

	reqBody := dto.ApproveTPSetRequest{
		Reason: "Approved for testing",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tp-sets/test-id/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp dto.ApproveTPSetResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.Message != "TP Set approved successfully" {
		t.Errorf("Expected message 'TP Set approved successfully', got %s", resp.Message)
	}
}

func TestApproveTPSet_Forbidden(t *testing.T) {
	mockService := &MockTPSetApplicationService{}
	handler := NewTPSetHandlerWithInterface(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add mock auth middleware with TEACHER role (not allowed to approve)
	router.Use(func(c *gin.Context) {
		authCtx := &middleware.AuthContext{
			UserID:      "test-user-id",
			SchoolID:    stringPtr("test-school-id"),
			Role:        "TEACHER",
			Permissions: []string{"tp_set:READ"},
		}
		c.Set(middleware.AuthContextKey, authCtx)
		c.Next()
	})

	router.POST("/tp-sets/:id/approve", handler.ApproveTPSet)

	reqBody := dto.ApproveTPSetRequest{
		Reason: "Approved for testing",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tp-sets/test-id/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", w.Code)
	}
}

func TestCreateTP_NotImplemented(t *testing.T) {
	mockService := &MockTPSetApplicationService{}
	router := setupTestRouter(mockService)

	reqBody := dto.CreateTPRequest{
		TPSetID:        "tp-set-id",
		SequenceNumber: 1,
		CPID:           "cp-id",
		SubjectID:      "subject-id",
		PhaseID:        "phase-id",
		ElementID:      "element-id",
		SubelementID:   "subelement-id",
		Status:         "DRAFT",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf("Expected status 501, got %d", w.Code)
	}
}

func TestListTPs_NotImplemented(t *testing.T) {
	mockService := &MockTPSetApplicationService{}
	router := setupTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/tps", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf("Expected status 501, got %d", w.Code)
	}
}

func TestGetTP_NotImplemented(t *testing.T) {
	mockService := &MockTPSetApplicationService{}
	router := setupTestRouter(mockService)

	req, _ := http.NewRequest("GET", "/tps/test-id", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf("Expected status 501, got %d", w.Code)
	}
}
