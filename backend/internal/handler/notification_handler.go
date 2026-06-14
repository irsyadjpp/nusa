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

// NotificationHandler handles HTTP requests for notification endpoints
type NotificationHandler struct {
	notificationService *service.NotificationService
}

// NewNotificationHandler creates a new notification handler
func NewNotificationHandler(notificationService *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// CreateNotification creates a new notification
// POST /api/v1/notifications
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateNotificationRequest{
		UserID:    req.UserID,
		Title:     req.Title,
		Message:   req.Message,
		Type:      domain.NotificationType(req.Type),
		ActionURL: req.ActionURL,
		Metadata:  req.Metadata,
	}

	notification, err := h.notificationService.CreateNotification(c.Request.Context(), domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := notification.ToNotificationResponse("")
	response := &dto.NotificationResponse{
		ID:        domainResponse.ID,
		UserID:    domainResponse.UserID,
		UserName:  domainResponse.UserName,
		Title:     domainResponse.Title,
		Message:   domainResponse.Message,
		Type:      dto.NotificationType(domainResponse.Type),
		IsRead:    domainResponse.IsRead,
		ReadAt:    domainResponse.ReadAt,
		ActionURL: domainResponse.ActionURL,
		Metadata:  domainResponse.Metadata,
		CreatedAt: domainResponse.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetNotification retrieves a notification by ID
// GET /api/v1/notifications/:id
func (h *NotificationHandler) GetNotification(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	notification, err := h.notificationService.GetNotification(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	domainResponse := notification.ToNotificationResponse("")
	response := &dto.NotificationResponse{
		ID:        domainResponse.ID,
		UserID:    domainResponse.UserID,
		UserName:  domainResponse.UserName,
		Title:     domainResponse.Title,
		Message:   domainResponse.Message,
		Type:      dto.NotificationType(domainResponse.Type),
		IsRead:    domainResponse.IsRead,
		ReadAt:    domainResponse.ReadAt,
		ActionURL: domainResponse.ActionURL,
		Metadata:  domainResponse.Metadata,
		CreatedAt: domainResponse.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ListNotifications retrieves notifications with filters and pagination
// GET /api/v1/notifications
func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var userID *string
	var notificationType *string
	var isRead *bool

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID = &userIDStr
	}
	if notificationTypeStr := c.Query("type"); notificationTypeStr != "" {
		notificationType = &notificationTypeStr
	}
	if isReadStr := c.Query("is_read"); isReadStr != "" {
		read, err := strconv.ParseBool(isReadStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid is_read parameter"})
			return
		}
		isRead = &read
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	notifications, total, err := h.notificationService.ListNotifications(c.Request.Context(), userID, notificationType, isRead, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	notificationResponses := make([]*dto.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		domainResponse := notification.ToNotificationResponse("")
		notificationResponses[i] = &dto.NotificationResponse{
			ID:        domainResponse.ID,
			UserID:    domainResponse.UserID,
			UserName:  domainResponse.UserName,
			Title:     domainResponse.Title,
			Message:   domainResponse.Message,
			Type:      dto.NotificationType(domainResponse.Type),
			IsRead:    domainResponse.IsRead,
			ReadAt:    domainResponse.ReadAt,
			ActionURL: domainResponse.ActionURL,
			Metadata:  domainResponse.Metadata,
			CreatedAt: domainResponse.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.NotificationListResponse{
		Notifications: notificationResponses,
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
	})
}

// MarkAsRead marks a notification as read
// PUT /api/v1/notifications/:id/read
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.notificationService.MarkAsRead(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

// DeleteNotification soft deletes a notification
// DELETE /api/v1/notifications/:id
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.notificationService.DeleteNotification(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetUnreadCount returns the count of unread notifications for a user
// GET /api/v1/notifications/unread/count
func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	count, err := h.notificationService.GetUnreadCount(c.Request.Context(), authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"unread_count": count})
}
