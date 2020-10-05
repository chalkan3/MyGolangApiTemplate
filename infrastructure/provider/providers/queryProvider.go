package providers

import (
	"go.uber.org/dig"
)

// QueryProvider is a provider for services
type QueryProvider struct {
}

// Provide is a helper Ioc
func (provider *QueryProvider) Provide(container *dig.Container) {

}

// NewQueryProvider IoC
func NewQueryProvider() *QueryProvider {
	return &QueryProvider{}
}
