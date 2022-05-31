package data

import "github.com/J-Obog/go-chat-server/api/service/message"

var messageStore []message.Message

func InitStore() {
	messageStore = make([]message.Message, 0)
}
