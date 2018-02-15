// Code generated by go-swagger; DO NOT EDIT.

package user

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
)

// NewServeUserPropertiesParams creates a new ServeUserPropertiesParams object
// with the default values initialized.
func NewServeUserPropertiesParams() *ServeUserPropertiesParams {
	var ()
	return &ServeUserPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeUserPropertiesParamsWithTimeout creates a new ServeUserPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeUserPropertiesParamsWithTimeout(timeout time.Duration) *ServeUserPropertiesParams {
	var ()
	return &ServeUserPropertiesParams{

		timeout: timeout,
	}
}

// NewServeUserPropertiesParamsWithContext creates a new ServeUserPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeUserPropertiesParamsWithContext(ctx context.Context) *ServeUserPropertiesParams {
	var ()
	return &ServeUserPropertiesParams{

		Context: ctx,
	}
}

// NewServeUserPropertiesParamsWithHTTPClient creates a new ServeUserPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeUserPropertiesParamsWithHTTPClient(client *http.Client) *ServeUserPropertiesParams {
	var ()
	return &ServeUserPropertiesParams{
		HTTPClient: client,
	}
}

/*ServeUserPropertiesParams contains all the parameters to send to the API endpoint
for the serve user properties operation typically these are written to a http.Request
*/
type ServeUserPropertiesParams struct {

	/*Fields*/
	Fields *string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve user properties params
func (o *ServeUserPropertiesParams) WithTimeout(timeout time.Duration) *ServeUserPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve user properties params
func (o *ServeUserPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve user properties params
func (o *ServeUserPropertiesParams) WithContext(ctx context.Context) *ServeUserPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve user properties params
func (o *ServeUserPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve user properties params
func (o *ServeUserPropertiesParams) WithHTTPClient(client *http.Client) *ServeUserPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve user properties params
func (o *ServeUserPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve user properties params
func (o *ServeUserPropertiesParams) WithFields(fields *string) *ServeUserPropertiesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve user properties params
func (o *ServeUserPropertiesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserLocator adds the userLocator to the serve user properties params
func (o *ServeUserPropertiesParams) WithUserLocator(userLocator string) *ServeUserPropertiesParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the serve user properties params
func (o *ServeUserPropertiesParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeUserPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
