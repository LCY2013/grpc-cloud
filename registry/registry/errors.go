package registry

import "errors"

var (
	// ErrRegistry registry error
	ErrRegistry = errors.New("registry error")
	// ErrServiceNotFound service not found
	ErrServiceNotFound = errors.New("service not found")
	// ErrServiceAlreadyExists service already exists
	ErrServiceAlreadyExists = errors.New("service already exists")
	// ErrServiceIDRequired service id required
	ErrServiceIDRequired = errors.New("service id required")
	// ErrServiceNameRequired service name required
	ErrServiceNameRequired = errors.New("service name required")
	// ErrServiceAddressRequired service address required
	ErrServiceAddressRequired = errors.New("service address required")
	// ErrServicePortRequired service port required
	ErrServicePortRequired = errors.New("service port required")
)
