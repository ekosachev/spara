package models

import "gorm.io/gorm"

type TrainingPlan struct {
	gorm.Model
	Name   string
	UserID uint
}
