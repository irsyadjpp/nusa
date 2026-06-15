package domain

import (
	"fmt"
	"time"
)

// Announcement represents a school announcement
type Announcement struct {
	ID             string     `json:"id" db:"id"`
	SchoolID       string     `json:"school_id" db:"school_id"`
	Title          string     `json:"title" db:"title"`
	Content        string     `json:"content" db:"content"`
	Priority       string     `json:"priority" db:"priority"`
	TargetAudience string     `json:"target_audience" db:"target_audience"`
	PublishedBy    string     `json:"published_by" db:"published_by"`
	PublishedAt    time.Time  `json:"published_at" db:"published_at"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty" db:"expires_at"`
	IsActive       bool       `json:"is_active" db:"is_active"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// AnnouncementPriority represents the priority level of an announcement
type AnnouncementPriority string

const (
	AnnouncementPriorityLow    AnnouncementPriority = "LOW"
	AnnouncementPriorityNormal AnnouncementPriority = "NORMAL"
	AnnouncementPriorityHigh   AnnouncementPriority = "HIGH"
	AnnouncementPriorityUrgent AnnouncementPriority = "URGENT"
)

// TargetAudience represents the target audience for an announcement
type TargetAudience string

const (
	TargetAudienceAll      TargetAudience = "ALL"
	TargetAudienceTeachers TargetAudience = "TEACHERS"
	TargetAudienceStudents TargetAudience = "STUDENTS"
	TargetAudienceParents  TargetAudience = "PARENTS"
	TargetAudienceAdmin    TargetAudience = "ADMIN"
)

// CreateAnnouncementRequest represents the request to create an announcement
type CreateAnnouncementRequest struct {
	SchoolID       string               `json:"school_id" binding:"required"`
	Title          string               `json:"title" binding:"required,max=255"`
	Content        string               `json:"content" binding:"required"`
	Priority       AnnouncementPriority `json:"priority" binding:"required,oneof=LOW NORMAL HIGH URGENT"`
	TargetAudience TargetAudience       `json:"target_audience" binding:"required,oneof=ALL TEACHERS STUDENTS PARENTS ADMIN"`
	ExpiresAt      *time.Time           `json:"expires_at,omitempty"`
}

// UpdateAnnouncementRequest represents the request to update an announcement
type UpdateAnnouncementRequest struct {
	Title          *string               `json:"title,omitempty" binding:"omitempty,max=255"`
	Content        *string               `json:"content,omitempty"`
	Priority       *AnnouncementPriority `json:"priority,omitempty" binding:"omitempty,oneof=LOW NORMAL HIGH URGENT"`
	TargetAudience *TargetAudience       `json:"target_audience,omitempty" binding:"omitempty,oneof=ALL TEACHERS STUDENTS PARENTS ADMIN"`
	ExpiresAt      *time.Time            `json:"expires_at,omitempty"`
	IsActive       *bool                 `json:"is_active,omitempty"`
}

// AnnouncementResponse represents the announcement data returned to clients
type AnnouncementResponse struct {
	ID              string               `json:"id"`
	SchoolID        string               `json:"school_id"`
	SchoolName      *string              `json:"school_name,omitempty"`
	Title           string               `json:"title"`
	Content         string               `json:"content"`
	Priority        AnnouncementPriority `json:"priority"`
	TargetAudience  TargetAudience       `json:"target_audience"`
	PublishedBy     string               `json:"published_by"`
	PublishedByName *string              `json:"published_by_name,omitempty"`
	PublishedAt     string               `json:"published_at"`
	ExpiresAt       *string              `json:"expires_at,omitempty"`
	IsActive        bool                 `json:"is_active"`
	CreatedAt       string               `json:"created_at"`
	UpdatedAt       string               `json:"updated_at"`
}

// ToAnnouncementResponse converts Announcement to AnnouncementResponse
func (a *Announcement) ToAnnouncementResponse(schoolName, publishedByName string) *AnnouncementResponse {
	var schoolNamePtr, publishedByNamePtr, expiresAtPtr *string

	if schoolName != "" {
		schoolNamePtr = &schoolName
	}
	if publishedByName != "" {
		publishedByNamePtr = &publishedByName
	}
	if a.ExpiresAt != nil {
		expiresAtStr := a.ExpiresAt.Format(time.RFC3339)
		expiresAtPtr = &expiresAtStr
	}

	return &AnnouncementResponse{
		ID:              a.ID,
		SchoolID:        a.SchoolID,
		SchoolName:      schoolNamePtr,
		Title:           a.Title,
		Content:         a.Content,
		Priority:        AnnouncementPriority(a.Priority),
		TargetAudience:  TargetAudience(a.TargetAudience),
		PublishedBy:     a.PublishedBy,
		PublishedByName: publishedByNamePtr,
		PublishedAt:     a.PublishedAt.Format(time.RFC3339),
		ExpiresAt:       expiresAtPtr,
		IsActive:        a.IsActive,
		CreatedAt:       a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       a.UpdatedAt.Format(time.RFC3339),
	}
}

// IsExpired checks if the announcement has expired
func (a *Announcement) IsExpired() bool {
	if a.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*a.ExpiresAt)
}

// Validate validates the announcement entity
func (a *Announcement) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("id is required")
	}
	if a.SchoolID == "" {
		return fmt.Errorf("school_id is required")
	}
	if a.Title == "" {
		return fmt.Errorf("title is required")
	}
	if a.Content == "" {
		return fmt.Errorf("content is required")
	}
	if a.Priority == "" {
		return fmt.Errorf("priority is required")
	}
	validPriorities := map[string]bool{
		string(AnnouncementPriorityLow):    true,
		string(AnnouncementPriorityNormal): true,
		string(AnnouncementPriorityHigh):   true,
		string(AnnouncementPriorityUrgent): true,
	}
	if !validPriorities[a.Priority] {
		return fmt.Errorf("invalid priority: %s", a.Priority)
	}
	if a.TargetAudience == "" {
		return fmt.Errorf("target_audience is required")
	}
	validAudiences := map[string]bool{
		string(TargetAudienceAll):      true,
		string(TargetAudienceTeachers): true,
		string(TargetAudienceStudents): true,
		string(TargetAudienceParents):  true,
		string(TargetAudienceAdmin):    true,
	}
	if !validAudiences[a.TargetAudience] {
		return fmt.Errorf("invalid target_audience: %s", a.TargetAudience)
	}
	if a.PublishedBy == "" {
		return fmt.Errorf("published_by is required")
	}
	return nil
}
