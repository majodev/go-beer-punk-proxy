// Code generated by go-swagger; DO NOT EDIT.

package beers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetBeerRouteParams creates a new GetBeerRouteParams object
// no default values defined in spec.
func NewGetBeerRouteParams() GetBeerRouteParams {

	return GetBeerRouteParams{}
}

// GetBeerRouteParams contains all the bound params for the get beer route operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetBeerRoute
type GetBeerRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*BeerID
	  Required: true
	  Minimum: 1
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBeerRouteParams() beforehand.
func (o *GetBeerRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBeerRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	// id
	// Required: true
	// Parameter is provided by construction from the route

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetBeerRouteParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

// validateID carries on validations for parameter ID
func (o *GetBeerRouteParams) validateID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("id", "path", o.ID, 1, false); err != nil {
		return err
	}

	return nil
}
