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

func resourceIpamVlan() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamVlanCreate,
		ReadContext:   resourceIpamVlanRead,
		UpdateContext: resourceIpamVlanUpdate,
		DeleteContext: resourceIpamVlanDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"vid": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: stringInSlice([]string{
					models.VLANStatusValueActive,
					models.VLANStatusValueDeprecated,
					models.VLANStatusValueReserved,
				}),
				Default: models.VLANStatusValueActive,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIpamVlanCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))

	params := &ipam.IpamVlansCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableVLAN{
		Name: &name,
		Vid:  &vid,
		Tags: expandTags(d.Get("tags").([]interface{})),
	}

	if v, ok := d.GetOk("site_id"); ok {
		siteID := int64(v.(int))
		params.Data.Site = &siteID
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("role_id"); ok {
		roleID := int64(v.(int))
		params.Data.Role = &roleID
	}

	resp, err := c.Ipam.IpamVlansCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create prefix: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamVlanRead(ctx, d, m)

	return diags
}

func resourceIpamVlanRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	vlanID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVlansReadParams{
		Context: ctx,
		ID:      vlanID,
	}

	resp, err := c.Ipam.IpamVlansRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get vlan: %v", err)
	}

	d.Set("name", resp.Payload.Name)
	d.Set("vid", resp.Payload.Vid)

	if resp.Payload.Site != nil {
		d.Set("site_id", resp.Payload.Site.ID)
	}

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.Role != nil {
		d.Set("role_id", resp.Payload.Role.ID)
	}

	d.Set("tags", flattenTags(resp.Payload.Tags))

	return diags
}

func resourceIpamVlanUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	vlanID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))

	params := &ipam.IpamVlansPartialUpdateParams{
		Context: ctx,
		ID:      vlanID,
	}

	params.Data = &models.WritableVLAN{
		Name: &name,
		Vid:  &vid,
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		params.Data.Site = &siteID
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("role_id") {
		roleID := int64(d.Get("role_id").(int))
		params.Data.Role = &roleID
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	_, err = c.Ipam.IpamVlansPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update vlan: %v", err)
	}

	return resourceIpamVlanRead(ctx, d, m)
}

func resourceIpamVlanDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	vlanID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVlansDeleteParams{
		Context: ctx,
		ID:      vlanID,
	}

	_, err = c.Ipam.IpamVlansDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete vlan: %v", err)
	}

	d.SetId("")

	return diags
}
