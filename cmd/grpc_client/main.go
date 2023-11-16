package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/romanfomindev/microservices-chat-server/internal/client/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var accessToken = flag.String("a", "", "access token")

func main() {
	flag.Parse()

	ctx := context.Background()

	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", auth.ServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to dial GRPC client: %v", err)
		return
	}

	client := auth.NewAuth(conn)

	err = client.CheckAccess(ctx, "/api/v1/send_message")
	if err != nil {
		log.Fatal("Access denied")
	}

	log.Fatal("Access granted")
}
