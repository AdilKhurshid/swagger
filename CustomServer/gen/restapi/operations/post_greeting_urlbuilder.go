// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// PostGreetingURL generates an URL for the post greeting operation
type PostGreetingURL struct {
	Department *string
	Age        *int32
	Enrollment *int64
	Name       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostGreetingURL) WithBasePath(bp string) *PostGreetingURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostGreetingURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostGreetingURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/hello"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var department string
	if o.Department != nil {
		department = *o.Department
	}
	if department != "" {
		qs.Set("Department", department)
	}

	var age string
	if o.Age != nil {
		age = swag.FormatInt32(*o.Age)
	}
	if age != "" {
		qs.Set("age", age)
	}

	var enrollment string
	if o.Enrollment != nil {
		enrollment = swag.FormatInt64(*o.Enrollment)
	}
	if enrollment != "" {
		qs.Set("enrollment", enrollment)
	}

	var name string
	if o.Name != nil {
		name = *o.Name
	}
	if name != "" {
		qs.Set("name", name)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostGreetingURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostGreetingURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostGreetingURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostGreetingURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostGreetingURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PostGreetingURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}