// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetStudentListOKCode is the HTTP code returned for type GetStudentListOK
const GetStudentListOKCode int = 200

/*GetStudentListOK returns list of Student

swagger:response getStudentListOK
*/
type GetStudentListOK struct {

	/*contains list of Student
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetStudentListOK creates GetStudentListOK with default headers values
func NewGetStudentListOK() *GetStudentListOK {

	return &GetStudentListOK{}
}

// WithPayload adds the payload to the get student list o k response
func (o *GetStudentListOK) WithPayload(payload string) *GetStudentListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get student list o k response
func (o *GetStudentListOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStudentListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}