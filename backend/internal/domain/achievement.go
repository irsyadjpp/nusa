package domain

import (
	"time"
)

// Achievement represents calculated student achievement (NOT a persistent entity)
// This is calculated by AchievementService at runtime
type Achievement struct {
	StudentID                  string                     `json:"student_id"`
	StudentName                string                     `json:"student_name"`
	TPID                       string                     `json:"tp_id"`
	TPTitle                    string                     `json:"tp_title"`
	CompetencyCode             string                     `json:"competency_code"`
	OverallScore               float64                    `json:"overall_score"`
	EvaluationPerformanceLevel EvaluationPerformanceLevel `json:"performance_level"`
	MasteryStatus              MasteryStatus              `json:"mastery_status"`
	CompetencyBreakdown        []CompetencyAchievement    `json:"competency_breakdown"`
	EvidenceCount              int                        `json:"evidence_count"`
	EvaluationCount            int                        `json:"evaluation_count"`
	CalculatedAt               time.Time                  `json:"calculated_at"`
}

// CompetencyAchievement represents achievement for a specific competency
type CompetencyAchievement struct {
	CompetencyName             string                     `json:"competency_name"`
	Score                      float64                    `json:"score"`
	EvaluationPerformanceLevel EvaluationPerformanceLevel `json:"performance_level"`
	MasteryStatus              MasteryStatus              `json:"mastery_status"`
	EvidenceCount              int                        `json:"evidence_count"`
}

// MasteryStatus represents the mastery status
type MasteryStatus string

const (
	MasteryStatusNotStarted MasteryStatus = "NOT_STARTED"
	MasteryStatusInProgress MasteryStatus = "IN_PROGRESS"
	MasteryStatusAchieved   MasteryStatus = "ACHIEVED"
	MasteryStatusExceeding  MasteryStatus = "EXCEEDING"
)

// CompetencyProgress represents student progress across competencies
type CompetencyProgress struct {
	StudentID          string                   `json:"student_id"`
	StudentName        string                   `json:"student_name"`
	SubjectID          string                   `json:"subject_id"`
	SubjectName        string                   `json:"subject_name"`
	PhaseID            string                   `json:"phase_id"`
	PhaseName          string                   `json:"phase_name"`
	OverallProgress    float64                  `json:"overall_progress"` // 0-100
	CompetencyProgress []CompetencyProgressItem `json:"competency_progress"`
	CalculatedAt       time.Time                `json:"calculated_at"`
}

// CompetencyProgressItem represents progress for a single competency
type CompetencyProgressItem struct {
	CPCode           string        `json:"cp_code"`
	CPDescription    string        `json:"cp_description"`
	Progress         float64       `json:"progress"` // 0-100
	Status           MasteryStatus `json:"status"`
	TPCount          int           `json:"tp_count"`
	CompletedTPCount int           `json:"completed_tp_count"`
}

// AchievementSummary represents a summary of achievements for reporting
type AchievementSummary struct {
	ReportID            string               `json:"report_id"`
	StudentID           string               `json:"student_id"`
	StudentName         string               `json:"student_name"`
	ClassID             string               `json:"class_id"`
	ClassName           string               `json:"class_name"`
	ReportPeriod        interface{}          `json:"report_period"`
	OverallAchievement  float64              `json:"overall_achievement"`
	SubjectBreakdown    []SubjectAchievement `json:"subject_breakdown"`
	Strengths           []string             `json:"strengths"`
	AreasForImprovement []string             `json:"areas_for_improvement"`
	Recommendations     []string             `json:"recommendations"`
	CalculatedAt        time.Time            `json:"calculated_at"`
}

// SubjectAchievement represents achievement per subject
type SubjectAchievement struct {
	SubjectID                  string                     `json:"subject_id"`
	SubjectName                string                     `json:"subject_name"`
	AverageScore               float64                    `json:"average_score"`
	EvaluationPerformanceLevel EvaluationPerformanceLevel `json:"performance_level"`
	CompetencyCount            int                        `json:"competency_count"`
	MasteredCount              int                        `json:"mastered_count"`
}

