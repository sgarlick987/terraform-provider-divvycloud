package divvycloud

import "github.com/hashicorp/terraform/helper/schema"

func resourceDivvycloudEventDrivenHarvestingProducer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudEventDrivenHarvestingProducerCreate,
		Read:   resourceDivvycloudEventDrivenHarvestingProducerRead,
		Update: resourceDivvycloudEventDrivenHarvestingProducerUpdate,
		Delete: resourceDivvycloudEventDrivenHarvestingProducerDelete,

		Schema: map[string]*schema.Schema{

		},
	}
}

func resourceDivvycloudEventDrivenHarvestingProducerCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingProducerRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingProducerUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceDivvycloudEventDrivenHarvestingProducerDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
