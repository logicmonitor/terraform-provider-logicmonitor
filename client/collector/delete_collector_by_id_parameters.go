// Code generated by go-swagger; DO NOT EDIT.

package collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"terraform-provider-logicmonitor/common"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteCollectorByIDParams creates a new DeleteCollectorByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCollectorByIDParams() *DeleteCollectorByIDParams {
	return &DeleteCollectorByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewDeleteCollectorByIDParamsWithTimeout creates a new DeleteCollectorByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteCollectorByIDParamsWithTimeout(timeout time.Duration) *DeleteCollectorByIDParams {
	return &DeleteCollectorByIDParams{
		timeout: timeout,
	}
}

// NewDeleteCollectorByIDParamsWithContext creates a new DeleteCollectorByIDParams object
// with the ability to set a context for a request.
func NewDeleteCollectorByIDParamsWithContext(ctx context.Context) *DeleteCollectorByIDParams {
	return &DeleteCollectorByIDParams{
		Context: ctx,
	}
}

// NewDeleteCollectorByIDParamsWithHTTPClient creates a new DeleteCollectorByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCollectorByIDParamsWithHTTPClient(client *http.Client) *DeleteCollectorByIDParams {
	return &DeleteCollectorByIDParams{
		HTTPClient: client,
	}
}

/* DeleteCollectorByIDParams contains all the parameters to send to the API endpoint
   for the delete collector by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteCollectorByIDParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectorByIDParams) WithDefaults() *DeleteCollectorByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectorByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete collector by Id params
func (o *DeleteCollectorByIDParams) WithTimeout(timeout time.Duration) *DeleteCollectorByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete collector by Id params
func (o *DeleteCollectorByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete collector by Id params
func (o *DeleteCollectorByIDParams) WithContext(ctx context.Context) *DeleteCollectorByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete collector by Id params
func (o *DeleteCollectorByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete collector by Id params
func (o *DeleteCollectorByIDParams) WithHTTPClient(client *http.Client) *DeleteCollectorByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete collector by Id params
func (o *DeleteCollectorByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete collector by Id params
func (o *DeleteCollectorByIDParams) WithID(id int32) *DeleteCollectorByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete collector by Id params
func (o *DeleteCollectorByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCollectorByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
