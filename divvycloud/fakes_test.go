package divvycloud

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sgarlick987/godivvycloud/client"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
	"github.com/sgarlick987/godivvycloud/client/organizations"
)

//As far as I can tell the only interface we have when using goswagger generated client is the ClientTransport
//we configure a fake one for our divvyclient, but since it might be used to make more than a single call
//we store the various parameters to a call in maps keyed by the operation ID
type FakeDivvycloudClientTransport struct {
	//checks if a operation was called
	CalledMap map[string]bool
	//the parameters an operation was called with
	ParamsMap map[string]runtime.ClientRequestWriter
	//the responses we want the operation to return for testing
	ResponseMap map[string]interface{}
	//the errors we want the operation to return for testing
	ErrorMap map[string]error
}

//this function fills out our FakeDivvycloudClientTransport attributes keyed on the ClientOperation.ID
//we populate our called map and parameter map to check in our tests
//we then return our configured fake response and error values
func (n *FakeDivvycloudClientTransport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	id := operation.ID
	n.CalledMap[id] = true
	n.ParamsMap[id] = operation.Params
	return n.ResponseMap[id], n.ErrorMap[id]
}

//sets up a DivvyCloudV2 client using our fake transport for testing
func setupFakeClient(transport runtime.ClientTransport, token string) *ClientTokenWrapper {
	return &ClientTokenWrapper{
		Token: token,
		DivvyCloudV2: &client.DivvyCloudV2{
			EventDrivenHarvesting: event_driven_harvesting.New(transport, strfmt.Default),
			Organizations: organizations.New(transport, strfmt.Default),
		},
	}
}
