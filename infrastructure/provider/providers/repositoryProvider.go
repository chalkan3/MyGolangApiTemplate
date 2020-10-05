package providers

import (
	"go.uber.org/dig"
)

// RespositoryProvider is a provider for services
type RespositoryProvider struct {
}

// Provide is a helper Ioc
func (provider *RespositoryProvider) Provide(container *dig.Container) {

}

// NewRespositoryProvider IoC
func NewRespositoryProvider() *RespositoryProvider {
	return &RespositoryProvider{}
}
