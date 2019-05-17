package divvycloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sgarlick987/godivvycloud/client/event_driven_harvesting"
	"github.com/sgarlick987/godivvycloud/models"
	"log"
	"strconv"
)

func resourceDivvycloudEventDrivenHarvestingProducer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudEventDrivenHarvestingProducerCreate,
		Read:   resourceDivvycloudEventDrivenHarvestingProducerRead,
		Delete: resourceDivvycloudEventDrivenHarvestingProducerDelete,

		Schema: map[string]*schema.Schema{

			"consumer_cloud_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of the consumer cloud to create this producer under",
			},
			"cloud_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of the cloud to be added as a producer for event driven harvesting",
			},
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Organization to Enable Event Driven Harvesting",
			},
			"enable_all_types": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enable all types including new ones in the future",
				Default:     true,
				ForceNew:    true,
			},
		},
	}
}

func resourceDivvycloudEventDrivenHarvestingProducerCreate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	log.Printf("[DEBUG] creating event driven harvest producer")
	consumerCloudId := d.Get("consumer_cloud_id").(string)
	cloudId := d.Get("cloud_id").(string)
	regions := []string{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-central-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"sa-east-1",
		"ap-south-1",
		"ca-central-1",
		"eu-north-1",
	}
	resourceTypes := []string{
		"bigdatainstance",
		"autoscalinggroup",
		"autoscalinglaunchconfiguration",
		"instance",
		"volume",
		"snapshot",
		"storagecontainer",
		"resourceaccesslist",
		"loadbalancer",
		"privatenetwork",
		"privatesubnet",
		"secret",
		"serviceuser",
		"servicegroup",
		"servicerole",
		"serviceaccesskey",
		"serviceencryptionkey",
		"sshkeypair",
		"networkinterface",
		"internetgateway",
		"natgateway",
		"routetable",
		"servicepolicy",
		"dbinstance",
		"dbcluster",
		"dbsnapshot",
		"identityprovider",
		"serviceloggroup",
		"hypervisor",
		"networkflowlog",
		"networkpeer",
		"awsplacementgroup",
		"mcinstance",
		"esinstance",
		"notificationtopic",
		"notificationsubscription",
		"stacktemplate",
		"divvyorganizationservice",
		"restapikey",
	}
	enableAllTypes := d.Get("enable_all_types").(bool)
	badges := []string{}
	byBadge := false
	byConsumer := false

	//this is just us converting the cloudId to its divvy native name and to an int32 for the generate godivvycloud client
	orgServiceId, err := strconv.Atoi(cloudId)

	if err != nil {
		log.Fatal(err)
	}

	body := &models.AddProducerRequest{
		OrganizationServiceIds: []int32{int32(orgServiceId)},
		Regions:                regions,
		ResourceTypes:          resourceTypes,
		EnableAllTypes:         &enableAllTypes,
		Propagate: &models.Propagate{
			Badges:     badges,
			ByBadge:    &byBadge,
			ByConsumer: &byConsumer,
		},
	}

	if _, err := c.PublicCloudEventdrivenharvestConsumerProducersAddByOrganizationidPost(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestConsumerProducersAddByOrganizationidPostParams().
			WithXAuthToken(token).
			WithOrganizationid(consumerCloudId).
			WithBody(body)); err != nil {
		return err
	}

	d.SetId(fmt.Sprint(cloudId))

	return resourceDivvycloudEventDrivenHarvestingProducerRead(d, meta)
}
func resourceDivvycloudEventDrivenHarvestingProducerRead(d *schema.ResourceData, meta interface{}) error {
	//token := meta.(*ClientTokenWrapper).Token
	//c := meta.(*ClientTokenWrapper).EventDrivenHarvesting
	return nil
}

func resourceDivvycloudEventDrivenHarvestingProducerDelete(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).EventDrivenHarvesting

	cloudId := d.Get("cloud_id").(string)
	resourceId := fmt.Sprintf("divvyorganizationservice:%s", cloudId)


	body := &models.RemoveProducerRequest{
		ResourceIds: []string{
			resourceId,
		},
	}

	if _, err := c.PublicCloudEventdrivenharvestProducersDisablePost(
		event_driven_harvesting.NewPublicCloudEventdrivenharvestProducersDisablePostParams().
			WithXAuthToken(token).
			WithBody(body)); err != nil {
		return err
	}

	return nil
}
