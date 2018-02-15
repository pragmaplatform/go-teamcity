// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGetPropertiesParams creates a new GetPropertiesParams object
// with the default values initialized.
func NewGetPropertiesParams() *GetPropertiesParams {
	var ()
	return &GetPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPropertiesParamsWithTimeout creates a new GetPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPropertiesParamsWithTimeout(timeout time.Duration) *GetPropertiesParams {
	var ()
	return &GetPropertiesParams{

		timeout: timeout,
	}
}

// NewGetPropertiesParamsWithContext creates a new GetPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPropertiesParamsWithContext(ctx context.Context) *GetPropertiesParams {
	var ()
	return &GetPropertiesParams{

		Context: ctx,
	}
}

// NewGetPropertiesParamsWithHTTPClient creates a new GetPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPropertiesParamsWithHTTPClient(client *http.Client) *GetPropertiesParams {
	var ()
	return &GetPropertiesParams{
		HTTPClient: client,
	}
}

/*GetPropertiesParams contains all the parameters to send to the API endpoint
for the get properties operation typically these are written to a http.Request
*/
type GetPropertiesParams struct {

	/*Fields*/
	Fields *string
	/*GroupLocator*/
	GroupLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get properties params
func (o *GetPropertiesParams) WithTimeout(timeout time.Duration) *GetPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get properties params
func (o *GetPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get properties params
func (o *GetPropertiesParams) WithContext(ctx context.Context) *GetPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get properties params
func (o *GetPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get properties params
func (o *GetPropertiesParams) WithHTTPClient(client *http.Client) *GetPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get properties params
func (o *GetPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get properties params
func (o *GetPropertiesParams) WithFields(fields *string) *GetPropertiesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get properties params
func (o *GetPropertiesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupLocator adds the groupLocator to the get properties params
func (o *GetPropertiesParams) WithGroupLocator(groupLocator string) *GetPropertiesParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the get properties params
func (o *GetPropertiesParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupLocator
	if err := r.SetPathParam("groupLocator", o.GroupLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
