package chat

import (
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	"github.com/romanfomindev/platform_common/pkg/db"
)

type Chat struct {
	chatRepo     repositories.Chat
	chatUserRepo repositories.ChatUser
	txManager    db.TxManager
}

func NewChatService(chatRepo repositories.Chat, chatUserRepo repositories.ChatUser, txManager db.TxManager) services.ChatService {
	return &Chat{
		chatRepo:     chatRepo,
		chatUserRepo: chatUserRepo,
		txManager:    txManager,
	}
}
