// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GetProductResponse v1 get product response
//
// swagger:model v1GetProductResponse
type V1GetProductResponse struct {

	// product
	Product *V1Product `json:"product,omitempty"`
}

// Validate validates this v1 get product response
func (m *V1GetProductResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetProductResponse) validateProduct(formats strfmt.Registry) error {
	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get product response based on the context it is used
func (m *V1GetProductResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetProductResponse) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

	if m.Product != nil {

		if swag.IsZero(m.Product) { // not required
			return nil
		}

		if err := m.Product.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetProductResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetProductResponse) UnmarshalBinary(b []byte) error {
	var res V1GetProductResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
