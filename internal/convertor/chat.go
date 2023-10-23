package convertor

import "github.com/romanfomindev/microservices-chat-server/internal/models"
import desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"

func ToUserChatFromDesc(userChat *desc.CreateRequest) models.ChatUser {
	return models.ChatUser{
		Usernames: userChat.GetUsernames(),
	}
}
