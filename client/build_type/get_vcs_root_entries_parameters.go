// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewGetVcsRootEntriesParams creates a new GetVcsRootEntriesParams object
// with the default values initialized.
func NewGetVcsRootEntriesParams() *GetVcsRootEntriesParams {
	var ()
	return &GetVcsRootEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootEntriesParamsWithTimeout creates a new GetVcsRootEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVcsRootEntriesParamsWithTimeout(timeout time.Duration) *GetVcsRootEntriesParams {
	var ()
	return &GetVcsRootEntriesParams{

		timeout: timeout,
	}
}

// NewGetVcsRootEntriesParamsWithContext creates a new GetVcsRootEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVcsRootEntriesParamsWithContext(ctx context.Context) *GetVcsRootEntriesParams {
	var ()
	return &GetVcsRootEntriesParams{

		Context: ctx,
	}
}

// NewGetVcsRootEntriesParamsWithHTTPClient creates a new GetVcsRootEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVcsRootEntriesParamsWithHTTPClient(client *http.Client) *GetVcsRootEntriesParams {
	var ()
	return &GetVcsRootEntriesParams{
		HTTPClient: client,
	}
}

/*GetVcsRootEntriesParams contains all the parameters to send to the API endpoint
for the get vcs root entries operation typically these are written to a http.Request
*/
type GetVcsRootEntriesParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vcs root entries params
func (o *GetVcsRootEntriesParams) WithTimeout(timeout time.Duration) *GetVcsRootEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root entries params
func (o *GetVcsRootEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root entries params
func (o *GetVcsRootEntriesParams) WithContext(ctx context.Context) *GetVcsRootEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root entries params
func (o *GetVcsRootEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root entries params
func (o *GetVcsRootEntriesParams) WithHTTPClient(client *http.Client) *GetVcsRootEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root entries params
func (o *GetVcsRootEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get vcs root entries params
func (o *GetVcsRootEntriesParams) WithBtLocator(btLocator string) *GetVcsRootEntriesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get vcs root entries params
func (o *GetVcsRootEntriesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get vcs root entries params
func (o *GetVcsRootEntriesParams) WithFields(fields *string) *GetVcsRootEntriesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root entries params
func (o *GetVcsRootEntriesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
