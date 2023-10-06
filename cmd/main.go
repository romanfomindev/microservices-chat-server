package main

import (
	"fmt"
	handlers "github.com/romanfomindev/microservices-chat-server/internal/handlers/chat_api_v1"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chatApiService := &handlers.ChatApiService{}
	desc.RegisterChatApiServer(s, chatApiService)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
