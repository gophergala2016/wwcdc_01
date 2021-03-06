package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// FindReviewsHandlerFunc turns a function with the right signature into a find reviews handler
type FindReviewsHandlerFunc func() middleware.Responder

// Handle executing the request and returning a response
func (fn FindReviewsHandlerFunc) Handle() middleware.Responder {
	return fn()
}

// FindReviewsHandler interface for that can handle valid find reviews params
type FindReviewsHandler interface {
	Handle() middleware.Responder
}

// NewFindReviews creates a new http.Handler for the find reviews operation
func NewFindReviews(ctx *middleware.Context, handler FindReviewsHandler) *FindReviews {
	return &FindReviews{Context: ctx, Handler: handler}
}

/*FindReviews swagger:route GET /reviews findReviews

Returns recent reviews

*/
type FindReviews struct {
	Context *middleware.Context
	Handler FindReviewsHandler
}

func (o *FindReviews) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle() // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
