package providers

import (
	"go.uber.org/dig"
)

// ServiceProvider is a provider for services
type ServiceProvider struct {
}

// Provide is a helper Ioc
func (provider *ServiceProvider) Provide(container *dig.Container) {

}

// NewServiceProvider IoC
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
