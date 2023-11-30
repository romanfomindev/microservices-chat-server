package stream

import (
	"log"
	"sync"

	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
)

type Chat struct {
	streams map[string]desc.ChatApi_ConnectChatServer
	m       sync.RWMutex
}

type Implementation struct {
	chats  map[uint64]*Chat
	mxChat sync.RWMutex

	channels  map[uint64]chan *desc.Message
	mxChannel sync.RWMutex
}

func NewImplementation() *Implementation {
	return &Implementation{
		chats:    make(map[uint64]*Chat),
		channels: make(map[uint64]chan *desc.Message),
	}
}

func (i *Implementation) ConnectChat(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) error {
	i.mxChannel.RLock()
	_, ok := i.channels[chatId]
	i.mxChannel.RUnlock()

	if !ok {
		i.channels[chatId] = make(chan *desc.Message, 100)
	}

	i.mxChat.Lock()
	if _, okChat := i.chats[chatId]; !okChat {
		i.chats[chatId] = &Chat{
			streams: make(map[string]desc.ChatApi_ConnectChatServer),
		}
	}
	i.mxChat.Unlock()

	i.chats[chatId].m.Lock()
	i.chats[chatId].streams[email] = stream
	i.chats[chatId].m.Unlock()

	for {
		select {
		case msg, okCh := <-i.channels[chatId]:
			if !okCh {
				return nil
			}

			for email, st := range i.chats[chatId].streams {
				log.Println(email)
				if email != msg.GetFrom() {
					if err := st.Send(msg); err != nil {
						return err
					}
				}
			}

		case <-stream.Context().Done():
			i.chats[chatId].m.Lock()
			delete(i.chats[chatId].streams, email)
			i.chats[chatId].m.Unlock()
			return nil
		}
	}
}

func (i *Implementation) SendMessage(chatId uint64, message *desc.Message) error {
	i.mxChannel.RLock()
	chatChan, ok := i.channels[chatId]
	i.mxChannel.RUnlock()

	if !ok {
		i.channels[chatId] = make(chan *desc.Message, 100)
		chatChan = i.channels[chatId]
	}

	chatChan <- message

	return nil
}
