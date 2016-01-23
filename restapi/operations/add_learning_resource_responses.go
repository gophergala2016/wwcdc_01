package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*AddLearningResourceOK learning-resource response

swagger:response addLearningResourceOK
*/
type AddLearningResourceOK struct {

	// In: body
	Payload *models.LearningResource `json:"body,omitempty"`
}

// NewAddLearningResourceOK creates AddLearningResourceOK with default headers values
func NewAddLearningResourceOK() *AddLearningResourceOK {
	return &AddLearningResourceOK{}
}

// WithPayload adds the payload to the add learning resource o k response
func (o *AddLearningResourceOK) WithPayload(payload *models.LearningResource) *AddLearningResourceOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddLearningResourceOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddLearningResourceDefault unexpected error

swagger:response addLearningResourceDefault
*/
type AddLearningResourceDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewAddLearningResourceDefault creates AddLearningResourceDefault with default headers values
func NewAddLearningResourceDefault(code int) *AddLearningResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &AddLearningResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add learning resource default response
func (o *AddLearningResourceDefault) WithStatusCode(code int) *AddLearningResourceDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the add learning resource default response
func (o *AddLearningResourceDefault) WithPayload(payload *models.ErrorModel) *AddLearningResourceDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddLearningResourceDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
