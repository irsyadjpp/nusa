package assessment

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/service"
	"github.com/nusa/backend/pkg/response"
)

type Handler struct {
	assessmentService *service.AssessmentService
	authService       *service.ResourceAuthorizationService
}

func NewHandler(
	assessmentService *service.AssessmentService,
	authService *service.ResourceAuthorizationService,
) *Handler {
	return &Handler{
		assessmentService: assessmentService,
		authService:       authService,
	}
}

// Assessment handlers
func (h *Handler) CreateAssessment(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req domain.CreateAssessmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	assessment, err := h.assessmentService.CreateAssessment(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, assessment)
}

func (h *Handler) ListAssessments(c *gin.Context) {
	ctx := context.Background()

	var modulAjarID *string
	if id := c.Query("modul_ajar_id"); id != "" {
		modulAjarID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	assessments, total, err := h.assessmentService.ListAssessments(ctx, modulAjarID, nil, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"assessments": assessments, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetAssessment(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	assessment, err := h.assessmentService.GetAssessment(ctx, id)
	if err != nil {
		response.Error(c, 404, "Not found")
		return
	}

	response.Success(c, assessment)
}

// Rubric handlers
func (h *Handler) CreateRubric(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req domain.CreateRubricRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	rubric, err := h.assessmentService.CreateRubric(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, rubric)
}

func (h *Handler) ListRubrics(c *gin.Context) {
	ctx := context.Background()

	var assessmentID *string
	if id := c.Query("assessment_id"); id != "" {
		assessmentID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	rubrics, total, err := h.assessmentService.ListRubrics(ctx, assessmentID, nil, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"rubrics": rubrics, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetRubric(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	rubric, err := h.assessmentService.GetRubric(ctx, id)
	if err != nil {
		response.Error(c, 404, "Rubric not found")
		return
	}

	response.Success(c, rubric)
}

func (h *Handler) UpdateRubric(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var req domain.UpdateRubricRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	rubric, err := h.assessmentService.UpdateRubric(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, rubric)
}

func (h *Handler) DeleteRubric(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := h.assessmentService.DeleteRubric(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Rubric deleted"})
}

// Evidence handlers
func (h *Handler) CreateEvidence(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req domain.CreateEvidenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	evidence, err := h.assessmentService.CreateEvidence(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, evidence)
}

func (h *Handler) ListEvidences(c *gin.Context) {
	ctx := context.Background()

	var studentID, assessmentID *string
	if id := c.Query("student_id"); id != "" {
		studentID = &id
	}
	if id := c.Query("assessment_id"); id != "" {
		assessmentID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	evidences, total, err := h.assessmentService.ListEvidences(ctx, studentID, assessmentID, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"evidences": evidences, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) UpdateEvidence(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	// Authorization check
	if err := h.authService.AuthorizeEvidenceOwnership(ctx, authCtx.UserID, id); err != nil {
		response.Error(c, 403, "Access denied")
		return
	}

	var req domain.UpdateEvidenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	evidence, err := h.assessmentService.UpdateEvidence(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, evidence)
}

func (h *Handler) DeleteEvidence(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	// Authorization check
	if err := h.authService.AuthorizeEvidenceOwnership(ctx, authCtx.UserID, id); err != nil {
		response.Error(c, 403, "Access denied")
		return
	}

	err := h.assessmentService.DeleteEvidence(ctx, id)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Evidence deleted"})
}

// Evaluation handlers
func (h *Handler) CreateEvaluation(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req domain.CreateEvaluationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	evaluation, err := h.assessmentService.CreateEvaluation(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, evaluation)
}

func (h *Handler) UpdateEvaluation(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	var req domain.UpdateEvaluationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	evaluation, err := h.assessmentService.UpdateEvaluation(ctx, id, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, evaluation)
}

func (h *Handler) ListEvaluations(c *gin.Context) {
	ctx := context.Background()

	var studentID *string
	if id := c.Query("student_id"); id != "" {
		studentID = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	evaluations, total, err := h.assessmentService.ListEvaluations(ctx, studentID, nil, nil, page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"evaluations": evaluations, "total": total, "page": page, "page_size": pageSize})
}

func (h *Handler) GetEvaluation(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	evaluation, err := h.assessmentService.GetEvaluation(ctx, id)
	if err != nil {
		response.Error(c, 404, "Evaluation not found")
		return
	}

	response.Success(c, evaluation)
}

func (h *Handler) GetEvaluationHistory(c *gin.Context) {
	ctx := context.Background()
	evidenceID := c.Param("evidence_id")

	history, err := h.assessmentService.GetEvaluationHistory(ctx, evidenceID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, history)
}

func (h *Handler) GetEvaluationFeedbackHistory(c *gin.Context) {
	ctx := context.Background()
	evaluationID := c.Param("id")

	history, err := h.assessmentService.GetEvaluationFeedbackHistory(ctx, evaluationID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, history)
}

// UpdateAssessment updates an assessment (EP-03)
func (h *Handler) UpdateAssessment(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	// Authorization check
	if err := h.authService.AuthorizeAssessmentOwnership(ctx, authCtx.UserID, id); err != nil {
		response.Error(c, 403, "Access denied")
		return
	}

	var req domain.UpdateAssessmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	assessment, err := h.assessmentService.UpdateAssessment(ctx, id, &req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, assessment)
}

// ApproveAssessment approves an assessment (EP-04)
func (h *Handler) ApproveAssessment(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)
	id := c.Param("id")

	err := h.assessmentService.ApproveAssessment(ctx, id, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Assessment approved"})
}

// UploadEvidence handles evidence file upload (EP-05)
func (h *Handler) UploadEvidence(c *gin.Context) {
	ctx := context.Background()
	authCtx := middleware.GetAuthContext(c)

	var req domain.CreateEvidenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	evidence, err := h.assessmentService.UploadEvidence(ctx, &req, authCtx.UserID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, evidence)
}

// GetEvidenceByID retrieves evidence by ID (EP-06)
func (h *Handler) GetEvidence(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	evidence, err := h.assessmentService.GetEvidenceByID(ctx, id)
	if err != nil {
		response.Error(c, 404, "Evidence not found")
		return
	}

	response.Success(c, evidence)
}
