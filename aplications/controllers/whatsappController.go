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

	msg := []*models.Message{}
	msgIgor := "msg from igor"
	msgClaudio := "msg from gustavo"
	msgGustavo := "msg from claudio"
	msg = append(msg, &models.Message{
		Body:      &msgIgor,
		ID:        1,
		Processed: true,
	})

	msg = append(msg, &models.Message{
		Body:      &msgClaudio,
		ID:        1,
		Processed: false,
	})

	msg = append(msg, &models.Message{
		Body:      &msgGustavo,
		ID:        1,
		Processed: true,
	})

	return apiops.NewGetMessagesOK().WithPayload(msg)
}

// CreateMessage create message
func (wc *WhatsAppController) CreateMessage(params apiops.SendMessageParams) middleware.Responder {
	message := wc.service.CreateNewMessage(params.Body)
	return apiops.NewSendMessageCreated().WithPayload(message)
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
