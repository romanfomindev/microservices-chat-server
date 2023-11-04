package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	handlers "github.com/romanfomindev/microservices-chat-server/internal/handlers/chat_api_v1"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"github.com/stretchr/testify/require"

	serviceMock "github.com/romanfomindev/microservices-chat-server/internal/services/mocks"
)

func TestCreateHandler(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) services.ChatService

	type args struct {
		ctx     context.Context
		request *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userNames = []string{gofakeit.FirstName(), gofakeit.FirstName(), gofakeit.FirstName()}
		id        = gofakeit.Uint64()
		name      = gofakeit.Name()

		req = &desc.CreateRequest{
			ChatName:  name,
			Usernames: userNames,
		}

		res = &desc.CreateResponse{
			Id: id,
		}

		errService = errors.New("error service")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:     ctx,
				request: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) services.ChatService {
				mock := serviceMock.NewChatServiceMock(t)
				mock.CreateMock.Return(id, nil)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:     ctx,
				request: req,
			},
			want: nil,
			err:  errService,
			chatServiceMock: func(mc *minimock.Controller) services.ChatService {
				mock := serviceMock.NewChatServiceMock(t)
				mock.CreateMock.Return(0, errService)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			handler := handlers.NewChatService(chatServiceMock)
			response, err := handler.Create(ctx, req)

			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
