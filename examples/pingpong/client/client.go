package main

import (
	"context"
	"net"

	"github.com/eyotang/tao"
	"github.com/eyotang/tao/examples/pingpong"
	"github.com/leesper/holmes"
)

var (
	rspChan = make(chan string)
)

func main() {
	defer holmes.Start().Stop()

	tao.Register(pingpong.PingPongMessage, pingpong.DeserializeMessage, ProcessPingPongMessage)

	c, err := net.Dial("tcp", "127.0.0.1:12346")
	if err != nil {
		holmes.Fatalln(err)
	}

	conn := tao.NewClientConn(0, c)
	defer conn.Close()

	conn.Start()
	req := pingpong.Message{
		Info: "ping",
	}
	for {
		conn.Write(req)
		holmes.Infoln(<-rspChan)
	}
}

// ProcessPingPongMessage handles business logic.
func ProcessPingPongMessage(ctx context.Context, conn tao.WriteCloser) {
	rsp := tao.MessageFromContext(ctx).(pingpong.Message)
	rspChan <- rsp.Info
}
