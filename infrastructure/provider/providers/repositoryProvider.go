package providers

import (
	"go.uber.org/dig"
	"o2b.com.br/whatsAppApi/infrastructure/provider/repository"
)

// RespositoryProvider is a provider for services
type RespositoryProvider struct {
}

// Provide is a helper Ioc
func (provider *RespositoryProvider) Provide(container *dig.Container) {
	container.Provide(repository.NewWhatsAppRepository)
}

// NewRespositoryProvider IoC
func NewRespositoryProvider() *RespositoryProvider {
	return &RespositoryProvider{}
}
