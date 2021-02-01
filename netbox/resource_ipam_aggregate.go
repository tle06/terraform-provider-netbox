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

func resourceIpamAggregate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamAggregateCreate,
		ReadContext:   resourceIpamAggregateRead,
		UpdateContext: resourceIpamAggregateUpdate,
		DeleteContext: resourceIpamAggregateDelete,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isCIDR,
			},

			"rir_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"family": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIpamAggregateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	prefix := d.Get("prefix").(string)

	params := &ipam.IpamAggregatesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableAggregate{
		Prefix: &prefix,
	}

	if v, ok := d.GetOk("rir_id"); ok {
		rirID := int64(v.(int))
		params.Data.Rir = &rirID
	}

	resp, err := c.Ipam.IpamAggregatesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create aggregate: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamAggregateRead(ctx, d, m)

	return diags
}

func resourceIpamAggregateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamAggregatesReadParams{
		Context: ctx,
		ID:      prefixID,
	}

	resp, err := c.Ipam.IpamAggregatesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get aggregate: %v", err)
	}

	d.Set("family", resp.Payload.Family.Label)
	d.Set("prefix", resp.Payload.Prefix)

	return diags
}

func resourceIpamAggregateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	prefix := d.Get("prefix").(string)

	params := &ipam.IpamAggregatesPartialUpdateParams{
		Context: ctx,
		ID:      prefixID,
	}

	params.Data = &models.WritableAggregate{
		Prefix: &prefix,
	}

	_, err = c.Ipam.IpamAggregatesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update aggregate: %v", err)
	}

	return resourceIpamAggregateRead(ctx, d, m)
}

func resourceIpamAggregateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamAggregatesDeleteParams{
		Context: ctx,
		ID:      prefixID,
	}

	_, err = c.Ipam.IpamAggregatesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete aggregate: %v", err)
	}

	d.SetId("")

	return diags
}
