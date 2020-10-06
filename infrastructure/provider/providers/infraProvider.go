package providers

import (
	"go.uber.org/dig"
	"o2b.com.br/whatsAppApi/aplications/controllers"
	"o2b.com.br/whatsAppApi/infrastructure/provider/server"
)

// InfraProvider is a provider for services
type InfraProvider struct {
}

// Provide is a helper Ioc
func (provider *InfraProvider) Provide(container *dig.Container) {
	container.Provide(controllers.NewHandleController)
	container.Provide(server.NewServer)

}

// NewInfraProvider IoC
func NewInfraProvider() *InfraProvider {
	return &InfraProvider{}
}
