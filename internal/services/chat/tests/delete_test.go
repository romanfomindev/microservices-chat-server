package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-chat-server/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-chat-server/internal/services/chat"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type chatRepositoryMockFunc func(mc *minimock.Controller) repositories.Chat
	type chatUserRepositoryMockFunc func(mc *minimock.Controller) repositories.ChatUser

	type args struct {
		ctx    context.Context
		chatId uint64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Uint64()
		repoErr = errors.New("repo error")
	)

	tests := []struct {
		name                   string
		args                   args
		err                    error
		chatRepositoryMock     chatRepositoryMockFunc
		chatUserRepositoryMock chatUserRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				chatId: id,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repositories.Chat {
				mock := repoMock.NewChatMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repositories.ChatUser {
				mock := repoMock.NewChatUserMock(mc)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:    ctx,
				chatId: id,
			},
			err: repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repositories.Chat {
				mock := repoMock.NewChatMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repositories.ChatUser {
				mock := repoMock.NewChatUserMock(mc)

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
			err := service.Delete(tt.args.ctx, tt.args.chatId)

			require.Equal(t, tt.err, err)
		})
	}
}
