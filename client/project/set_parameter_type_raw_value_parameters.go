// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewSetParameterTypeRawValueParams creates a new SetParameterTypeRawValueParams object
// with the default values initialized.
func NewSetParameterTypeRawValueParams() *SetParameterTypeRawValueParams {
	var ()
	return &SetParameterTypeRawValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetParameterTypeRawValueParamsWithTimeout creates a new SetParameterTypeRawValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetParameterTypeRawValueParamsWithTimeout(timeout time.Duration) *SetParameterTypeRawValueParams {
	var ()
	return &SetParameterTypeRawValueParams{

		timeout: timeout,
	}
}

// NewSetParameterTypeRawValueParamsWithContext creates a new SetParameterTypeRawValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetParameterTypeRawValueParamsWithContext(ctx context.Context) *SetParameterTypeRawValueParams {
	var ()
	return &SetParameterTypeRawValueParams{

		Context: ctx,
	}
}

// NewSetParameterTypeRawValueParamsWithHTTPClient creates a new SetParameterTypeRawValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetParameterTypeRawValueParamsWithHTTPClient(client *http.Client) *SetParameterTypeRawValueParams {
	var ()
	return &SetParameterTypeRawValueParams{
		HTTPClient: client,
	}
}

/*SetParameterTypeRawValueParams contains all the parameters to send to the API endpoint
for the set parameter type raw value operation typically these are written to a http.Request
*/
type SetParameterTypeRawValueParams struct {

	/*Body*/
	Body string
	/*Name*/
	Name string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithTimeout(timeout time.Duration) *SetParameterTypeRawValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithContext(ctx context.Context) *SetParameterTypeRawValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithHTTPClient(client *http.Client) *SetParameterTypeRawValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithBody(body string) *SetParameterTypeRawValueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetBody(body string) {
	o.Body = body
}

// WithName adds the name to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithName(name string) *SetParameterTypeRawValueParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) WithProjectLocator(projectLocator string) *SetParameterTypeRawValueParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set parameter type raw value params
func (o *SetParameterTypeRawValueParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetParameterTypeRawValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
