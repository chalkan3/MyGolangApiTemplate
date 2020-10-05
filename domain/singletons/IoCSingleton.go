package singletons

import (
	"sync"

	i "o2b.com.br/whatsAppApi/infrastructure/provider/IoC"
)

// GetIoCSingleton get the ioc singleton
func GetIoCSingleton() *i.IoC {
	var once sync.Once
	var iocSingleton *i.IoC

	once.Do(func() {
		iocSingleton = i.NewIoC()
	})
	return iocSingleton
}
