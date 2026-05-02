package services

import (
	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/repositories"
)

type UserService struct {
	BaseServise[dto.User]
}

func NewUserService(repo repositories.BaseRepository[dto.User]) UserService {
	return UserService{
		BaseServise: BaseServise[dto.User]{repo: repo},
	}
}
