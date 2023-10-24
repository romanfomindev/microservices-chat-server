package chat

import (
	"context"
	"github.com/romanfomindev/microservices-chat-server/internal/client/db"

	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type Chat struct {
	db db.Client
}

func NewChatRepository(db db.Client) repositories.Chat {
	return &Chat{
		db: db,
	}
}

func (r *Chat) Create(ctx context.Context, name string) (uint64, error) {
	var lastInsertId uint64

	sqlStatement := "INSERT INTO chats (name) VALUES ($1) RETURNING id"

	q := db.Query{
		Name:     "chat.Create",
		QueryRaw: sqlStatement,
	}

	err := r.db.DB().QueryRowContext(ctx, q, name).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (r *Chat) Delete(ctx context.Context, id uint64) error {
	sqlStatement := "DELETE FROM chats  WHERE id = $1"

	q := db.Query{
		Name:     "chat.Create",
		QueryRaw: sqlStatement,
	}

	_, err := r.db.DB().ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
