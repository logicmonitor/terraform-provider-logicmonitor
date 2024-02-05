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

// NewGetWebsiteGroupByIDParams creates a new GetWebsiteGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteGroupByIDParams() *GetWebsiteGroupByIDParams {
	return &GetWebsiteGroupByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewGetWebsiteGroupByIDParamsWithTimeout creates a new GetWebsiteGroupByIDParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteGroupByIDParamsWithTimeout(timeout time.Duration) *GetWebsiteGroupByIDParams {
	return &GetWebsiteGroupByIDParams{
		timeout: timeout,
	}
}

// NewGetWebsiteGroupByIDParamsWithContext creates a new GetWebsiteGroupByIDParams object
// with the ability to set a context for a request.
func NewGetWebsiteGroupByIDParamsWithContext(ctx context.Context) *GetWebsiteGroupByIDParams {
	return &GetWebsiteGroupByIDParams{
		Context: ctx,
	}
}

// NewGetWebsiteGroupByIDParamsWithHTTPClient creates a new GetWebsiteGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteGroupByIDParamsWithHTTPClient(client *http.Client) *GetWebsiteGroupByIDParams {
	return &GetWebsiteGroupByIDParams{
		HTTPClient: client,
	}
}

/*
GetWebsiteGroupByIDParams contains all the parameters to send to the API endpoint

	for the get website group by Id operation.

	Typically these are written to a http.Request.
*/
type GetWebsiteGroupByIDParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteGroupByIDParams) WithDefaults() *GetWebsiteGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteGroupByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) WithTimeout(timeout time.Duration) *GetWebsiteGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) WithContext(ctx context.Context) *GetWebsiteGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) WithHTTPClient(client *http.Client) *GetWebsiteGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) WithID(id int32) *GetWebsiteGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website group by Id params
func (o *GetWebsiteGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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