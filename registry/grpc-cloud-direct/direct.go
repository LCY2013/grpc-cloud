package grpc_cloud_direct

import (
	"bytes"
	"encoding/json"
	"github.com/LCY2013/grpc-cloud/logger"
	"github.com/LCY2013/grpc-cloud/registry/registry"
	"net/http"
)

// DirectRegistry direct registry
type DirectRegistry struct {
	gatewayAddr string
	serviceMap  map[string]*registry.Service
	watch       chan []*registry.Service
}

func (d *DirectRegistry) Register(serviceName string, serviceID string, serviceAddress string, servicePort int) error {
	serviceInfo := &registry.Service{
		Name:    serviceName,
		ID:      serviceID,
		Address: serviceAddress,
		Port:    servicePort,
	}
	service, _ := json.Marshal(serviceInfo)
	rsp, err := http.Post(d.gatewayAddr+"/register", "application/json", bytes.NewBuffer(service))
	if err != nil {
		return err
	}
	logger.Log.Info(rsp)
	return nil
}

func (d *DirectRegistry) Deregister(serviceID string) error {
	serviceInfo := &registry.Service{
		ID: serviceID,
	}
	service, _ := json.Marshal(serviceInfo)
	rsp, err := http.Post(d.gatewayAddr+"/deregister", "application/json", bytes.NewBuffer(service))
	if err != nil {
		return err
	}
	logger.Log.Info(rsp)
	return nil
}

func (d *DirectRegistry) GetService(serviceName string) ([]*registry.Service, error) {
	return nil, nil
}

func (d *DirectRegistry) GetServiceByID(serviceID string) (*registry.Service, error) {
	return d.serviceMap[serviceID], nil
}

func (d *DirectRegistry) ListServices(page, pageSize int) ([]*registry.Service, error) {
	return nil, nil
}

func (d *DirectRegistry) Watch() (<-chan []*registry.Service, error) {
	if d.watch == nil {
		d.watch = make(chan []*registry.Service)
	}
	return d.watch, nil
}
