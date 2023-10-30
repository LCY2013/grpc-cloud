package main

import (
	"context"
	"github.com/LCY2013/grpc-cloud/logger"
	v1 "github.com/LCY2013/grpc-cloud/proto/gen/go/grpc/echo/service/v1"
	grpc_cloud_direct "github.com/LCY2013/grpc-cloud/registry/grpc-cloud-direct"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {
	svr := grpc.NewServer()
	v1.RegisterEchoServiceServer(svr, &GrpcGatewayServerProc{})
	reflection.Register(svr)
	go func() {
		time.Sleep(5 * time.Second)
		client := grpc_cloud_direct.NewDirectRegistryClient("http://127.0.0.1:8080")
		err := client.Register("echo-service", "echo-service", "127.0.0.1", 8888)
		if err != nil {
			logger.Log.Error(err)
		}
	}()
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	//port := l.Addr().(*net.TCPAddr).Port
	err = svr.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

type GrpcGatewayServerProc struct {
	v1.EchoServiceServer
}

func (g *GrpcGatewayServerProc) Echo(ctx context.Context, message *v1.StringMessage) (*v1.StringMessage, error) {
	return &v1.StringMessage{
		Value: "hello, " + message.Value,
	}, nil
}

func (g *GrpcGatewayServerProc) EchoCustomer(ctx context.Context, message *v1.StringMessage) (*v1.StringMessage, error) {
	return &v1.StringMessage{
		Value: "* hello, " + message.Value,
	}, nil
}
