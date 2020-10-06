package repository

import (
	"database/sql"

	"o2b.com.br/whatsAppApi/domain/querys"

	"o2b.com.br/whatsAppApi/infrastructure/database"
	"o2b.com.br/whatsAppApi/models"
)

// WhatsAppRepository repostiroy
type WhatsAppRepository struct {
	connectionDB *sql.DB
}

// CreateNewMessage create new message
func (r *WhatsAppRepository) CreateNewMessage(message *models.Message) *models.Message {
	insertQuery, err := r.connectionDB.Prepare(querys.InsertNewMessage())
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(message.Body, false)
	return message

}

// NewWhatsAppRepository IoC
func NewWhatsAppRepository(_connectionDB *database.Database) *WhatsAppRepository {
	return &WhatsAppRepository{
		connectionDB: _connectionDB.GetConnection(),
	}
}
