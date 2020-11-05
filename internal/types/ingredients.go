// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Ingredients Ingredients
//
// swagger:model ingredients
type Ingredients struct {

	// hops
	// Required: true
	Hops []*Hop `json:"hops"`

	// malt
	// Required: true
	Malt []*Malt `json:"malt"`

	// yeast
	Yeast string `json:"yeast,omitempty"`
}

// Validate validates this ingredients
func (m *Ingredients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ingredients) validateHops(formats strfmt.Registry) error {

	if err := validate.Required("hops", "body", m.Hops); err != nil {
		return err
	}

	for i := 0; i < len(m.Hops); i++ {
		if swag.IsZero(m.Hops[i]) { // not required
			continue
		}

		if m.Hops[i] != nil {
			if err := m.Hops[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Ingredients) validateMalt(formats strfmt.Registry) error {

	if err := validate.Required("malt", "body", m.Malt); err != nil {
		return err
	}

	for i := 0; i < len(m.Malt); i++ {
		if swag.IsZero(m.Malt[i]) { // not required
			continue
		}

		if m.Malt[i] != nil {
			if err := m.Malt[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("malt" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ingredients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ingredients) UnmarshalBinary(b []byte) error {
	var res Ingredients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
