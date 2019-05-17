package divvycloud

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
	"github.com/sgarlick987/godivvycloud/models"
	"testing"
)

/**
TODO: dry up tests
*/

func TestResourceDivvycloudEventDrivenHarvestingDeleteError(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestDisable": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestDisable": errors.New("generic error deleting event driven harvesting"),
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestDisable": nil,
		},
	}

	token := "123456"

	c := setupFakeClient(transport, token)
	d := &schema.ResourceData{}

	err := resourceDivvycloudEventDrivenHarvestingDelete(d, c)

	if err == nil {
		t.Error("should return an error when PublicCloudEventdrivenharvestDisable fails")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingDelete(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestDisable": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestDisable": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestDisable": event_driven_harvesting.NewPublicCloudEventdrivenharvestDisableOK(),
		},
	}
	token := "123456"

	c := setupFakeClient(transport, token)
	d := &schema.ResourceData{}

	err := resourceDivvycloudEventDrivenHarvestingDelete(d, c)

	if err != nil {
		t.Error("success should return nil error")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestDisable"] {
		t.Error("success should call PublicCloudEventdrivenharvestDisable")
	}

	params := transport.ParamsMap["PublicCloudEventdrivenharvestDisable"].(*event_driven_harvesting.PublicCloudEventdrivenharvestDisableParams)

	if params.XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingUpdateHasChangeEnable(t *testing.T) {
	enabledOk := &event_driven_harvesting.PublicCloudEventdrivenharvestEnableOK{}
	status := true

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestEnable": false,
			"PublicCloudEventdrivenharvest":       false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestEnable": nil,
			"PublicCloudEventdrivenharvest":       nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestEnable": enabledOk,
			"PublicCloudEventdrivenharvest": &event_driven_harvesting.PublicCloudEventdrivenharvestOK{
				Payload: &models.EventDrivenHarvestingStatus{
					EventDrivenHarvestEnabled: &status,
				},
			},
		},
	}

	token := "123456"

	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"enabled":         "false",
			"organization_id": "1",
		},
	}

	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"enabled":         true,
			"organization_id": "1",
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	resource := resourceDivvycloudEventDrivenHarvesting()

	diff, err := resource.Diff(state, conf, nil)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	c := setupFakeClient(transport, token)

	newState, err := resource.Apply(state, diff, c)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	data := resource.Data(newState)

	if data.Id() != "1" {
		t.Error("id should be set to organization id on successful update")
	}

	if transport.
		ParamsMap["PublicCloudEventdrivenharvestEnable"].(*event_driven_harvesting.PublicCloudEventdrivenharvestEnableParams).
		XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if transport.
		ParamsMap["PublicCloudEventdrivenharvest"].(*event_driven_harvesting.PublicCloudEventdrivenharvestParams).
		XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestEnable"] {
		t.Error("create should be called")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvest"] {
		t.Error("read should be called")
	}

	if transport.ErrorMap["PublicCloudEventdrivenharvest"] != nil {
		t.Error("read should not return an error")
	}

	if transport.ErrorMap["PublicCloudEventdrivenharvestEnable"] != nil {
		t.Error("enable should not return an error")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingUpdateHasChangeDisable(t *testing.T) {
	disabledOk := &event_driven_harvesting.PublicCloudEventdrivenharvestDisableOK{}

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestDisable": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestDisable": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestDisable": disabledOk,
		},
	}

	token := "123456"

	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"id":              "1",
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"id":              "1",
			"enabled":         false,
			"organization_id": "1",
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	resource := resourceDivvycloudEventDrivenHarvesting()

	diff, err := resource.Diff(state, conf, nil)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	c := setupFakeClient(transport, token)

	newState, err := resource.Apply(state, diff, c)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	data := resource.Data(newState)

	if data.Id() != "1" {
		t.Error("id should be set to organization id on successful update")
	}

	if transport.
		ParamsMap["PublicCloudEventdrivenharvestDisable"].(*event_driven_harvesting.PublicCloudEventdrivenharvestDisableParams).
		XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestDisable"] {
		t.Error("delete should be called")
	}

	if transport.ErrorMap["PublicCloudEventdrivenharvestDisable"] != nil {
		t.Error("Disable should not return an error")
	}
}
func TestResourceDivvycloudEventDrivenHarvestingUpdateHasNoChange(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestEnable":  false,
			"PublicCloudEventdrivenharvestDisable": false,
		},
		ErrorMap:    map[string]error{},
		ResponseMap: map[string]interface{}{},
	}

	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	d := resource.Data(state)

	c := setupFakeClient(transport, "")

	if err := resourceDivvycloudEventDrivenHarvestingUpdate(d, c); err != nil {
		t.Error("nil should be returned when no changes")
	}

	if d.Id() != "1" {
		t.Error("id should be set to organization id on no changes")
	}

	if transport.CalledMap["PublicCloudEventdrivenharvestDisable"] {
		t.Error("delete should not be called when no changes")
	}

	if transport.CalledMap["PublicCloudEventdrivenharvestEnable"] {
		t.Error("create should not be called when no changes")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingCreateEnabled(t *testing.T) {
	enabledOk := &event_driven_harvesting.PublicCloudEventdrivenharvestEnableOK{}
	status := true

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestEnable": false,
			"PublicCloudEventdrivenharvest":       false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestEnable": nil,
			"PublicCloudEventdrivenharvest":       nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestEnable": enabledOk,
			"PublicCloudEventdrivenharvest": &event_driven_harvesting.PublicCloudEventdrivenharvestOK{
				Payload: &models.EventDrivenHarvestingStatus{
					EventDrivenHarvestEnabled: &status,
				},
			},
		},
	}
	token := "123456"
	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	c := setupFakeClient(transport, token)
	d := resource.Data(state)

	err := resourceDivvycloudEventDrivenHarvestingCreate(d, c)

	if err != nil {
		t.Error("success should return nil error")
	}

	if !transport.CalledMap["PublicCloudEventdrivenharvestEnable"] {
		t.Error("success should call PublicCloudEventdrivenharvestEnable")
	}

	params := transport.ParamsMap["PublicCloudEventdrivenharvestEnable"].(*event_driven_harvesting.PublicCloudEventdrivenharvestEnableParams)

	if params.XAuthToken != token {
		t.Error("should be called with XAuthToken from ClientTokenWrapper")
	}

	if d.Id() != "1" {
		t.Error("id should set to organization_id on success")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingCreateDisabled(t *testing.T) {
	enabledOk := &event_driven_harvesting.PublicCloudEventdrivenharvestEnableOK{}
	status := false

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestEnable": false,
			"PublicCloudEventdrivenharvest":       false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestEnable": nil,
			"PublicCloudEventdrivenharvest":       nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestEnable": enabledOk,
			"PublicCloudEventdrivenharvest": &event_driven_harvesting.PublicCloudEventdrivenharvestOK{
				Payload: &models.EventDrivenHarvestingStatus{
					EventDrivenHarvestEnabled: &status,
				},
			},
		},
	}
	token := "123456"
	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"enabled":         "false",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	c := setupFakeClient(transport, token)
	d := resource.Data(state)

	err := resourceDivvycloudEventDrivenHarvestingCreate(d, c)

	if err != nil {
		t.Error("success should return nil error")
	}

	if transport.CalledMap["PublicCloudEventdrivenharvestEnable"] {
		t.Error("should not call PublicCloudEventdrivenharvestEnable when disabled")
	}

	if d.Id() != "1" {
		t.Error("id should set to organization_id on success")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingCreateError(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvestEnable": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvestEnable": errors.New("error creating event driven harvesting"),
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvestEnable": nil,
		},
	}
	token := "123456"
	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	c := setupFakeClient(transport, token)
	d := resource.Data(state)

	err := resourceDivvycloudEventDrivenHarvestingCreate(d, c)

	if err == nil {
		t.Error("create failure should return error")
	}

	if d.Id() != "" {
		t.Error("id should not be set on creation error")
	}
}

func TestResourceDivvycloudEventDrivenHarvestingRead(t *testing.T) {
	status := true
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvest": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvest": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvest": &event_driven_harvesting.PublicCloudEventdrivenharvestOK{
				Payload: &models.EventDrivenHarvestingStatus{
					EventDrivenHarvestEnabled: &status,
				},
			},
		},
	}
	token := "123456"
	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	c := setupFakeClient(transport, token)
	d := resource.Data(state)

	err := resourceDivvycloudEventDrivenHarvestingRead(d, c)

	if err != nil {
		t.Error("successful read should not return error")
	}

	if d.Id() != "1" {
		t.Error("id should be set on successful read")
	}
}
func TestResourceDivvycloudEventDrivenHarvestingReadError(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudEventdrivenharvest": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudEventdrivenharvest": errors.New("error reading resource"),
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudEventdrivenharvest": nil,
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: "1",
		Attributes: map[string]string{
			"enabled":         "true",
			"organization_id": "1",
		},
	}

	resource := resourceDivvycloudEventDrivenHarvesting()
	c := setupFakeClient(transport, token)
	d := resource.Data(state)

	if err := resourceDivvycloudEventDrivenHarvestingRead(d, c); err == nil {
		t.Error("error should be returned for read failure")
	}
}
