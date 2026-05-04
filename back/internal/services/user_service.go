package services

import (
	"context"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
	"github.com/ekosachev/spara/internal/utils"
)

type UserService struct {
	BaseServise[dto.User]
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{
		BaseServise: BaseServise[dto.User]{repo: repo},
		repo:        repo,
	}
}

func (s *UserService) Create(ctx context.Context, user dto.User) (*dto.User, error) {
	hashedPassword, err := utils.HashPassword(user.PasswordHash)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, dto.User{
		Model:        user.Model,
		Email:        user.Email,
		PasswordHash: hashedPassword,
		Username:     user.Username,
	})
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	return s.repo.GetByEmail(ctx, email)
}
