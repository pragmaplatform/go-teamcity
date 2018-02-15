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

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewUpdateVcsRootEntryParams creates a new UpdateVcsRootEntryParams object
// with the default values initialized.
func NewUpdateVcsRootEntryParams() *UpdateVcsRootEntryParams {
	var ()
	return &UpdateVcsRootEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVcsRootEntryParamsWithTimeout creates a new UpdateVcsRootEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVcsRootEntryParamsWithTimeout(timeout time.Duration) *UpdateVcsRootEntryParams {
	var ()
	return &UpdateVcsRootEntryParams{

		timeout: timeout,
	}
}

// NewUpdateVcsRootEntryParamsWithContext creates a new UpdateVcsRootEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVcsRootEntryParamsWithContext(ctx context.Context) *UpdateVcsRootEntryParams {
	var ()
	return &UpdateVcsRootEntryParams{

		Context: ctx,
	}
}

// NewUpdateVcsRootEntryParamsWithHTTPClient creates a new UpdateVcsRootEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVcsRootEntryParamsWithHTTPClient(client *http.Client) *UpdateVcsRootEntryParams {
	var ()
	return &UpdateVcsRootEntryParams{
		HTTPClient: client,
	}
}

/*UpdateVcsRootEntryParams contains all the parameters to send to the API endpoint
for the update vcs root entry operation typically these are written to a http.Request
*/
type UpdateVcsRootEntryParams struct {

	/*Body*/
	Body *models.VcsRootEntry
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithTimeout(timeout time.Duration) *UpdateVcsRootEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithContext(ctx context.Context) *UpdateVcsRootEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithHTTPClient(client *http.Client) *UpdateVcsRootEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithBody(body *models.VcsRootEntry) *UpdateVcsRootEntryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetBody(body *models.VcsRootEntry) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithBtLocator(btLocator string) *UpdateVcsRootEntryParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithFields(fields *string) *UpdateVcsRootEntryParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootLocator adds the vcsRootLocator to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) WithVcsRootLocator(vcsRootLocator string) *UpdateVcsRootEntryParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the update vcs root entry params
func (o *UpdateVcsRootEntryParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVcsRootEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
