package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deiscc/workflow-manager/pkg/swagger/models"
)

// NewPublishDoctorInfoParams creates a new PublishDoctorInfoParams object
// with the default values initialized.
func NewPublishDoctorInfoParams() *PublishDoctorInfoParams {
	var ()
	return &PublishDoctorInfoParams{}
}

/*PublishDoctorInfoParams contains all the parameters to send to the API endpoint
for the publish doctor info operation typically these are written to a http.Request
*/
type PublishDoctorInfoParams struct {

	/*Body*/
	Body *models.DoctorInfo
	/*UUID
	  A universal Id to represent a sepcific request or report

	*/
	UUID string
}

// WithBody adds the body to the publish doctor info params
func (o *PublishDoctorInfoParams) WithBody(body *models.DoctorInfo) *PublishDoctorInfoParams {
	o.Body = body
	return o
}

// WithUUID adds the uuid to the publish doctor info params
func (o *PublishDoctorInfoParams) WithUUID(uuid string) *PublishDoctorInfoParams {
	o.UUID = uuid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PublishDoctorInfoParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.DoctorInfo)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
