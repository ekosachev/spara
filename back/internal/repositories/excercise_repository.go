package repositories

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/models"
	"gorm.io/gorm"
)

type ExcerciseRepository struct {
	db *gorm.DB
}

func NewExcerciseRepository(db *gorm.DB) ExcerciseRepository {
	return ExcerciseRepository{db}
}

func (r ExcerciseRepository) Create(ctx context.Context, excercise dto.Excercise) (*dto.Excercise, error) {
	exceciseModel := models.Excercise{
		Name:        excercise.Name,
		Description: excercise.Description,
	}
	err := r.db.Model(models.Excercise{}).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&exceciseModel).Error
	})

	if err != nil {
		return nil, err
	} else {
		return &dto.Excercise{
			Model: dto.Model{
				ID:        exceciseModel.ID,
				CreatedAt: exceciseModel.CreatedAt,
				UpdatedAt: exceciseModel.UpdatedAt,
			},
			Name:        exceciseModel.Name,
			Description: exceciseModel.Description,
		}, nil
	}
}
