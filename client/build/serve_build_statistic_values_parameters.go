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
)

// NewServeBuildStatisticValuesParams creates a new ServeBuildStatisticValuesParams object
// with the default values initialized.
func NewServeBuildStatisticValuesParams() *ServeBuildStatisticValuesParams {
	var ()
	return &ServeBuildStatisticValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildStatisticValuesParamsWithTimeout creates a new ServeBuildStatisticValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildStatisticValuesParamsWithTimeout(timeout time.Duration) *ServeBuildStatisticValuesParams {
	var ()
	return &ServeBuildStatisticValuesParams{

		timeout: timeout,
	}
}

// NewServeBuildStatisticValuesParamsWithContext creates a new ServeBuildStatisticValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildStatisticValuesParamsWithContext(ctx context.Context) *ServeBuildStatisticValuesParams {
	var ()
	return &ServeBuildStatisticValuesParams{

		Context: ctx,
	}
}

// NewServeBuildStatisticValuesParamsWithHTTPClient creates a new ServeBuildStatisticValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildStatisticValuesParamsWithHTTPClient(client *http.Client) *ServeBuildStatisticValuesParams {
	var ()
	return &ServeBuildStatisticValuesParams{
		HTTPClient: client,
	}
}

/*ServeBuildStatisticValuesParams contains all the parameters to send to the API endpoint
for the serve build statistic values operation typically these are written to a http.Request
*/
type ServeBuildStatisticValuesParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) WithTimeout(timeout time.Duration) *ServeBuildStatisticValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) WithContext(ctx context.Context) *ServeBuildStatisticValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) WithHTTPClient(client *http.Client) *ServeBuildStatisticValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) WithBuildLocator(buildLocator string) *ServeBuildStatisticValuesParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) WithFields(fields *string) *ServeBuildStatisticValuesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve build statistic values params
func (o *ServeBuildStatisticValuesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildStatisticValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
