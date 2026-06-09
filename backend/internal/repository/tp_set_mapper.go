package repository

import (
	"github.com/nusa/backend/internal/domain"
)

// MapTPSetDBModelToDomain converts TPSetDBModel to domain.TPSet
func MapTPSetDBModelToDomain(dbModel *TPSetDBModel) *domain.TPSet {
	if dbModel == nil {
		return nil
	}

	return &domain.TPSet{
		ID:               dbModel.ID,
		CPID:             dbModel.CPID,
		VersionNo:        dbModel.VersionNo,
		Status:           domain.WorkflowStatus(dbModel.Status),
		GenerationSource: domain.GenerationSource(dbModel.GenerationSource),
		GenerationReason: dbModel.GenerationReason,
		GeneratedBy:      dbModel.GeneratedBy,
		AIGenerationID:   dbModel.AIGenerationID,
		ApprovedBy:       dbModel.ApprovedBy,
		ApprovedAt:       dbModel.ApprovedAt,
		CreatedAt:        dbModel.CreatedAt,
		UpdatedAt:        dbModel.UpdatedAt,
	}
}

// MapTPSetDomainToDBModel converts domain.TPSet to TPSetDBModel
func MapTPSetDomainToDBModel(domainModel *domain.TPSet) *TPSetDBModel {
	if domainModel == nil {
		return nil
	}

	return &TPSetDBModel{
		ID:               domainModel.ID,
		CPID:             domainModel.CPID,
		VersionNo:        domainModel.VersionNo,
		Status:           string(domainModel.Status),
		GenerationSource: string(domainModel.GenerationSource),
		GenerationReason: domainModel.GenerationReason,
		GeneratedBy:      domainModel.GeneratedBy,
		AIGenerationID:   domainModel.AIGenerationID,
		ApprovedBy:       domainModel.ApprovedBy,
		ApprovedAt:       domainModel.ApprovedAt,
		CreatedAt:        domainModel.CreatedAt,
		UpdatedAt:        domainModel.UpdatedAt,
	}
}

// MapTPDBModelToDomain converts TPDBModel to domain.TP
func MapTPDBModelToDomain(dbModel *TPDBModel) *domain.TP {
	if dbModel == nil {
		return nil
	}

	return &domain.TP{
		ID:                 dbModel.ID,
		TPSetID:            dbModel.TPSetID,
		SequenceNumber:     dbModel.SequenceNumber,
		CPID:               dbModel.CPID,
		SubjectID:          dbModel.SubjectID,
		PhaseID:            dbModel.PhaseID,
		ElementID:          dbModel.ElementID,
		SubelementID:       dbModel.SubelementID,
		UserID:             dbModel.UserID,
		Status:             domain.WorkflowStatus(dbModel.Status),
		Title:              dbModel.Title,
		LearningObjectives: dbModel.LearningObjectives,
		TimeAllocation:     dbModel.TimeAllocation,
		Prerequisites:      dbModel.Prerequisites,
		EstimatedWeeks:     dbModel.EstimatedWeeks,
		SuccessCriteria:    dbModel.SuccessCriteria,
		VersionNo:          dbModel.VersionNo,
		IsCurrentVersion:   dbModel.IsCurrentVersion,
		ParentVersionID:    dbModel.ParentVersionID,
		CreatedAt:          dbModel.CreatedAt,
		UpdatedAt:          dbModel.UpdatedAt,
	}
}

// MapTPDomainToDBModel converts domain.TP to TPDBModel
func MapTPDomainToDBModel(domainModel *domain.TP) *TPDBModel {
	if domainModel == nil {
		return nil
	}

	return &TPDBModel{
		ID:                 domainModel.ID,
		TPSetID:            domainModel.TPSetID,
		SequenceNumber:     domainModel.SequenceNumber,
		CPID:               domainModel.CPID,
		SubjectID:          domainModel.SubjectID,
		PhaseID:            domainModel.PhaseID,
		ElementID:          domainModel.ElementID,
		SubelementID:       domainModel.SubelementID,
		UserID:             domainModel.UserID,
		Status:             string(domainModel.Status),
		Title:              domainModel.Title,
		LearningObjectives: domainModel.LearningObjectives,
		TimeAllocation:     domainModel.TimeAllocation,
		Prerequisites:      domainModel.Prerequisites,
		EstimatedWeeks:     domainModel.EstimatedWeeks,
		SuccessCriteria:    domainModel.SuccessCriteria,
		VersionNo:          domainModel.VersionNo,
		IsCurrentVersion:   domainModel.IsCurrentVersion,
		ParentVersionID:    domainModel.ParentVersionID,
		CreatedAt:          domainModel.CreatedAt,
		UpdatedAt:          domainModel.UpdatedAt,
	}
}

// MapTPSetDBModelsToDomain converts slice of TPSetDBModel to slice of domain.TPSet
func MapTPSetDBModelsToDomain(dbModels []*TPSetDBModel) []*domain.TPSet {
	if dbModels == nil {
		return nil
	}

	domainModels := make([]*domain.TPSet, 0, len(dbModels))
	for _, dbModel := range dbModels {
		domainModels = append(domainModels, MapTPSetDBModelToDomain(dbModel))
	}
	return domainModels
}

// MapTPDBModelsToDomain converts slice of TPDBModel to slice of domain.TP
func MapTPDBModelsToDomain(dbModels []*TPDBModel) []*domain.TP {
	if dbModels == nil {
		return nil
	}

	domainModels := make([]*domain.TP, 0, len(dbModels))
	for _, dbModel := range dbModels {
		domainModels = append(domainModels, MapTPDBModelToDomain(dbModel))
	}
	return domainModels
}
