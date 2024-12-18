// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomersServiceChangeSmsNumberBody customers service change sms number body
//
// swagger:model CustomersServiceChangeSmsNumberBody
type CustomersServiceChangeSmsNumberBody struct {

	// sms number
	SmsNumber string `json:"smsNumber,omitempty"`
}

// Validate validates this customers service change sms number body
func (m *CustomersServiceChangeSmsNumberBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customers service change sms number body based on context it is used
func (m *CustomersServiceChangeSmsNumberBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomersServiceChangeSmsNumberBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomersServiceChangeSmsNumberBody) UnmarshalBinary(b []byte) error {
	var res CustomersServiceChangeSmsNumberBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
