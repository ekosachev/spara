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

	trainingPlanExcerciseModels := make([]models.TrainingPlanExcercises, len(trainingPlan.Excercises))

	for i, v := range trainingPlan.Excercises {
		trainingPlanExcerciseModel := models.TrainingPlanExcercises{
			TrainingPlanID: int(trainingPlanModel.ID),
			ExcerciseID:    v.ExcerciseID,
			Order:          v.Order,
		}

		err := gorm.G[models.TrainingPlanExcercises](r.db).Create(ctx, &trainingPlanExcerciseModel)

		if err != nil {
			return nil, err
		}

		trainingPlanExcerciseModels[i] = trainingPlanExcerciseModel
	}

	if err != nil {
		return nil, err
	}

	trainingPlanDTO := trainingPlanDtoFromModel(trainingPlanModel)

	// trainingPlanDTO.Excercises = trainingPlan.Excercises
	excercises, err := gorm.G[models.TrainingPlanExcercises](r.db).Where("training_plan_id = ?", trainingPlanModel.ID).Find(ctx)

	if err != nil {
		return nil, err
	}

	trainingPlanDTO.Excercises = make([]dto.TrainingPlanExcercise, len(excercises))
	for i, e := range excercises {
		trainingPlanDTO.Excercises[i] = dto.TrainingPlanExcercise{
			ExcerciseID: e.ExcerciseID,
			Order:       e.Order,
		}
	}

	return &trainingPlanDTO, nil
}

func (r TrainingPlanRepository) GetByID(ctx context.Context, id uint) (*dto.TrainingPlan, error) {
	trainingPlanModel, err := gorm.G[models.TrainingPlan](r.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	dto := trainingPlanDtoFromModel(trainingPlanModel)
	return &dto, nil
}
