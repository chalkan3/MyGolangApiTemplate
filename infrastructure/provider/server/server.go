package server

import (
	"o2b.com.br/whatsAppApi/aplications/controllers"
	"o2b.com.br/whatsAppApi/restapi/operations"
)

// Server entrypoint for aplication
type Server struct {
	handleController *controllers.HandleController
	api              *operations.WhatsAppAPIAPI
}

// Controllers my server run
func (s *Server) Controllers() {
	s.handleController.Handle(s.api)
}

// BindAPI bind a api from swagger
func (s *Server) BindAPI(ap *operations.WhatsAppAPIAPI) *Server {
	s.api = ap
	return s
}

// NewServer Server IoC
func NewServer(_handleController *controllers.HandleController) *Server {
	return &Server{
		handleController: _handleController,
	}
}
