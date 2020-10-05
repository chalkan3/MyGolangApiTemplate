package providers

import (
	"go.uber.org/dig"
)

// RoutesProvider is a provider for services
type RoutesProvider struct {
}

// Provide is a helper Ioc
func (provider *RoutesProvider) Provide(container *dig.Container) {

}

// NewRoutesProvider IoC
func NewRoutesProvider() *RoutesProvider {
	return &RoutesProvider{}
}
