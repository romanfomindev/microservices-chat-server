package chat

import (
	"context"
)

func (s *Chat) CheckUserInChat(ctx context.Context, chatId uint64, email string) bool {
	return s.chatUserRepo.FindUserInChat(ctx, chatId, email)
}
