package grpc_cloud_direct

import (
	"encoding/json"
	"github.com/LCY2013/grpc-cloud/logger"
	"github.com/LCY2013/grpc-cloud/registry/registry"
	"io"
	"net/http"
)

// DirectRegistryServer direct registry service
type DirectRegistryServer struct {
	DirectRegistry
}

// NewDirectRegistryService new direct registry service
func NewDirectRegistryService() *DirectRegistryServer {
	return &DirectRegistryServer{
		DirectRegistry: DirectRegistry{},
	}
}

// Handle direct registry service
func (d *DirectRegistryServer) Handle() {
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		service, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Log.Error(err)
		}
		var serviceInfo *registry.Service
		if err = json.Unmarshal(service, &serviceInfo); err != nil {
			logger.Log.Error(err)
		}
		d.serviceMap[serviceInfo.ID] = serviceInfo
		d.watch <- []*registry.Service{serviceInfo}
		if _, err = writer.Write([]byte("registry success")); err != nil {
			logger.Log.Error(err)
		}
	})
	http.HandleFunc("/deregister", func(writer http.ResponseWriter, request *http.Request) {
		service, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Log.Error(err)
		}
		var serviceInfo *registry.Service
		if err = json.Unmarshal(service, &serviceInfo); err != nil {
			logger.Log.Error(err)
		}
		delete(d.serviceMap, serviceInfo.ID)
		if _, err = writer.Write([]byte("deregister success")); err != nil {
			logger.Log.Error(err)
		}
	})
}

// DirectRegistryClient direct registry client
type DirectRegistryClient struct {
	DirectRegistry
}

// NewDirectRegistryClient new direct registry client
func NewDirectRegistryClient(gatewayAddr string) *DirectRegistryClient {
	return &DirectRegistryClient{
		DirectRegistry: DirectRegistry{
			gatewayAddr: gatewayAddr,
		},
	}
}
