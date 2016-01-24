package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// FindUserByIDHandlerFunc turns a function with the right signature into a find user by id handler
type FindUserByIDHandlerFunc func(FindUserByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindUserByIDHandlerFunc) Handle(params FindUserByIDParams) middleware.Responder {
	return fn(params)
}

// FindUserByIDHandler interface for that can handle valid find user by id params
type FindUserByIDHandler interface {
	Handle(FindUserByIDParams) middleware.Responder
}

// NewFindUserByID creates a new http.Handler for the find user by id operation
func NewFindUserByID(ctx *middleware.Context, handler FindUserByIDHandler) *FindUserByID {
	return &FindUserByID{Context: ctx, Handler: handler}
}

/*FindUserByID swagger:route GET /users/{id} findUserById

Returns a user based on a single ID, if the user has access to the user

*/
type FindUserByID struct {
	Context *middleware.Context
	Params  FindUserByIDParams
	Handler FindUserByIDHandler
}

func (o *FindUserByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewFindUserByIDParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
