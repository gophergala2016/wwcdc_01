package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*DeleteReviewNoContent review deleted

swagger:response deleteReviewNoContent
*/
type DeleteReviewNoContent struct {
}

// NewDeleteReviewNoContent creates DeleteReviewNoContent with default headers values
func NewDeleteReviewNoContent() *DeleteReviewNoContent {
	return &DeleteReviewNoContent{}
}

// WriteResponse to the client
func (o *DeleteReviewNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DeleteReviewDefault unexpected error

swagger:response deleteReviewDefault
*/
type DeleteReviewDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewDeleteReviewDefault creates DeleteReviewDefault with default headers values
func NewDeleteReviewDefault(code int) *DeleteReviewDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteReviewDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete review default response
func (o *DeleteReviewDefault) WithStatusCode(code int) *DeleteReviewDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the delete review default response
func (o *DeleteReviewDefault) WithPayload(payload *models.ErrorModel) *DeleteReviewDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DeleteReviewDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}