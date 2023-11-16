package auth

import (
	"context"

	accessDesc "github.com/romanfomindev/microservices-chat-server/pkg/access_v1"
	"google.golang.org/grpc"
)

const ServicePort = 50052

type Auth struct {
	connection *grpc.ClientConn
}

func NewAuth(connection *grpc.ClientConn) *Auth {
	return &Auth{
		connection: connection,
	}
}

func (a *Auth) CheckAccess(ctx context.Context, endpoint string) error {
	cl := accessDesc.NewAccessServiceClient(a.connection)

	_, err := cl.Check(ctx, &accessDesc.CheckRequest{
		EndpointAddress: endpoint,
	})

	return err
}
