package providers

import (
	"go.uber.org/dig"
	"o2b.com.br/whatsAppApi/infrastructure/database"
)

// EntityProvider is a provider for services
type EntityProvider struct {
}

// Provide is a helper Ioc
func (provider *EntityProvider) Provide(container *dig.Container) {
	container.Provide(database.NewDatabase)
	container.Provide(database.NewRabbitMQ)
}

// NewEntityProvider IoC
func NewEntityProvider() *EntityProvider {
	return &EntityProvider{}
}
