package providerinterfaces

import "go.uber.org/dig"

// IProvider is a interface of an provider
type IProvider interface {
	Provide(*dig.Container)
}
