// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddDevice adds a new device
*/
func (a *Client) AddDevice(params *AddDeviceParams) (*AddDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addDevice",
		Method:             "POST",
		PathPattern:        "/device/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddDeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddDeviceOK), nil

}

/*
DeleteDeviceByID deletes a device
*/
func (a *Client) DeleteDeviceByID(params *DeleteDeviceByIDParams) (*DeleteDeviceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDeviceById",
		Method:             "DELETE",
		PathPattern:        "/device/devices/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeviceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDeviceByIDOK), nil

}

/*
GetDeviceByID gets device by id
*/
func (a *Client) GetDeviceByID(params *GetDeviceByIDParams) (*GetDeviceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceById",
		Method:             "GET",
		PathPattern:        "/device/devices/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceByIDOK), nil

}

/*
GetDeviceList gets device list
*/
func (a *Client) GetDeviceList(params *GetDeviceListParams) (*GetDeviceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceList",
		Method:             "GET",
		PathPattern:        "/device/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceListOK), nil

}

/*
PatchDevice updates a device
*/
func (a *Client) PatchDevice(params *PatchDeviceParams) (*PatchDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDevice",
		Method:             "PATCH",
		PathPattern:        "/device/devices/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchDeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchDeviceOK), nil

}

/*
UpdateDevice updates a device
*/
func (a *Client) UpdateDevice(params *UpdateDeviceParams) (*UpdateDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDevice",
		Method:             "PUT",
		PathPattern:        "/device/devices/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDeviceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
