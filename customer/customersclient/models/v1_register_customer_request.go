// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RegisterCustomerRequest v1 register customer request
//
// swagger:model v1RegisterCustomerRequest
type V1RegisterCustomerRequest struct {

	// name
	Name string `json:"name,omitempty"`

	// sms number
	SmsNumber string `json:"smsNumber,omitempty"`
}

// Validate validates this v1 register customer request
func (m *V1RegisterCustomerRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 register customer request based on context it is used
func (m *V1RegisterCustomerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RegisterCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegisterCustomerRequest) UnmarshalBinary(b []byte) error {
	var res V1RegisterCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
