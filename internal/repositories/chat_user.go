package repositories

import "context"

type ChatUser interface {
	CreateBatch(ctx context.Context, chatId uint64, usernames []string) error
}
