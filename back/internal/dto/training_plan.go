package dto

type TrainingPlan struct {
	Model
	Name   string
	UserID uint
}

type CreateTrainingPlanRequest struct {
	Name string `json:"name" binding:"required"`
}

type TrainingPlanResponse struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
