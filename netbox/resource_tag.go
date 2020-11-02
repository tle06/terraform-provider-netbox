package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/extras"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},

			"color": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &extras.ExtrasTagsCreateParams{
		Context: ctx,
	}

	params.Data = &models.Tag{
		Name: &name,
		Slug: &slug,
	}

	if v, ok := d.GetOk("color"); ok {
		params.Data.Color = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	resp, err := c.Extras.ExtrasTagsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create tag: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceTagRead(ctx, d, m)

	return diags
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &extras.ExtrasTagsReadParams{
		Context: ctx,
		ID:      id,
	}

	resp, err := c.Extras.ExtrasTagsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get tag: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("slug", resp.Payload.Slug)
	d.Set("color", resp.Payload.Color)
	d.Set("description", resp.Payload.Description)

	return diags
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &extras.ExtrasTagsPartialUpdateParams{
		Context: ctx,
		ID:      id,
	}

	params.Data = &models.Tag{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("color") {
		params.Data.Color = d.Get("color").(string)
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	_, err = c.Extras.ExtrasTagsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update tag: %v", err)
	}

	return resourceTagRead(ctx, d, m)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &extras.ExtrasTagsDeleteParams{
		Context: ctx,
		ID:      id,
	}

	_, err = c.Extras.ExtrasTagsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete tag: %v", err)
	}

	d.SetId("")

	return diags
}
