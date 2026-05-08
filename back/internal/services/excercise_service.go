package services

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

type ExcerciseService struct {
	BaseServise[dto.Excercise]
	repo repositories.ExcerciseRepository
}

func NewExcerciseService(repo repositories.ExcerciseRepository) ExcerciseService {
	return ExcerciseService{
		BaseServise: BaseServise[dto.Excercise]{repo: repo},
		repo:        repo,
	}
}

func (s *ExcerciseService) GetAll(ctx context.Context) ([]dto.Excercise, error) {
	return s.repo.GetAll(ctx)
}
