// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutGreetingHandlerFunc turns a function with the right signature into a put greeting handler
type PutGreetingHandlerFunc func(PutGreetingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutGreetingHandlerFunc) Handle(params PutGreetingParams) middleware.Responder {
	return fn(params)
}

// PutGreetingHandler interface for that can handle valid put greeting params
type PutGreetingHandler interface {
	Handle(PutGreetingParams) middleware.Responder
}

// NewPutGreeting creates a new http.Handler for the put greeting operation
func NewPutGreeting(ctx *middleware.Context, handler PutGreetingHandler) *PutGreeting {
	return &PutGreeting{Context: ctx, Handler: handler}
}

/*PutGreeting swagger:route PUT /hello putGreeting

PutGreeting put greeting API

*/
type PutGreeting struct {
	Context *middleware.Context
	Handler PutGreetingHandler
}

func (o *PutGreeting) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutGreetingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
