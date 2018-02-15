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

// NewGetProjectsOrderParams creates a new GetProjectsOrderParams object
// with the default values initialized.
func NewGetProjectsOrderParams() *GetProjectsOrderParams {
	var ()
	return &GetProjectsOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsOrderParamsWithTimeout creates a new GetProjectsOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsOrderParamsWithTimeout(timeout time.Duration) *GetProjectsOrderParams {
	var ()
	return &GetProjectsOrderParams{

		timeout: timeout,
	}
}

// NewGetProjectsOrderParamsWithContext creates a new GetProjectsOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsOrderParamsWithContext(ctx context.Context) *GetProjectsOrderParams {
	var ()
	return &GetProjectsOrderParams{

		Context: ctx,
	}
}

// NewGetProjectsOrderParamsWithHTTPClient creates a new GetProjectsOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsOrderParamsWithHTTPClient(client *http.Client) *GetProjectsOrderParams {
	var ()
	return &GetProjectsOrderParams{
		HTTPClient: client,
	}
}

/*GetProjectsOrderParams contains all the parameters to send to the API endpoint
for the get projects order operation typically these are written to a http.Request
*/
type GetProjectsOrderParams struct {

	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get projects order params
func (o *GetProjectsOrderParams) WithTimeout(timeout time.Duration) *GetProjectsOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects order params
func (o *GetProjectsOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects order params
func (o *GetProjectsOrderParams) WithContext(ctx context.Context) *GetProjectsOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects order params
func (o *GetProjectsOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects order params
func (o *GetProjectsOrderParams) WithHTTPClient(client *http.Client) *GetProjectsOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects order params
func (o *GetProjectsOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the get projects order params
func (o *GetProjectsOrderParams) WithField(field string) *GetProjectsOrderParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the get projects order params
func (o *GetProjectsOrderParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the get projects order params
func (o *GetProjectsOrderParams) WithProjectLocator(projectLocator string) *GetProjectsOrderParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get projects order params
func (o *GetProjectsOrderParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
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
