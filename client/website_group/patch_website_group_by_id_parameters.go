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

	"terraform-provider-logicmonitor/models"
)

// NewPatchWebsiteGroupByIDParams creates a new PatchWebsiteGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchWebsiteGroupByIDParams() *PatchWebsiteGroupByIDParams {
	return &PatchWebsiteGroupByIDParams{
		timeout: common.DefaultTimeout,
	}
}

// NewPatchWebsiteGroupByIDParamsWithTimeout creates a new PatchWebsiteGroupByIDParams object
// with the ability to set a timeout on a request.
func NewPatchWebsiteGroupByIDParamsWithTimeout(timeout time.Duration) *PatchWebsiteGroupByIDParams {
	return &PatchWebsiteGroupByIDParams{
		timeout: timeout,
	}
}

// NewPatchWebsiteGroupByIDParamsWithContext creates a new PatchWebsiteGroupByIDParams object
// with the ability to set a context for a request.
func NewPatchWebsiteGroupByIDParamsWithContext(ctx context.Context) *PatchWebsiteGroupByIDParams {
	return &PatchWebsiteGroupByIDParams{
		Context: ctx,
	}
}

// NewPatchWebsiteGroupByIDParamsWithHTTPClient creates a new PatchWebsiteGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchWebsiteGroupByIDParamsWithHTTPClient(client *http.Client) *PatchWebsiteGroupByIDParams {
	return &PatchWebsiteGroupByIDParams{
		HTTPClient: client,
	}
}

/*
PatchWebsiteGroupByIDParams contains all the parameters to send to the API endpoint

	for the patch website group by Id operation.

	Typically these are written to a http.Request.
*/
type PatchWebsiteGroupByIDParams struct {

	// Body.
	Body *models.WebsiteGroup

	// ID.
	//
	// Format: int32
	ID int32

	// OpType.
	//
	// Default: "refresh"
	OpType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWebsiteGroupByIDParams) WithDefaults() *PatchWebsiteGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWebsiteGroupByIDParams) SetDefaults() {
	var (
		opTypeDefault = string("refresh")
	)

	val := PatchWebsiteGroupByIDParams{
		OpType: &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithTimeout(timeout time.Duration) *PatchWebsiteGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithContext(ctx context.Context) *PatchWebsiteGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithHTTPClient(client *http.Client) *PatchWebsiteGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithBody(body *models.WebsiteGroup) *PatchWebsiteGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetBody(body *models.WebsiteGroup) {
	o.Body = body
}

// WithID adds the id to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithID(id int32) *PatchWebsiteGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) WithOpType(opType *string) *PatchWebsiteGroupByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the patch website group by Id params
func (o *PatchWebsiteGroupByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWebsiteGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.OpType != nil {

		// query param opType
		var qrOpType string

		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {

			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
