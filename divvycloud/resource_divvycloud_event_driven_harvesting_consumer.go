package divvycloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
)

func resourceDivvycloudEventDrivenHarvestingConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudEventDrivenHarvestingConsumerCreate,
		Read:   resourceDivvycloudEventDrivenHarvestingConsumerRead,
		Delete: resourceDivvycloudEventDrivenHarvestingConsumerDelete,

		Schema: map[string]*schema.Schema{
			"cloud_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of the cloud to be added as a consumer for event driven harvesting",
			},
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Organization to Enable Event Driven Harvesting Consumer",
			},
		},
	}
}

func resourceDivvycloudEventDrivenHarvestingConsumerCreate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	cloudId := d.Get("cloud_id").(string)
	organizationId := d.Get("organization_id").(string)

	if _, err := c.PublicCloudEventdrivenharvestByOrganizationidPost(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestByOrganizationidPostParams().
			WithXAuthToken(token).
			WithOrganizationid(cloudId)); err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", organizationId, cloudId))

	return nil
}

func resourceDivvycloudEventDrivenHarvestingConsumerRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDivvycloudEventDrivenHarvestingConsumerDelete(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	cloudId := d.Get("cloud_id").(string)

	if _, err := c.PublicCloudEventdrivenharvestDisableConsumerByOrganizationidDelete(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestDisableConsumerByOrganizationidDeleteParams().
			WithXAuthToken(token).
			WithOrganizationid(cloudId)); err != nil {
		return err
	}

	return nil
}
