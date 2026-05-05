package dto

type Excercise struct {
	Model
	Name        string
	Description string
}

type CreateExcerciseRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type ExcerciseResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
