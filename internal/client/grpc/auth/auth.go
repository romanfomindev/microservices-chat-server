package auth

import (
	"context"

	accessDesc "github.com/romanfomindev/microservices-chat-server/pkg/access_v1"
	"google.golang.org/grpc"
)

const ServicePort = 50052

type Auth struct {
	client accessDesc.AccessServiceClient
}

func NewAuth(connection *grpc.ClientConn) *Auth {
	return &Auth{
		client: accessDesc.NewAccessServiceClient(connection),
	}
}

func (a *Auth) CheckAccess(ctx context.Context, endpoint string) error {
	_, err := a.client.Check(ctx, &accessDesc.CheckRequest{
		EndpointAddress: endpoint,
	})

	return err
}
