package auth

import (
	"context"
	"fmt"
	accessDesc "github.com/romanfomindev/microservices-chat-server/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func CheckAccess(ctx context.Context, endpoint string, accessToken string) bool {
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to dial GRPC client: %v", err)
		return false
	}

	cl := accessDesc.NewAccessServiceClient(conn)

	_, err = cl.Check(ctx, &accessDesc.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		log.Printf("CheckAccess error: %v", err)
		return false
	}

	return true
}
