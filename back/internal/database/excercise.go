package database

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

func seedExcercises(repo *repositories.ExcerciseRepository, ctx context.Context) error {
	excercises := []dto.Excercise{
		dto.Excercise{Name: "Standing Barbell Curl"},
		dto.Excercise{Name: "Incline Dumbbell Curl"},
		dto.Excercise{Name: "Dumbbell Bench Press"},
		dto.Excercise{Name: "Incline Bench Press"},
		dto.Excercise{Name: "Bent Over Row"},
		dto.Excercise{Name: "Seated Cable Row"},
	}

	for _, e := range excercises {
		existing, _ := repo.GetByName(ctx, e.Name)
		if existing != nil {
			continue
		}
		_, err := repo.Create(ctx, e)
		if err != nil {
			return err
		}
	}
	return nil
}
