package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*FindReviewByIDOK review response

swagger:response findReviewByIdOK
*/
type FindReviewByIDOK struct {

	// In: body
	Payload *models.Review `json:"body,omitempty"`
}

// NewFindReviewByIDOK creates FindReviewByIDOK with default headers values
func NewFindReviewByIDOK() *FindReviewByIDOK {
	return &FindReviewByIDOK{}
}

// WithPayload adds the payload to the find review by id o k response
func (o *FindReviewByIDOK) WithPayload(payload *models.Review) *FindReviewByIDOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindReviewByIDOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindReviewByIDDefault unexpected error

swagger:response findReviewByIdDefault
*/
type FindReviewByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewFindReviewByIDDefault creates FindReviewByIDDefault with default headers values
func NewFindReviewByIDDefault(code int) *FindReviewByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &FindReviewByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find review by id default response
func (o *FindReviewByIDDefault) WithStatusCode(code int) *FindReviewByIDDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the find review by id default response
func (o *FindReviewByIDDefault) WithPayload(payload *models.ErrorModel) *FindReviewByIDDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindReviewByIDDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
