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

// UpdateMessage update processed msg
func (r *WhatsAppRepository) UpdateMessage(message *models.Message) {
	r.connectionDB.Exec(querys.UpdateMessage(), message.ID)
}

// CreateNewMessage create new message
func (r *WhatsAppRepository) CreateNewMessage(message *models.Message) *models.Message {
	insertQuery, err := r.connectionDB.Prepare(querys.InsertNewMessage())
	if err != nil {
		panic(err.Error())
	}

	res, _ := insertQuery.Exec(message.Idd, false)

	id, _ := res.LastInsertId()
	message.ID = id
	message.Processed = false
	return message

}

// NewWhatsAppRepository IoC
func NewWhatsAppRepository(_connectionDB *database.Database) *WhatsAppRepository {
	return &WhatsAppRepository{
		connectionDB: _connectionDB.GetConnection(),
	}
}
