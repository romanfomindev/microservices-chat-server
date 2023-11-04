package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-chat-server/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-chat-server/internal/services/chat"
	"github.com/romanfomindev/platform_common/pkg/db"
	"github.com/stretchr/testify/require"
)

type txManagerMock struct {
}

func (tx txManagerMock) ReadCommitted(ctx context.Context, f db.Handler) error {
	return f(ctx)
}

func TestCreate(t *testing.T) {
	type chatRepositoryMockFunc func(mc *minimock.Controller) repositories.Chat
	type chatUserRepositoryMockFunc func(mc *minimock.Controller) repositories.ChatUser

	type args struct {
		ctx      context.Context
		chat     models.Chat
		chatUser models.ChatUser
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatName  = gofakeit.Name()
		chatId    = gofakeit.Uint64()
		userNames = []string{gofakeit.FirstName(), gofakeit.FirstName(), gofakeit.FirstName()}
		chatModel = models.Chat{
			Name: chatName,
		}

		chatUser = models.ChatUser{
			Usernames: userNames,
		}
		errCreateChat     = errors.New("error create chat")
		errAddUsersToChat = errors.New("error add users")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                   string
		args                   args
		want                   uint64
		err                    error
		chatRepositoryMock     chatRepositoryMockFunc
		chatUserRepositoryMock chatUserRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:      ctx,
				chat:     chatModel,
				chatUser: chatUser,
			},
			want: chatId,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repositories.Chat {
				mock := repoMock.NewChatMock(mc)
				mock.CreateMock.Expect(ctx, chatName).Return(chatId, nil)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repositories.ChatUser {
				mock := repoMock.NewChatUserMock(mc)
				chatUser.ChatId = chatId
				mock.CreateBatchMock.Return(nil)

				return mock
			},
		},
		{
			name: "error create chat case",
			args: args{
				ctx:      ctx,
				chat:     chatModel,
				chatUser: chatUser,
			},
			want: 0,
			err:  errCreateChat,
			chatRepositoryMock: func(mc *minimock.Controller) repositories.Chat {
				mock := repoMock.NewChatMock(mc)
				mock.CreateMock.Expect(ctx, chatName).Return(0, errCreateChat)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repositories.ChatUser {
				mock := repoMock.NewChatUserMock(mc)

				return mock
			},
		},
		{
			name: "error add user to chat case",
			args: args{
				ctx:      ctx,
				chat:     chatModel,
				chatUser: chatUser,
			},
			want: 0,
			err:  errAddUsersToChat,
			chatRepositoryMock: func(mc *minimock.Controller) repositories.Chat {
				mock := repoMock.NewChatMock(mc)
				mock.CreateMock.Expect(ctx, chatName).Return(chatId, nil)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repositories.ChatUser {
				mock := repoMock.NewChatUserMock(mc)
				chatUser.ChatId = chatId
				mock.CreateBatchMock.Expect(ctx, chatUser).Return(errAddUsersToChat)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatRepoMock := tt.chatRepositoryMock(mc)
			chatUserRepoMock := tt.chatUserRepositoryMock(mc)
			txManager := txManagerMock{}
			service := chat.NewChatService(chatRepoMock, chatUserRepoMock, txManager)
			newId, err := service.Create(tt.args.ctx, tt.args.chat, tt.args.chatUser)
			require.Equal(t, tt.want, newId)
			require.Equal(t, tt.err, err)
		})
	}

}
