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

// BoilVolume BoilVolume
//
// swagger:model boilVolume
type BoilVolume struct {

	// unit
	// Required: true
	Unit *Unit `json:"unit"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this boil volume
func (m *BoilVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BoilVolume) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this boil volume based on the context it is used
func (m *BoilVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BoilVolume) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BoilVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoilVolume) UnmarshalBinary(b []byte) error {
	var res BoilVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
