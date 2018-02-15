// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewServeSourceFileParams creates a new ServeSourceFileParams object
// with the default values initialized.
func NewServeSourceFileParams() *ServeSourceFileParams {
	var ()
	return &ServeSourceFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeSourceFileParamsWithTimeout creates a new ServeSourceFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeSourceFileParamsWithTimeout(timeout time.Duration) *ServeSourceFileParams {
	var ()
	return &ServeSourceFileParams{

		timeout: timeout,
	}
}

// NewServeSourceFileParamsWithContext creates a new ServeSourceFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeSourceFileParamsWithContext(ctx context.Context) *ServeSourceFileParams {
	var ()
	return &ServeSourceFileParams{

		Context: ctx,
	}
}

// NewServeSourceFileParamsWithHTTPClient creates a new ServeSourceFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeSourceFileParamsWithHTTPClient(client *http.Client) *ServeSourceFileParams {
	var ()
	return &ServeSourceFileParams{
		HTTPClient: client,
	}
}

/*ServeSourceFileParams contains all the parameters to send to the API endpoint
for the serve source file operation typically these are written to a http.Request
*/
type ServeSourceFileParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*FileName*/
	FileName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve source file params
func (o *ServeSourceFileParams) WithTimeout(timeout time.Duration) *ServeSourceFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve source file params
func (o *ServeSourceFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve source file params
func (o *ServeSourceFileParams) WithContext(ctx context.Context) *ServeSourceFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve source file params
func (o *ServeSourceFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve source file params
func (o *ServeSourceFileParams) WithHTTPClient(client *http.Client) *ServeSourceFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve source file params
func (o *ServeSourceFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve source file params
func (o *ServeSourceFileParams) WithBuildLocator(buildLocator string) *ServeSourceFileParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve source file params
func (o *ServeSourceFileParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFileName adds the fileName to the serve source file params
func (o *ServeSourceFileParams) WithFileName(fileName string) *ServeSourceFileParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the serve source file params
func (o *ServeSourceFileParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *ServeSourceFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param fileName
	if err := r.SetPathParam("fileName", o.FileName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
