// Code generated by go-swagger; DO NOT EDIT.

package pokemondescription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ozum.safaoglu/pokemon-api/api/swagger/models"
)

// GetV1NameOKCode is the HTTP code returned for type GetV1NameOK
const GetV1NameOKCode int = 200

/*GetV1NameOK Describes a pokemon

swagger:response getV1NameOK
*/
type GetV1NameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Description `json:"body,omitempty"`
}

// NewGetV1NameOK creates GetV1NameOK with default headers values
func NewGetV1NameOK() *GetV1NameOK {

	return &GetV1NameOK{}
}

// WithPayload adds the payload to the get v1 name o k response
func (o *GetV1NameOK) WithPayload(payload *models.Description) *GetV1NameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 name o k response
func (o *GetV1NameOK) SetPayload(payload *models.Description) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1NameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetV1NameDefault generic error response

swagger:response getV1NameDefault
*/
type GetV1NameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetV1NameDefault creates GetV1NameDefault with default headers values
func NewGetV1NameDefault(code int) *GetV1NameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetV1NameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get v1 name default response
func (o *GetV1NameDefault) WithStatusCode(code int) *GetV1NameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get v1 name default response
func (o *GetV1NameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get v1 name default response
func (o *GetV1NameDefault) WithPayload(payload *models.Error) *GetV1NameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 name default response
func (o *GetV1NameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1NameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
