package divvycloud

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sgarlick987/godivvycloud/client/add_cloud_account"
	"github.com/sgarlick987/godivvycloud/client/clouds"
	"github.com/sgarlick987/godivvycloud/models"
	"strconv"
	"testing"
)

func TestResourceDivvycloudAwsCloudAccountStsDelete(t *testing.T) {
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudDeleteByCloudResourceIdPost": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudDeleteByCloudResourceIdPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudDeleteByCloudResourceIdPost": clouds.NewPublicCloudDeleteByCloudResourceIDPostOK(),
		},
	}

	token := "12345"
	resourceId := "divvyorgsvc:1"

	state := &terraform.InstanceState{
		ID: "1/1",
		Attributes: map[string]string{
			"resource_id": resourceId,
		},
	}

	r := resourceDivvycloudAwsCloudAccountSts()
	d := r.Data(state)
	c := setupFakeClient(transport, token)

	err := resourceDivvycloudAwsCloudAccountStsDelete(d, c)

	if err != nil {
		t.Error("no error should be returned for successful delete")
	}

	if !transport.CalledMap["PublicCloudDeleteByCloudResourceIdPost"] {
		t.Error("PublicCloudDeleteByCloudResourceIdPost should be called for delete")
	}

	params := transport.ParamsMap["PublicCloudDeleteByCloudResourceIdPost"].(*clouds.PublicCloudDeleteByCloudResourceIDPostParams)

	if params.XAuthToken != token {
		t.Error("should be called with the supplied token")
	}

	if params.CloudResourceID != resourceId {
		t.Error("should be called with the resourceId")
	}
}

func TestResourceDivvycloudAwsCloudAccountStsDeleteError(t *testing.T) {
	returnedErr := errors.New("error deleting")
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PublicCloudDeleteByCloudResourceIdPost": false,
		},
		ErrorMap: map[string]error{
			"PublicCloudDeleteByCloudResourceIdPost": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PublicCloudDeleteByCloudResourceIdPost": nil,
		},
	}

	token := "12345"
	resourceId := "divvyorgsvc:1"

	state := &terraform.InstanceState{
		ID: "1/1",
		Attributes: map[string]string{
			"resource_id": resourceId,
		},
	}

	r := resourceDivvycloudAwsCloudAccountSts()
	d := r.Data(state)
	c := setupFakeClient(transport, token)

	err := resourceDivvycloudAwsCloudAccountStsDelete(d, c)

	if err != returnedErr {
		t.Error("error should be returned for failed delete")
	}
}

func TestResourceDivvycloudAwsCloudAccountStsCreate(t *testing.T) {
	organizationId := "1"
	cloudId := int32(1)
	accountId := "123456"
	//id := "1/1"
	roleArn := "arn"
	resourceId := "divvysvc:1"
	name := "mycloud"
	cloudType := "AWS"
	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeCloudAddPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeCloudAddPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeCloudAddPost": &add_cloud_account.PrototypeCloudAddPostOK{
				Payload: &models.AddAWSCloudAccount{
					ID:          &cloudId,
					ResourceID:  &resourceId,
					CloudTypeID: &cloudType,
					Name:        &name,
				},
			},
		},
	}

	token := "12345"

	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"name":            name,
			"role_arn":        roleArn,
			"account_id":      accountId,
			"organization_id": organizationId,
		},
	}

	r := resourceDivvycloudAwsCloudAccountSts()
	d := r.Data(state)
	c := setupFakeClient(transport, token)

	err := resourceDivvycloudAwsCloudAccountStsCreate(d, c)

	if err != nil {
		t.Error("error should be nil on successful create")
	}

	if d.Id() != "1/1" {
		t.Error("id should be set on successful create")
	}

	if d.Get("cloud_id").(int) != int(cloudId) {
		t.Error("cloud_id should be set to state")
	}

	if d.Get("name").(string) != name {
		t.Error("name should be set to state")
	}

	if d.Get("cloud_type_id").(string) != cloudType {
		t.Error("name should be set to state")
	}

	if d.Get("resource_id").(string) != resourceId {
		t.Error("name should be set to state")
	}

	if d.Get("organization_id").(string) != organizationId {
		t.Error("organization_id should be set to state")
	}

	if d.Get("role_arn").(string) != roleArn {
		t.Error("roleArn should be set to state")
	}

	if d.Get("account_id").(string) != accountId {
		t.Error("accountId should be set to state")
	}

	if !transport.CalledMap["PrototypeCloudAddPost"] {
		t.Error("PrototypeCloudAddPost should be called on create")
	}

	params := transport.ParamsMap["PrototypeCloudAddPost"].(*add_cloud_account.PrototypeCloudAddPostParams)

	if params.XAuthToken != token {
		t.Error("should be called with supplied token")
	}

	if *params.Body.CreationParams.Name != name {
		t.Error("should be called with supplied name")
	}

	if *params.Body.CreationParams.RoleArn != roleArn {
		t.Error("should be called with supplied roleArn")
	}

	if *params.Body.CreationParams.AccountNumber != accountId {
		t.Error("should be called with supplied account number")
	}
}

