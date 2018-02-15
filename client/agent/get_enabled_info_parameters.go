// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewGetEnabledInfoParams creates a new GetEnabledInfoParams object
// with the default values initialized.
func NewGetEnabledInfoParams() *GetEnabledInfoParams {
	var ()
	return &GetEnabledInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnabledInfoParamsWithTimeout creates a new GetEnabledInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnabledInfoParamsWithTimeout(timeout time.Duration) *GetEnabledInfoParams {
	var ()
	return &GetEnabledInfoParams{

		timeout: timeout,
	}
}

// NewGetEnabledInfoParamsWithContext creates a new GetEnabledInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnabledInfoParamsWithContext(ctx context.Context) *GetEnabledInfoParams {
	var ()
	return &GetEnabledInfoParams{

		Context: ctx,
	}
}

// NewGetEnabledInfoParamsWithHTTPClient creates a new GetEnabledInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnabledInfoParamsWithHTTPClient(client *http.Client) *GetEnabledInfoParams {
	var ()
	return &GetEnabledInfoParams{
		HTTPClient: client,
	}
}

/*GetEnabledInfoParams contains all the parameters to send to the API endpoint
for the get enabled info operation typically these are written to a http.Request
*/
type GetEnabledInfoParams struct {

	/*AgentLocator*/
	AgentLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get enabled info params
func (o *GetEnabledInfoParams) WithTimeout(timeout time.Duration) *GetEnabledInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get enabled info params
func (o *GetEnabledInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get enabled info params
func (o *GetEnabledInfoParams) WithContext(ctx context.Context) *GetEnabledInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get enabled info params
func (o *GetEnabledInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get enabled info params
func (o *GetEnabledInfoParams) WithHTTPClient(client *http.Client) *GetEnabledInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get enabled info params
func (o *GetEnabledInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the get enabled info params
func (o *GetEnabledInfoParams) WithAgentLocator(agentLocator string) *GetEnabledInfoParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the get enabled info params
func (o *GetEnabledInfoParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithFields adds the fields to the get enabled info params
func (o *GetEnabledInfoParams) WithFields(fields *string) *GetEnabledInfoParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get enabled info params
func (o *GetEnabledInfoParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnabledInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
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
