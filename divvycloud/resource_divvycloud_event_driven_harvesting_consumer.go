package divvycloud

import "github.com/hashicorp/terraform/helper/schema"

func resourceDivvycloudEventDrivenHarvestingConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudEventDrivenHarvestingConsumerCreate,
		Read:   resourceDivvycloudEventDrivenHarvestingConsumerRead,
		Update: resourceDivvycloudEventDrivenHarvestingConsumerUpdate,
		Delete: resourceDivvycloudEventDrivenHarvestingConsumerDelete,

		Schema: map[string]*schema.Schema{

		},
	}
}

func resourceDivvycloudEventDrivenHarvestingConsumerCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingConsumerRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingConsumerUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingConsumerDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