// ClassAchievement represents achievement summary for a class
type ClassAchievement struct {
	ClassID                 string                             `json:"class_id"`
	ClassName               string                             `json:"class_name"`
	SubjectID               string                             `json:"subject_id"`
	SubjectName             string                             `json:"subject_name"`
	TotalStudents           int                                `json:"total_students"`
	ClassAverage            float64                            `json:"class_average"`
	PerformanceDistribution map[EvaluationPerformanceLevel]int `json:"performance_distribution"`
	MasteryDistribution     map[MasteryStatus]int              `json:"mastery_distribution"`
	StudentAchievements     []StudentClassAchievement          `json:"student_achievements"`
	CalculatedAt            time.Time                          `json:"calculated_at"`
}

// StudentClassAchievement represents a student's achievement within a class
type StudentClassAchievement struct {
	StudentID                  string                     `json:"student_id"`
	StudentName                string                     `json:"student_name"`
	AverageScore               float64                    `json:"average_score"`
	EvaluationPerformanceLevel EvaluationPerformanceLevel `json:"performance_level"`
	MasteryStatus              MasteryStatus              `json:"mastery_status"`
	Rank                       int                        `json:"rank"`
}

// AchievementService is a domain service for calculating achievements at runtime
// NO persistence - all calculations are done on-demand
type AchievementService struct{}

// NewAchievementService creates a new AchievementService
func NewAchievementService() *AchievementService {
	return &AchievementService{}
}

// CalculateStudentAchievement calculates achievement for a specific student and TP
// This is a runtime calculation based on evaluations
func (s *AchievementService) CalculateStudentAchievement(
	studentID string,
	tpID string,
	evaluations []Evaluation,
	tpSuccessCriteria interface{},
) (*Achievement, error) {
	// Calculate overall score from evaluations
	var totalScore float64
	var maxScore float64

	for _, eval := range evaluations {
		totalScore += float64(eval.TotalScore)
		maxScore += float64(eval.MaxScore)
	}

	var overallScore float64
	if maxScore > 0 {
		overallScore = (totalScore / maxScore) * 100
	}

	// Determine performance level
	performanceLevel := s.determineEvaluationPerformanceLevel(overallScore)

	// Determine mastery status
	masteryStatus := s.determineMasteryStatus(overallScore, performanceLevel)

	// Build competency breakdown
	competencyBreakdown := s.buildCompetencyBreakdown(evaluations)

	return &Achievement{
		StudentID:                  studentID,
		TPID:                       tpID,
		OverallScore:               overallScore,
		EvaluationPerformanceLevel: performanceLevel,
		MasteryStatus:              masteryStatus,
		CompetencyBreakdown:        competencyBreakdown,
		EvidenceCount:              len(evaluations),
		EvaluationCount:            len(evaluations),
		CalculatedAt:               time.Now(),
	}, nil
}

