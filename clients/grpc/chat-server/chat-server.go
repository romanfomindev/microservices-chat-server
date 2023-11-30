package chat_server

import (
	"context"

	chatserverDesc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ChatServer struct {
	connection *grpc.ClientConn
}

func NewChatServer(connection *grpc.ClientConn) *ChatServer {
	return &ChatServer{
		connection: connection,
	}
}

func (c *ChatServer) Create(ctx context.Context, accessToken, name string, usernames []string) (uint64, error) {
	cl := chatserverDesc.NewChatApiClient(c.connection)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	r, err := cl.Create(ctx, &chatserverDesc.CreateRequest{
		ChatName:  name,
		Usernames: usernames,
	})
	if err != nil {
		return 0, err
	}

	return r.GetId(), nil
}

func (c *ChatServer) Delete(ctx context.Context, accessToken string, id uint64) error {
	cl := chatserverDesc.NewChatApiClient(c.connection)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := cl.Delete(ctx, &chatserverDesc.DeleteRequest{
		Id: id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatServer) SendMessage(ctx context.Context, accessToken, from, text string) error {
	cl := chatserverDesc.NewChatApiClient(c.connection)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := cl.SendMessage(ctx, &chatserverDesc.SendMessageRequest{
		From: from,
		Text: text,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *ChatServer) Connect(ctx context.Context, accessToken, from, text string) error {
	cl := chatserverDesc.NewChatApiClient(c.connection)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	//_, err :=

	//if err != nil {
	//	return err
	//}

	return nil
}
