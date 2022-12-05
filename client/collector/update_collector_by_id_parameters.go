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

	"terraform-provider-logicmonitor/models"
)

// NewUpdateCollectorByIDParams creates a new UpdateCollectorByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCollectorByIDParams() *UpdateCollectorByIDParams {
	return &UpdateCollectorByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewUpdateCollectorByIDParamsWithTimeout creates a new UpdateCollectorByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateCollectorByIDParamsWithTimeout(timeout time.Duration) *UpdateCollectorByIDParams {
	return &UpdateCollectorByIDParams{
		timeout: timeout,
	}
}

// NewUpdateCollectorByIDParamsWithContext creates a new UpdateCollectorByIDParams object
// with the ability to set a context for a request.
func NewUpdateCollectorByIDParamsWithContext(ctx context.Context) *UpdateCollectorByIDParams {
	return &UpdateCollectorByIDParams{
		Context: ctx,
	}
}

// NewUpdateCollectorByIDParamsWithHTTPClient creates a new UpdateCollectorByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCollectorByIDParamsWithHTTPClient(client *http.Client) *UpdateCollectorByIDParams {
	return &UpdateCollectorByIDParams{
		HTTPClient: client,
	}
}

/* UpdateCollectorByIDParams contains all the parameters to send to the API endpoint
   for the update collector by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateCollectorByIDParams struct {

	// Body.
	Body *models.Collector

	// ForceUpdateFailedOverDevices.
	ForceUpdateFailedOverDevices *bool

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectorByIDParams) WithDefaults() *UpdateCollectorByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectorByIDParams) SetDefaults() {
	var (
		forceUpdateFailedOverDevicesDefault = bool(false)
	)

	val := UpdateCollectorByIDParams{
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithTimeout(timeout time.Duration) *UpdateCollectorByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithContext(ctx context.Context) *UpdateCollectorByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithHTTPClient(client *http.Client) *UpdateCollectorByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithBody(body *models.Collector) *UpdateCollectorByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetBody(body *models.Collector) {
	o.Body = body
}

// WithForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) *UpdateCollectorByIDParams {
	o.SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices)
	return o
}

// SetForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) {
	o.ForceUpdateFailedOverDevices = forceUpdateFailedOverDevices
}

// WithID adds the id to the update collector by Id params
func (o *UpdateCollectorByIDParams) WithID(id int32) *UpdateCollectorByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update collector by Id params
func (o *UpdateCollectorByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCollectorByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ForceUpdateFailedOverDevices != nil {

		// query param forceUpdateFailedOverDevices
		var qrForceUpdateFailedOverDevices bool

		if o.ForceUpdateFailedOverDevices != nil {
			qrForceUpdateFailedOverDevices = *o.ForceUpdateFailedOverDevices
		}
		qForceUpdateFailedOverDevices := swag.FormatBool(qrForceUpdateFailedOverDevices)
		if qForceUpdateFailedOverDevices != "" {

			if err := r.SetQueryParam("forceUpdateFailedOverDevices", qForceUpdateFailedOverDevices); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
