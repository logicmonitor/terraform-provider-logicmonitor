// Code generated by go-swagger; DO NOT EDIT.

package device

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

// NewGetDeviceByIDParams creates a new GetDeviceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceByIDParams() *GetDeviceByIDParams {
	return &GetDeviceByIDParams{
		timeout: 3 * cr.DefaultTimeout,
	}
}

// NewGetDeviceByIDParamsWithTimeout creates a new GetDeviceByIDParams object
// with the ability to set a timeout on a request.
func NewGetDeviceByIDParamsWithTimeout(timeout time.Duration) *GetDeviceByIDParams {
	return &GetDeviceByIDParams{
		timeout: timeout,
	}
}

// NewGetDeviceByIDParamsWithContext creates a new GetDeviceByIDParams object
// with the ability to set a context for a request.
func NewGetDeviceByIDParamsWithContext(ctx context.Context) *GetDeviceByIDParams {
	return &GetDeviceByIDParams{
		Context: ctx,
	}
}

// NewGetDeviceByIDParamsWithHTTPClient creates a new GetDeviceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceByIDParamsWithHTTPClient(client *http.Client) *GetDeviceByIDParams {
	return &GetDeviceByIDParams{
		HTTPClient: client,
	}
}

/* GetDeviceByIDParams contains all the parameters to send to the API endpoint
   for the get device by Id operation.

   Typically these are written to a http.Request.
*/
type GetDeviceByIDParams struct {

	// End.
	//
	// Format: int64
	End *int64

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	// NetflowFilter.
	NetflowFilter *string

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceByIDParams) WithDefaults() *GetDeviceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device by Id params
func (o *GetDeviceByIDParams) WithTimeout(timeout time.Duration) *GetDeviceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device by Id params
func (o *GetDeviceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device by Id params
func (o *GetDeviceByIDParams) WithContext(ctx context.Context) *GetDeviceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device by Id params
func (o *GetDeviceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device by Id params
func (o *GetDeviceByIDParams) WithHTTPClient(client *http.Client) *GetDeviceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device by Id params
func (o *GetDeviceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get device by Id params
func (o *GetDeviceByIDParams) WithEnd(end *int64) *GetDeviceByIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get device by Id params
func (o *GetDeviceByIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get device by Id params
func (o *GetDeviceByIDParams) WithFields(fields *string) *GetDeviceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device by Id params
func (o *GetDeviceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get device by Id params
func (o *GetDeviceByIDParams) WithID(id int32) *GetDeviceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device by Id params
func (o *GetDeviceByIDParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the get device by Id params
func (o *GetDeviceByIDParams) WithNetflowFilter(netflowFilter *string) *GetDeviceByIDParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get device by Id params
func (o *GetDeviceByIDParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithStart adds the start to the get device by Id params
func (o *GetDeviceByIDParams) WithStart(start *int64) *GetDeviceByIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get device by Id params
func (o *GetDeviceByIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
