package divvycloud

import "github.com/hashicorp/terraform/helper/schema"

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
		},
	}
}

func resourceDivvycloudEventDrivenHarvestingCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
