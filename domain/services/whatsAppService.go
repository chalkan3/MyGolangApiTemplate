package services

import (
	"o2b.com.br/whatsAppApi/infrastructure/provider/repository"
	"o2b.com.br/whatsAppApi/models"
)

// WhatsAppService is my service of whatsapp
type WhatsAppService struct {
	repository *repository.WhatsAppRepository
}

// CreateNewMessage create new message
func (s *WhatsAppService) CreateNewMessage(message *models.Message) *models.Message {
	return s.repository.CreateNewMessage(message)
}

// NewWhatsAppService IoC
func NewWhatsAppService(_repository *repository.WhatsAppRepository) *WhatsAppService {
	return &WhatsAppService{
		repository: _repository,
	}
}
