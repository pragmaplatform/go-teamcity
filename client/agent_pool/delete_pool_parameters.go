// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

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

// NewDeletePoolParams creates a new DeletePoolParams object
// with the default values initialized.
func NewDeletePoolParams() *DeletePoolParams {
	var ()
	return &DeletePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePoolParamsWithTimeout creates a new DeletePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePoolParamsWithTimeout(timeout time.Duration) *DeletePoolParams {
	var ()
	return &DeletePoolParams{

		timeout: timeout,
	}
}

// NewDeletePoolParamsWithContext creates a new DeletePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePoolParamsWithContext(ctx context.Context) *DeletePoolParams {
	var ()
	return &DeletePoolParams{

		Context: ctx,
	}
}

// NewDeletePoolParamsWithHTTPClient creates a new DeletePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePoolParamsWithHTTPClient(client *http.Client) *DeletePoolParams {
	var ()
	return &DeletePoolParams{
		HTTPClient: client,
	}
}

/*DeletePoolParams contains all the parameters to send to the API endpoint
for the delete pool operation typically these are written to a http.Request
*/
type DeletePoolParams struct {

	/*AgentPoolLocator*/
	AgentPoolLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pool params
func (o *DeletePoolParams) WithTimeout(timeout time.Duration) *DeletePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pool params
func (o *DeletePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pool params
func (o *DeletePoolParams) WithContext(ctx context.Context) *DeletePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pool params
func (o *DeletePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pool params
func (o *DeletePoolParams) WithHTTPClient(client *http.Client) *DeletePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pool params
func (o *DeletePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the delete pool params
func (o *DeletePoolParams) WithAgentPoolLocator(agentPoolLocator string) *DeletePoolParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the delete pool params
func (o *DeletePoolParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
