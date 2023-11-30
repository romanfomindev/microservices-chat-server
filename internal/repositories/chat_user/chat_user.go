package chat_user

import (
	"context"
	"fmt"
	"strings"

	"github.com/romanfomindev/microservices-chat-server/internal/models"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
	"github.com/romanfomindev/platform_common/pkg/db"
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

func (r *ChatUser) FindUserInChat(ctx context.Context, chatId uint64, email string) bool {
	sqlStatement := `
		SELECT EXISTS (
			SELECT 1 FROM chat_users WHERE chat_id = $1 AND username = $2
		)
	`
	q := db.Query{
		Name:     "chat_users.FindUserInChat",
		QueryRaw: sqlStatement,
	}

	var exist bool

	err := r.db.DB().QueryRowContext(ctx, q, chatId, email).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}
