package services

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

type TrainingPlanServcie struct {
	BaseServise[dto.TrainingPlan]
	repo repositories.TrainingPlanRepository
}

func NewTrainingPlanService(repo repositories.TrainingPlanRepository) TrainingPlanServcie {
	return TrainingPlanServcie{
		BaseServise: BaseServise[dto.TrainingPlan]{repo: repo},
		repo:        repo,
	}
}

func (s *TrainingPlanServcie) Create(ctx context.Context, trainingPlan dto.TrainingPlan) (*dto.TrainingPlan, error) {

	return s.repo.Create(ctx, trainingPlan)
}
