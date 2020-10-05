package providers

import (
	"go.uber.org/dig"
)

// ControllerProvider is a provider for services
type ControllerProvider struct {
}

// Provide is a helper Ioc
func (provider *ControllerProvider) Provide(container *dig.Container) {

}

// NewControllerProvider IoC
func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{}
}
