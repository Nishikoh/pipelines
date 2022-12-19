// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetExperimentV1Params creates a new GetExperimentV1Params object
// with the default values initialized.
func NewGetExperimentV1Params() *GetExperimentV1Params {
	var ()
	return &GetExperimentV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExperimentV1ParamsWithTimeout creates a new GetExperimentV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExperimentV1ParamsWithTimeout(timeout time.Duration) *GetExperimentV1Params {
	var ()
	return &GetExperimentV1Params{

		timeout: timeout,
	}
}

// NewGetExperimentV1ParamsWithContext creates a new GetExperimentV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetExperimentV1ParamsWithContext(ctx context.Context) *GetExperimentV1Params {
	var ()
	return &GetExperimentV1Params{

		Context: ctx,
	}
}

// NewGetExperimentV1ParamsWithHTTPClient creates a new GetExperimentV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExperimentV1ParamsWithHTTPClient(client *http.Client) *GetExperimentV1Params {
	var ()
	return &GetExperimentV1Params{
		HTTPClient: client,
	}
}

/*GetExperimentV1Params contains all the parameters to send to the API endpoint
for the get experiment v1 operation typically these are written to a http.Request
*/
type GetExperimentV1Params struct {

	/*ID
	  The ID of the experiment to be retrieved.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get experiment v1 params
func (o *GetExperimentV1Params) WithTimeout(timeout time.Duration) *GetExperimentV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get experiment v1 params
func (o *GetExperimentV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get experiment v1 params
func (o *GetExperimentV1Params) WithContext(ctx context.Context) *GetExperimentV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get experiment v1 params
func (o *GetExperimentV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get experiment v1 params
func (o *GetExperimentV1Params) WithHTTPClient(client *http.Client) *GetExperimentV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get experiment v1 params
func (o *GetExperimentV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get experiment v1 params
func (o *GetExperimentV1Params) WithID(id string) *GetExperimentV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get experiment v1 params
func (o *GetExperimentV1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetExperimentV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}