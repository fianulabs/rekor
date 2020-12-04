// Code generated by go-swagger; DO NOT EDIT.

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLogInfoHandlerFunc turns a function with the right signature into a get log info handler
type GetLogInfoHandlerFunc func(GetLogInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogInfoHandlerFunc) Handle(params GetLogInfoParams) middleware.Responder {
	return fn(params)
}

// GetLogInfoHandler interface for that can handle valid get log info params
type GetLogInfoHandler interface {
	Handle(GetLogInfoParams) middleware.Responder
}

// NewGetLogInfo creates a new http.Handler for the get log info operation
func NewGetLogInfo(ctx *middleware.Context, handler GetLogInfoHandler) *GetLogInfo {
	return &GetLogInfo{Context: ctx, Handler: handler}
}

/*GetLogInfo swagger:route GET /api/v1/log tlog getLogInfo

Get information about the current state of the transparency log

Returns the current root hash and size of the merkle tree used to store the log entries.

*/
type GetLogInfo struct {
	Context *middleware.Context
	Handler GetLogInfoHandler
}

func (o *GetLogInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLogInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
