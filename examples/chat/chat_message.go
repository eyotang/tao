package chat

import (
	"context"

	"github.com/eyotang/tao"
	"github.com/leesper/holmes"
)

const (
	// ChatMessage is the message number of chat message.
	ChatMessage     int32 = 1
	ChatMessageName       = "ChatMessage"
)

// Message defines the chat message.
type Message struct {
	Content string
}

// RequestCommand returns the request command.
func (cm Message) RequestCommand() int32 {
	return ChatMessage
}

// ResponseCommand returns the response command.
func (cm Message) ResponseCommand() int32 {
	return ChatMessage
}

// RequestName returns the request name.
func (cm Message) RequestName() string {
	return ChatMessageName
}

// ResponseName returns the response name.
func (cm Message) ResponseName() string {
	return ChatMessageName
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

func (cm Message) Len() int64 {
	return int64(len(cm.Content))
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	holmes.Infof("ProcessMessage")
	s, ok := tao.ServerFromContext(ctx)
	if ok {
		msg := tao.MessageFromContext(ctx)
		s.Broadcast(msg)
	}
}
