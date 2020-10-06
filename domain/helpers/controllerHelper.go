package helpers

import (
	"reflect"

	interfacecontroller "o2b.com.br/whatsAppApi/aplications/controllers/interface"
)

// ControllersHelper is a healper for controller
type ControllersHelper struct {
	controllers map[string]interfacecontroller.IController
}

// AppendNewController create new controller
func (ch *ControllersHelper) AppendNewController(funcName string, controller interfacecontroller.IController) {
	ch.controllers[funcName] = controller
}

// GetController get a controller by string
func (ch *ControllersHelper) GetController(funcName string) interfacecontroller.IController {
	return ch.controllers[funcName]
}

// InitiateControllers initiate new controller
func (ch *ControllersHelper) InitiateControllers(funcName string) reflect.Type {
	return reflect.TypeOf(ch.GetController(funcName))
}

// NewControllerHelper IoC
func NewControllerHelper() *ControllersHelper {
	return &ControllersHelper{
		controllers: make(map[string]interfacecontroller.IController),
	}
}
