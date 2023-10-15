package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/romanfomindev/microservices-chat-server/internal/config"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type ChatUser struct {
	conn *pgx.Conn
}

func NewChatUser(ctx context.Context, cfg config.PGConfig) (repositories.ChatUser, error) {
	conn, err := pgx.Connect(ctx, cfg.DSN())
	if err != nil {
		return nil, err
	}

	return &ChatUser{
		conn: conn,
	}, nil
}

func (r *ChatUser) CreateBatch(ctx context.Context, chatId uint64, usernames []string) error {
	valueStrings := make([]string, 0, len(usernames))
	valueArgs := make([]interface{}, 0, len(usernames)*2)
	i := 0
	for _, username := range usernames {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, chatId)
		valueArgs = append(valueArgs, username)
		i++
	}
	statement := fmt.Sprintf("INSERT INTO chat_users (chat_id, username) VALUES %s", strings.Join(valueStrings, ","))

	_, err := r.conn.Exec(ctx, statement, valueArgs...)

	return err
}
