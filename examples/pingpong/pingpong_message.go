package pingpong

import "github.com/eyotang/tao"

const (
	// PingPongMessage defines message number.
	PingPongMessage     int32 = 1
	PingPongMessageName       = "PingPongMessage"
)

// Message defines message format.
type Message struct {
	Info string
}

// RequestCommand returns the request command.
func (pp Message) RequestCommand() int32 {
	return PingPongMessage
}

func (pp Message) ResponseCommand() int32 {
	return PingPongMessage
}

func (pp Message) RequestName() string {
	return PingPongMessageName
}

func (pp Message) ResponseName() string {
	return PingPongMessageName
}

// Serialize serializes Message into bytes.
func (pp Message) Serialize() ([]byte, error) {
	return []byte(pp.Info), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	info := string(data)
	msg := Message{
		Info: info,
	}
	return msg, nil
}

// func ProcessPingPongMessage(ctx tao.Context, conn tao.Connection) {
//   if serverConn, ok := conn.(*tao.ServerConnection); ok {
//     if serverConn.GetOwner() != nil {
//       connections := serverConn.GetOwner().GetAllConnections()
//       for v := range connections.IterValues() {
//         c := v.(tao.Connection)
//         c.Write(ctx.Message())
//       }
//     }
//   }
// }
