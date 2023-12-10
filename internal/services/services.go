package services

import (
	"context"

	"github.com/romanfomindev/microservices-chat-server/internal/models"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
)

type ChatService interface {
	Create(ctx context.Context, chat models.Chat, chatUsers models.ChatUser) (uint64, error)
	Delete(ctx context.Context, chatId uint64) error
	CheckUserInChat(ctx context.Context, chatId uint64, email string) bool
}

type Implementation interface {
	ConnectChat(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) error
	SendMessage(chatId uint64, message *desc.Message) error
}
