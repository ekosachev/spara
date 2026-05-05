package models

import "gorm.io/gorm"

type TrainingPlan struct {
	gorm.Model
	Name       string
	UserID     uint
	Excercices []*Excercise `gorm:"many2many:training_plan_excercises"`
}

type TrainingPlanExcercises struct {
	TrainingPlanID int `gorm:"primaryKey"`
	ExcerciseID    int `gorm:"primaryKey"`
	Order          int
}
