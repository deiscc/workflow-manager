package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new operations API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
CreateClusterDetails creates a cluster with all components
*/
func (a *Client) CreateClusterDetails(params *CreateClusterDetailsParams) (*CreateClusterDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterDetailsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createClusterDetails",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterDetailsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterDetailsOK), nil
}

/*
GetClusterByID reads a cluster details
*/
func (a *Client) GetClusterByID(params *GetClusterByIDParams) (*GetClusterByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterByIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getClusterById",
		Method:             "GET",
		PathPattern:        "/clusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterByIDOK), nil
}

/*
GetClustersByAge lists clusters
*/
func (a *Client) GetClustersByAge(params *GetClustersByAgeParams) (*GetClustersByAgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersByAgeParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getClustersByAge",
		Method:             "GET",
		PathPattern:        "/clusters/age",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClustersByAgeReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClustersByAgeOK), nil
}

/*
GetClustersCount reads the count of the known deis clusters
*/
func (a *Client) GetClustersCount(params *GetClustersCountParams) (*GetClustersCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersCountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getClustersCount",
		Method:             "GET",
		PathPattern:        "/clusters/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClustersCountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClustersCountOK), nil
}

/*
GetComponentByName lists the releases of a component
*/
func (a *Client) GetComponentByName(params *GetComponentByNameParams) (*GetComponentByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentByNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getComponentByName",
		Method:             "GET",
		PathPattern:        "/versions/{train}/{component}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetComponentByNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentByNameOK), nil
}

/*
GetComponentByRelease reads the specified release of a component
*/
func (a *Client) GetComponentByRelease(params *GetComponentByReleaseParams) (*GetComponentByReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentByReleaseParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getComponentByRelease",
		Method:             "GET",
		PathPattern:        "/versions/{train}/{component}/{release}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetComponentByReleaseReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentByReleaseOK), nil
}

/*
GetComponentsByLatestRelease lists the latest release version of the components
*/
func (a *Client) GetComponentsByLatestRelease(params *GetComponentsByLatestReleaseParams) (*GetComponentsByLatestReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentsByLatestReleaseParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getComponentsByLatestRelease",
		Method:             "POST",
		PathPattern:        "/versions/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetComponentsByLatestReleaseReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentsByLatestReleaseOK), nil
}

/*
PublishComponentRelease publishes a new release for the component
*/
func (a *Client) PublishComponentRelease(params *PublishComponentReleaseParams) (*PublishComponentReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublishComponentReleaseParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "publishComponentRelease",
		Method:             "POST",
		PathPattern:        "/versions/{train}/{component}/{release}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublishComponentReleaseReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PublishComponentReleaseOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}