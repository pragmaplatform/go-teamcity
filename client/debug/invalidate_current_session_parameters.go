// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewInvalidateCurrentSessionParams creates a new InvalidateCurrentSessionParams object
// with the default values initialized.
func NewInvalidateCurrentSessionParams() *InvalidateCurrentSessionParams {

	return &InvalidateCurrentSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidateCurrentSessionParamsWithTimeout creates a new InvalidateCurrentSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvalidateCurrentSessionParamsWithTimeout(timeout time.Duration) *InvalidateCurrentSessionParams {

	return &InvalidateCurrentSessionParams{

		timeout: timeout,
	}
}

// NewInvalidateCurrentSessionParamsWithContext creates a new InvalidateCurrentSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvalidateCurrentSessionParamsWithContext(ctx context.Context) *InvalidateCurrentSessionParams {

	return &InvalidateCurrentSessionParams{

		Context: ctx,
	}
}

// NewInvalidateCurrentSessionParamsWithHTTPClient creates a new InvalidateCurrentSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvalidateCurrentSessionParamsWithHTTPClient(client *http.Client) *InvalidateCurrentSessionParams {

	return &InvalidateCurrentSessionParams{
		HTTPClient: client,
	}
}

/*InvalidateCurrentSessionParams contains all the parameters to send to the API endpoint
for the invalidate current session operation typically these are written to a http.Request
*/
type InvalidateCurrentSessionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invalidate current session params
func (o *InvalidateCurrentSessionParams) WithTimeout(timeout time.Duration) *InvalidateCurrentSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidate current session params
func (o *InvalidateCurrentSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidate current session params
func (o *InvalidateCurrentSessionParams) WithContext(ctx context.Context) *InvalidateCurrentSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidate current session params
func (o *InvalidateCurrentSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidate current session params
func (o *InvalidateCurrentSessionParams) WithHTTPClient(client *http.Client) *InvalidateCurrentSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidate current session params
func (o *InvalidateCurrentSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidateCurrentSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
