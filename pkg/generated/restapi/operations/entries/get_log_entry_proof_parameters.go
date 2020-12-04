// Code generated by go-swagger; DO NOT EDIT.

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetLogEntryProofParams creates a new GetLogEntryProofParams object
// no default values defined in spec.
func NewGetLogEntryProofParams() GetLogEntryProofParams {

	return GetLogEntryProofParams{}
}

// GetLogEntryProofParams contains all the bound params for the get log entry proof operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLogEntryProof
type GetLogEntryProofParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the UUID of the entry for which the inclusion proof information should be returned
	  Required: true
	  In: path
	*/
	EntryUUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLogEntryProofParams() beforehand.
func (o *GetLogEntryProofParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEntryUUID, rhkEntryUUID, _ := route.Params.GetOK("entryUUID")
	if err := o.bindEntryUUID(rEntryUUID, rhkEntryUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEntryUUID binds and validates parameter EntryUUID from path.
func (o *GetLogEntryProofParams) bindEntryUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.EntryUUID = raw

	return nil
}
