// Code generated by go-swagger; DO NOT EDIT.

package pokemons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHandlerFunc turns a function with the right signature into a get handler
type GetHandlerFunc func(GetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHandlerFunc) Handle(params GetParams) middleware.Responder {
	return fn(params)
}

// GetHandler interface for that can handle valid get params
type GetHandler interface {
	Handle(GetParams) middleware.Responder
}

// NewGet creates a new http.Handler for the get operation
func NewGet(ctx *middleware.Context, handler GetHandler) *Get {
	return &Get{Context: ctx, Handler: handler}
}

/*Get swagger:route GET / pokemons get

Get get API

*/
type Get struct {
	Context *middleware.Context
	Handler GetHandler
}

func (o *Get) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