// CalculateCompetencyProgress calculates progress across competencies for a student
func (s *AchievementService) CalculateCompetencyProgress(
	studentID string,
	studentName string,
	subjectID string,
	subjectName string,
	phaseID string,
	phaseName string,
	tps []TP,
	evaluations []Evaluation,
) (*CompetencyProgress, error) {
	// Group evaluations by TP
	tpEvaluations := make(map[string][]Evaluation)
	for _, eval := range evaluations {
		tpEvaluations[eval.EvidenceID] = append(tpEvaluations[eval.EvidenceID], eval)
	}

	// Calculate progress for each TP
	var progressItems []CompetencyProgressItem
	var totalProgress float64
	totalTPs := len(tps)
	completedTPs := 0

	for _, tp := range tps {
		tpEvals := tpEvaluations[tp.ID]
		progress := s.calculateTPProgress(tpEvals)

		if progress >= 80 { // Threshold for "completed"
			completedTPs++
		}

		totalProgress += progress

		progressItems = append(progressItems, CompetencyProgressItem{
			CPCode:           "", // Would be populated from CP
			CPDescription:    "", // Would be populated from CP
			Progress:         progress,
			Status:           s.determineMasteryStatusFromProgress(progress),
			TPCount:          1,
			CompletedTPCount: 0, // Would be calculated
		})
	}

	var overallProgress float64
	if totalTPs > 0 {
		overallProgress = totalProgress / float64(totalTPs)
	}

	return &CompetencyProgress{
		StudentID:          studentID,
		StudentName:        studentName,
		SubjectID:          subjectID,
		SubjectName:        subjectName,
		PhaseID:            phaseID,
		PhaseName:          phaseName,
		OverallProgress:    overallProgress,
		CompetencyProgress: progressItems,
		CalculatedAt:       time.Now(),
	}, nil
}

// GenerateAchievementSummary generates a comprehensive achievement summary for a student
func (s *AchievementService) GenerateAchievementSummary(
	studentID string,
	studentName string,
	classID string,
	className string,
	reportPeriod interface{},
	achievements []Achievement,
) (*AchievementSummary, error) {
	if len(achievements) == 0 {
		return &AchievementSummary{
			StudentID:           studentID,
			StudentName:         studentName,
			ClassID:             classID,
			ClassName:           className,
			ReportPeriod:        reportPeriod,
			OverallAchievement:  0,
			SubjectBreakdown:    []SubjectAchievement{},
			Strengths:           []string{},
			AreasForImprovement: []string{},
			Recommendations:     []string{},
			CalculatedAt:        time.Now(),
		}, nil
	}

	// Calculate overall achievement
	var totalScore float64
	for _, ach := range achievements {
		totalScore += ach.OverallScore
	}
	overallAchievement := totalScore / float64(len(achievements))

	// Build subject breakdown
	subjectBreakdown := s.buildSubjectBreakdown(achievements)

	// Identify strengths and areas for improvement
	strengths, areasForImprovement := s.identifyStrengthsAndWeaknesses(achievements)

	// Generate recommendations
	recommendations := s.generateRecommendations(achievements)

	return &AchievementSummary{
		StudentID:           studentID,
		StudentName:         studentName,
		ClassID:             classID,
		ClassName:           className,
		ReportPeriod:        reportPeriod,
		OverallAchievement:  overallAchievement,
		SubjectBreakdown:    subjectBreakdown,
		Strengths:           strengths,
		AreasForImprovement: areasForImprovement,
		Recommendations:     recommendations,
		CalculatedAt:        time.Now(),
	}, nil
}

// GenerateClassAchievement generates achievement summary for an entire class
func (s *AchievementService) GenerateClassAchievement(
	classID string,
	className string,
	subjectID string,
	subjectName string,
	studentAchievements []Achievement,
) (*ClassAchievement, error) {
	if len(studentAchievements) == 0 {
		return &ClassAchievement{
			ClassID:                 classID,
			ClassName:               className,
			SubjectID:               subjectID,
			SubjectName:             subjectName,
			TotalStudents:           0,
			ClassAverage:            0,
			PerformanceDistribution: map[EvaluationPerformanceLevel]int{},
			MasteryDistribution:     map[MasteryStatus]int{},
			StudentAchievements:     []StudentClassAchievement{},
			CalculatedAt:            time.Now(),
		}, nil
	}

	// Calculate class average
	var totalScore float64
	for _, ach := range studentAchievements {
		totalScore += ach.OverallScore
	}
	classAverage := totalScore / float64(len(studentAchievements))

	// Build performance distribution
	perfDist := make(map[EvaluationPerformanceLevel]int)
	masteryDist := make(map[MasteryStatus]int)

	for _, ach := range studentAchievements {
		perfDist[ach.EvaluationPerformanceLevel]++
		masteryDist[ach.MasteryStatus]++
	}

	// Build student achievements with ranking
	studentClassAchievements := s.buildStudentClassAchievements(studentAchievements)

	return &ClassAchievement{
		ClassID:                 classID,
		ClassName:               className,
		SubjectID:               subjectID,
		SubjectName:             subjectName,
		TotalStudents:           len(studentAchievements),
		ClassAverage:            classAverage,
		PerformanceDistribution: perfDist,
		MasteryDistribution:     masteryDist,
		StudentAchievements:     studentClassAchievements,
		CalculatedAt:            time.Now(),
	}, nil
}

