package main

import (
	"github.com/LCY2013/grpc-cloud/grpc-gateway/registry"
	httprouter "github.com/LCY2013/grpc-cloud/grpc-gateway/route"
	"github.com/LCY2013/grpc-cloud/logger"
	"net/http"
)

func main() {
	//grpcgateway.ListServiceMethod(context.Background(), "127.0.0.1:8888")
	registry.Init()
	httprouter.Init()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Error(err)
	}
}
