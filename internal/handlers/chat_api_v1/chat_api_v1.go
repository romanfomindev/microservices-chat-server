package handlers

import (
	"context"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type ChatApiService struct {
	desc.UnimplementedChatApiServer
}

func (s *ChatApiService) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	id := int64(123)

	log.Printf("usernames: %+v", request.GetUsernames())

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *ChatApiService) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("ID: %+v", request.GetId())

	return &emptypb.Empty{}, nil
}

func (s *ChatApiService) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("From: %s, Text: %s, Timestamp: %+v", request.GetFrom(), request.GetText(), request.GetTimestamp())

	return &emptypb.Empty{}, nil
}
