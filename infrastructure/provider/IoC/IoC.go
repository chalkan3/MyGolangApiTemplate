package IoC

import (
	"o2b.com.br/whatsAppApi/infrastructure/provider"
)

// IoC inversion of control
type IoC struct {
	container *provider.ContainerProvider
}

// GetContainer get main ioc container
func (ioc *IoC) GetContainer() *provider.ContainerProvider {
	return ioc.container
}

// NewIoC create a new ioc
func NewIoC() *IoC {
	return &IoC{
		container: provider.NewIocProvider(),
	}
}
