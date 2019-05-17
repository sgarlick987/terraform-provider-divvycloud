package divvycloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sgarlick987/godivvycloud/client/add_cloud_account"
	"github.com/sgarlick987/godivvycloud/client/clouds"
	"github.com/sgarlick987/godivvycloud/models"
)

func resourceDivvycloudAwsCloudAccountSts() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudAwsCloudAccountStsCreate,
		Read:   resourceDivvycloudAwsCloudAccountStsRead,
		Update: resourceDivvycloudAwsCloudAccountStsUpdate,
		Delete: resourceDivvycloudAwsCloudAccountStsDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud to be created.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS Account Id for the cloud to be created.",
			},
			"role_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS Account Id for the cloud to be created.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Id of the cloud created.",
			},
			"cloud_id": {

				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Integer id of the cloud created.",
			},
			"cloud_type_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Type Id of the cloud created.",
			},
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Organization to Enable Event Driven Harvesting",
			},
		},
	}
}

func resourceDivvycloudAwsCloudAccountStsCreate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).AddCloudAccount

	sessionName := "divvycloud-devops"
	authenticationType := "instance_assume_role"
	cloudType := "AWS"
	accountId := d.Get("account_id").(string)
	name := d.Get("name").(string)
	roleArn := d.Get("role_arn").(string)
	organizationId := d.Get("organization_id").(string)

	sts := &models.AddAWSCloudAccountInstanceAssumeRoleRequest{
		CreationParams: &models.CreationParams1{
			AccountNumber:      &accountId,
			RoleArn:            &roleArn,
			Name:               &name,
			SessionName:        &sessionName,
			AuthenticationType: &authenticationType,
			CloudType:          &cloudType,
		},
	}

	params := add_cloud_account.NewPrototypeCloudAddPostParams().
		WithXAuthToken(token).
		WithBody(sts)

	resp, err := c.PrototypeCloudAddPost(params)

	if err != nil {
		return err
	}

	resourceId := resp.Payload.ResourceID

	d.SetId(fmt.Sprintf("%s/%s", organizationId, *resourceId))
	if err := d.Set("resource_id", *resourceId); err != nil {
		return err
	}
	if err := d.Set("cloud_type_id", *resp.Payload.CloudTypeID); err != nil {
		return err
	}
	if err := d.Set("cloud_id", *resp.Payload.ID); err != nil {
		return err
	}

	return resourceDivvycloudAwsCloudAccountStsRead(d, meta)
}

func resourceDivvycloudAwsCloudAccountStsRead(d *schema.ResourceData, meta interface{}) error {
	//token := meta.(*ClientTokenWrapper).Token
	//c := meta.(*ClientTokenWrapper).Resources
	//
	//resourceId := d.Get("resource_id").(string)
	//params := resources.NewPublicResourceDetailByResourceIDGetParams().
	//	WithXAuthToken(token).
	//	WithResourceID(resourceId)
	//
	//_, err := c.PublicResourceDetailByResourceIDGet(params)
	//
	//if err != nil {
	//	return err
	//}

	//_ := resp.Payload.Details

	return nil
}

func resourceDivvycloudAwsCloudAccountStsUpdate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).Clouds

	sessionName := "divvycloud-devops"
	authenticationType := "instance_assume_role"
	cloudType := "AWS"
	accountId := d.Get("account_id").(string)
	name := d.Get("name").(string)
	roleArn := d.Get("role_arn").(string)

	sts := &models.CreationParams1{
		AccountNumber:      &accountId,
		RoleArn:            &roleArn,
		Name:               &name,
		SessionName:        &sessionName,
		AuthenticationType: &authenticationType,
		CloudType:          &cloudType,
	}

	params := clouds.NewPrototypeCloudUpdateByCloudIDPostParams().
		WithXAuthToken(token).
		WithCloudID(int64(d.Get("cloud_id").(int))).
		WithBody(sts)

	if _, err := c.PrototypeCloudUpdateByCloudIDPost(params); err != nil {
		return err
	}

	return resourceDivvycloudAwsCloudAccountStsRead(d, meta)
}

func resourceDivvycloudAwsCloudAccountStsDelete(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).Clouds

	params := clouds.NewPublicCloudDeleteByCloudResourceIDPostParams().
		WithXAuthToken(token).
		WithCloudResourceID(d.Get("resource_id").(string))

	if _, err := c.PublicCloudDeleteByCloudResourceIDPost(params); err != nil {
		return err
	}

	return nil
}
