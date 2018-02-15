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

// NewReplaceStepParametersParams creates a new ReplaceStepParametersParams object
// with the default values initialized.
func NewReplaceStepParametersParams() *ReplaceStepParametersParams {
	var ()
	return &ReplaceStepParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStepParametersParamsWithTimeout creates a new ReplaceStepParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceStepParametersParamsWithTimeout(timeout time.Duration) *ReplaceStepParametersParams {
	var ()
	return &ReplaceStepParametersParams{

		timeout: timeout,
	}
}

// NewReplaceStepParametersParamsWithContext creates a new ReplaceStepParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceStepParametersParamsWithContext(ctx context.Context) *ReplaceStepParametersParams {
	var ()
	return &ReplaceStepParametersParams{

		Context: ctx,
	}
}

// NewReplaceStepParametersParamsWithHTTPClient creates a new ReplaceStepParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceStepParametersParamsWithHTTPClient(client *http.Client) *ReplaceStepParametersParams {
	var ()
	return &ReplaceStepParametersParams{
		HTTPClient: client,
	}
}

/*ReplaceStepParametersParams contains all the parameters to send to the API endpoint
for the replace step parameters operation typically these are written to a http.Request
*/
type ReplaceStepParametersParams struct {

	/*Body*/
	Body *models.Properties
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace step parameters params
func (o *ReplaceStepParametersParams) WithTimeout(timeout time.Duration) *ReplaceStepParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace step parameters params
func (o *ReplaceStepParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace step parameters params
func (o *ReplaceStepParametersParams) WithContext(ctx context.Context) *ReplaceStepParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace step parameters params
func (o *ReplaceStepParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace step parameters params
func (o *ReplaceStepParametersParams) WithHTTPClient(client *http.Client) *ReplaceStepParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace step parameters params
func (o *ReplaceStepParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace step parameters params
func (o *ReplaceStepParametersParams) WithBody(body *models.Properties) *ReplaceStepParametersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace step parameters params
func (o *ReplaceStepParametersParams) SetBody(body *models.Properties) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace step parameters params
func (o *ReplaceStepParametersParams) WithBtLocator(btLocator string) *ReplaceStepParametersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace step parameters params
func (o *ReplaceStepParametersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace step parameters params
func (o *ReplaceStepParametersParams) WithFields(fields *string) *ReplaceStepParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace step parameters params
func (o *ReplaceStepParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStepID adds the stepID to the replace step parameters params
func (o *ReplaceStepParametersParams) WithStepID(stepID string) *ReplaceStepParametersParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the replace step parameters params
func (o *ReplaceStepParametersParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStepParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stepId
	if err := r.SetPathParam("stepId", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
