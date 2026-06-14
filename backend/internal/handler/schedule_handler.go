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

// ScheduleHandler handles HTTP requests for schedule endpoints
type ScheduleHandler struct {
	scheduleService *service.ScheduleService
}

// NewScheduleHandler creates a new schedule handler
func NewScheduleHandler(scheduleService *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
	}
}

// CreateSchedule creates a new schedule
// POST /api/v1/schedules
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateScheduleRequest{
		ClassID:   req.ClassID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Room:      req.Room,
	}

	schedule, err := h.scheduleService.CreateSchedule(c.Request.Context(), domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := schedule.ToScheduleResponse("")
	response := &dto.ScheduleResponse{
		ID:        domainResponse.ID,
		ClassID:   domainResponse.ClassID,
		ClassName: domainResponse.ClassName,
		DayOfWeek: domainResponse.DayOfWeek,
		DayName:   domainResponse.DayName,
		StartTime: domainResponse.StartTime,
		EndTime:   domainResponse.EndTime,
		Room:      domainResponse.Room,
		IsActive:  domainResponse.IsActive,
		CreatedAt: domainResponse.CreatedAt,
		UpdatedAt: domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetSchedule retrieves a schedule by ID
// GET /api/v1/schedules/:id
func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
	id := c.Param("id")

	schedule, err := h.scheduleService.GetSchedule(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	domainResponse := schedule.ToScheduleResponse("")
	response := &dto.ScheduleResponse{
		ID:        domainResponse.ID,
		ClassID:   domainResponse.ClassID,
		ClassName: domainResponse.ClassName,
		DayOfWeek: domainResponse.DayOfWeek,
		DayName:   domainResponse.DayName,
		StartTime: domainResponse.StartTime,
		EndTime:   domainResponse.EndTime,
		Room:      domainResponse.Room,
		IsActive:  domainResponse.IsActive,
		CreatedAt: domainResponse.CreatedAt,
		UpdatedAt: domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ListSchedules retrieves schedules with filters and pagination
// GET /api/v1/schedules
func (h *ScheduleHandler) ListSchedules(c *gin.Context) {
	var classID *string
	var dayOfWeek, isActive *int

	if classIDStr := c.Query("class_id"); classIDStr != "" {
		classID = &classIDStr
	}
	if dayOfWeekStr := c.Query("day_of_week"); dayOfWeekStr != "" {
		day, err := strconv.Atoi(dayOfWeekStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day_of_week parameter"})
			return
		}
		dayOfWeek = &day
	}
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		active, err := strconv.ParseBool(isActiveStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid is_active parameter"})
			return
		}
		isActiveInt := 0
		if active {
			isActiveInt = 1
		}
		isActive = &isActiveInt
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var isActiveBool *bool
	if isActive != nil {
		b := *isActive == 1
		isActiveBool = &b
	}

	schedules, total, err := h.scheduleService.ListSchedules(c.Request.Context(), classID, dayOfWeek, isActiveBool, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scheduleResponses := make([]*dto.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		domainResponse := schedule.ToScheduleResponse("")
		scheduleResponses[i] = &dto.ScheduleResponse{
			ID:        domainResponse.ID,
			ClassID:   domainResponse.ClassID,
			ClassName: domainResponse.ClassName,
			DayOfWeek: domainResponse.DayOfWeek,
			DayName:   domainResponse.DayName,
			StartTime: domainResponse.StartTime,
			EndTime:   domainResponse.EndTime,
			Room:      domainResponse.Room,
			IsActive:  domainResponse.IsActive,
			CreatedAt: domainResponse.CreatedAt,
			UpdatedAt: domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ScheduleListResponse{
		Schedules: scheduleResponses,
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
	})
}

// UpdateSchedule updates schedule information
// PUT /api/v1/schedules/:id
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	var req dto.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.UpdateScheduleRequest{
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Room:      req.Room,
		IsActive:  req.IsActive,
	}

	schedule, err := h.scheduleService.UpdateSchedule(c.Request.Context(), id, domainReq, authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := schedule.ToScheduleResponse("")
	response := &dto.ScheduleResponse{
		ID:        domainResponse.ID,
		ClassID:   domainResponse.ClassID,
		ClassName: domainResponse.ClassName,
		DayOfWeek: domainResponse.DayOfWeek,
		DayName:   domainResponse.DayName,
		StartTime: domainResponse.StartTime,
		EndTime:   domainResponse.EndTime,
		Room:      domainResponse.Room,
		IsActive:  domainResponse.IsActive,
		CreatedAt: domainResponse.CreatedAt,
		UpdatedAt: domainResponse.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSchedule soft deletes a schedule
// DELETE /api/v1/schedules/:id
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.scheduleService.DeleteSchedule(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetClassSchedules retrieves all schedules for a class
// GET /api/v1/classes/:classId/schedules
func (h *ScheduleHandler) GetClassSchedules(c *gin.Context) {
	classID := c.Param("classId")

	schedules, err := h.scheduleService.GetClassSchedules(c.Request.Context(), classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scheduleResponses := make([]*dto.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		domainResponse := schedule.ToScheduleResponse("")
		scheduleResponses[i] = &dto.ScheduleResponse{
			ID:        domainResponse.ID,
			ClassID:   domainResponse.ClassID,
			ClassName: domainResponse.ClassName,
			DayOfWeek: domainResponse.DayOfWeek,
			DayName:   domainResponse.DayName,
			StartTime: domainResponse.StartTime,
			EndTime:   domainResponse.EndTime,
			Room:      domainResponse.Room,
			IsActive:  domainResponse.IsActive,
			CreatedAt: domainResponse.CreatedAt,
			UpdatedAt: domainResponse.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ScheduleListResponse{
		Schedules: scheduleResponses,
		Total:     len(scheduleResponses),
		Page:      1,
		PageSize:  len(scheduleResponses),
	})
}
