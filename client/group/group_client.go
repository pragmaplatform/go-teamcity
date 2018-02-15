// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddRoleSimple add role simple API
*/
func (a *Client) AddRoleSimple(params *AddRoleSimpleParams) (*AddRoleSimpleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleSimpleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRoleSimple",
		Method:             "POST",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddRoleSimpleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddRoleSimpleOK), nil

}

/*
DeleteGroup delete group API
*/
func (a *Client) DeleteGroup(params *DeleteGroupParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroup",
		Method:             "DELETE",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetParentGroups get parent groups API
*/
func (a *Client) GetParentGroups(params *GetParentGroupsParams) (*GetParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParentGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getParentGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetParentGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetParentGroupsOK), nil

}

/*
GetProperties get properties API
*/
func (a *Client) GetProperties(params *GetPropertiesParams) (*GetPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProperties",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPropertiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPropertiesOK), nil

}

/*
ServeGroup serve group API
*/
func (a *Client) ServeGroup(params *ServeGroupParams) (*ServeGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveGroup",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeGroupOK), nil

}

/*
ServeGroups serve groups API
*/
func (a *Client) ServeGroups(params *ServeGroupsParams) (*ServeGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeGroupsOK), nil

}

/*
SetParentGroups set parent groups API
*/
func (a *Client) SetParentGroups(params *SetParentGroupsParams) (*SetParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetParentGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setParentGroups",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetParentGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetParentGroupsOK), nil

}

/*
SetRoles set roles API
*/
func (a *Client) SetRoles(params *SetRolesParams) (*SetRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setRoles",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetRolesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
