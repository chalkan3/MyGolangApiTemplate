// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMessagesHandlerFunc turns a function with the right signature into a get messages handler
type GetMessagesHandlerFunc func(GetMessagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMessagesHandlerFunc) Handle(params GetMessagesParams) middleware.Responder {
	return fn(params)
}

// GetMessagesHandler interface for that can handle valid get messages params
type GetMessagesHandler interface {
	Handle(GetMessagesParams) middleware.Responder
}

// NewGetMessages creates a new http.Handler for the get messages operation
func NewGetMessages(ctx *middleware.Context, handler GetMessagesHandler) *GetMessages {
	return &GetMessages{Context: ctx, Handler: handler}
}

/*GetMessages swagger:route GET / api getMessages

GetMessages get messages API

*/
type GetMessages struct {
	Context *middleware.Context
	Handler GetMessagesHandler
}

func (o *GetMessages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMessagesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
