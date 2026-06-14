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

// MessageHandler handles HTTP requests for message endpoints
type MessageHandler struct {
	messageService *service.MessageService
}

// NewMessageHandler creates a new message handler
func NewMessageHandler(messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
	}
}

// CreateMessage creates a new message
// POST /api/v1/messages
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var req dto.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := &domain.CreateMessageRequest{
		SenderID:       req.SenderID,
		ReceiverID:     req.ReceiverID,
		Subject:        req.Subject,
		Content:        req.Content,
		ParentMessageID: req.ParentMessageID,
	}

	message, err := h.messageService.CreateMessage(c.Request.Context(), domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	domainResponse := message.ToMessageResponse("", "")
	response := &dto.MessageResponse{
		ID:             domainResponse.ID,
		SenderID:       domainResponse.SenderID,
		SenderName:     domainResponse.SenderName,
		ReceiverID:     domainResponse.ReceiverID,
		ReceiverName:   domainResponse.ReceiverName,
		Subject:        domainResponse.Subject,
		Content:        domainResponse.Content,
		IsRead:         domainResponse.IsRead,
		ReadAt:         domainResponse.ReadAt,
		ParentMessageID: domainResponse.ParentMessageID,
		CreatedAt:      domainResponse.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// GetMessage retrieves a message by ID
// GET /api/v1/messages/:id
func (h *MessageHandler) GetMessage(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	message, err := h.messageService.GetMessage(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	domainResponse := message.ToMessageResponse("", "")
	response := &dto.MessageResponse{
		ID:             domainResponse.ID,
		SenderID:       domainResponse.SenderID,
		SenderName:     domainResponse.SenderName,
		ReceiverID:     domainResponse.ReceiverID,
		ReceiverName:   domainResponse.ReceiverName,
		Subject:        domainResponse.Subject,
		Content:        domainResponse.Content,
		IsRead:         domainResponse.IsRead,
		ReadAt:         domainResponse.ReadAt,
		ParentMessageID: domainResponse.ParentMessageID,
		CreatedAt:      domainResponse.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ListMessages retrieves messages with filters and pagination
// GET /api/v1/messages
func (h *MessageHandler) ListMessages(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	var senderID, receiverID *string
	var isRead *bool

	if senderIDStr := c.Query("sender_id"); senderIDStr != "" {
		senderID = &senderIDStr
	}
	if receiverIDStr := c.Query("receiver_id"); receiverIDStr != "" {
		receiverID = &receiverIDStr
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

	messages, total, err := h.messageService.ListMessages(c.Request.Context(), senderID, receiverID, isRead, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messageResponses := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		domainResponse := message.ToMessageResponse("", "")
		messageResponses[i] = &dto.MessageResponse{
			ID:             domainResponse.ID,
			SenderID:       domainResponse.SenderID,
			SenderName:     domainResponse.SenderName,
			ReceiverID:     domainResponse.ReceiverID,
			ReceiverName:   domainResponse.ReceiverName,
			Subject:        domainResponse.Subject,
			Content:        domainResponse.Content,
			IsRead:         domainResponse.IsRead,
			ReadAt:         domainResponse.ReadAt,
			ParentMessageID: domainResponse.ParentMessageID,
			CreatedAt:      domainResponse.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.MessageListResponse{
		Messages: messageResponses,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// MarkAsRead marks a message as read
// PUT /api/v1/messages/:id/read
func (h *MessageHandler) MarkAsRead(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.messageService.MarkAsRead(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message marked as read"})
}

// DeleteMessage soft deletes a message
// DELETE /api/v1/messages/:id
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	id := c.Param("id")

	if err := h.messageService.DeleteMessage(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetConversation retrieves messages between two users
// GET /api/v1/conversations/:userId1/:userId2
func (h *MessageHandler) GetConversation(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	userID1 := c.Param("userId1")
	userID2 := c.Param("userId2")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	messages, err := h.messageService.GetConversation(c.Request.Context(), userID1, userID2, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messageResponses := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		domainResponse := message.ToMessageResponse("", "")
		messageResponses[i] = &dto.MessageResponse{
			ID:             domainResponse.ID,
			SenderID:       domainResponse.SenderID,
			SenderName:     domainResponse.SenderName,
			ReceiverID:     domainResponse.ReceiverID,
			ReceiverName:   domainResponse.ReceiverName,
			Subject:        domainResponse.Subject,
			Content:        domainResponse.Content,
			IsRead:         domainResponse.IsRead,
			ReadAt:         domainResponse.ReadAt,
			ParentMessageID: domainResponse.ParentMessageID,
			CreatedAt:      domainResponse.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.MessageListResponse{
		Messages: messageResponses,
		Total:    len(messageResponses),
		Page:     page,
		PageSize: pageSize,
	})
}

// GetUnreadCount returns the count of unread messages for a user
// GET /api/v1/messages/unread/count
func (h *MessageHandler) GetUnreadCount(c *gin.Context) {
	authCtx := middleware.GetAuthContext(c)
	if authCtx == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	count, err := h.messageService.GetUnreadCount(c.Request.Context(), authCtx.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"unread_count": count})
}
