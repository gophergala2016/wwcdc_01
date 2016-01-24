package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// DeleteReviewHandlerFunc turns a function with the right signature into a delete review handler
type DeleteReviewHandlerFunc func(DeleteReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReviewHandlerFunc) Handle(params DeleteReviewParams) middleware.Responder {
	return fn(params)
}

// DeleteReviewHandler interface for that can handle valid delete review params
type DeleteReviewHandler interface {
	Handle(DeleteReviewParams) middleware.Responder
}

// NewDeleteReview creates a new http.Handler for the delete review operation
func NewDeleteReview(ctx *middleware.Context, handler DeleteReviewHandler) *DeleteReview {
	return &DeleteReview{Context: ctx, Handler: handler}
}

/*DeleteReview swagger:route DELETE /reviews/{id} deleteReview

deletes a single review based on the ID supplied

*/
type DeleteReview struct {
	Context *middleware.Context
	Params  DeleteReviewParams
	Handler DeleteReviewHandler
}

func (o *DeleteReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewDeleteReviewParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}