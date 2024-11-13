// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AddProductResponse v1 add product response
//
// swagger:model v1AddProductResponse
type V1AddProductResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this v1 add product response
func (m *V1AddProductResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 add product response based on context it is used
func (m *V1AddProductResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AddProductResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AddProductResponse) UnmarshalBinary(b []byte) error {
	var res V1AddProductResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
