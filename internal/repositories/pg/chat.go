package pg

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type Chat struct {
	conn *pgx.Conn
}

func NewChatRepository(conn *pgx.Conn) (repositories.Chat, error) {
	return &Chat{
		conn: conn,
	}, nil
}

func (r *Chat) Create(ctx context.Context, name string) (uint64, error) {
	var lastInsertId uint64

	sqlStatement := "INSERT INTO chats (name) VALUES ($1) RETURNING id"

	err := r.conn.QueryRow(ctx, sqlStatement, name).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (r *Chat) Delete(ctx context.Context, id uint64) error {
	sqlStatement := "DELETE FROM chats  WHERE id = $1"

	_, err := r.conn.Exec(ctx, sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
