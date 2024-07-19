// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/to-com/poc-td/app/swagger/restmodel"
)

// ListToteAssignmentsOKCode is the HTTP code returned for type ListToteAssignmentsOK
const ListToteAssignmentsOKCode int = 200

/*ListToteAssignmentsOK Successful response.

swagger:response listToteAssignmentsOK
*/
type ListToteAssignmentsOK struct {

	/*
	  In: Body
	*/
	Payload *restmodel.ListToteAssignmentsResponse `json:"body,omitempty"`
}

// NewListToteAssignmentsOK creates ListToteAssignmentsOK with default headers values
func NewListToteAssignmentsOK() *ListToteAssignmentsOK {

	return &ListToteAssignmentsOK{}
}

// WithPayload adds the payload to the list tote assignments o k response
func (o *ListToteAssignmentsOK) WithPayload(payload *restmodel.ListToteAssignmentsResponse) *ListToteAssignmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list tote assignments o k response
func (o *ListToteAssignmentsOK) SetPayload(payload *restmodel.ListToteAssignmentsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListToteAssignmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListToteAssignmentsBadRequestCode is the HTTP code returned for type ListToteAssignmentsBadRequest
const ListToteAssignmentsBadRequestCode int = 400

/*ListToteAssignmentsBadRequest Bad Request.

swagger:response listToteAssignmentsBadRequest
*/
type ListToteAssignmentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *restmodel.Response400 `json:"body,omitempty"`
}

// NewListToteAssignmentsBadRequest creates ListToteAssignmentsBadRequest with default headers values
func NewListToteAssignmentsBadRequest() *ListToteAssignmentsBadRequest {

	return &ListToteAssignmentsBadRequest{}
}

// WithPayload adds the payload to the list tote assignments bad request response
func (o *ListToteAssignmentsBadRequest) WithPayload(payload *restmodel.Response400) *ListToteAssignmentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list tote assignments bad request response
func (o *ListToteAssignmentsBadRequest) SetPayload(payload *restmodel.Response400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListToteAssignmentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}