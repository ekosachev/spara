package models

import (
	"github.com/ekosachev/spara/internal/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Username     string
	PasswordHash string
}

func (u *User) DTO() dto.User {
	return dto.User{
		Model: dto.Model{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
