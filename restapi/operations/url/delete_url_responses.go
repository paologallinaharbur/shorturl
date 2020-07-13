// Code generated by go-swagger; DO NOT EDIT.

package url

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/paologallinaharbur/shorturl/models"
)

// DeleteURLNoContentCode is the HTTP code returned for type DeleteURLNoContent
const DeleteURLNoContentCode int = 204

/*DeleteURLNoContent deleted

swagger:response deleteUrlNoContent
*/
type DeleteURLNoContent struct {
}

// NewDeleteURLNoContent creates DeleteURLNoContent with default headers values
func NewDeleteURLNoContent() *DeleteURLNoContent {

	return &DeleteURLNoContent{}
}

// WriteResponse to the client
func (o *DeleteURLNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteURLBadRequestCode is the HTTP code returned for type DeleteURLBadRequest
const DeleteURLBadRequestCode int = 400

/*DeleteURLBadRequest bad request

swagger:response deleteUrlBadRequest
*/
type DeleteURLBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteURLBadRequest creates DeleteURLBadRequest with default headers values
func NewDeleteURLBadRequest() *DeleteURLBadRequest {

	return &DeleteURLBadRequest{}
}

// WithPayload adds the payload to the delete Url bad request response
func (o *DeleteURLBadRequest) WithPayload(payload *models.Error) *DeleteURLBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Url bad request response
func (o *DeleteURLBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteURLBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteURLInternalServerErrorCode is the HTTP code returned for type DeleteURLInternalServerError
const DeleteURLInternalServerErrorCode int = 500

/*DeleteURLInternalServerError internal server error

swagger:response deleteUrlInternalServerError
*/
type DeleteURLInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteURLInternalServerError creates DeleteURLInternalServerError with default headers values
func NewDeleteURLInternalServerError() *DeleteURLInternalServerError {

	return &DeleteURLInternalServerError{}
}

// WithPayload adds the payload to the delete Url internal server error response
func (o *DeleteURLInternalServerError) WithPayload(payload *models.Error) *DeleteURLInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Url internal server error response
func (o *DeleteURLInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteURLInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
