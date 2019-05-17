package divvycloud

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sgarlick987/godivvycloud/client/organizations"
	"github.com/sgarlick987/godivvycloud/models"
	"strconv"
	"testing"
)

func TestResourceDivvycloudOrganizationDelete(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete": nil,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete":
			&organizations.PrototypeDomainOrganizationDeleteByOrganizationResourceIDDeleteOK{},
		},
	}

	token := "123456"
	orgResourceId := "divvyorg:1"
	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"name":            "myorg",
			"organization_id": "1",
			"resource_id":     orgResourceId,
		},
	}
	r := resourceDivvycloudOrganization()
	d := r.Data(state)

	c := setupFakeClient(transport, token)
	err := resourceDivvycloudOrganizationDelete(d, c)

	if err != nil {
		t.Error("error should be nil on successful delete")
	}

	if !transport.CalledMap["PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete"] {
		t.Error("PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete should be called")
	}

	params := transport.ParamsMap["PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete"].(*organizations.PrototypeDomainOrganizationDeleteByOrganizationResourceIDDeleteParams)

	if params.XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if params.OrganizationResourceID != orgResourceId {
		t.Error("delete should be called with the organizations resource id")
	}

}

func TestResourceDivvycloudOrganizationDeleteError(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete": errors.New("failed to delete org"),
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete": nil,
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"name":            "myorg",
			"organization_id": "1",
			"resource_id":     "divvyorg:1",
		},
	}

	r := resourceDivvycloudOrganization()
	d := r.Data(state)

	c := setupFakeClient(transport, token)
	err := resourceDivvycloudOrganizationDelete(d, c)

	if err == nil {
		t.Error("error should be returned on failed delete")
	}

	if !transport.CalledMap["PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete"] {
		t.Error("PrototypeDomainOrganizationDeleteByOrganizationResourceIdDelete should be called")
	}

}

func TestResourceDivvycloudOrganizationCreate(t *testing.T) {
	organizationId := 1
	resourceId := "divvyorg:1"
	name := "myorg"

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationCreatePost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationCreatePost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationCreatePost":
			&organizations.PrototypeDomainOrganizationCreatePostOK{
				Payload: &models.DomainOrganizationCreateResponse{
					ResourceID:     resourceId,
					OrganizationID: int32(organizationId),
				},
			},
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"name": name,
		},
	}

	r := resourceDivvycloudOrganization()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudOrganizationCreate(d, c)

	if err != nil {
		t.Error("error should be nil on successful create")
	}

	if !transport.CalledMap["PrototypeDomainOrganizationCreatePost"] {
		t.Error("PrototypeDomainOrganizationCreatePost should be called")
	}

	params := transport.ParamsMap["PrototypeDomainOrganizationCreatePost"].(*organizations.PrototypeDomainOrganizationCreatePostParams)

	if params.XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if *params.Body.OrganizationName != name {
		t.Error("create should be called with the organizations name")
	}

	if d.Id() != resourceId {
		t.Error("id should be set to returned resourceId")
	}

	if d.Get("resource_id").(string) != resourceId {
		t.Error("resource id should be set to state")
	}

	if d.Get("organization_id").(string) != strconv.Itoa(organizationId) {
		t.Error("organization id should be set to state")
	}

}

func TestResourceDivvycloudOrganizationCreateError(t *testing.T) {
	returnErr := errors.New("error creating")
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationCreatePost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationCreatePost": returnErr,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationCreatePost": nil,
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"name": "myorg",
		},
	}

	r := resourceDivvycloudOrganization()
	d := r.Data(state)
	c := setupFakeClient(transport, token)
	err := resourceDivvycloudOrganizationCreate(d, c)

	if err != returnErr {
		t.Error("error should be returned on failed create")
	}

	if !transport.CalledMap["PrototypeDomainOrganizationCreatePost"] {
		t.Error("PrototypeDomainOrganizationCreatePost should be called")
	}

	if d.Id() != "" {
		t.Error("id should not be set on failed create")
	}
}

func TestResourceDivvycloudOrganizationUpdate(t *testing.T) {
	organizationId := 1
	resourceId := "divvyorg:1"
	name := "myorg"
	newName := "notmyorg"

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost":
			&organizations.PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK{},
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: resourceId,
		Attributes: map[string]string{
			"resource_id":     resourceId,
			"organization_id": strconv.Itoa(organizationId),
			"name":            name,
		},
	}

	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"name": newName,
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	r := resourceDivvycloudOrganization()
	diff, err := r.Diff(state, conf, nil)

	c := setupFakeClient(transport, token)

	newState, err := r.Apply(state, diff, c)

	if err != nil {
		t.Error("error should not be returned on successful update")
	}

	d := r.Data(newState)

	if !transport.CalledMap["PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost"] {
		t.Error("PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost should be called")
	}

	if d.Get("name").(string) != newName {
		t.Error("name should be updated o the new name")
	}

	if d.Get("resource_id").(string) != resourceId {
		t.Error("resource id should be same after update")
	}

	if d.Get("organization_id").(string) != strconv.Itoa(organizationId) {
		t.Error("organization id should be same after update")
	}

	if d.Id() != resourceId {
		t.Error("id should be set on successful update")
	}
}

func TestResourceDivvycloudOrganizationUpdateError(t *testing.T) {
	organizationId := 1
	resourceId := "divvyorg:1"
	name := "myorg"
	newName := "notmyorg"
	returnedErr := errors.New("failed to update")

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeDomainOrganizationUpdateByOrganizationResourceIdPost": nil,
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: resourceId,
		Attributes: map[string]string{
			"resource_id":     resourceId,
			"organization_id": strconv.Itoa(organizationId),
			"name":            name,
		},
	}

	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"name":            newName,
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	r := resourceDivvycloudOrganization()
	diff, err := r.Diff(state, conf, nil)

	c := setupFakeClient(transport, token)

	_, err = r.Apply(state, diff, c)

	if err != returnedErr {
		t.Error("error should be returned on failed update")
	}
}
