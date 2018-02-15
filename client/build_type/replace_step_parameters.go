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

// NewReplaceStepParams creates a new ReplaceStepParams object
// with the default values initialized.
func NewReplaceStepParams() *ReplaceStepParams {
	var ()
	return &ReplaceStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStepParamsWithTimeout creates a new ReplaceStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceStepParamsWithTimeout(timeout time.Duration) *ReplaceStepParams {
	var ()
	return &ReplaceStepParams{

		timeout: timeout,
	}
}

// NewReplaceStepParamsWithContext creates a new ReplaceStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceStepParamsWithContext(ctx context.Context) *ReplaceStepParams {
	var ()
	return &ReplaceStepParams{

		Context: ctx,
	}
}

// NewReplaceStepParamsWithHTTPClient creates a new ReplaceStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceStepParamsWithHTTPClient(client *http.Client) *ReplaceStepParams {
	var ()
	return &ReplaceStepParams{
		HTTPClient: client,
	}
}

/*ReplaceStepParams contains all the parameters to send to the API endpoint
for the replace step operation typically these are written to a http.Request
*/
type ReplaceStepParams struct {

	/*Body*/
	Body *models.Step
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

// WithTimeout adds the timeout to the replace step params
func (o *ReplaceStepParams) WithTimeout(timeout time.Duration) *ReplaceStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace step params
func (o *ReplaceStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace step params
func (o *ReplaceStepParams) WithContext(ctx context.Context) *ReplaceStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace step params
func (o *ReplaceStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace step params
func (o *ReplaceStepParams) WithHTTPClient(client *http.Client) *ReplaceStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace step params
func (o *ReplaceStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace step params
func (o *ReplaceStepParams) WithBody(body *models.Step) *ReplaceStepParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace step params
func (o *ReplaceStepParams) SetBody(body *models.Step) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace step params
func (o *ReplaceStepParams) WithBtLocator(btLocator string) *ReplaceStepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace step params
func (o *ReplaceStepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace step params
func (o *ReplaceStepParams) WithFields(fields *string) *ReplaceStepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace step params
func (o *ReplaceStepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStepID adds the stepID to the replace step params
func (o *ReplaceStepParams) WithStepID(stepID string) *ReplaceStepParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the replace step params
func (o *ReplaceStepParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
