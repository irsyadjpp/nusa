package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nusa/backend/internal/domain"
)

func TestAchievementService_CalculateStudentAchievement_Success(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	tpTitle := "Test TP"
	tp := &domain.TP{
		ID:    "tp-1",
		Title: &tpTitle,
	}

	mockTPRepo.On("GetTPByID", mock.Anything, "tp-1").Return(tp, nil)

	result, err := service.CalculateStudentAchievement(context.Background(), "student-1", "Student Name", "tp-1")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Student Name", result.StudentName)
	assert.Equal(t, tpTitle, result.TPTitle)

	mockTPRepo.AssertExpectations(t)
}

func TestAchievementService_CalculateStudentAchievement_TPNotFound(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	mockTPRepo.On("GetTPByID", mock.Anything, "tp-1").Return(nil, errors.New("not found"))

	result, err := service.CalculateStudentAchievement(context.Background(), "student-1", "Student Name", "tp-1")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get TP")

	mockTPRepo.AssertExpectations(t)
}

func TestAchievementService_CalculateCompetencyProgress_Success(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	tps := []*domain.TP{
		{ID: "tp-1", Title: ptr("TP 1")},
		{ID: "tp-2", Title: ptr("TP 2")},
	}

	mockTPRepo.On("ListTPs", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tps, nil)

	result, err := service.CalculateCompetencyProgress(context.Background(), "student-1", "Student Name", "subject-1", "Subject Name", "phase-1", "Phase Name")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Student Name", result.StudentName)
	assert.Equal(t, "Subject Name", result.SubjectName)

	mockTPRepo.AssertExpectations(t)
}

func TestAchievementService_CalculateCompetencyProgress_ListError(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	mockTPRepo.On("ListTPs", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("database error"))

	result, err := service.CalculateCompetencyProgress(context.Background(), "student-1", "Student Name", "subject-1", "Subject Name", "phase-1", "Phase Name")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get TPs")

	mockTPRepo.AssertExpectations(t)
}

func TestAchievementService_GenerateAchievementSummary_Success(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	result, err := service.GenerateAchievementSummary(context.Background(), "student-1", "Student Name", "class-1", "Class Name", "report-period")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Student Name", result.StudentName)
	assert.Equal(t, "class-1", result.ClassID)
}

func TestAchievementService_GenerateClassAchievement_Success(t *testing.T) {
	mockAssessmentRepo := new(MockAssessmentRepository)
	mockTPRepo := new(MockTPRepository)
	service := NewAchievementService(mockAssessmentRepo, mockTPRepo)

	result, err := service.GenerateClassAchievement(context.Background(), "class-1", "Class Name", "subject-1", "Subject Name")

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "class-1", result.ClassID)
	assert.Equal(t, "Class Name", result.ClassName)
}
