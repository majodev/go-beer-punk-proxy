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

// Beer Beer
//
// swagger:model beer
type Beer struct {

	// abv
	// Required: true
	Abv *float64 `json:"abv"`

	// attenuation level
	// Required: true
	AttenuationLevel *float64 `json:"attenuation_level"`

	// boil volume
	// Required: true
	BoilVolume *BoilVolume `json:"boil_volume"`

	// brewers tips
	// Required: true
	BrewersTips *string `json:"brewers_tips"`

	// contributed by
	// Required: true
	ContributedBy *string `json:"contributed_by"`

	// description
	// Required: true
	Description *string `json:"description"`

	// ebc
	Ebc interface{} `json:"ebc,omitempty"`

	// first brewed
	// Required: true
	FirstBrewed *string `json:"first_brewed"`

	// food pairing
	// Required: true
	FoodPairing []string `json:"food_pairing"`

	// ibu
	Ibu interface{} `json:"ibu,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// image url
	// Format: uri
	ImageURL strfmt.URI `json:"image_url,omitempty"`

	// ingredients
	// Required: true
	Ingredients *Ingredients `json:"ingredients"`

	// method
	// Required: true
	Method *Method `json:"method"`

	// name
	// Required: true
	Name *string `json:"name"`

	// ph
	Ph interface{} `json:"ph,omitempty"`

	// srm
	Srm interface{} `json:"srm,omitempty"`

	// tagline
	// Required: true
	Tagline *string `json:"tagline"`

	// target fg
	// Required: true
	TargetFg *int64 `json:"target_fg"`

	// target og
	// Required: true
	TargetOg *float64 `json:"target_og"`

	// volume
	// Required: true
	Volume *BoilVolume `json:"volume"`
}

// Validate validates this beer
func (m *Beer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttenuationLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoilVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBrewersTips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstBrewed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFoodPairing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetFg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetOg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Beer) validateAbv(formats strfmt.Registry) error {

	if err := validate.Required("abv", "body", m.Abv); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateAttenuationLevel(formats strfmt.Registry) error {

	if err := validate.Required("attenuation_level", "body", m.AttenuationLevel); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateBoilVolume(formats strfmt.Registry) error {

	if err := validate.Required("boil_volume", "body", m.BoilVolume); err != nil {
		return err
	}

	if m.BoilVolume != nil {
		if err := m.BoilVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boil_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boil_volume")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) validateBrewersTips(formats strfmt.Registry) error {

	if err := validate.Required("brewers_tips", "body", m.BrewersTips); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateContributedBy(formats strfmt.Registry) error {

	if err := validate.Required("contributed_by", "body", m.ContributedBy); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateFirstBrewed(formats strfmt.Registry) error {

	if err := validate.Required("first_brewed", "body", m.FirstBrewed); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateFoodPairing(formats strfmt.Registry) error {

	if err := validate.Required("food_pairing", "body", m.FoodPairing); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateImageURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageURL) { // not required
		return nil
	}

	if err := validate.FormatOf("image_url", "body", "uri", m.ImageURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateIngredients(formats strfmt.Registry) error {

	if err := validate.Required("ingredients", "body", m.Ingredients); err != nil {
		return err
	}

	if m.Ingredients != nil {
		if err := m.Ingredients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingredients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ingredients")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateTagline(formats strfmt.Registry) error {

	if err := validate.Required("tagline", "body", m.Tagline); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateTargetFg(formats strfmt.Registry) error {

	if err := validate.Required("target_fg", "body", m.TargetFg); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateTargetOg(formats strfmt.Registry) error {

	if err := validate.Required("target_og", "body", m.TargetOg); err != nil {
		return err
	}

	return nil
}

func (m *Beer) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this beer based on the context it is used
func (m *Beer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoilVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngredients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Beer) contextValidateBoilVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.BoilVolume != nil {
		if err := m.BoilVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boil_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boil_volume")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) contextValidateIngredients(ctx context.Context, formats strfmt.Registry) error {

	if m.Ingredients != nil {
		if err := m.Ingredients.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingredients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ingredients")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *Beer) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Beer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Beer) UnmarshalBinary(b []byte) error {
	var res Beer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
