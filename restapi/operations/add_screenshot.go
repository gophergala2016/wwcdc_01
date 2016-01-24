package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// AddScreenshotHandlerFunc turns a function with the right signature into a add screenshot handler
type AddScreenshotHandlerFunc func(AddScreenshotParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddScreenshotHandlerFunc) Handle(params AddScreenshotParams) middleware.Responder {
	return fn(params)
}

// AddScreenshotHandler interface for that can handle valid add screenshot params
type AddScreenshotHandler interface {
	Handle(AddScreenshotParams) middleware.Responder
}

// NewAddScreenshot creates a new http.Handler for the add screenshot operation
func NewAddScreenshot(ctx *middleware.Context, handler AddScreenshotHandler) *AddScreenshot {
	return &AddScreenshot{Context: ctx, Handler: handler}
}

/*AddScreenshot swagger:route POST /screenshot/{id} addScreenshot

Adds a screenshot to the learning resource id specified

*/
type AddScreenshot struct {
	Context *middleware.Context
	Params  AddScreenshotParams
	Handler AddScreenshotHandler
}

func (o *AddScreenshot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewAddScreenshotParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
