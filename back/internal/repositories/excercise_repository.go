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

func dtoFromModel(model models.Excercise) dto.Excercise {
	return dto.Excercise{
		Model: dto.Model{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:        model.Name,
		Description: model.Description,
	}
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
		dto := dtoFromModel(exceciseModel)
		return &dto, nil
	}
}

func (r ExcerciseRepository) GetByName(ctx context.Context, name string) (*dto.Excercise, error) {
	excercise, err := gorm.G[models.Excercise](r.db).Where(models.Excercise{Name: name}).First(ctx)

	if err != nil {
		return nil, err
	} else {
		dto := dtoFromModel(excercise)
		return &dto, nil
	}
}
