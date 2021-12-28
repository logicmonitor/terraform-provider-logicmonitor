// Code generated by go-swagger; DO NOT EDIT.

package collector

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
	"github.com/go-openapi/swag"
)

// NewMiscGetCollectorDownloadTokenParams creates a new MiscGetCollectorDownloadTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMiscGetCollectorDownloadTokenParams() *MiscGetCollectorDownloadTokenParams {
	return &MiscGetCollectorDownloadTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMiscGetCollectorDownloadTokenParamsWithTimeout creates a new MiscGetCollectorDownloadTokenParams object
// with the ability to set a timeout on a request.
func NewMiscGetCollectorDownloadTokenParamsWithTimeout(timeout time.Duration) *MiscGetCollectorDownloadTokenParams {
	return &MiscGetCollectorDownloadTokenParams{
		timeout: timeout,
	}
}

// NewMiscGetCollectorDownloadTokenParamsWithContext creates a new MiscGetCollectorDownloadTokenParams object
// with the ability to set a context for a request.
func NewMiscGetCollectorDownloadTokenParamsWithContext(ctx context.Context) *MiscGetCollectorDownloadTokenParams {
	return &MiscGetCollectorDownloadTokenParams{
		Context: ctx,
	}
}

// NewMiscGetCollectorDownloadTokenParamsWithHTTPClient creates a new MiscGetCollectorDownloadTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewMiscGetCollectorDownloadTokenParamsWithHTTPClient(client *http.Client) *MiscGetCollectorDownloadTokenParams {
	return &MiscGetCollectorDownloadTokenParams{
		HTTPClient: client,
	}
}

/* MiscGetCollectorDownloadTokenParams contains all the parameters to send to the API endpoint
   for the misc get collector download token operation.

   Typically these are written to a http.Request.
*/
type MiscGetCollectorDownloadTokenParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the misc get collector download token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MiscGetCollectorDownloadTokenParams) WithDefaults() *MiscGetCollectorDownloadTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the misc get collector download token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MiscGetCollectorDownloadTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) WithTimeout(timeout time.Duration) *MiscGetCollectorDownloadTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) WithContext(ctx context.Context) *MiscGetCollectorDownloadTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) WithHTTPClient(client *http.Client) *MiscGetCollectorDownloadTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) WithID(id int32) *MiscGetCollectorDownloadTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the misc get collector download token params
func (o *MiscGetCollectorDownloadTokenParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MiscGetCollectorDownloadTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
