package services

import (
	"context"

	"github.com/ekosachev/spara/internal/repositories"
)

type BaseServise[T any] struct {
	repo repositories.BaseRepository[T]
}

func (s *BaseServise[T]) Create(ctx context.Context, dto T) (*T, error) {
	return s.repo.Create(ctx, dto)
}
