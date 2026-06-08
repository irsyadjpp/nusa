package achievement

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/service"
)

// Handler handles HTTP requests for achievement endpoints
type Handler struct {
	achievementService *service.AchievementService
}

// NewHandler creates a new achievement handler
func NewHandler(achievementService *service.AchievementService) *Handler {
	return &Handler{
		achievementService: achievementService,
	}
}

// GetStudentAchievement retrieves achievement for a specific student and TP
// GET /students/:id/achievement
func (h *Handler) GetStudentAchievement(c *gin.Context) {
	studentID := c.Param("id")
	tpID := c.Query("tp_id")

	if tpID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tp_id query parameter is required"})
		return
	}

	// In a real implementation, we would get student name from a user service
	studentName := "Student Name"

	achievement, err := h.achievementService.CalculateStudentAchievement(c.Request.Context(), studentID, studentName, tpID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, achievement)
}

// GetStudentProgress retrieves competency progress for a specific student
// GET /students/:id/progress
func (h *Handler) GetStudentProgress(c *gin.Context) {
	studentID := c.Param("id")
	subjectID := c.Query("subject_id")
	phaseID := c.Query("phase_id")

	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject_id query parameter is required"})
		return
	}

	// In a real implementation, we would get student name from a user service
	studentName := "Student Name"
	subjectName := "Subject Name"
	phaseName := "Phase Name"

	progress, err := h.achievementService.CalculateCompetencyProgress(
		c.Request.Context(),
		studentID,
		studentName,
		subjectID,
		subjectName,
		phaseID,
		phaseName,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// GetClassAchievement retrieves achievement summary for an entire class
// GET /classes/:id/achievement
func (h *Handler) GetClassAchievement(c *gin.Context) {
	classID := c.Param("id")
	subjectID := c.Query("subject_id")

	if subjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject_id query parameter is required"})
		return
	}

	// In a real implementation, we would get class name from a class service
	className := "Class Name"
	subjectName := "Subject Name"

	classAchievement, err := h.achievementService.GenerateClassAchievement(
		c.Request.Context(),
		classID,
		className,
		subjectID,
		subjectName,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classAchievement)
}

// GetReportAchievementSummary retrieves achievement summary for a report
// GET /reports/:id/achievement-summary
func (h *Handler) GetReportAchievementSummary(c *gin.Context) {
	_ = c.Param("id") // reportID - used for future report lookup
	studentID := c.Query("student_id")
	classID := c.Query("class_id")

	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id query parameter is required"})
		return
	}

	if classID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "class_id query parameter is required"})
		return
	}

	// In a real implementation, we would get student name and class name from respective services
	studentName := "Student Name"
	className := "Class Name"
	reportPeriod := map[string]interface{}{
		"start_date": "2024-01-01",
		"end_date":   "2024-12-31",
	}

	summary, err := h.achievementService.GenerateAchievementSummary(
		c.Request.Context(),
		studentID,
		studentName,
		classID,
		className,
		reportPeriod,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
