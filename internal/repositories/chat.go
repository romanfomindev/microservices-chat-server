package repositories

import "context"

type Chat interface {
	Create(ctx context.Context, name string) (uint64, error)
	Delete(ctx context.Context, id uint64) error
}
