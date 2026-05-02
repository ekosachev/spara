package repositories

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, dto dto.User) error {
	return r.db.Model(models.User{}).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(models.User{
			Email:        dto.Email,
			Username:     dto.Username,
			PasswordHash: dto.PasswordHash,
		}).Error
	})
}
