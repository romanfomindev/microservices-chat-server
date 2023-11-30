package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	handlers "github.com/romanfomindev/microservices-chat-server/internal/handlers/chat_api_v1"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	serviceMock "github.com/romanfomindev/microservices-chat-server/internal/services/mocks"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteHandler(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) services.ChatService
	type implimentationMockFunc func(mc *minimock.Controller) services.Implementation

	type args struct {
		ctx     context.Context
		request *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Uint64()

		req = &desc.DeleteRequest{
			Id: id,
		}

		res        = &emptypb.Empty{}
		errService = errors.New("service error")
	)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		chatServiceMock    chatServiceMockFunc
		implimentationMock implimentationMockFunc
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
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
			implimentationMock: func(mc *minimock.Controller) services.Implementation {
				mock := serviceMock.NewImplementationMock(t)
				mock.ConnectChatMock.Return(nil)

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
				mock.DeleteMock.Expect(ctx, id).Return(errService)

				return mock
			},
			implimentationMock: func(mc *minimock.Controller) services.Implementation {
				mock := serviceMock.NewImplementationMock(t)
				mock.ConnectChatMock.Return(nil)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			implimentationMock := tt.implimentationMock(mc)
			handler := handlers.NewChatService(chatServiceMock, implimentationMock)
			response, err := handler.Delete(ctx, req)

			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
