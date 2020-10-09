package services

import (
	"o2b.com.br/whatsAppApi/domain/helpers"
	"o2b.com.br/whatsAppApi/infrastructure/database"
	"o2b.com.br/whatsAppApi/infrastructure/provider/repository"
	"o2b.com.br/whatsAppApi/models"
)

// WhatsAppService is my service of whatsapp
type WhatsAppService struct {
	repository *repository.WhatsAppRepository
	rabbitMQ   *database.RabbitMQ
}

func (s *WhatsAppService) publishToProcessQueue(message *models.Message) {
	s.rabbitMQ.Publish(helpers.JSONStringfy(message))
}

// CreateNewMessage create new message
func (s *WhatsAppService) CreateNewMessage(message *models.Message) *models.Message {
	savedMessage := s.repository.CreateNewMessage(message)
	s.publishToProcessQueue(message)
	return savedMessage
}

// UpdateMessage create new message
func (s *WhatsAppService) UpdateMessage(message *models.Message) {
	s.repository.UpdateMessage(message)
}

// NewWhatsAppService IoC
func NewWhatsAppService(_repository *repository.WhatsAppRepository, _rabbitMQ *database.RabbitMQ) *WhatsAppService {
	return &WhatsAppService{
		repository: _repository,
		rabbitMQ:   _rabbitMQ,
	}
}
