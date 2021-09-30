// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostStudentParams creates a new PostStudentParams object
// no default values defined in spec.
func NewPostStudentParams() PostStudentParams {

	return PostStudentParams{}
}

// PostStudentParams contains all the bound params for the post student operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostStudent
type PostStudentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*defaults to CS if not given
	  In: query
	*/
	Department *string
	/*Should be greater than 20
	  In: query
	*/
	Age *int32
	/*It is int64
	  In: query
	*/
	Enrollment *int64
	/*defaults to Adil if not given
	  In: query
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostStudentParams() beforehand.
func (o *PostStudentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDepartment, qhkDepartment, _ := qs.GetOK("Department")
	if err := o.bindDepartment(qDepartment, qhkDepartment, route.Formats); err != nil {
		res = append(res, err)
	}

	qAge, qhkAge, _ := qs.GetOK("age")
	if err := o.bindAge(qAge, qhkAge, route.Formats); err != nil {
		res = append(res, err)
	}

	qEnrollment, qhkEnrollment, _ := qs.GetOK("enrollment")
	if err := o.bindEnrollment(qEnrollment, qhkEnrollment, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDepartment binds and validates parameter Department from query.
func (o *PostStudentParams) bindDepartment(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Department = &raw

	return nil
}

// bindAge binds and validates parameter Age from query.
func (o *PostStudentParams) bindAge(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("age", "query", "int32", raw)
	}
	o.Age = &value

	return nil
}

// bindEnrollment binds and validates parameter Enrollment from query.
func (o *PostStudentParams) bindEnrollment(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("enrollment", "query", "int64", raw)
	}
	o.Enrollment = &value

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *PostStudentParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}
