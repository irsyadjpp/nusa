package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/domain"
	"github.com/nusa/backend/internal/handler/dto"
	"github.com/nusa/backend/internal/middleware"
	"github.com/nusa/backend/internal/service"
)

// AnnouncementHandler handles HTTP requests for announcement endpoints
type AnnouncementHandler struct {
	announcementService *service.AnnouncementService
}

// NewAnnouncementHandler creates a new announcement handler
func NewAnnouncementHandler(announcementService *service.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{
		announcementService: announcementService,
	}
}

// CreateAnnouncement creates a new announcement
// POST /api/v1/announcements
func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateAnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateAnnouncementRequest{
		SchoolID:       req.SchoolID,
		Title:          req.Title,
		Content:        req.Content,
		Priority:       domain.AnnouncementPriority(req.Priority),
		TargetAudience: domain.TargetAudience(req.TargetAudience),
		ExpiresAt:      req.ExpiresAt,
	}

	announcement, err := h.announcementService.CreateAnnouncement(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := announcement.ToAnnouncementResponse("", "")
	response := &dto.AnnouncementResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		Title:           domainResponse.Title,
		Content:         domainResponse.Content,
		Priority:        dto.AnnouncementPriority(domainResponse.Priority),
		TargetAudience:  dto.TargetAudience(domainResponse.TargetAudience),
		PublishedBy:     domainResponse.PublishedBy,
		PublishedByName: domainResponse.PublishedByName,
		PublishedAt:     domainResponse.PublishedAt,
		ExpiresAt:       domainResponse.ExpiresAt,
		IsActive:        domainResponse.IsActive,
		CreatedAt:       domainResponse.CreatedAt,
		UpdatedAt:       domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAnnouncement retrieves an announcement by ID
// GET /api/v1/announcements/:id
func (h *AnnouncementHandler) GetAnnouncement(c *gin.Context) {
	id := c.Param("id")

	announcement, err := h.announcementService.GetAnnouncement(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
		return
	}

	domainResponse := announcement.ToAnnouncementResponse("", "")
	response := &dto.AnnouncementResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		Title:           domainResponse.Title,
		Content:         domainResponse.Content,
		Priority:        dto.AnnouncementPriority(domainResponse.Priority),
		TargetAudience:  dto.TargetAudience(domainResponse.TargetAudience),
		PublishedBy:     domainResponse.PublishedBy,
		PublishedByName: domainResponse.PublishedByName,
		PublishedAt:     domainResponse.PublishedAt,
		ExpiresAt:       domainResponse.ExpiresAt,
		IsActive:        domainResponse.IsActive,
		CreatedAt:       domainResponse.CreatedAt,
		UpdatedAt:       domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ListAnnouncements retrieves announcements with filters and pagination
// GET /api/v1/announcements
func (h *AnnouncementHandler) ListAnnouncements(c *gin.Context) {
	var schoolID, priority, targetAudience *string
	var isActive *bool

	if schoolIDStr := c.Query("school_id"); schoolIDStr != "" {
		schoolID = &schoolIDStr
	}
	if priorityStr := c.Query("priority"); priorityStr != "" {
		priority = &priorityStr
	}
	if targetAudienceStr := c.Query("target_audience"); targetAudienceStr != "" {
		targetAudience = &targetAudienceStr
	}
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		active, err := strconv.ParseBool(isActiveStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid is_active parameter"})
			return
		}
		isActive = &active
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	announcements, total, err := h.announcementService.ListAnnouncements(c.Request.Context(), schoolID, priority, targetAudience, isActive, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	announcementResponses := make([]*dto.AnnouncementResponse, len(announcements))
	for i, announcement := range announcements {
		domainResponse := announcement.ToAnnouncementResponse("", "")
		announcementResponses[i] = &dto.AnnouncementResponse{
			ID:              domainResponse.ID,
			SchoolID:        domainResponse.SchoolID,
			SchoolName:      domainResponse.SchoolName,
			Title:           domainResponse.Title,
			Content:         domainResponse.Content,
			Priority:        dto.AnnouncementPriority(domainResponse.Priority),
			TargetAudience:  dto.TargetAudience(domainResponse.TargetAudience),
			PublishedBy:     domainResponse.PublishedBy,
			PublishedByName: domainResponse.PublishedByName,
			PublishedAt:     domainResponse.PublishedAt,
			ExpiresAt:       domainResponse.ExpiresAt,
			IsActive:        domainResponse.IsActive,
			CreatedAt:       domainResponse.CreatedAt,
			UpdatedAt:       domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.AnnouncementListResponse{
		Announcements: announcementResponses,
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
	})
}

// UpdateAnnouncement updates announcement information
// PUT /api/v1/announcements/:id
func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateAnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateAnnouncementRequest{
		Title:          req.Title,
		Content:        req.Content,
		Priority:       (*domain.AnnouncementPriority)(req.Priority),
		TargetAudience: (*domain.TargetAudience)(req.TargetAudience),
		ExpiresAt:      req.ExpiresAt,
		IsActive:       req.IsActive,
	}

	announcement, err := h.announcementService.UpdateAnnouncement(c.Request.Context(), id, domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := announcement.ToAnnouncementResponse("", "")
	response := &dto.AnnouncementResponse{
		ID:              domainResponse.ID,
		SchoolID:        domainResponse.SchoolID,
		SchoolName:      domainResponse.SchoolName,
		Title:           domainResponse.Title,
		Content:         domainResponse.Content,
		Priority:        dto.AnnouncementPriority(domainResponse.Priority),
		TargetAudience:  dto.TargetAudience(domainResponse.TargetAudience),
		PublishedBy:     domainResponse.PublishedBy,
		PublishedByName: domainResponse.PublishedByName,
		PublishedAt:     domainResponse.PublishedAt,
		ExpiresAt:       domainResponse.ExpiresAt,
		IsActive:        domainResponse.IsActive,
		CreatedAt:       domainResponse.CreatedAt,
		UpdatedAt:       domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAnnouncement soft deletes an announcement
// DELETE /api/v1/announcements/:id
func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.announcementService.DeleteAnnouncement(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetSchoolAnnouncements retrieves all announcements for a school
// GET /api/v1/schools/:schoolId/announcements
func (h *AnnouncementHandler) GetSchoolAnnouncements(c *gin.Context) {
	schoolID := c.Param("schoolId")

	announcements, err := h.announcementService.GetSchoolAnnouncements(c.Request.Context(), schoolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	announcementResponses := make([]*dto.AnnouncementResponse, len(announcements))
	for i, announcement := range announcements {
		domainResponse := announcement.ToAnnouncementResponse("", "")
		announcementResponses[i] = &dto.AnnouncementResponse{
			ID:              domainResponse.ID,
			SchoolID:        domainResponse.SchoolID,
			SchoolName:      domainResponse.SchoolName,
			Title:           domainResponse.Title,
			Content:         domainResponse.Content,
			Priority:        dto.AnnouncementPriority(domainResponse.Priority),
			TargetAudience:  dto.TargetAudience(domainResponse.TargetAudience),
			PublishedBy:     domainResponse.PublishedBy,
			PublishedByName: domainResponse.PublishedByName,
			PublishedAt:     domainResponse.PublishedAt,
			ExpiresAt:       domainResponse.ExpiresAt,
			IsActive:        domainResponse.IsActive,
			CreatedAt:       domainResponse.CreatedAt,
			UpdatedAt:       domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.AnnouncementListResponse{
		Announcements: announcementResponses,
		Total:         len(announcementResponses),
		Page:          1,
		PageSize:      len(announcementResponses),
	})
}
