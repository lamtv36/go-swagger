// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPetListParams creates a new PetListParams object
// no default values defined in spec.
func NewPetListParams() PetListParams {

	return PetListParams{}
}

// PetListParams contains all the bound params for the pet list operation
// typically these are obtained from a http.Request
//
// swagger:parameters PetList
type PetListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Status values that need to be considered for filter
	  Required: true
	  In: query
	  Collection Format: multi
	*/
	Status []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPetListParams() beforehand.
func (o *PetListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindStatus binds and validates array parameter Status from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *PetListParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("status", "query", rawData)
	}

	// CollectionFormat: multi
	statusIC := rawData

	if len(statusIC) == 0 {
		return errors.Required("status", "query", statusIC)
	}

	var statusIR []string
	for i, statusIV := range statusIC {
		statusI := statusIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "status", i), "query", statusI, []interface{}{"available", "pending", "sold"}, true); err != nil {
			return err
		}

		statusIR = append(statusIR, statusI)
	}

	o.Status = statusIR

	return nil
}
