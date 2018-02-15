// Code generated by go-swagger; DO NOT EDIT.

package mute

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

// NewGetMutesParams creates a new GetMutesParams object
// with the default values initialized.
func NewGetMutesParams() *GetMutesParams {
	var ()
	return &GetMutesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMutesParamsWithTimeout creates a new GetMutesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMutesParamsWithTimeout(timeout time.Duration) *GetMutesParams {
	var ()
	return &GetMutesParams{

		timeout: timeout,
	}
}

// NewGetMutesParamsWithContext creates a new GetMutesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMutesParamsWithContext(ctx context.Context) *GetMutesParams {
	var ()
	return &GetMutesParams{

		Context: ctx,
	}
}

// NewGetMutesParamsWithHTTPClient creates a new GetMutesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMutesParamsWithHTTPClient(client *http.Client) *GetMutesParams {
	var ()
	return &GetMutesParams{
		HTTPClient: client,
	}
}

/*GetMutesParams contains all the parameters to send to the API endpoint
for the get mutes operation typically these are written to a http.Request
*/
type GetMutesParams struct {

	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mutes params
func (o *GetMutesParams) WithTimeout(timeout time.Duration) *GetMutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mutes params
func (o *GetMutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mutes params
func (o *GetMutesParams) WithContext(ctx context.Context) *GetMutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mutes params
func (o *GetMutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mutes params
func (o *GetMutesParams) WithHTTPClient(client *http.Client) *GetMutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mutes params
func (o *GetMutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get mutes params
func (o *GetMutesParams) WithFields(fields *string) *GetMutesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get mutes params
func (o *GetMutesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get mutes params
func (o *GetMutesParams) WithLocator(locator *string) *GetMutesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get mutes params
func (o *GetMutesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetMutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string
		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {
			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
