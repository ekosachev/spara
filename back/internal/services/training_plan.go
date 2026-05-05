package services

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

type TrainingPlanServcie struct {
	BaseServise[dto.TrainingPlan]
	repo          repositories.TrainingPlanRepository
	userRepo      repositories.UserRepository
	excerciseRepo repositories.ExcerciseRepository
}

func NewTrainingPlanService(repo repositories.TrainingPlanRepository, userRepo repositories.UserRepository, excerciseRepo repositories.ExcerciseRepository) TrainingPlanServcie {
	return TrainingPlanServcie{
		BaseServise:   BaseServise[dto.TrainingPlan]{repo: repo},
		repo:          repo,
		userRepo:      userRepo,
		excerciseRepo: excerciseRepo,
	}
}

func (s *TrainingPlanServcie) Create(ctx context.Context, trainingPlan dto.TrainingPlan) (*dto.TrainingPlan, error) {
	if _, err := s.userRepo.GetByID(ctx, trainingPlan.UserID); err != nil {
		return nil, err
	}

	for _, v := range trainingPlan.Excercises {
		excerciseID := v.ExcerciseID
		_, err := s.excerciseRepo.GetByID(ctx, uint(excerciseID))
		if err != nil {
			return nil, err
		}
	}

	return s.repo.Create(ctx, trainingPlan)
}
