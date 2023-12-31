package chat

import "context"

func (s *Chat) Delete(ctx context.Context, chatId uint64) error {
	err := s.chatRepo.Delete(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}
