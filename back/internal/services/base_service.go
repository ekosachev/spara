package services

import (
	"context"

	"github.com/ekosachev/spara/internal/repositories"
)

type BaseServise[T any] struct {
	Repo repositories.BaseRepository[T]
}

func (s *BaseServise[T]) Create(ctx context.Context, dto T) error {
	return s.Repo.Create(ctx, dto)
}
