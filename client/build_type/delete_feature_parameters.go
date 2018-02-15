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

// NewDeleteFeatureParams creates a new DeleteFeatureParams object
// with the default values initialized.
func NewDeleteFeatureParams() *DeleteFeatureParams {
	var ()
	return &DeleteFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeatureParamsWithTimeout creates a new DeleteFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFeatureParamsWithTimeout(timeout time.Duration) *DeleteFeatureParams {
	var ()
	return &DeleteFeatureParams{

		timeout: timeout,
	}
}

// NewDeleteFeatureParamsWithContext creates a new DeleteFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFeatureParamsWithContext(ctx context.Context) *DeleteFeatureParams {
	var ()
	return &DeleteFeatureParams{

		Context: ctx,
	}
}

// NewDeleteFeatureParamsWithHTTPClient creates a new DeleteFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFeatureParamsWithHTTPClient(client *http.Client) *DeleteFeatureParams {
	var ()
	return &DeleteFeatureParams{
		HTTPClient: client,
	}
}

/*DeleteFeatureParams contains all the parameters to send to the API endpoint
for the delete feature operation typically these are written to a http.Request
*/
type DeleteFeatureParams struct {

	/*BtLocator*/
	BtLocator string
	/*FeatureID*/
	FeatureID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feature params
func (o *DeleteFeatureParams) WithTimeout(timeout time.Duration) *DeleteFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feature params
func (o *DeleteFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feature params
func (o *DeleteFeatureParams) WithContext(ctx context.Context) *DeleteFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feature params
func (o *DeleteFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feature params
func (o *DeleteFeatureParams) WithHTTPClient(client *http.Client) *DeleteFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feature params
func (o *DeleteFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the delete feature params
func (o *DeleteFeatureParams) WithBtLocator(btLocator string) *DeleteFeatureParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete feature params
func (o *DeleteFeatureParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFeatureID adds the featureID to the delete feature params
func (o *DeleteFeatureParams) WithFeatureID(featureID string) *DeleteFeatureParams {
	o.SetFeatureID(featureID)
	return o
}

// SetFeatureID adds the featureId to the delete feature params
func (o *DeleteFeatureParams) SetFeatureID(featureID string) {
	o.FeatureID = featureID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param featureId
	if err := r.SetPathParam("featureId", o.FeatureID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
