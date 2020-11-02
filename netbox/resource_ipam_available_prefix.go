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

func resourceIpamAvailablePrefix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamAvailablePrefixCreate,
		ReadContext:   resourceIpamAvailablePrefixRead,
		DeleteContext: resourceIpamAvailablePrefixDelete,

		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"family": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceIpamAvailablePrefixCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &ipam.IpamPrefixesAvailablePrefixesCreateParams{
		Context: ctx,
		ID:      int64(d.Get("prefix_id").(int)),
	}

	prefixLength := int64(d.Get("prefix_length").(int))
	params.Data = &models.PrefixLength{
		PrefixLength: &prefixLength,
	}

	resp, err := c.Ipam.IpamPrefixesAvailablePrefixesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create available prefix: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamAvailablePrefixRead(ctx, d, m)

	return diags
}

func resourceIpamAvailablePrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamPrefixesReadParams{
		Context: ctx,
		ID:      prefixID,
	}

	resp, err := c.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get prefix: %v", err)
	}

	d.Set("family", resp.Payload.Family.Value)
	d.Set("prefix", resp.Payload.Prefix)

	return diags
}

func resourceIpamAvailablePrefixDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamPrefixesDeleteParams{
		Context: ctx,
		ID:      prefixID,
	}

	_, err = c.Ipam.IpamPrefixesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete prefix: %v", err)
	}

	d.SetId("")

	return diags
}
