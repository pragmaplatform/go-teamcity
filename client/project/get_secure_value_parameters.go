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

// NewGetSecureValueParams creates a new GetSecureValueParams object
// with the default values initialized.
func NewGetSecureValueParams() *GetSecureValueParams {
	var ()
	return &GetSecureValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecureValueParamsWithTimeout creates a new GetSecureValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecureValueParamsWithTimeout(timeout time.Duration) *GetSecureValueParams {
	var ()
	return &GetSecureValueParams{

		timeout: timeout,
	}
}

// NewGetSecureValueParamsWithContext creates a new GetSecureValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecureValueParamsWithContext(ctx context.Context) *GetSecureValueParams {
	var ()
	return &GetSecureValueParams{

		Context: ctx,
	}
}

// NewGetSecureValueParamsWithHTTPClient creates a new GetSecureValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecureValueParamsWithHTTPClient(client *http.Client) *GetSecureValueParams {
	var ()
	return &GetSecureValueParams{
		HTTPClient: client,
	}
}

/*GetSecureValueParams contains all the parameters to send to the API endpoint
for the get secure value operation typically these are written to a http.Request
*/
type GetSecureValueParams struct {

	/*ProjectLocator*/
	ProjectLocator string
	/*Token*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get secure value params
func (o *GetSecureValueParams) WithTimeout(timeout time.Duration) *GetSecureValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get secure value params
func (o *GetSecureValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get secure value params
func (o *GetSecureValueParams) WithContext(ctx context.Context) *GetSecureValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get secure value params
func (o *GetSecureValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get secure value params
func (o *GetSecureValueParams) WithHTTPClient(client *http.Client) *GetSecureValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get secure value params
func (o *GetSecureValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectLocator adds the projectLocator to the get secure value params
func (o *GetSecureValueParams) WithProjectLocator(projectLocator string) *GetSecureValueParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get secure value params
func (o *GetSecureValueParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WithToken adds the token to the get secure value params
func (o *GetSecureValueParams) WithToken(token string) *GetSecureValueParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get secure value params
func (o *GetSecureValueParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecureValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
