package echo

import (
	"context"

	"github.com/eyotang/tao"
	"github.com/leesper/holmes"
)

const (
	EchoMessage     int32 = 1
	EchoMessageName       = "EchoMessage"
)

// Message defines the echo message.
type Message struct {
	Content string
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	return []byte(em.Content), nil
}

// RequestCommand returns request command.
func (em Message) RequestCommand() int32 {
	return EchoMessage
}

// ResponseCommand returns response command.
func (em Message) ResponseCommand() int32 {
	return EchoMessage
}

// RequestName returns request name.
func (em Message) RequestName() string {
	return EchoMessageName
}

// ResponseName returns response name.
func (em Message) ResponseName() string {
	return EchoMessageName
}

func (em Message) Len() int64 {
	return int64(len(em.Content))
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	msg := string(data)
	echo := Message{
		Content: msg,
	}
	return echo, nil
}

// ProcessMessage process the logic of echo message.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	msg := tao.MessageFromContext(ctx).(Message)
	holmes.Infof("receving message %s\n", msg.Content)
	conn.Write(msg)
}
