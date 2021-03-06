package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*ErrorModel ErrorModel error model

swagger:model ErrorModel
*/
type ErrorModel struct {

	/* Code code

	Required: true
	*/
	Code int32 `json:"code,omitempty"`

	/* Message message

	Required: true
	*/
	Message string `json:"message,omitempty"`
}

// Validate validates this error model
func (m *ErrorModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorModel) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", int32(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *ErrorModel) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}
