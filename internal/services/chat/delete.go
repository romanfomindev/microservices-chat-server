package chat

import "context"

func (s *ChatService) Delete(ctx context.Context, chatId uint64) error {
	err := s.chatRepo.Delete(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}
