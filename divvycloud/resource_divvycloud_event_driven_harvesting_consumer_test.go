package divvycloud

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
	"testing"
)

func TestResourceDivvycloudEventDrivenHarvestingConsumerRead(t *testing.T) {
	err := resourceDivvycloudEventDrivenHarvestingConsumerRead(nil, nil)
	if err != nil {
		t.Error("successful read should return nil error")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingConsumerDelete(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": event_driven_harvesting.NewPublicCloudEventdrivenharvestDisableConsumerByOrganizationidDeleteOK(),
		},
	}

	organizationId := "1"
	cloudId := "1"
	id := fmt.Sprintf("%s/%s", organizationId, cloudId)
	state := &terraform.InstanceState{
		ID: id,
		Attributes: map[string]string{
			"cloud_id":        cloudId,
			"organization_id": organizationId,
		},
	}

	token := "123456"
	r := resourceDivvycloudEventDrivenHarvestingConsumer()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudEventDrivenHarvestingConsumerDelete(d, c)

	if err != nil {
		t.Error("no error should be returned for successful delete")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete"] {
		t.Error("PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete should be called for delete")
	}

	params := transport.
		ParamsMap["PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete"].(*event_driven_harvesting.PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDeleteParams)

	if params.XAuthToken != token {
		t.Error("should be called with supplied token")
	}

	//yes orgid != cloud id, this seems to be a misnamed parameter in the swagger doc for the generated client
	if params.Organizationid != cloudId {
		t.Error("given cloud id should be use for delete")
	}

	if d.Id() != id {
		t.Error("id should be set to orgid/cloudid")
	}

	if d.Get("cloud_id").(string) != cloudId {
		t.Error("cloud id should be set on state")
	}

	if d.Get("organization_id").(string) != organizationId {
		t.Error("organization id should be set on state")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingConsumerDeleteError(t *testing.T) {
	returnedErr := errors.New("error deleting consumer")
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete": nil,
		},
	}

	organizationId := "1"
	cloudId := "1"
	id := fmt.Sprintf("%s/%s", organizationId, cloudId)
	state := &terraform.InstanceState{
		ID: id,
		Attributes: map[string]string{
			"cloud_id":        cloudId,
			"organization_id": organizationId,
		},
	}

	token := "123456"
	r := resourceDivvycloudEventDrivenHarvestingConsumer()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudEventDrivenHarvestingConsumerDelete(d, c)

	if err != returnedErr {
		t.Error("error should be returned for failed delete")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingConsumerCreate(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestByOrganizationidPost": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestByOrganizationidPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestByOrganizationidPost":
			event_driven_harvesting.NewPublicCloudEventdrivenharvestByOrganizationidPostOK(),
		},
	}

	organizationId := "1"
	cloudId := "1"
	id := fmt.Sprintf("%s/%s", organizationId, cloudId)

	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"cloud_id":        cloudId,
			"organization_id": organizationId,
		},
	}

	token := "123456"
	r := resourceDivvycloudEventDrivenHarvestingConsumer()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudEventDrivenHarvestingConsumerCreate(d, c)

	if err != nil {
		t.Error("error should not be returned on successful create")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestByOrganizationidPost"] {
		t.Error("PublicCloudEventdrivenharvestByOrganizationidPost should be called on create")
	}

	if d.Id() != id {
		t.Error("id should be set on successful create")
	}

	if d.Get("cloud_id").(string) != cloudId {
		t.Error("cloud_id should be set on state on successful create")
	}

	if d.Get("organization_id").(string) != cloudId {
		t.Error("organization_id should be set on state on successful create")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingConsumerCreateError(t *testing.T) {
	returnedErr := errors.New("error on create")
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestByOrganizationidPost": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestByOrganizationidPost": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestByOrganizationidPost": nil,
		},
	}

	organizationId := "1"
	cloudId := "1"

	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"cloud_id":        cloudId,
			"organization_id": organizationId,
		},
	}

	token := "123456"
	r := resourceDivvycloudEventDrivenHarvestingConsumer()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudEventDrivenHarvestingConsumerCreate(d, c)

	if err != returnedErr {
		t.Error("error should be returned on failed create")
	}
}
