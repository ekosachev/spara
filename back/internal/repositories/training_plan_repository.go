package repositories

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/models"
	"gorm.io/gorm"
)

type TrainingPlanRepository struct {
	db *gorm.DB
}

func trainingPlanDtoFromModel(model models.TrainingPlan) dto.TrainingPlan {
	return dto.TrainingPlan{
		Model: dto.Model{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Name:   model.Name,
		UserID: model.UserID,
	}
}

func NewTrainingPlanRepository(db *gorm.DB) TrainingPlanRepository {
	return TrainingPlanRepository{db}
}

func (r TrainingPlanRepository) Create(ctx context.Context, trainingPlan dto.TrainingPlan) (*dto.TrainingPlan, error) {
	trainingPlanModel := models.TrainingPlan{
		Name:   trainingPlan.Name,
		UserID: trainingPlan.UserID,
	}
	err := r.db.Model(models.TrainingPlan{}).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&trainingPlanModel).Error
	})

	if err != nil {
		return nil, err
	} else {
		dto := trainingPlanDtoFromModel(trainingPlanModel)
		return &dto, nil
	}
}

func (r TrainingPlanRepository) GetByID(ctx context.Context, id uint) (*dto.TrainingPlan, error) {
	trainingPlanModel, err := gorm.G[models.TrainingPlan](r.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	dto := trainingPlanDtoFromModel(trainingPlanModel)
	return &dto, nil
}
