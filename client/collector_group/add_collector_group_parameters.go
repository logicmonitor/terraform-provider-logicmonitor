// Code generated by go-swagger; DO NOT EDIT.

package collector_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"terraform-provider-logicmonitor/models"
)

// NewAddCollectorGroupParams creates a new AddCollectorGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddCollectorGroupParams() *AddCollectorGroupParams {
	return &AddCollectorGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddCollectorGroupParamsWithTimeout creates a new AddCollectorGroupParams object
// with the ability to set a timeout on a request.
func NewAddCollectorGroupParamsWithTimeout(timeout time.Duration) *AddCollectorGroupParams {
	return &AddCollectorGroupParams{
		timeout: timeout,
	}
}

// NewAddCollectorGroupParamsWithContext creates a new AddCollectorGroupParams object
// with the ability to set a context for a request.
func NewAddCollectorGroupParamsWithContext(ctx context.Context) *AddCollectorGroupParams {
	return &AddCollectorGroupParams{
		Context: ctx,
	}
}

// NewAddCollectorGroupParamsWithHTTPClient creates a new AddCollectorGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddCollectorGroupParamsWithHTTPClient(client *http.Client) *AddCollectorGroupParams {
	return &AddCollectorGroupParams{
		HTTPClient: client,
	}
}

/* AddCollectorGroupParams contains all the parameters to send to the API endpoint
   for the add collector group operation.

   Typically these are written to a http.Request.
*/
type AddCollectorGroupParams struct {

	// Body.
	Body *models.CollectorGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add collector group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCollectorGroupParams) WithDefaults() *AddCollectorGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add collector group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCollectorGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add collector group params
func (o *AddCollectorGroupParams) WithTimeout(timeout time.Duration) *AddCollectorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add collector group params
func (o *AddCollectorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add collector group params
func (o *AddCollectorGroupParams) WithContext(ctx context.Context) *AddCollectorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add collector group params
func (o *AddCollectorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add collector group params
func (o *AddCollectorGroupParams) WithHTTPClient(client *http.Client) *AddCollectorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add collector group params
func (o *AddCollectorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add collector group params
func (o *AddCollectorGroupParams) WithBody(body *models.CollectorGroup) *AddCollectorGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add collector group params
func (o *AddCollectorGroupParams) SetBody(body *models.CollectorGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddCollectorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}