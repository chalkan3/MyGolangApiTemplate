package provider

import (
	"go.uber.org/dig"
	constants "o2b.com.br/whatsAppApi/domain/constants"
	providerinterfaces "o2b.com.br/whatsAppApi/infrastructure/provider/interfaces"
	"o2b.com.br/whatsAppApi/infrastructure/provider/providers"
	"o2b.com.br/whatsAppApi/infrastructure/provider/server"
	"o2b.com.br/whatsAppApi/restapi/operations"
)

// ContainerProvider controlls of IoC
type ContainerProvider struct {
	Container *dig.Container
	providers [constants.NumberProviders]providerinterfaces.IProvider
}

func buildContainerProvider() *dig.Container {
	return dig.New()
}

// Provide My Ioc
func (p *ContainerProvider) Provide() *dig.Container {
	for _, pro := range p.providers {
		pro.Provide(p.Container)
	}

	return p.Container
}

// Run EntryPoint
func (p *ContainerProvider) Run(api *operations.WhatsAppAPIAPI) {
	p.Container.Invoke(func(server *server.Server) {
		server.BindAPI(api).Controllers()
	})
}

// NewIocProvider Create provider
func NewIocProvider() *ContainerProvider {
	return &ContainerProvider{
		Container: buildContainerProvider(),
		providers: [constants.NumberProviders]providerinterfaces.IProvider{
			providers.NewEntityProvider(),
			providers.NewHelpersProvider(),
			providers.NewRespositoryProvider(),
			providers.NewServiceProvider(),
			providers.NewControllerProvider(),
			providers.NewRoutesProvider(),
			providers.NewInfraProvider(),
		},
	}
}
