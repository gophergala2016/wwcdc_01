package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// AddReviewHandlerFunc turns a function with the right signature into a add review handler
type AddReviewHandlerFunc func(AddReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddReviewHandlerFunc) Handle(params AddReviewParams) middleware.Responder {
	return fn(params)
}

// AddReviewHandler interface for that can handle valid add review params
type AddReviewHandler interface {
	Handle(AddReviewParams) middleware.Responder
}

// NewAddReview creates a new http.Handler for the add review operation
func NewAddReview(ctx *middleware.Context, handler AddReviewHandler) *AddReview {
	return &AddReview{Context: ctx, Handler: handler}
}

/*AddReview swagger:route POST /reviews addReview

Creates a new review.

*/
type AddReview struct {
	Context *middleware.Context
	Params  AddReviewParams
	Handler AddReviewHandler
}

func (o *AddReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewAddReviewParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}