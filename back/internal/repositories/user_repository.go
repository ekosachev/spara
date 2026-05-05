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

func userDtoFromModel(model models.User) dto.User {
	return dto.User{
		Model: dto.Model{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Email:        model.Email,
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
	}
}

func (r UserRepository) Create(ctx context.Context, user dto.User) (*dto.User, error) {
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
		dto := userDtoFromModel(userModel)
		return &dto, nil
	}
}

func (r UserRepository) GetByEmail(ctx context.Context, email string) (*dto.User, error) {

	userModel, err := gorm.G[models.User](r.db).Where(models.User{Email: email}).First(ctx)

	if err != nil {
		return nil, err
	}

	dto := userDtoFromModel(userModel)
	return &dto, nil
}

func (r UserRepository) GetByID(ctx context.Context, id uint) (*dto.User, error) {
	userModel, err := gorm.G[models.User](r.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	dto := userDtoFromModel(userModel)
	return &dto, nil
}
