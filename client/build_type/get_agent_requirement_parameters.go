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

// NewGetAgentRequirementParams creates a new GetAgentRequirementParams object
// with the default values initialized.
func NewGetAgentRequirementParams() *GetAgentRequirementParams {
	var ()
	return &GetAgentRequirementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentRequirementParamsWithTimeout creates a new GetAgentRequirementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAgentRequirementParamsWithTimeout(timeout time.Duration) *GetAgentRequirementParams {
	var ()
	return &GetAgentRequirementParams{

		timeout: timeout,
	}
}

// NewGetAgentRequirementParamsWithContext creates a new GetAgentRequirementParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAgentRequirementParamsWithContext(ctx context.Context) *GetAgentRequirementParams {
	var ()
	return &GetAgentRequirementParams{

		Context: ctx,
	}
}

// NewGetAgentRequirementParamsWithHTTPClient creates a new GetAgentRequirementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAgentRequirementParamsWithHTTPClient(client *http.Client) *GetAgentRequirementParams {
	var ()
	return &GetAgentRequirementParams{
		HTTPClient: client,
	}
}

/*GetAgentRequirementParams contains all the parameters to send to the API endpoint
for the get agent requirement operation typically these are written to a http.Request
*/
type GetAgentRequirementParams struct {

	/*AgentRequirementLocator*/
	AgentRequirementLocator string
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get agent requirement params
func (o *GetAgentRequirementParams) WithTimeout(timeout time.Duration) *GetAgentRequirementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent requirement params
func (o *GetAgentRequirementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent requirement params
func (o *GetAgentRequirementParams) WithContext(ctx context.Context) *GetAgentRequirementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent requirement params
func (o *GetAgentRequirementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent requirement params
func (o *GetAgentRequirementParams) WithHTTPClient(client *http.Client) *GetAgentRequirementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent requirement params
func (o *GetAgentRequirementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentRequirementLocator adds the agentRequirementLocator to the get agent requirement params
func (o *GetAgentRequirementParams) WithAgentRequirementLocator(agentRequirementLocator string) *GetAgentRequirementParams {
	o.SetAgentRequirementLocator(agentRequirementLocator)
	return o
}

// SetAgentRequirementLocator adds the agentRequirementLocator to the get agent requirement params
func (o *GetAgentRequirementParams) SetAgentRequirementLocator(agentRequirementLocator string) {
	o.AgentRequirementLocator = agentRequirementLocator
}

// WithBtLocator adds the btLocator to the get agent requirement params
func (o *GetAgentRequirementParams) WithBtLocator(btLocator string) *GetAgentRequirementParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get agent requirement params
func (o *GetAgentRequirementParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get agent requirement params
func (o *GetAgentRequirementParams) WithFields(fields *string) *GetAgentRequirementParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get agent requirement params
func (o *GetAgentRequirementParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentRequirementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentRequirementLocator
	if err := r.SetPathParam("agentRequirementLocator", o.AgentRequirementLocator); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
