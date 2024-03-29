// Code generated by go-swagger; DO NOT EDIT.

package escalation_chain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new escalation chain API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for escalation chain API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddEscalationChain adds escalation chain
*/
func (a *Client) AddEscalationChain(params *AddEscalationChainParams) (*AddEscalationChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEscalationChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addEscalationChain",
		Method:             "POST",
		PathPattern:        "/setting/alert/chains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddEscalationChainReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddEscalationChainOK), nil

}

/*
DeleteEscalationChainByID deletes escalation chain
*/
func (a *Client) DeleteEscalationChainByID(params *DeleteEscalationChainByIDParams) (*DeleteEscalationChainByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEscalationChainByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEscalationChainById",
		Method:             "DELETE",
		PathPattern:        "/setting/alert/chains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEscalationChainByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEscalationChainByIDOK), nil

}

/*
GetEscalationChainByID gets escalation chain by id
*/
func (a *Client) GetEscalationChainByID(params *GetEscalationChainByIDParams) (*GetEscalationChainByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEscalationChainByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEscalationChainById",
		Method:             "GET",
		PathPattern:        "/setting/alert/chains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEscalationChainByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEscalationChainByIDOK), nil

}

/*
GetEscalationChainList gets escalation chain list
*/
func (a *Client) GetEscalationChainList(params *GetEscalationChainListParams) (*GetEscalationChainListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEscalationChainListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEscalationChainList",
		Method:             "GET",
		PathPattern:        "/setting/alert/chains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEscalationChainListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEscalationChainListOK), nil

}

/*
PatchEscalationChainByID updates escalation chain
*/
func (a *Client) PatchEscalationChainByID(params *PatchEscalationChainByIDParams) (*PatchEscalationChainByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEscalationChainByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchEscalationChainById",
		Method:             "PATCH",
		PathPattern:        "/setting/alert/chains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEscalationChainByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchEscalationChainByIDOK), nil

}

/*
UpdateEscalationChainByID updates escalation chain
*/
func (a *Client) UpdateEscalationChainByID(params *UpdateEscalationChainByIDParams) (*UpdateEscalationChainByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEscalationChainByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEscalationChainById",
		Method:             "PUT",
		PathPattern:        "/setting/alert/chains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEscalationChainByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateEscalationChainByIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
