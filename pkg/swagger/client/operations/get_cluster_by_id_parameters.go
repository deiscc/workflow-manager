package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetClusterByIDParams creates a new GetClusterByIDParams object
// with the default values initialized.
func NewGetClusterByIDParams() *GetClusterByIDParams {
	var ()
	return &GetClusterByIDParams{}
}

/*GetClusterByIDParams contains all the parameters to send to the API endpoint
for the get cluster by id operation typically these are written to a http.Request
*/
type GetClusterByIDParams struct {

	/*ID*/
	ID string
}

// WithID adds the id to the get cluster by id params
func (o *GetClusterByIDParams) WithID(id string) *GetClusterByIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterByIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}