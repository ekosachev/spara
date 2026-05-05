package dto

type TrainingPlan struct {
	Model
	Name       string
	UserID     uint
	Excercises []TrainingPlanExcercise
}

type TrainingPlanExcercise struct {
	Order     int
	Excercise Excercise
}

type TrainingPlanExcerciseRequest struct {
	Order       int `json:"order" binding:"required"`
	ExcerciseID int `json:"excercise_id" binding:"required"`
}

type CreateTrainingPlanRequest struct {
	Name       string                         `json:"name" binding:"required"`
	Excercises []TrainingPlanExcerciseRequest `json:"excercises" binding:"omitempty"`
}

type TrainingPlanResponse struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
