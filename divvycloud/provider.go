package divvycloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sgarlick987/godivvycloud/client"
	"net/url"
)

func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full address for the Divvycloud install.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username for connecting to the Divvycloud install.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password for connecting to the Divvycloud install.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"divvycloud_aws_cloud_account_sts":            resourceDivvycloudAwsCloudAccountSts(),
			"divvycloud_organization":                     resourceDivvycloudOrganization(),
			"divvycloud_event_driven_harvesting":          resourceDivvycloudEventDrivenHarvesting(),
			"divvycloud_event_driven_harvesting_consumer": resourceDivvycloudEventDrivenHarvestingConsumer(),
			"divvycloud_event_driven_harvesting_producer": resourceDivvycloudEventDrivenHarvestingProducer(),
		},
		ConfigureFunc: providerConfigure,
	}

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	addressUrl, err := url.Parse(address)

	if err != nil {
		return nil, err
	}

	transport := &client.TransportConfig{
		BasePath: addressUrl.Path,
		Host:     addressUrl.Host,
		Schemes:  []string{addressUrl.Scheme},
	}

	c := &LoginConfig{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Client:   client.NewHTTPClientWithConfig(nil, transport),
	}

	return c.WrappedClient()

}
