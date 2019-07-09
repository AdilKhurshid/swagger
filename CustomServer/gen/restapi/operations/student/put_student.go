// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutStudentHandlerFunc turns a function with the right signature into a put student handler
type PutStudentHandlerFunc func(PutStudentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutStudentHandlerFunc) Handle(params PutStudentParams) middleware.Responder {
	return fn(params)
}

// PutStudentHandler interface for that can handle valid put student params
type PutStudentHandler interface {
	Handle(PutStudentParams) middleware.Responder
}

// NewPutStudent creates a new http.Handler for the put student operation
func NewPutStudent(ctx *middleware.Context, handler PutStudentHandler) *PutStudent {
	return &PutStudent{Context: ctx, Handler: handler}
}

/*PutStudent swagger:route PUT /{id} student putStudent

PutStudent put student API

*/
type PutStudent struct {
	Context *middleware.Context
	Handler PutStudentHandler
}

func (o *PutStudent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutStudentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
