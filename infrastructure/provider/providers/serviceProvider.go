package providers

import (
	"go.uber.org/dig"
	"o2b.com.br/whatsAppApi/domain/services"
)

// ServiceProvider is a provider for services
type ServiceProvider struct {
}

// Provide is a helper Ioc
func (provider *ServiceProvider) Provide(container *dig.Container) {
	container.Provide(services.NewWhatsAppService)
}

// NewServiceProvider IoC
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
