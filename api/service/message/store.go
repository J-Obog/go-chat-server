package message

var messageStore []Message

func InitMessageStore() {
	messageStore = make([]Message, 0)
}
