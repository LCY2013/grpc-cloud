package registry

import (
	"context"
	"fmt"
	grpcgateway "github.com/LCY2013/grpc-cloud/grpc-gateway/proto"
	httprouter "github.com/LCY2013/grpc-cloud/grpc-gateway/route"
	"github.com/LCY2013/grpc-cloud/logger"
	grpc_cloud_direct "github.com/LCY2013/grpc-cloud/registry/grpc-cloud-direct"
)

func Init() {
	logger.Log.Info("registry init")
	service := grpc_cloud_direct.NewDirectRegistryService()
	service.Handle()
	go route(service)
	logger.Log.Info("registry finish")
}

func route(service *grpc_cloud_direct.DirectRegistryServer) {
	watch, err := service.Watch()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	for {
		select {
		case services := <-watch:
			for _, srv := range services {
				logger.Log.Infof("service id %s, name %s, address %s, port %d", srv.ID, srv.Name, srv.Address, srv.Port)
				annotations, err := grpcgateway.ListServiceMethod(context.Background(), fmt.Sprintf("%s:%d", srv.Address, srv.Port))
				if err != nil {
					logger.Log.Error(err)
					continue
				}
				for uri, anno := range annotations {
					httprouter.R.AddPath(uri, anno)
				}
			}
		}
	}
}
