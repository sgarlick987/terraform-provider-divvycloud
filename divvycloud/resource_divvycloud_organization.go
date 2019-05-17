package divvycloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sgarlick987/godivvycloud/client/organizations"
	"github.com/sgarlick987/godivvycloud/models"
	"strconv"
)

func resourceDivvycloudOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceDivvycloudOrganizationCreate,
		Read:   resourceDivvycloudOrganizationRead,
		Update: resourceDivvycloudOrganizationUpdate,
		Delete: resourceDivvycloudOrganizationDelete,

		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Id of the created Organization.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Organization to be created.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Id of the created Organization.",
			},
		},
	}
}

func resourceDivvycloudOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).Organizations

	name := d.Get("name").(string)

	params := organizations.NewPrototypeDomainOrganizationCreatePostParams().
		WithXAuthToken(token).
		WithBody(&models.CreateneworganizationRequest{
			OrganizationName: &name,
		})

	resp, err := c.PrototypeDomainOrganizationCreatePost(params)

	if err != nil {
		return err
	}

	payload := resp.Payload
	resourceId := payload.ResourceID

	d.SetId(resourceId)

	if err := d.Set("resource_id", resourceId); err != nil {
		return err
	}

	if err := d.Set("organization_id", strconv.Itoa(int(payload.OrganizationID))); err != nil {
		return err
	}

	return resourceDivvycloudOrganizationRead(d, meta)
}

func resourceDivvycloudOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	//token := meta.(*ClientTokenWrapper).Token
	//c := meta.(*ClientTokenWrapper).Resources

	return nil
}

func resourceDivvycloudOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).Organizations

	name := d.Get("name").(string)
	resourceId := d.Get("resource_id").(string)

	params := organizations.NewPrototypeDomainOrganizationUpdateByOrganizationResourceIDPostParams().
		WithXAuthToken(token).
		WithOrganizationResourceID(resourceId).
		WithBody(&models.EditorganizationnameRequest{
			OrganizationName: &name,
		})

	if _, err := c.PrototypeDomainOrganizationUpdateByOrganizationResourceIDPost(params); err != nil {
		return err
	}

	return resourceDivvycloudOrganizationRead(d, meta)
}

func resourceDivvycloudOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	token := meta.(*ClientTokenWrapper).Token
	c := meta.(*ClientTokenWrapper).Organizations

	params := organizations.NewPrototypeDomainOrganizationDeleteByOrganizationResourceIDDeleteParams().
		WithXAuthToken(token).
		WithOrganizationResourceID(d.Get("resource_id").(string))

	if _, err := c.PrototypeDomainOrganizationDeleteByOrganizationResourceIDDelete(params); err != nil {
		return err
	}

	return nil
}
