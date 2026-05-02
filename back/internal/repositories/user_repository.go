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

func (r *UserRepository) Create(ctx context.Context, user dto.User) (*dto.User, error) {
	userModel := models.User{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}
	err := r.db.Model(models.User{}).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&userModel).Error
	})

	if err != nil {
		return nil, err
	} else {
		return &dto.User{
			Model: dto.Model{
				ID:        userModel.ID,
				CreatedAt: userModel.CreatedAt,
				UpdatedAt: userModel.UpdatedAt,
			},
			Email:        userModel.Email,
			Username:     userModel.Username,
			PasswordHash: userModel.PasswordHash,
		}, nil
	}
}
