package models

import "gorm.io/gorm"

type Excercise struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
}
