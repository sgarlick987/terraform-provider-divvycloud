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
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	log.Printf("[DEBUG] creating event driven harvest")
	enabled := d.Get("enabled").(bool)
	organizationId := d.Get("organization_id").(string)

	d.SetId(organizationId)

	if enabled {
		if _, err := c.PublicCloudEventdrivenharvestEnable(
			event_driven_harvesting.NewPublicCloudEventdrivenharvestEnableParams().WithXAuthToken(token)); err != nil {
			return err
		}
	}
	return resourceDivvycloudEventDrivenHarvestingRead(d, meta)
}

func resourceDivvycloudEventDrivenHarvestingRead(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	log.Printf("[DEBUG] reading event driven harvest")
	ok, err := c.PublicCloudEventdrivenharvest(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestParams().WithXAuthToken(token))

	if err != nil {
		return err
	}

	if err = d.Set("enabled", ok.Payload.EventDrivenHarvestEnabled); err != nil {
		return err
	}

	return nil
}

func resourceDivvycloudEventDrivenHarvestingUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] updating event driven harvest")

	if d.HasChange("enabled") {
		enabled := d.Get("enabled").(bool)

		if enabled {
			if err := resourceDivvycloudEventDrivenHarvestingCreate(d, meta); err != nil {
				return err
			}
		} else {
			if err := resourceDivvycloudEventDrivenHarvestingDelete(d, meta); err != nil {
				return err
			}
		}
	}

	return resourceDivvycloudEventDrivenHarvestingRead(d, meta)
}

func resourceDivvycloudEventDrivenHarvestingDelete(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	log.Printf("[DEBUG] deleting event driven harvest")

	if _, err := c.PublicCloudEventdrivenharvestDisable(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestDisableParams().WithXAuthToken(token)); err != nil {
		return err
	}

	return nil
}
