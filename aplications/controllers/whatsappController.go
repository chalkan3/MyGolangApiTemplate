package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	apiops "o2b.com.br/whatsAppApi/restapi/operations/api"
)

// WhatsAppController whatsApp controller
type WhatsAppController struct {
}

// Get get messages
func (wc *WhatsAppController) Get(params apiops.GetMessagesParams) middleware.Responder {
	return apiops.NewGetMessagesOK()
}

// NewWhatsAppController IoC
func NewWhatsAppController() *WhatsAppController {
	return &WhatsAppController{}
}
