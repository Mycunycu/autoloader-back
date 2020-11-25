package services

type Authorization interface {
}

type Services struct {
	Authorization
}

// GetServices - ...
func GetServices() *Services {
	return &Services{}
}
