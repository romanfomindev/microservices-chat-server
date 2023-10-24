package chat_user

import (
	"context"
	"fmt"
	"github.com/romanfomindev/microservices-chat-server/internal/client/db"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
	"strings"

	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
)

type ChatUser struct {
	db db.Client
}

func NewChatUser(db db.Client) repositories.ChatUser {
	return &ChatUser{
		db: db,
	}
}

func (r *ChatUser) CreateBatch(ctx context.Context, chatUsers models.ChatUser) error {
	valueStrings := make([]string, 0, len(chatUsers.Usernames))
	valueArgs := make([]interface{}, 0, len(chatUsers.Usernames)*2)
	i := 0
	for _, username := range chatUsers.Usernames {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, chatUsers.ChatId)
		valueArgs = append(valueArgs, username)
		i++
	}

	sqlStatement := fmt.Sprintf("INSERT INTO chat_users (chat_id, username) VALUES %s", strings.Join(valueStrings, ","))

	q := db.Query{
		Name:     "chat.Create",
		QueryRaw: sqlStatement,
	}

	_, err := r.db.DB().ExecContext(ctx, q, valueArgs...)

	return err
}
