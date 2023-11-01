package chat

import (
	"context"

	"github.com/romanfomindev/microservices-chat-server/internal/models"
)

func (s *Chat) Create(ctx context.Context, chat models.Chat, chatUsers models.ChatUser) (uint64, error) {
	var chatId uint64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err := s.chatRepo.Create(ctx, chat.Name)
		if err != nil {
			return err
		}

		data := models.ChatUser{
			ChatId:    id,
			Usernames: chatUsers.Usernames,
		}

		err = s.chatUserRepo.CreateBatch(ctx, data)
		if err != nil {
			return err
		}

		chatId = id
		return nil
	})
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
