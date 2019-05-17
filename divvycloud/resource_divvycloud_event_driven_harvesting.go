package divvycloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
	"log"
)

func resourceDivvycloudEventDrivenHarvesting() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudEventDrivenHarvestingCreate,
		Read:   resourceDivvycloudEventDrivenHarvestingRead,
		Update: resourceDivvycloudEventDrivenHarvestingUpdate,
		Delete: resourceDivvycloudEventDrivenHarvestingDelete,

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Enable Event Driven Harvesting",
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

func resourceDivvycloudEventDrivenHarvestingCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] creating event driven harvest")

	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	enabled := d.Get("enabled").(bool)
	organizationId := d.Get("organization_id").(string)

	if enabled {
		if _, err := c.PublicCloudEventdrivenharvestEnable(
			event_driven_harvesting.NewPublicCloudEventdrivenharvestEnableParams().WithXAuthToken(token)); err != nil {
			return err
		}
	}

	d.SetId(organizationId)

	return resourceDivvycloudEventDrivenHarvestingRead(d, meta)
}

func resourceDivvycloudEventDrivenHarvestingRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] reading event driven harvest")

	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	ok, err := c.PublicCloudEventdrivenharvest(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestParams().WithXAuthToken(token))

	if err != nil {
		return err
	}

	return d.Set("enabled", ok.Payload.EventDrivenHarvestEnabled)
}

func resourceDivvycloudEventDrivenHarvestingUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] updating event driven harvest")

	if d.HasChange("enabled") {
		if d.Get("enabled").(bool) {
			return resourceDivvycloudEventDrivenHarvestingCreate(d, meta)
		} else {
			return resourceDivvycloudEventDrivenHarvestingDelete(d, meta)
		}
	}

	return nil
}

func resourceDivvycloudEventDrivenHarvestingDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] deleting event driven harvest")

	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	_, err := c.PublicCloudEventdrivenharvestDisable(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestDisableParams().WithXAuthToken(token))

	return err
}
