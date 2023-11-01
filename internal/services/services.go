package services

import (
	"context"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
)

type ChatService interface {
	Create(ctx context.Context, chat models.Chat, chatUsers models.ChatUser) (uint64, error)
	Delete(ctx context.Context, chatId uint64) error
}
