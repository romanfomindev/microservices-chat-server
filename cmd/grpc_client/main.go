package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	accessDesc "github.com/romanfomindev/microservices-chat-server/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var accessToken = flag.String("a", "", "access token")
var servicePort = 50052

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := accessDesc.NewAccessServiceClient(conn)

	_, err = cl.Check(ctx, &accessDesc.CheckRequest{
		EndpointAddress: "/api/v1/send_message",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Access granted")
}
