// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Authentication authentication
//
// swagger:model Authentication
type Authentication struct {

	// NTLM Authentication domain
	Domain string `json:"domain,omitempty"`

	// NTLM authentication password
	// Required: true
	Password *string `json:"password"`

	// Authentication type
	// Example: basic
	// Required: true
	Type *string `json:"type"`

	// NTLM  authentication userName
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this authentication
func (m *Authentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authentication) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *Authentication) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Authentication) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authentication based on context it is used
func (m *Authentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Authentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authentication) UnmarshalBinary(b []byte) error {
	var res Authentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
