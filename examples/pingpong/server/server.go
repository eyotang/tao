package main

import (
	"context"
	"net"
	"runtime"

	"github.com/eyotang/tao"
	"github.com/eyotang/tao/examples/pingpong"
	"github.com/leesper/holmes"
)

// PingPongServer defines pingpong server.
type PingPongServer struct {
	*tao.Server
}

// NewPingPongServer returns PingPongServer.
func NewPingPongServer() *PingPongServer {
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		holmes.Infoln("closing pingpong client")
	})

	return &PingPongServer{
		tao.NewServer(onConnect, onError, onClose),
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer holmes.Start().Stop()
	tao.MonitorOn(12345)
	tao.Register(pingpong.PingPongMessage, pingpong.DeserializeMessage, ProcessPingPongMessage)

	l, err := net.Listen("tcp", ":12346")
	if err != nil {
		holmes.Fatalln("listen error", err)
	}

	server := NewPingPongServer()

	server.Start(l)
}

// ProcessPingPongMessage handles business logic.
func ProcessPingPongMessage(ctx context.Context, conn tao.WriteCloser) {
	ping := tao.MessageFromContext(ctx).(pingpong.Message)
	holmes.Infoln(ping.Info)
	rsp := pingpong.Message{
		Info: "pong",
	}
	conn.Write(rsp)
}
