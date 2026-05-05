package services

import (
	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

type ExcerciseService struct {
	BaseServise[dto.Excercise]
	repo repositories.ExcerciseRepository
}

func NewExcerciseRepository(repo repositories.ExcerciseRepository) ExcerciseService {
	return ExcerciseService{
		BaseServise: BaseServise[dto.Excercise]{repo: repo},
		repo:        repo,
	}
}
