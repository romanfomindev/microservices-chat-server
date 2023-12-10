package app

import (
	"context"
	"log"

	"github.com/romanfomindev/microservices-chat-server/internal/config"
	"github.com/romanfomindev/microservices-chat-server/internal/config/env"
	handlers "github.com/romanfomindev/microservices-chat-server/internal/handlers/chat_api_v1"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories"
	chatRepo "github.com/romanfomindev/microservices-chat-server/internal/repositories/chat"
	"github.com/romanfomindev/microservices-chat-server/internal/repositories/chat_user"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	"github.com/romanfomindev/microservices-chat-server/internal/services/chat"
	streamService "github.com/romanfomindev/microservices-chat-server/internal/services/stream"
	"github.com/romanfomindev/platform_common/pkg/closer"
	"github.com/romanfomindev/platform_common/pkg/db"
	"github.com/romanfomindev/platform_common/pkg/db/pg"
	"github.com/romanfomindev/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	dbClient           db.Client
	pgConfig           config.PGConfig
	txManager          db.TxManager
	grpcConfig         config.GRPCConfig
	chatRepository     repositories.Chat
	chatUserRepository repositories.ChatUser
	chatService        services.ChatService
	chatHandlers       *handlers.ChatApiService
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN(), s.pgConfig.Timeout())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repositories.Chat {
	if s.chatUserRepository == nil {
		repo := chatRepo.NewChatRepository(s.DBClient(ctx))
		s.chatRepository = repo
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatUserRepository(ctx context.Context) repositories.ChatUser {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chat_user.NewChatUser(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) services.ChatService {
	if s.chatService == nil {
		s.chatService = chat.NewChatService(s.ChatRepository(ctx), s.ChatUserRepository(ctx), s.TxManager(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) ChatHandlers(ctx context.Context) *handlers.ChatApiService {
	if s.chatHandlers == nil {
		s.chatHandlers = handlers.NewChatService(s.ChatService(ctx), streamService.NewImplementation())
	}
	return s.chatHandlers
}
