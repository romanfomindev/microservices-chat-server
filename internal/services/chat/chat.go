package chat

import (
	"github.com/romanfomindev/microservices-chat-server/internal/client/db"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type ChatService struct {
	chatRepo     repositories.Chat
	chatUserRepo repositories.ChatUser
	txManager    db.TxManager
}

func NewChatService(chatRepo repositories.Chat, chatUserRepo repositories.ChatUser, txManager db.TxManager) *ChatService {
	return &ChatService{
		chatRepo:     chatRepo,
		chatUserRepo: chatUserRepo,
		txManager:    txManager,
	}
}
