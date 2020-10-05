package providers

import (
	"go.uber.org/dig"
)

// EntityProvider is a provider for services
type EntityProvider struct {
}

// Provide is a helper Ioc
func (provider *EntityProvider) Provide(container *dig.Container) {

}

// NewEntityProvider IoC
func NewEntityProvider() *EntityProvider {
	return &EntityProvider{}
}
