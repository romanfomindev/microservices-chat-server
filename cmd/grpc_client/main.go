package main

import (
	"context"
	"flag"
	"github.com/romanfomindev/microservices-chat-server/internal/client/grpc/auth"
	"log"
)

var accessToken = flag.String("a", "", "access token")

func main() {
	flag.Parse()

	ctx := context.Background()

	if auth.CheckAccess(ctx, "/api/v1/send_message", *accessToken) {
		log.Fatal("Access granted")
	} else {
		log.Fatal("Access denied")
	}
}
