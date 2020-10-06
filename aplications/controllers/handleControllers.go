package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"o2b.com.br/whatsAppApi/restapi/operations"
	apiops "o2b.com.br/whatsAppApi/restapi/operations/api"
)

// HandleController handle my controller for injection
type HandleController struct {
	whatsAppController *WhatsAppController
}

// Handle entrypoint of controllers
func (h *HandleController) Handle(api *operations.WhatsAppAPIAPI) {
	api.APIGetMessagesHandler = apiops.GetMessagesHandlerFunc(func(params apiops.GetMessagesParams) middleware.Responder {
		return h.whatsAppController.Get(params)
	})
}

// NewHandleController IoC
func NewHandleController(_whatsAppController *WhatsAppController) *HandleController {
	return &HandleController{
		whatsAppController: _whatsAppController,
	}
}
