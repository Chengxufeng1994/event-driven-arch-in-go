// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Store v1 store
//
// swagger:model v1Store
type V1Store struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// participating
	Participating bool `json:"participating,omitempty"`
}

// Validate validates this v1 store
func (m *V1Store) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 store based on context it is used
func (m *V1Store) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Store) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Store) UnmarshalBinary(b []byte) error {
	var res V1Store
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
