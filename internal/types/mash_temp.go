// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MashTemp MashTemp
//
// swagger:model mashTemp
type MashTemp struct {

	// duration
	Duration interface{} `json:"duration,omitempty"`

	// temp
	// Required: true
	Temp *BoilVolume `json:"temp"`
}

// Validate validates this mash temp
func (m *MashTemp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MashTemp) validateTemp(formats strfmt.Registry) error {

	if err := validate.Required("temp", "body", m.Temp); err != nil {
		return err
	}

	if m.Temp != nil {
		if err := m.Temp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("temp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mash temp based on the context it is used
func (m *MashTemp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MashTemp) contextValidateTemp(ctx context.Context, formats strfmt.Registry) error {

	if m.Temp != nil {
		if err := m.Temp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("temp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MashTemp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MashTemp) UnmarshalBinary(b []byte) error {
	var res MashTemp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
