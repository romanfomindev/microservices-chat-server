package chat

import "context"

func (m *ChatService) Delete(ctx context.Context, chatId uint64) error {
	err := m.chatRepo.Delete(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}
