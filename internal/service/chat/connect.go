package chat

import (
	"context"

	chat_server "github.com/romanfomindev/microservices-chat/clients/grpc/chat-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ChatService) Connect(ctx context.Context, accessToken, chatName string, username []string) error {
	conn, err := grpc.Dial(
		s.cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	client := chat_server.NewChatServer(conn)

	chatId, err := client.Create(ctx, accessToken, chatName, username)
	if err != nil {
		return err
	}

	return nil
}