// Helper methods

func (s *AchievementService) determineEvaluationPerformanceLevel(score float64) EvaluationPerformanceLevel {
	if score >= 90 {
		return EvaluationPerformanceLevelExcellent
	} else if score >= 75 {
		return EvaluationPerformanceLevelProficient
	} else if score >= 60 {
		return EvaluationPerformanceLevelDeveloping
	}
	return EvaluationPerformanceLevelBeginning
}

func (s *AchievementService) determineMasteryStatus(score float64, level EvaluationPerformanceLevel) MasteryStatus {
	if level == EvaluationPerformanceLevelExcellent {
		return MasteryStatusExceeding
	} else if level == EvaluationPerformanceLevelProficient {
		return MasteryStatusAchieved
	} else if score > 0 {
		return MasteryStatusInProgress
	}
	return MasteryStatusNotStarted
}

func (s *AchievementService) determineMasteryStatusFromProgress(progress float64) MasteryStatus {
	if progress >= 90 {
		return MasteryStatusExceeding
	} else if progress >= 75 {
		return MasteryStatusAchieved
	} else if progress > 0 {
		return MasteryStatusInProgress
	}
	return MasteryStatusNotStarted
}

func (s *AchievementService) buildCompetencyBreakdown(evaluations []Evaluation) []CompetencyAchievement {
	// In a real implementation, this would group by competency
	// For now, return a simplified version
	return []CompetencyAchievement{}
}

func (s *AchievementService) calculateTPProgress(evaluations []Evaluation) float64 {
	if len(evaluations) == 0 {
		return 0
	}

	var totalScore float64
	var maxScore float64

	for _, eval := range evaluations {
		totalScore += float64(eval.TotalScore)
		maxScore += float64(eval.MaxScore)
	}

	if maxScore > 0 {
		return (totalScore / maxScore) * 100
	}
	return 0
}

func (s *AchievementService) buildSubjectBreakdown(achievements []Achievement) []SubjectAchievement {
	// Group by subject and calculate averages
	// Simplified implementation
	return []SubjectAchievement{}
}

func (s *AchievementService) identifyStrengthsAndWeaknesses(achievements []Achievement) ([]string, []string) {
	var strengths []string
	var weaknesses []string

	for _, ach := range achievements {
		if ach.EvaluationPerformanceLevel == EvaluationPerformanceLevelExcellent {
			strengths = append(strengths, ach.TPTitle)
		} else if ach.EvaluationPerformanceLevel == EvaluationPerformanceLevelBeginning {
			weaknesses = append(weaknesses, ach.TPTitle)
		}
	}

	return strengths, weaknesses
}

func (s *AchievementService) generateRecommendations(achievements []Achievement) []string {
	var recommendations []string

	for _, ach := range achievements {
		if ach.EvaluationPerformanceLevel == EvaluationPerformanceLevelDeveloping || ach.EvaluationPerformanceLevel == EvaluationPerformanceLevelBeginning {
			recommendations = append(recommendations, "Focus on improving "+ach.TPTitle)
		}
	}

	return recommendations
}

func (s *AchievementService) buildStudentClassAchievements(achievements []Achievement) []StudentClassAchievement {
	// Sort by score and assign ranks
	// Simplified implementation
	return []StudentClassAchievement{}
}
