package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// AddUserHandlerFunc turns a function with the right signature into a add user handler
type AddUserHandlerFunc func(AddUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddUserHandlerFunc) Handle(params AddUserParams) middleware.Responder {
	return fn(params)
}

// AddUserHandler interface for that can handle valid add user params
type AddUserHandler interface {
	Handle(AddUserParams) middleware.Responder
}

// NewAddUser creates a new http.Handler for the add user operation
func NewAddUser(ctx *middleware.Context, handler AddUserHandler) *AddUser {
	return &AddUser{Context: ctx, Handler: handler}
}

/*AddUser swagger:route POST /users addUser

Registers a new user.

*/
type AddUser struct {
	Context *middleware.Context
	Params  AddUserParams
	Handler AddUserHandler
}

func (o *AddUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewAddUserParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
