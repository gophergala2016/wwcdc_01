package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*AddReviewOK review response

swagger:response addReviewOK
*/
type AddReviewOK struct {

	// In: body
	Payload *models.Review `json:"body,omitempty"`
}

// NewAddReviewOK creates AddReviewOK with default headers values
func NewAddReviewOK() *AddReviewOK {
	return &AddReviewOK{}
}

// WithPayload adds the payload to the add review o k response
func (o *AddReviewOK) WithPayload(payload *models.Review) *AddReviewOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddReviewOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddReviewDefault unexpected error

swagger:response addReviewDefault
*/
type AddReviewDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewAddReviewDefault creates AddReviewDefault with default headers values
func NewAddReviewDefault(code int) *AddReviewDefault {
	if code <= 0 {
		code = 500
	}

	return &AddReviewDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add review default response
func (o *AddReviewDefault) WithStatusCode(code int) *AddReviewDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the add review default response
func (o *AddReviewDefault) WithPayload(payload *models.ErrorModel) *AddReviewDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddReviewDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
