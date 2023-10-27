package registry

// IRegistry registry interface
type IRegistry interface {
	// Register register service
	Register(serviceName string, serviceID string, serviceAddress string, servicePort int) error
	// Deregister deregister service
	Deregister(serviceID string) error
	// GetService get service
	GetService(serviceName string) ([]*Service, error)
	// GetServiceByID get service by id
	GetServiceByID(serviceID string) (*Service, error)
	// ListServices list service
	ListServices(page, pageSize int) ([]*Service, error)
	// Watch service
	Watch() (<-chan []*Service, error)
}
