// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Hop Hop
//
// swagger:model hop
type Hop struct {

	// add
	// Required: true
	Add Add `json:"add"`

	// amount
	// Required: true
	Amount *BoilVolume `json:"amount"`

	// attribute
	// Required: true
	Attribute Attribute `json:"attribute"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this hop
func (m *Hop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttribute(formats); err != nil {
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

func (m *Hop) validateAdd(formats strfmt.Registry) error {

	if err := m.Add.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("add")
		}
		return err
	}

	return nil
}

func (m *Hop) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *Hop) validateAttribute(formats strfmt.Registry) error {

	if err := m.Attribute.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attribute")
		}
		return err
	}

	return nil
}

func (m *Hop) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hop) UnmarshalBinary(b []byte) error {
	var res Hop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}