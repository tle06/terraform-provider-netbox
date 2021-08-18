package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimRegion() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimRegionCreate,
		ReadContext:   resourceDcimRegionRead,
		UpdateContext: resourceDcimRegionUpdate,
		DeleteContext: resourceDcimRegionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringLenBetween(0, 50),
			},

			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},

			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"description": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringLenBetween(0, 200),
			},
		},
	}
}

func resourceDcimRegionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	var name = d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimRegionsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableRegion{
		Name: &name,
		Slug: &slug,
	}

	if v, ok := d.GetOk("parent_id"); ok {
		parentID := int64(v.(int))
		params.Data.Parent = &parentID
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	resp, err := c.Dcim.DcimRegionsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create region: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceDcimRegionRead(ctx, d, m)

	return diags
}

func resourceDcimRegionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	regionID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimRegionsReadParams{
		Context: ctx,
		ID:      regionID,
	}

	resp, err := c.Dcim.DcimRegionsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get region: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("slug", resp.Payload.Slug)

	if resp.Payload.Parent != nil {
		d.Set("parent_id", resp.Payload.Parent.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	return diags
}

func resourceDcimRegionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	regionID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimRegionsPartialUpdateParams{
		Context: ctx,
		ID:      regionID,
	}

	params.Data = &models.WritableRegion{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("parent_id") {
		parentID := int64(d.Get("parent_id").(int))
		params.Data.Parent = &parentID
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	_, err = c.Dcim.DcimRegionsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update region: %v", err)
	}

	return resourceDcimRegionRead(ctx, d, m)
}

func resourceDcimRegionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	regionID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimRegionsDeleteParams{
		Context: ctx,
		ID:      regionID,
	}

	_, err = c.Dcim.DcimRegionsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete region: %v", err)
	}

	d.SetId("")

	return diags
}
