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

// NewReplaceFeatureParametersParams creates a new ReplaceFeatureParametersParams object
// with the default values initialized.
func NewReplaceFeatureParametersParams() *ReplaceFeatureParametersParams {
	var ()
	return &ReplaceFeatureParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceFeatureParametersParamsWithTimeout creates a new ReplaceFeatureParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceFeatureParametersParamsWithTimeout(timeout time.Duration) *ReplaceFeatureParametersParams {
	var ()
	return &ReplaceFeatureParametersParams{

		timeout: timeout,
	}
}

// NewReplaceFeatureParametersParamsWithContext creates a new ReplaceFeatureParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceFeatureParametersParamsWithContext(ctx context.Context) *ReplaceFeatureParametersParams {
	var ()
	return &ReplaceFeatureParametersParams{

		Context: ctx,
	}
}

// NewReplaceFeatureParametersParamsWithHTTPClient creates a new ReplaceFeatureParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceFeatureParametersParamsWithHTTPClient(client *http.Client) *ReplaceFeatureParametersParams {
	var ()
	return &ReplaceFeatureParametersParams{
		HTTPClient: client,
	}
}

/*ReplaceFeatureParametersParams contains all the parameters to send to the API endpoint
for the replace feature parameters operation typically these are written to a http.Request
*/
type ReplaceFeatureParametersParams struct {

	/*Body*/
	Body *models.Properties
	/*BtLocator*/
	BtLocator string
	/*FeatureID*/
	FeatureID string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithTimeout(timeout time.Duration) *ReplaceFeatureParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithContext(ctx context.Context) *ReplaceFeatureParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithHTTPClient(client *http.Client) *ReplaceFeatureParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithBody(body *models.Properties) *ReplaceFeatureParametersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetBody(body *models.Properties) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithBtLocator(btLocator string) *ReplaceFeatureParametersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFeatureID adds the featureID to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithFeatureID(featureID string) *ReplaceFeatureParametersParams {
	o.SetFeatureID(featureID)
	return o
}

// SetFeatureID adds the featureId to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetFeatureID(featureID string) {
	o.FeatureID = featureID
}

// WithFields adds the fields to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) WithFields(fields *string) *ReplaceFeatureParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace feature parameters params
func (o *ReplaceFeatureParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceFeatureParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param featureId
	if err := r.SetPathParam("featureId", o.FeatureID); err != nil {
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
