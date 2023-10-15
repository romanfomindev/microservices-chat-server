package main

import (
	"context"
	"log"
	"net"

	"github.com/romanfomindev/microservices-chat-server/internal/config"
	"github.com/romanfomindev/microservices-chat-server/internal/config/env"
	handlers "github.com/romanfomindev/microservices-chat-server/internal/handlers/chat_api_v1"
	"github.com/romanfomindev/microservices-chat-server/internal/managers"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories/pg"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string = ".env"

func main() {
	ctx := context.Background()
	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	chatRepo, err := pg.NewChatRepository(ctx, pgConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	chatUserRepo, err := pg.NewChatUser(ctx, pgConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	manager := managers.NewChatManager(chatRepo, chatUserRepo)

	chatApiService := handlers.NewChatService(manager)
	desc.RegisterChatApiServer(s, chatApiService)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
