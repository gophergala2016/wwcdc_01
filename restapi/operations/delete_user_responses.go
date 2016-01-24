package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/gophergala2016/wwcdc_01/models"
)

/*DeleteUserNoContent user deleted

swagger:response deleteUserNoContent
*/
type DeleteUserNoContent struct {
}

// NewDeleteUserNoContent creates DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DeleteUserDefault unexpected error

swagger:response deleteUserDefault
*/
type DeleteUserDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewDeleteUserDefault creates DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user default response
func (o *DeleteUserDefault) WithStatusCode(code int) *DeleteUserDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the delete user default response
func (o *DeleteUserDefault) WithPayload(payload *models.ErrorModel) *DeleteUserDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DeleteUserDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}