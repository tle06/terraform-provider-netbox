package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpamRir() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamRirCreate,
		ReadContext:   resourceIpamRirRead,
		UpdateContext: resourceIpamRirUpdate,
		DeleteContext: resourceIpamRirDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},

		},
	}
}

func resourceIpamRirCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &ipam.IpamRirsCreateParams{
		Context: ctx,
	}

	params.Data = &models.RIR{
		Name: &name,
		Slug: &slug,
	}

	resp, err := c.Ipam.IpamRirsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create rir: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamRirRead(ctx, d, m)

	return diags
}

func resourceIpamRirRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	rirID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamRirsReadParams{
		Context: ctx,
		ID:      rirID,
	}

	resp, err := c.Ipam.IpamRirsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get rir: %v", err)
	}
	d.Set("slug", resp.Payload.Slug)

	return diags
}

func resourceIpamRirUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	rirID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &ipam.IpamRirsPartialUpdateParams{
		Context: ctx,
		ID:      rirID,
	}

	params.Data = &models.RIR{
		Name: &name,
		Slug: &slug,
	}

	_, err = c.Ipam.IpamRirsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update rir: %v", err)
	}

	return resourceIpamRirRead(ctx, d, m)
}

func resourceIpamRirDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	rirID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamRirsDeleteParams{
		Context: ctx,
		ID:      rirID,
	}

	_, err = c.Ipam.IpamRirsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete rir: %v", err)
	}

	d.SetId("")

	return diags
}
