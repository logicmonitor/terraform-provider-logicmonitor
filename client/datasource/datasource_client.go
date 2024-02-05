// Code generated by go-swagger; DO NOT EDIT.

package datasource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new datasource API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for datasource API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddDatasourceByID adds datasource
*/
func (a *Client) AddDatasourceByID(params *AddDatasourceByIDParams) (*AddDatasourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddDatasourceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addDatasourceById",
		Method:             "POST",
		PathPattern:        "/setting/datasources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddDatasourceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddDatasourceByIDOK), nil

}

/*
DeleteDatasourceByID deletes datasource
*/
func (a *Client) DeleteDatasourceByID(params *DeleteDatasourceByIDParams) (*DeleteDatasourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatasourceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDatasourceById",
		Method:             "DELETE",
		PathPattern:        "/setting/datasources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDatasourceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDatasourceByIDOK), nil

}

/*
GetDatasourceByID gets datasource by id
*/
func (a *Client) GetDatasourceByID(params *GetDatasourceByIDParams) (*GetDatasourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatasourceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDatasourceById",
		Method:             "GET",
		PathPattern:        "/setting/datasources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatasourceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDatasourceByIDOK), nil

}

/*
GetDatasourceList gets datasource list
*/
func (a *Client) GetDatasourceList(params *GetDatasourceListParams) (*GetDatasourceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatasourceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDatasourceList",
		Method:             "GET",
		PathPattern:        "/setting/datasources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatasourceListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDatasourceListOK), nil

}

/*
PatchDatasourceByID updates datasource
*/
func (a *Client) PatchDatasourceByID(params *PatchDatasourceByIDParams) (*PatchDatasourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDatasourceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDatasourceById",
		Method:             "PATCH",
		PathPattern:        "/setting/datasources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchDatasourceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchDatasourceByIDOK), nil

}

/*
UpdateDatasourceByID updates datasource
*/
func (a *Client) UpdateDatasourceByID(params *UpdateDatasourceByIDParams) (*UpdateDatasourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDatasourceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDatasourceById",
		Method:             "PUT",
		PathPattern:        "/setting/datasources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDatasourceByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDatasourceByIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}