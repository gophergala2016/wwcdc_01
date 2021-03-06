package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// FindRecommendationForLearningResourceHandlerFunc turns a function with the right signature into a find recommendation for learning resource handler
type FindRecommendationForLearningResourceHandlerFunc func(FindRecommendationForLearningResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindRecommendationForLearningResourceHandlerFunc) Handle(params FindRecommendationForLearningResourceParams) middleware.Responder {
	return fn(params)
}

// FindRecommendationForLearningResourceHandler interface for that can handle valid find recommendation for learning resource params
type FindRecommendationForLearningResourceHandler interface {
	Handle(FindRecommendationForLearningResourceParams) middleware.Responder
}

// NewFindRecommendationForLearningResource creates a new http.Handler for the find recommendation for learning resource operation
func NewFindRecommendationForLearningResource(ctx *middleware.Context, handler FindRecommendationForLearningResourceHandler) *FindRecommendationForLearningResource {
	return &FindRecommendationForLearningResource{Context: ctx, Handler: handler}
}

/*FindRecommendationForLearningResource swagger:route GET /learning-resources/{id}/recommendations findRecommendationForLearningResource

Returns recommendations based on the user viewing and the current learning resource being viewed

*/
type FindRecommendationForLearningResource struct {
	Context *middleware.Context
	Params  FindRecommendationForLearningResourceParams
	Handler FindRecommendationForLearningResourceHandler
}

func (o *FindRecommendationForLearningResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewFindRecommendationForLearningResourceParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
