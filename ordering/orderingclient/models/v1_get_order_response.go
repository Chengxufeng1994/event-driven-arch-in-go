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

// V1GetOrderResponse v1 get order response
//
// swagger:model v1GetOrderResponse
type V1GetOrderResponse struct {

	// order
	Order *V1Order `json:"order,omitempty"`
}

// Validate validates this v1 get order response
func (m *V1GetOrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetOrderResponse) validateOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get order response based on the context it is used
func (m *V1GetOrderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetOrderResponse) contextValidateOrder(ctx context.Context, formats strfmt.Registry) error {

	if m.Order != nil {

		if swag.IsZero(m.Order) { // not required
			return nil
		}

		if err := m.Order.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetOrderResponse) UnmarshalBinary(b []byte) error {
	var res V1GetOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
