package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/gophergala2016/wwcdc_01/models"
)

// NewAuthUserParams creates a new AuthUserParams object
// with the default values initialized.
func NewAuthUserParams() AuthUserParams {
	var ()
	return AuthUserParams{}
}

// AuthUserParams contains all the bound params for the auth user operation
// typically these are obtained from a http.Request
//
// swagger:parameters authUser
type AuthUserParams struct {
	/*User to login
	  Required: true
	  In: body
	*/
	User *models.UserAuth
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AuthUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body models.UserAuth
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("user", "body", "", err))
	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.User = &body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
