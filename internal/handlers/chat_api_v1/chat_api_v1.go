package handlers

import (
	"context"
	"log"

	"github.com/romanfomindev/microservices-chat-server/internal/convertor"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatApiService struct {
	desc.UnimplementedChatApiServer
	Service services.ChatService
}

func NewChatService(service services.ChatService) *ChatApiService {
	return &ChatApiService{
		Service: service,
	}
}

func (s *ChatApiService) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("usernames: %+v", request.GetUsernames())

	chatModel := models.Chat{
		Name: request.GetChatName(),
	}
	chatId, err := s.Service.Create(ctx, chatModel, convertor.ToUserChatFromDesc(request))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}

func (s *ChatApiService) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("ID: %+v", request.GetId())
	err := s.Service.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatApiService) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("From: %s, Text: %s, Timestamp: %+v", request.GetFrom(), request.GetText(), request.GetTimestamp())

	return &emptypb.Empty{}, nil
}
