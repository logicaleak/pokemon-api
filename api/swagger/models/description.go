// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Description description
//
// swagger:model description
type Description struct {

	// description
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this description
func (m *Description) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Description) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Description) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Description) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Description) UnmarshalBinary(b []byte) error {
	var res Description
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
