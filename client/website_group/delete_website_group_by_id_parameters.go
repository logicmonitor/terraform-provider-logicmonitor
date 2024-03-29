// Code generated by go-swagger; DO NOT EDIT.

package website_group

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

// NewDeleteWebsiteGroupByIDParams creates a new DeleteWebsiteGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWebsiteGroupByIDParams() *DeleteWebsiteGroupByIDParams {
	return &DeleteWebsiteGroupByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewDeleteWebsiteGroupByIDParamsWithTimeout creates a new DeleteWebsiteGroupByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteWebsiteGroupByIDParamsWithTimeout(timeout time.Duration) *DeleteWebsiteGroupByIDParams {
	return &DeleteWebsiteGroupByIDParams{
		timeout: timeout,
	}
}

// NewDeleteWebsiteGroupByIDParamsWithContext creates a new DeleteWebsiteGroupByIDParams object
// with the ability to set a context for a request.
func NewDeleteWebsiteGroupByIDParamsWithContext(ctx context.Context) *DeleteWebsiteGroupByIDParams {
	return &DeleteWebsiteGroupByIDParams{
		Context: ctx,
	}
}

// NewDeleteWebsiteGroupByIDParamsWithHTTPClient creates a new DeleteWebsiteGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWebsiteGroupByIDParamsWithHTTPClient(client *http.Client) *DeleteWebsiteGroupByIDParams {
	return &DeleteWebsiteGroupByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteWebsiteGroupByIDParams contains all the parameters to send to the API endpoint

	for the delete website group by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteWebsiteGroupByIDParams struct {

	// DeleteChildren.
	//
	// Format: int32
	DeleteChildren *int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebsiteGroupByIDParams) WithDefaults() *DeleteWebsiteGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebsiteGroupByIDParams) SetDefaults() {
	var (
		deleteChildrenDefault = int32(0)
	)

	val := DeleteWebsiteGroupByIDParams{
		DeleteChildren: &deleteChildrenDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) WithTimeout(timeout time.Duration) *DeleteWebsiteGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) WithContext(ctx context.Context) *DeleteWebsiteGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) WithHTTPClient(client *http.Client) *DeleteWebsiteGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteChildren adds the deleteChildren to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) WithDeleteChildren(deleteChildren *int32) *DeleteWebsiteGroupByIDParams {
	o.SetDeleteChildren(deleteChildren)
	return o
}

// SetDeleteChildren adds the deleteChildren to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) SetDeleteChildren(deleteChildren *int32) {
	o.DeleteChildren = deleteChildren
}

// WithID adds the id to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) WithID(id int32) *DeleteWebsiteGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete website group by Id params
func (o *DeleteWebsiteGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWebsiteGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteChildren != nil {

		// query param deleteChildren
		var qrDeleteChildren int32

		if o.DeleteChildren != nil {
			qrDeleteChildren = *o.DeleteChildren
		}
		qDeleteChildren := swag.FormatInt32(qrDeleteChildren)
		if qDeleteChildren != "" {

			if err := r.SetQueryParam("deleteChildren", qDeleteChildren); err != nil {
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
