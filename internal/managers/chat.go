package managers

import (
	"context"

	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type ChatManager struct {
	chatRepo     repositories.Chat
	chatUserRepo repositories.ChatUser
}

func NewChatManager(chatRepo repositories.Chat, chatUserRepo repositories.ChatUser) *ChatManager {
	return &ChatManager{
		chatRepo:     chatRepo,
		chatUserRepo: chatUserRepo,
	}
}

func (m *ChatManager) Create(ctx context.Context, chatName string, usernames []string) (uint64, error) {

	/** TODO в дальнейшем транзакции обернуть */
	chatId, err := m.chatRepo.Create(ctx, chatName)
	if err != nil {
		return 0, err
	}

	err = m.chatUserRepo.CreateBatch(ctx, chatId, usernames)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}

func (m *ChatManager) Delete(ctx context.Context, chatId uint64) error {
	err := m.chatRepo.Delete(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}
