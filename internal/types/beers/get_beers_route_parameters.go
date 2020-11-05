// Code generated by go-swagger; DO NOT EDIT.

package beers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetBeersRouteParams creates a new GetBeersRouteParams object
// with the default values initialized.
func NewGetBeersRouteParams() GetBeersRouteParams {

	var (
		// initialize parameters with default values

		pageDefault    = int64(1)
		perPageDefault = int64(25)
	)

	return GetBeersRouteParams{
		Page: &pageDefault,

		PerPage: &perPageDefault,
	}
}

// GetBeersRouteParams contains all the bound params for the get beers route operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetBeersRoute
type GetBeersRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Returns all beers with ABV greater than the supplied number
	  In: query
	*/
	AbvGt *int64 `query:"abv_gt"`
	/*Returns all beers with ABV less than the supplied number
	  In: query
	*/
	AbvLt *int64 `query:"abv_lt"`
	/*Returns all beers matching the supplied name (this will match partial strings as well so e.g punk will return Punk IPA), if you need to add spaces just add an underscore (_).
	  In: query
	*/
	BeerName *string `query:"beer_name"`
	/*Returns all beers brewed after this date, the date format is mm-yyyy e.g 10-2011
	  In: query
	*/
	BrewedAfter *strfmt.DateTime `query:"brewed_after"`
	/*Returns all beers brewed before this date, the date format is mm-yyyy e.g 10-2011
	  In: query
	*/
	BrewedBefore *strfmt.DateTime `query:"brewed_before"`
	/*Returns all beers with EBC greater than the supplied number
	  In: query
	*/
	EbcGt *int64 `query:"ebc_gt"`
	/*Returns all beers with EBC less than the supplied number
	  In: query
	*/
	EbcLt *int64 `query:"ebc_lt"`
	/*Returns all beers matching the supplied food string, this performs a fuzzy match, if you need to add spaces just add an underscore (_).
	  In: query
	*/
	Food *string `query:"food"`
	/*Returns all beers matching the supplied hops name, this performs a fuzzy match, if you need to add spaces just add an underscore (_).
	  In: query
	*/
	Hops *string `query:"hops"`
	/*Returns all beers with IBU greater than the supplied number
	  In: query
	*/
	IbuGt *int64 `query:"ibu_gt"`
	/*Returns all beers with IBU less than the supplied number
	  In: query
	*/
	IbuLt *int64 `query:"ibu_lt"`
	/*Returns all beers matching the supplied ID's. You can pass in multiple ID's by separating them with a | symbol.
	  In: query
	*/
	Ids *string `query:"ids"`
	/*Returns all beers matching the supplied malt name, this performs a fuzzy match, if you need to add spaces just add an underscore (_).
	  In: query
	*/
	Malt *string `query:"malt"`
	/*page
	  Minimum: 1
	  In: query
	  Default: 1
	*/
	Page *int64 `query:"page"`
	/*per_page
	  Maximum: 80
	  Minimum: 1
	  In: query
	  Default: 25
	*/
	PerPage *int64 `query:"per_page"`
	/*Returns all beers matching the supplied yeast name, this performs a fuzzy match, if you need to add spaces just add an underscore (_).
	  In: query
	*/
	Yeast *string `query:"yeast"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBeersRouteParams() beforehand.
func (o *GetBeersRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAbvGt, qhkAbvGt, _ := qs.GetOK("abv_gt")
	if err := o.bindAbvGt(qAbvGt, qhkAbvGt, route.Formats); err != nil {
		res = append(res, err)
	}

	qAbvLt, qhkAbvLt, _ := qs.GetOK("abv_lt")
	if err := o.bindAbvLt(qAbvLt, qhkAbvLt, route.Formats); err != nil {
		res = append(res, err)
	}

	qBeerName, qhkBeerName, _ := qs.GetOK("beer_name")
	if err := o.bindBeerName(qBeerName, qhkBeerName, route.Formats); err != nil {
		res = append(res, err)
	}

	qBrewedAfter, qhkBrewedAfter, _ := qs.GetOK("brewed_after")
	if err := o.bindBrewedAfter(qBrewedAfter, qhkBrewedAfter, route.Formats); err != nil {
		res = append(res, err)
	}

	qBrewedBefore, qhkBrewedBefore, _ := qs.GetOK("brewed_before")
	if err := o.bindBrewedBefore(qBrewedBefore, qhkBrewedBefore, route.Formats); err != nil {
		res = append(res, err)
	}

	qEbcGt, qhkEbcGt, _ := qs.GetOK("ebc_gt")
	if err := o.bindEbcGt(qEbcGt, qhkEbcGt, route.Formats); err != nil {
		res = append(res, err)
	}

	qEbcLt, qhkEbcLt, _ := qs.GetOK("ebc_lt")
	if err := o.bindEbcLt(qEbcLt, qhkEbcLt, route.Formats); err != nil {
		res = append(res, err)
	}

	qFood, qhkFood, _ := qs.GetOK("food")
	if err := o.bindFood(qFood, qhkFood, route.Formats); err != nil {
		res = append(res, err)
	}

	qHops, qhkHops, _ := qs.GetOK("hops")
	if err := o.bindHops(qHops, qhkHops, route.Formats); err != nil {
		res = append(res, err)
	}

	qIbuGt, qhkIbuGt, _ := qs.GetOK("ibu_gt")
	if err := o.bindIbuGt(qIbuGt, qhkIbuGt, route.Formats); err != nil {
		res = append(res, err)
	}

	qIbuLt, qhkIbuLt, _ := qs.GetOK("ibu_lt")
	if err := o.bindIbuLt(qIbuLt, qhkIbuLt, route.Formats); err != nil {
		res = append(res, err)
	}

	qIds, qhkIds, _ := qs.GetOK("ids")
	if err := o.bindIds(qIds, qhkIds, route.Formats); err != nil {
		res = append(res, err)
	}

	qMalt, qhkMalt, _ := qs.GetOK("malt")
	if err := o.bindMalt(qMalt, qhkMalt, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPerPage, qhkPerPage, _ := qs.GetOK("per_page")
	if err := o.bindPerPage(qPerPage, qhkPerPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qYeast, qhkYeast, _ := qs.GetOK("yeast")
	if err := o.bindYeast(qYeast, qhkYeast, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBeersRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	// abv_gt
	// Required: false
	// AllowEmptyValue: false

	// abv_lt
	// Required: false
	// AllowEmptyValue: false

	// beer_name
	// Required: false
	// AllowEmptyValue: false

	// brewed_after
	// Required: false
	// AllowEmptyValue: false

	if err := o.validateBrewedAfter(formats); err != nil {
		res = append(res, err)
	}

	// brewed_before
	// Required: false
	// AllowEmptyValue: false

	if err := o.validateBrewedBefore(formats); err != nil {
		res = append(res, err)
	}

	// ebc_gt
	// Required: false
	// AllowEmptyValue: false

	// ebc_lt
	// Required: false
	// AllowEmptyValue: false

	// food
	// Required: false
	// AllowEmptyValue: false

	// hops
	// Required: false
	// AllowEmptyValue: false

	// ibu_gt
	// Required: false
	// AllowEmptyValue: false

	// ibu_lt
	// Required: false
	// AllowEmptyValue: false

	// ids
	// Required: false
	// AllowEmptyValue: false

	// malt
	// Required: false
	// AllowEmptyValue: false

	// page
	// Required: false
	// AllowEmptyValue: false

	if err := o.validatePage(formats); err != nil {
		res = append(res, err)
	}

	// per_page
	// Required: false
	// AllowEmptyValue: false

	if err := o.validatePerPage(formats); err != nil {
		res = append(res, err)
	}

	// yeast
	// Required: false
	// AllowEmptyValue: false

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAbvGt binds and validates parameter AbvGt from query.
func (o *GetBeersRouteParams) bindAbvGt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("abv_gt", "query", "int64", raw)
	}
	o.AbvGt = &value

	return nil
}

// bindAbvLt binds and validates parameter AbvLt from query.
func (o *GetBeersRouteParams) bindAbvLt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("abv_lt", "query", "int64", raw)
	}
	o.AbvLt = &value

	return nil
}

// bindBeerName binds and validates parameter BeerName from query.
func (o *GetBeersRouteParams) bindBeerName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.BeerName = &raw

	return nil
}

// bindBrewedAfter binds and validates parameter BrewedAfter from query.
func (o *GetBeersRouteParams) bindBrewedAfter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("brewed_after", "query", "strfmt.DateTime", raw)
	}
	o.BrewedAfter = (value.(*strfmt.DateTime))

	if err := o.validateBrewedAfter(formats); err != nil {
		return err
	}

	return nil
}

// validateBrewedAfter carries on validations for parameter BrewedAfter
func (o *GetBeersRouteParams) validateBrewedAfter(formats strfmt.Registry) error {

	// Required: false
	if o.BrewedAfter == nil {
		return nil
	}

	if err := validate.FormatOf("brewed_after", "query", "date-time", (*o.BrewedAfter).String(), formats); err != nil {
		return err
	}
	return nil
}

// bindBrewedBefore binds and validates parameter BrewedBefore from query.
func (o *GetBeersRouteParams) bindBrewedBefore(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("brewed_before", "query", "strfmt.DateTime", raw)
	}
	o.BrewedBefore = (value.(*strfmt.DateTime))

	if err := o.validateBrewedBefore(formats); err != nil {
		return err
	}

	return nil
}

// validateBrewedBefore carries on validations for parameter BrewedBefore
func (o *GetBeersRouteParams) validateBrewedBefore(formats strfmt.Registry) error {

	// Required: false
	if o.BrewedBefore == nil {
		return nil
	}

	if err := validate.FormatOf("brewed_before", "query", "date-time", (*o.BrewedBefore).String(), formats); err != nil {
		return err
	}
	return nil
}

// bindEbcGt binds and validates parameter EbcGt from query.
func (o *GetBeersRouteParams) bindEbcGt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ebc_gt", "query", "int64", raw)
	}
	o.EbcGt = &value

	return nil
}

// bindEbcLt binds and validates parameter EbcLt from query.
func (o *GetBeersRouteParams) bindEbcLt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ebc_lt", "query", "int64", raw)
	}
	o.EbcLt = &value

	return nil
}

// bindFood binds and validates parameter Food from query.
func (o *GetBeersRouteParams) bindFood(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Food = &raw

	return nil
}

// bindHops binds and validates parameter Hops from query.
func (o *GetBeersRouteParams) bindHops(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Hops = &raw

	return nil
}

// bindIbuGt binds and validates parameter IbuGt from query.
func (o *GetBeersRouteParams) bindIbuGt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ibu_gt", "query", "int64", raw)
	}
	o.IbuGt = &value

	return nil
}

// bindIbuLt binds and validates parameter IbuLt from query.
func (o *GetBeersRouteParams) bindIbuLt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ibu_lt", "query", "int64", raw)
	}
	o.IbuLt = &value

	return nil
}

// bindIds binds and validates parameter Ids from query.
func (o *GetBeersRouteParams) bindIds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Ids = &raw

	return nil
}

// bindMalt binds and validates parameter Malt from query.
func (o *GetBeersRouteParams) bindMalt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Malt = &raw

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetBeersRouteParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetBeersRouteParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	if err := o.validatePage(formats); err != nil {
		return err
	}

	return nil
}

// validatePage carries on validations for parameter Page
func (o *GetBeersRouteParams) validatePage(formats strfmt.Registry) error {

	// Required: false
	if o.Page == nil {
		return nil
	}

	if err := validate.MinimumInt("page", "query", int64(*o.Page), 1, false); err != nil {
		return err
	}

	return nil
}

// bindPerPage binds and validates parameter PerPage from query.
func (o *GetBeersRouteParams) bindPerPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetBeersRouteParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("per_page", "query", "int64", raw)
	}
	o.PerPage = &value

	if err := o.validatePerPage(formats); err != nil {
		return err
	}

	return nil
}

// validatePerPage carries on validations for parameter PerPage
func (o *GetBeersRouteParams) validatePerPage(formats strfmt.Registry) error {

	// Required: false
	if o.PerPage == nil {
		return nil
	}

	if err := validate.MinimumInt("per_page", "query", int64(*o.PerPage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("per_page", "query", int64(*o.PerPage), 80, false); err != nil {
		return err
	}

	return nil
}

// bindYeast binds and validates parameter Yeast from query.
func (o *GetBeersRouteParams) bindYeast(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Yeast = &raw

	return nil
}