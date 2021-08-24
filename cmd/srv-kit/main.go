package main

import (
	"github.com/Evolt0/srv-kit/cmd/srv-kit/global"
	"github.com/Evolt0/srv-kit/pkg/apis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var monitoredSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

func main() {
	config := &global.Config{}
	config.Init()
	logrus.Println(config)
	go func() {
		ls, err := net.Listen("tcp", config.Port)
		if err != nil {
			logrus.Fatalf("failed to tcp listen! %v", err)
		}
		gs := grpc.NewServer()

		apis.InitGRPCRouter(gs)
		err = gs.Serve(ls)
		logrus.Fatalf("failed to run project! %v", err)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, monitoredSignals...)
	select {
	case <-quit:
		logrus.Println("exit...")
	}
}
