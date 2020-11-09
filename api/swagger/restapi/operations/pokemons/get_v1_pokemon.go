// Code generated by go-swagger; DO NOT EDIT.

package pokemons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1PokemonHandlerFunc turns a function with the right signature into a get v1 pokemon handler
type GetV1PokemonHandlerFunc func(GetV1PokemonParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1PokemonHandlerFunc) Handle(params GetV1PokemonParams) middleware.Responder {
	return fn(params)
}

// GetV1PokemonHandler interface for that can handle valid get v1 pokemon params
type GetV1PokemonHandler interface {
	Handle(GetV1PokemonParams) middleware.Responder
}

// NewGetV1Pokemon creates a new http.Handler for the get v1 pokemon operation
func NewGetV1Pokemon(ctx *middleware.Context, handler GetV1PokemonHandler) *GetV1Pokemon {
	return &GetV1Pokemon{Context: ctx, Handler: handler}
}

/*GetV1Pokemon swagger:route GET /v1/pokemon pokemons getV1Pokemon

GetV1Pokemon get v1 pokemon API

*/
type GetV1Pokemon struct {
	Context *middleware.Context
	Handler GetV1PokemonHandler
}

func (o *GetV1Pokemon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1PokemonParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}