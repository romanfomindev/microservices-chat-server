package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/romanfomindev/microservices-chat-server/internal/convertor"
	"github.com/romanfomindev/microservices-chat-server/internal/models"
	"github.com/romanfomindev/microservices-chat-server/internal/services"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatApiService struct {
	desc.UnimplementedChatApiServer
	Service       services.ChatService
	StreamService services.Implementation
}

func NewChatService(service services.ChatService, streamService services.Implementation) *ChatApiService {
	return &ChatApiService{
		Service:       service,
		StreamService: streamService,
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
	email := ctx.Value("email").(string)
	checkUserInChat := s.Service.CheckUserInChat(ctx, request.ChatId, email)
	if !checkUserInChat {
		return nil, errors.New("user not in chat")
	}

	msg := request.GetMessage()
	msg.From = email
	err := s.StreamService.SendMessage(request.GetChatId(), msg)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatApiService) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatApi_ConnectChatServer) error {
	chatId := req.GetChatId()
	email := stream.Context().Value("email").(string)

	return s.StreamService.ConnectChat(chatId, email, stream)
}
