package mq

type Message struct {
	Topic string
	Body  []byte
}

func NewMessage(topic string, body []byte) *Message {
	return &Message{
		Topic: topic,
		Body:  body,
	}
}

type MessageExt struct {
	MsgID string
	Message
}
type SendResult struct {
	Status int
	MsgID  string
}

func NewSendResult(status int, msgID string) *SendResult {
	return &SendResult{
		Status: status,
		MsgID:  msgID,
	}
}
