package handlers

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"github.com/romanfomindev/microservices-chat-server/internal/managers"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatApiService struct {
	desc.UnimplementedChatApiServer
	Manager *managers.ChatManager
}

func NewChatService(manager *managers.ChatManager) desc.ChatApiServer {
	return &ChatApiService{
		Manager: manager,
	}
}

func (s *ChatApiService) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("usernames: %+v", request.GetUsernames())

	chatName := gofakeit.BeerName()
	chatId, err := s.Manager.Create(ctx, chatName, request.GetUsernames())
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}

func (s *ChatApiService) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("ID: %+v", request.GetId())
	err := s.Manager.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatApiService) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("From: %s, Text: %s, Timestamp: %+v", request.GetFrom(), request.GetText(), request.GetTimestamp())

	return &emptypb.Empty{}, nil
}
