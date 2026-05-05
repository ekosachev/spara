package database

import (
	"context"

	"github.com/ekosachev/spara/internal/repositories"
)

func Seed(excerciseRepo *repositories.ExcerciseRepository) error {
	ctx := context.Background()

	if err := seedExcercises(excerciseRepo, ctx); err != nil {
		return err
	}
	return nil
}
