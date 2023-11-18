package main

import (
	"context"
	"github.com/LCY2013/grpc-cloud/logger"
	grpc_cloud_plugin_validate "github.com/LCY2013/grpc-cloud/plugins/grpc-cloud-plugin-validate"
	v1 "github.com/LCY2013/grpc-cloud/proto/gen/go/grpc/person/service/v1"
	grpc_cloud_direct "github.com/LCY2013/grpc-cloud/registry/grpc-cloud-direct"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {
	svr := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpc_cloud_plugin_validate.WithValidatePbServerInterceptor,
		// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
		//otelgrpc.UnaryServerInterceptor(),
		//srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
		//logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		//selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
		//recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
	),
		grpc.ChainStreamInterceptor(
		//otelgrpc.StreamServerInterceptor(),
		//srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
		//logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		//selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
		//recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		))
	v1.RegisterPersonServiceServer(svr, &PersonServiceProc{})
	reflection.Register(svr)
	go func() {
		time.Sleep(5 * time.Second)
		client := grpc_cloud_direct.NewDirectRegistryClient("http://127.0.0.1:8080")
		err := client.Register("person-service", "person-service", "127.0.0.1", 8889)
		if err != nil {
			logger.Log.Error(err)
		}
	}()
	l, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Fatal(err)
	}
	//port := l.Addr().(*net.TCPAddr).Port
	err = svr.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

type PersonServiceProc struct {
	v1.PersonServiceServer
}

func (g *PersonServiceProc) GetPerson(ctx context.Context, message *v1.PersonMessage) (*v1.PersonMessage, error) {
	return &v1.PersonMessage{
		Name: "hello, " + message.Name,
		Age:  message.Age,
	}, nil
}
