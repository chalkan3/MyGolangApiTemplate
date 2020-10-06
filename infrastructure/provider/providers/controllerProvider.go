package providers

import (
	"go.uber.org/dig"
	"o2b.com.br/whatsAppApi/aplications/controllers"
)

// ControllerProvider is a provider for services
type ControllerProvider struct {
}

// Provide is a helper Ioc
func (provider *ControllerProvider) Provide(container *dig.Container) {
	container.Provide(controllers.NewWhatsAppController)
	container.Provide(controllers.NewHandleController)
}

// NewControllerProvider IoC
func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{}
}
