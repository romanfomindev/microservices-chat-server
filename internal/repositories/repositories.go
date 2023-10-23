package repositories

import (
	"context"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
)

type Chat interface {
	Create(ctx context.Context, name string) (uint64, error)
	Delete(ctx context.Context, id uint64) error
}

type ChatUser interface {
	CreateBatch(ctx context.Context, chatUsers models.ChatUser) error
}
