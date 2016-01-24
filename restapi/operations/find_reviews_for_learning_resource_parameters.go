package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewFindReviewsForLearningResourceParams creates a new FindReviewsForLearningResourceParams object
// with the default values initialized.
func NewFindReviewsForLearningResourceParams() FindReviewsForLearningResourceParams {
	var ()
	return FindReviewsForLearningResourceParams{}
}

// FindReviewsForLearningResourceParams contains all the bound params for the find reviews for learning resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters findReviewsForLearningResource
type FindReviewsForLearningResourceParams struct {
	/*ID of learning resource to fetch reviews for
	  Required: true
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindReviewsForLearningResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindReviewsForLearningResourceParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}