package providers

import (
	"go.uber.org/dig"
)

// HelpersProvider is a provider for services
type HelpersProvider struct {
}

// Provide is a helper Ioc
func (provider *HelpersProvider) Provide(container *dig.Container) {

}

// NewHelpersProvider IoC
func NewHelpersProvider() *HelpersProvider {
	return &HelpersProvider{}
}
