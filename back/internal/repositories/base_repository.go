package repositories

import "context"

type BaseRepository[T any] interface {
	Create(ctx context.Context, dto T) (*T, error)
	// GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id uint) (*T, error)
	// Update(ctx context.Context, id uint, dto T) (int, error)
	// Delete(ctx context.Context, id uint) (int, error)
}
