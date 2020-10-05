package server

import (
	"github.com/go-openapi/runtime/middleware"
	"o2b.com.br/whatsAppApi/restapi/operations"
	apiops "o2b.com.br/whatsAppApi/restapi/operations/api"
)

// Server entrypoint for aplication
type Server struct {
	api *operations.WhatsAppAPIAPI
}

// Controllers my server run
func (s *Server) Controllers() {
	// Controllers
	s.api.APIGetMessagesHandler = apiops.GetMessagesHandlerFunc(func(params apiops.GetMessagesParams) middleware.Responder {
		return apiops.NewGetMessagesOK()
	})

	if s.api.APIGetMessagesHandler == nil {
		s.api.APIGetMessagesHandler = apiops.GetMessagesHandlerFunc(func(params apiops.GetMessagesParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GetMessages has not yet been implemented")
		})
	}
}

// BindAPI bind a api from swagger
func (s *Server) BindAPI(ap *operations.WhatsAppAPIAPI) *Server {
	s.api = ap
	return s
}

// NewServer Server IoC
func NewServer() *Server {
	return &Server{}
}
