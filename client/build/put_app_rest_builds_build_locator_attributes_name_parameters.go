// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewPutAppRestBuildsBuildLocatorAttributesNameParams creates a new PutAppRestBuildsBuildLocatorAttributesNameParams object
// with the default values initialized.
func NewPutAppRestBuildsBuildLocatorAttributesNameParams() *PutAppRestBuildsBuildLocatorAttributesNameParams {
	var ()
	return &PutAppRestBuildsBuildLocatorAttributesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithTimeout creates a new PutAppRestBuildsBuildLocatorAttributesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithTimeout(timeout time.Duration) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	var ()
	return &PutAppRestBuildsBuildLocatorAttributesNameParams{

		timeout: timeout,
	}
}

// NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithContext creates a new PutAppRestBuildsBuildLocatorAttributesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithContext(ctx context.Context) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	var ()
	return &PutAppRestBuildsBuildLocatorAttributesNameParams{

		Context: ctx,
	}
}

// NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithHTTPClient creates a new PutAppRestBuildsBuildLocatorAttributesNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAppRestBuildsBuildLocatorAttributesNameParamsWithHTTPClient(client *http.Client) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	var ()
	return &PutAppRestBuildsBuildLocatorAttributesNameParams{
		HTTPClient: client,
	}
}

/*PutAppRestBuildsBuildLocatorAttributesNameParams contains all the parameters to send to the API endpoint
for the put app rest builds build locator attributes name operation typically these are written to a http.Request
*/
type PutAppRestBuildsBuildLocatorAttributesNameParams struct {

	/*Body*/
	Body *models.Property
	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithTimeout(timeout time.Duration) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithContext(ctx context.Context) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithHTTPClient(client *http.Client) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithBody(body *models.Property) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithBuildLocator(buildLocator string) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithFields(fields *string) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WithName(name string) *PutAppRestBuildsBuildLocatorAttributesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put app rest builds build locator attributes name params
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutAppRestBuildsBuildLocatorAttributesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
