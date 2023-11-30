package clients

import "context"

type Auth interface {
	Login(ctx context.Context, email, password string) (string, error)
	AccessToken(ctx context.Context, refreshToken string) (string, error)
	CheckAccess(ctx context.Context, endpoint string) error
}

type ChatServer interface {
	Create(ctx context.Context, name string) (uint64, error)
	Delete(ctx context.Context, chatId uint64) error
	SendMessage(ctx context.Context, chatId uint64, message string) error
}
