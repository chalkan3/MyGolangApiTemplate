package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"o2b.com.br/whatsAppApi/domain/services"
	"o2b.com.br/whatsAppApi/models"
	apiops "o2b.com.br/whatsAppApi/restapi/operations/api"
)

// WhatsAppController whatsApp controller
type WhatsAppController struct {
	service *services.WhatsAppService
}

// Get get messages
func (wc *WhatsAppController) Get(params apiops.GetMessagesParams) middleware.Responder {

	return apiops.NewGetMessagesOK()
}

// CreateMessage create message
func (wc *WhatsAppController) CreateMessage(params apiops.SendMessageParams) middleware.Responder {
	message := wc.service.CreateNewMessage(params.Body)
	if message != nil {
		return apiops.NewSendMessageCreated().WithPayload(&models.Resp{
			MessageQueued: true,
		})
	} else {
		return apiops.NewSendMessageCreated().WithPayload(&models.Resp{
			MessageQueued: false,
		})
	}

}

// UpdateMessage create message
func (wc *WhatsAppController) UpdateMessage(params apiops.SyncParams) middleware.Responder {
	wc.service.UpdateMessage(params.Body)
	return apiops.NewSyncCreated()
}

// NewWhatsAppController IoC
func NewWhatsAppController(_service *services.WhatsAppService) *WhatsAppController {
	return &WhatsAppController{
		service: _service,
	}
}
