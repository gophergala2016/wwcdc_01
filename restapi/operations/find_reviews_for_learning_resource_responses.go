package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*FindReviewsForLearningResourceOK learning resource response

swagger:response findReviewsForLearningResourceOK
*/
type FindReviewsForLearningResourceOK struct {

	// In: body
	Payload []*models.Review `json:"body,omitempty"`
}

// NewFindReviewsForLearningResourceOK creates FindReviewsForLearningResourceOK with default headers values
func NewFindReviewsForLearningResourceOK() *FindReviewsForLearningResourceOK {
	return &FindReviewsForLearningResourceOK{}
}

// WithPayload adds the payload to the find reviews for learning resource o k response
func (o *FindReviewsForLearningResourceOK) WithPayload(payload []*models.Review) *FindReviewsForLearningResourceOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindReviewsForLearningResourceOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindReviewsForLearningResourceDefault unexpected error

swagger:response findReviewsForLearningResourceDefault
*/
type FindReviewsForLearningResourceDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewFindReviewsForLearningResourceDefault creates FindReviewsForLearningResourceDefault with default headers values
func NewFindReviewsForLearningResourceDefault(code int) *FindReviewsForLearningResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &FindReviewsForLearningResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find reviews for learning resource default response
func (o *FindReviewsForLearningResourceDefault) WithStatusCode(code int) *FindReviewsForLearningResourceDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the find reviews for learning resource default response
func (o *FindReviewsForLearningResourceDefault) WithPayload(payload *models.ErrorModel) *FindReviewsForLearningResourceDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *FindReviewsForLearningResourceDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
