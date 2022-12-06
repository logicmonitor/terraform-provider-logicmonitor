// Code generated by go-swagger; DO NOT EDIT.

package device_group

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

// NewDeleteDeviceGroupByIDParams creates a new DeleteDeviceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceGroupByIDParams() *DeleteDeviceGroupByIDParams {
	return &DeleteDeviceGroupByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewDeleteDeviceGroupByIDParamsWithTimeout creates a new DeleteDeviceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceGroupByIDParamsWithTimeout(timeout time.Duration) *DeleteDeviceGroupByIDParams {
	return &DeleteDeviceGroupByIDParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceGroupByIDParamsWithContext creates a new DeleteDeviceGroupByIDParams object
// with the ability to set a context for a request.
func NewDeleteDeviceGroupByIDParamsWithContext(ctx context.Context) *DeleteDeviceGroupByIDParams {
	return &DeleteDeviceGroupByIDParams{
		Context: ctx,
	}
}

// NewDeleteDeviceGroupByIDParamsWithHTTPClient creates a new DeleteDeviceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceGroupByIDParamsWithHTTPClient(client *http.Client) *DeleteDeviceGroupByIDParams {
	return &DeleteDeviceGroupByIDParams{
		HTTPClient: client,
	}
}

/* DeleteDeviceGroupByIDParams contains all the parameters to send to the API endpoint
   for the delete device group by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteDeviceGroupByIDParams struct {

	// DeleteChildren.
	DeleteChildren *bool

	// DeleteHard.
	//
	// Default: true
	DeleteHard *bool

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceGroupByIDParams) WithDefaults() *DeleteDeviceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceGroupByIDParams) SetDefaults() {
	var (
		deleteChildrenDefault = bool(false)

		deleteHardDefault = bool(true)
	)

	val := DeleteDeviceGroupByIDParams{
		DeleteChildren: &deleteChildrenDefault,
		DeleteHard:     &deleteHardDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithTimeout(timeout time.Duration) *DeleteDeviceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithContext(ctx context.Context) *DeleteDeviceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithHTTPClient(client *http.Client) *DeleteDeviceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteChildren adds the deleteChildren to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithDeleteChildren(deleteChildren *bool) *DeleteDeviceGroupByIDParams {
	o.SetDeleteChildren(deleteChildren)
	return o
}

// SetDeleteChildren adds the deleteChildren to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetDeleteChildren(deleteChildren *bool) {
	o.DeleteChildren = deleteChildren
}

// WithDeleteHard adds the deleteHard to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithDeleteHard(deleteHard *bool) *DeleteDeviceGroupByIDParams {
	o.SetDeleteHard(deleteHard)
	return o
}

// SetDeleteHard adds the deleteHard to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetDeleteHard(deleteHard *bool) {
	o.DeleteHard = deleteHard
}

// WithID adds the id to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) WithID(id int32) *DeleteDeviceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device group by Id params
func (o *DeleteDeviceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteChildren != nil {

		// query param deleteChildren
		var qrDeleteChildren bool

		if o.DeleteChildren != nil {
			qrDeleteChildren = *o.DeleteChildren
		}
		qDeleteChildren := swag.FormatBool(qrDeleteChildren)
		if qDeleteChildren != "" {

			if err := r.SetQueryParam("deleteChildren", qDeleteChildren); err != nil {
				return err
			}
		}
	}

	if o.DeleteHard != nil {

		// query param deleteHard
		var qrDeleteHard bool

		if o.DeleteHard != nil {
			qrDeleteHard = *o.DeleteHard
		}
		qDeleteHard := swag.FormatBool(qrDeleteHard)
		if qDeleteHard != "" {

			if err := r.SetQueryParam("deleteHard", qDeleteHard); err != nil {
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