func TestResourceDivvycloudAwsCloudAccountStsCreateError(t *testing.T) {
	organizationId := "1"
	accountId := "123456"
	name := "mycloud"
	returnedErr := errors.New("error creating")

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeCloudAddPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeCloudAddPost": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeCloudAddPost": nil,
		},
	}

	token := "12345"

	state := &terraform.InstanceState{
		Attributes: map[string]string{
			"name":            name,
			"account_id":      accountId,
			"organization_id": organizationId,
		},
	}

	r := resourceDivvycloudAwsCloudAccountSts()
	d := r.Data(state)
	c := setupFakeClient(transport, token)

	err := resourceDivvycloudAwsCloudAccountStsCreate(d, c)

	if err != returnedErr {
		t.Error("error should be returned on failed create")
	}
}

func TestResourceDivvycloudAwsCloudAccountStsUpdate(t *testing.T) {
	organizationId := "1"
	cloudId := int32(1)
	accountId := "123456"
	id := "1/1"
	roleArn := "arn"
	resourceId := "divvysvc:1"
	name := "mycloud"
	cloudType := "AWS"

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeCloudUpdateByCloudIdPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeCloudUpdateByCloudIdPost": nil,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeCloudUpdateByCloudIdPost": clouds.NewPrototypeCloudUpdateByCloudIDPostCreated(),
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: id,
		Attributes: map[string]string{
			"cloud_id":        strconv.Itoa(int(cloudId)),
			"account_id":      accountId,
			"cloud_type_id":   cloudType,
			"role_arn":        roleArn,
			"resource_id":     resourceId,
			"organization_id": organizationId,
			"name":            name,
		},
	}

	newName := "newcloud"
	newAccountId := "newAccount"
	newRoleArn := "newrole"
	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"id":         id,
			"name":       newName,
			"role_arn":   newRoleArn,
			"account_id": newAccountId,
			"cloud_id":        strconv.Itoa(int(cloudId)),
			"cloud_type_id":   cloudType,
			"resource_id":     resourceId,
			"organization_id": organizationId,
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	r := resourceDivvycloudAwsCloudAccountSts()
	diff, err := r.Diff(state, conf, nil)

	c := setupFakeClient(transport, token)

	newState, err := r.Apply(state, diff, c)

	if err != nil {
		t.Error("error should not be returned on successful update")
	}

	d := r.Data(newState)

	if !transport.CalledMap["PrototypeCloudUpdateByCloudIdPost"] {
		t.Error("PrototypeCloudUpdateByCloudIdPost should be called")
	}

	if d.Id() != id {
		t.Error("id should be set")
	}

	if d.Get("cloud_id").(int) != int(cloudId) {
		t.Error("cloud_id should be set to state")
	}

	if d.Get("name").(string) != newName {
		t.Error("new name should be set to state")
	}

	if d.Get("cloud_type_id").(string) != cloudType {
		t.Error("name should be set to state")
	}

	if d.Get("resource_id").(string) != resourceId {
		t.Error("name should be set to state")
	}

	if d.Get("organization_id").(string) != organizationId {
		t.Error("organization_id should be set to state")
	}

	if d.Get("role_arn").(string) != newRoleArn {
		t.Error("new roleArn should be set to state")
	}

	if d.Get("account_id").(string) != newAccountId {
		t.Error("new accountId should be set to state")
	}

}

func TestResourceDivvycloudAwsCloudAccountStsUpdateError(t *testing.T) {
	organizationId := "1"
	cloudId := int32(1)
	accountId := "123456"
	id := "1/1"
	roleArn := "arn"
	resourceId := "divvysvc:1"
	name := "mycloud"
	cloudType := "AWS"
	returnedErr := errors.New("error updateing")

	transport := &FakeDivvycloudClientTransport{
		ParamsMap: map[string]runtime.ClientRequestWriter{},
		CalledMap: map[string]bool{
			"PrototypeCloudUpdateByCloudIdPost": false,
		},
		ErrorMap: map[string]error{
			"PrototypeCloudUpdateByCloudIdPost": returnedErr,
		},
		ResponseMap: map[string]interface{}{
			"PrototypeCloudUpdateByCloudIdPost": nil,
		},
	}

	token := "123456"
	state := &terraform.InstanceState{
		ID: id,
		Attributes: map[string]string{
			"cloud_id":        strconv.Itoa(int(cloudId)),
			"account_id":      accountId,
			"cloud_type_id":   cloudType,
			"role_arn":        roleArn,
			"resource_id":     resourceId,
			"organization_id": organizationId,
			"name":            name,
		},
	}

	newName := "newcloud"
	newAccountId := "newAccount"
	newRoleArn := "newrole"
	raw, err := config.NewRawConfig(
		map[string]interface{}{
			"id":         id,
			"name":       newName,
			"role_arn":   newRoleArn,
			"account_id": newAccountId,
			"cloud_id":        strconv.Itoa(int(cloudId)),
			"cloud_type_id":   cloudType,
			"resource_id":     resourceId,
			"organization_id": organizationId,
		})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	conf := terraform.NewResourceConfig(raw)
	r := resourceDivvycloudAwsCloudAccountSts()
	diff, err := r.Diff(state, conf, nil)

	c := setupFakeClient(transport, token)

	_, err = r.Apply(state, diff, c)

	if err != returnedErr {
		t.Error("error should be returned on failed update")
	}
}
