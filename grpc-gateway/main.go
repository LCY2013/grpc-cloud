package main

import (
	"github.com/LCY2013/grpc-cloud/grpc-gateway/registry"
	"github.com/LCY2013/grpc-cloud/logger"
	"net/http"
)

func main() {
	//grpcgateway.ListServiceMethod(context.Background(), "127.0.0.1:8888")
	registry.Init()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logger.Log.Info("hello world")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Error(err)
	}
}
