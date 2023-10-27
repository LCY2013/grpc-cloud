package registry

// Service info
type Service struct {
	// ID is the unique service id
	ID string `json:"id"`
	// Name is the service name
	Name string `json:"name"`
	// Address is the service address
	Address string `json:"address"`
	// Port is the service port
	Port int `json:"port"`
	// Tags is the service tags
	Tags []string `json:"tags"`
}
